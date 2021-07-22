package manage

import (
	"bytes"
	"context"
	"fmt"
	"installer/pkg/constants"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/duration"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// NodeOutputStructure struct.
type NodeOutputStructure struct {
	Name    string
	Status  string
	Roles   string
	Time    string
	Version string
}

type PVOutputStructure struct {
	Name          string
	Capacity      string
	AccessModes   string
	ReclaimPolicy string
	Status        string
	Claim         string
	StorageClass  string
	Reason        string
	Age           string
}

type NamespaceOutPutStructure struct {
	Name   string
	Status string
	Age    string
}

type NodeStatusOutPutStructure struct {
	Name string
	CPU string
	CPUPer string
	MEMORY string
	MEMORYPer string
}

type SCOutPutStructure struct {
	Name                 string
	Provisioner          string
	ReclaimPolicy        string
	VolumeBindingMode    string
	AllowVolumeExpansion string
}

// GetNodeInternalIP
func GetNodeInternalIP(node *v1.Node) string {
	for _, address := range node.Status.Addresses {
		if address.Type == v1.NodeInternalIP {
			return address.Address
		}
	}

	return "<none>"
}

// GetKubeInfo returns kubelet version, runtime version and master ip.
func GetKubeInfo(clientset *kubernetes.Clientset) (string, string, string) {
	var masterIP, containerRuntimeVersion string
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	if strings.Contains(nodes.Items[0].Status.NodeInfo.ContainerRuntimeVersion, "docker") {
		containerRuntimeVersion = FormatDockerVersion(nodes.Items[0].Status.NodeInfo.ContainerRuntimeVersion)
	}

	for i := 0; i < len(nodes.Items); i++ {
		if strings.Contains(strings.Join(FindNodeRoles(&nodes.Items[i]), ","), "master") {
			masterIP = GetNodeInternalIP(&nodes.Items[i])
			break
		}
	}

	return nodes.Items[0].Status.NodeInfo.KubeletVersion, containerRuntimeVersion, masterIP
}

// FormatDockerVersion.
func FormatDockerVersion(version string) string {
	var buffer bytes.Buffer
	str := regexp.MustCompile("[0-9]+").FindAllString(version, -1)
	buffer.WriteString("docker-ce-")

	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(str[i])
		buffer.WriteString(fmt.Sprintf("%02d", num))
		buffer.WriteString(".")
	}

	return buffer.String()[0 : len(buffer.String())-1]
}

// TranslateTimestampSince returns the elapsed time since timestamp in human-readable approximation.
func TranslateTimestampSince(timestamp metav1.Time) string {
	if timestamp.IsZero() {
		return "<unknown>"
	}

	return duration.HumanDuration(time.Since(timestamp.Time))
}

// FindNodeRoles returns node roles.
func FindNodeRoles(node *v1.Node) []string {
	roles := sets.NewString()
	for k, v := range node.Labels {
		switch {
		case strings.HasPrefix(k, constants.LabelNodeRolePrefix):
			if role := strings.TrimPrefix(k, constants.LabelNodeRolePrefix); len(role) > 0 {
				roles.Insert(role)
			}

		case k == constants.NodeLabelRole && v != "":
			roles.Insert(v)
		}
	}
	return roles.List()
}

// GetAccessModesAsString returns a string representation of an array of access modes.
// modes, when present, are always in the same order: RWO,ROX,RWX.
func GetAccessModesAsString(modes []v1.PersistentVolumeAccessMode) string {
	modes = removeDuplicateAccessModes(modes)
	var modesStr []string
	if containsAccessMode(modes, v1.ReadWriteOnce) {
		modesStr = append(modesStr, "RWO")
	}
	if containsAccessMode(modes, v1.ReadOnlyMany) {
		modesStr = append(modesStr, "ROX")
	}
	if containsAccessMode(modes, v1.ReadWriteMany) {
		modesStr = append(modesStr, "RWX")
	}
	return strings.Join(modesStr, ",")
}

// GetPersistentVolumeClass returns StorageClassName.
func GetPersistentVolumeClass(volume *v1.PersistentVolume) string {
	// Use beta annotation first
	if class, found := volume.Annotations[v1.BetaStorageClassAnnotation]; found {
		return class
	}

	return volume.Spec.StorageClassName
}

// removeDuplicateAccessModes returns an array of access modes without any duplicates
func removeDuplicateAccessModes(modes []v1.PersistentVolumeAccessMode) []v1.PersistentVolumeAccessMode {
	var accessModes []v1.PersistentVolumeAccessMode
	for _, m := range modes {
		if !containsAccessMode(accessModes, m) {
			accessModes = append(accessModes, m)
		}
	}
	return accessModes
}

func containsAccessMode(modes []v1.PersistentVolumeAccessMode, mode v1.PersistentVolumeAccessMode) bool {
	for _, m := range modes {
		if m == mode {
			return true
		}
	}
	return false
}
