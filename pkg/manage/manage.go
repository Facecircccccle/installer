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

type NodeOutputStructure struct {
	Name	string
	Status	string
	Roles	string
	Time	string
	Version	string
}

func GetNodeInternalIP(node *v1.Node) string {
	for _, address := range node.Status.Addresses {
		if address.Type == v1.NodeInternalIP {
			return address.Address
		}
	}

	return "<none>"
}

func GetKubeInfo(clientset *kubernetes.Clientset) (string, string, string){
	var masterIP, containerRuntimeVersion string
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	if strings.Contains(nodes.Items[0].Status.NodeInfo.ContainerRuntimeVersion, "docker"){
		containerRuntimeVersion = FormatDockerVersion(nodes.Items[0].Status.NodeInfo.ContainerRuntimeVersion)
	}

	for i := 0; i < len(nodes.Items); i++ {
		if strings.Contains(strings.Join(FindNodeRoles(&nodes.Items[i]), ","), "master"){
			masterIP = GetNodeInternalIP(&nodes.Items[i])
			break
		}
	}

	return nodes.Items[0].Status.NodeInfo.KubeletVersion, containerRuntimeVersion, masterIP
}

func FormatDockerVersion(version string) string {
	var buffer bytes.Buffer
	str := regexp.MustCompile("[0-9]+").FindAllString(version, -1)
	buffer.WriteString("docker-ce-")

	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(str[i])
		buffer.WriteString(fmt.Sprintf("%02d", num))
		buffer.WriteString(".")
	}

	return buffer.String()[0:len(buffer.String())-1]
}

func TranslateTimestampSince(timestamp metav1.Time) string {
	if timestamp.IsZero() {
		return "<unknown>"
	}

	return duration.HumanDuration(time.Since(timestamp.Time))
}

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