package setup

import (
	"encoding/json"
)

type Setup struct {
	MasterCount int
	NodeCount   int
	AccessCount int

	Masters  []Master
	Nodes    []Node
	Accesses []Access

	Docker     DockerSetup
	Etcd       EtcdSetup
	Kubernetes KubernetesSetup
}

type Master struct {
	IPAddr  string
	NetCard string
	//Name string
	//UserName string
	//code string
}

type Node struct {
	IPAddr string
	//Name     string
	//UserName string
	//code 	 string
	Allocate NodeAllocate
}

type Access struct {
	IPAddr string
	//Name string
	//UserName string
	//code string
}

type DockerSetup struct {
	RepositoryName string
	RepositoryIP   string
	Version        string
}

type EtcdSetup struct {
	Version string
	DataDir string
}

type FeatureGates struct {
	AllowExtTrafficLocalEndpoints   bool
	DynamicVolumeProvisioning       bool
	VolumeSubpath                   bool
	StorageObjectInUseProtection    bool
	SupportIPVSProxyMode            bool
	AdvancedAuditing                bool
	MountPropagation                bool
	CSIPersistentVolume             bool
	GCERegionalPersistentDisk       bool
	KubeletPluginsWatcher           bool
	VolumeScheduling                bool
	CustomPodDNS                    bool
	HugePages                       bool
	PersistentLocalVolumes          bool
	PodPriority                     bool
	PodReadinessGates               bool
	CustomResourcePublishOpenAPI    bool
	CustomResourceSubresources      bool
	CustomResourceValidation        bool
	CustomResourceWebhookConversion bool
	AttachVolumeLimit               bool
	CustomResourceDefaulting        bool
	NodeLease                       bool
	PodShareProcessNamespace        bool
	ResourceQuotaScopeSelectors     bool
	ScheduleDaemonSetPods           bool
	ServiceLoadBalancerFinalizer    bool
	TaintNodesByCondition           bool
	VolumeSubpathEnvExpansion       bool
	WatchBookmark                   bool

	BlockVolume                 bool
	CSIBlockVolume              bool
	ExternalPolicyForExternalIP bool
	TaintBasedEvictions         bool
	VolumePVCDataSource         bool
	WindowsGMSA                 bool
	WindowsRunAsUserName        bool

	DryRun                         bool
	EvenPodsSpread                 bool
	RotateKubeletClientCertificate bool
	StreamingProxyRedirects        bool
	ExecProbeTimeout               bool
	RuntimeClass                   bool
	SCTPSupport                    bool
	ServiceAppProtocol             bool
	StartupProbe                   bool
	SupportNodePidsLimit           bool
	SupportPodPidsLimit            bool
	TokenRequest                   bool
	TokenRequestProjection         bool
	VolumeSnapshotDataSource       bool

	CRIContainerLogRotation       bool
	EndpointSlice                 bool
	EndpointSliceNodeName         bool
	LegacyNodeRoleBehavior        bool
	NodeDisruptionExclusion       bool
	PodDisruptionBudget           bool
	RootCAConfigMap               bool
	ServiceAccountIssuerDiscovery bool
	ServiceNodeExclusion          bool
}

type KubernetesSetup struct {
	ClusterName           string
	ControllerManagerAddr string //0.0.0.0
	SchedulerAddr         string //0.0.0.0
	ControlPlaneEndpoint  string
	VirtualIP             string
	CertificatesDir       string
	ImageRepository       string
	Networking            KubeNetwork
	Version               string
	CertSANs              string //cluster.k8s.ebupt.com
	NetComponent          NetComponentSetup
	AdmissionPlugin       AdmissionPlugin
	FeatureGates          FeatureGates
}

type KubeNetwork struct {
	PodSubnet     string //10.244.0.0/16
	ServiceSubnet string //10.96.0.0/12
	DNSdomain     string
}

type NetComponentSetup struct {
	Component string
	Version   string
}

type AdmissionPlugin struct {
	NamespaceLifecycle         bool
	LimitRanger                bool
	ServiceAccount             bool
	DefaultStorageClass        bool
	DefaultTolerationSeconds   bool
	MutatingAdmissionWebhook   bool
	ValidatingAdmissionWebhook bool
	ResourceQuota              bool
}

type NodeAllocate struct {
	KubeReservedCPU string
	SysReservedCPU  string
	KubeMemory      string //500Mi
	SysMemory       string
	KubeStorage     string //10Gi
	SysStorage      string
	EvictionMemory  string //500Mi
	EvictionNodefs  string //10%
}

func NewSampleNodeAllocate() NodeAllocate {
	return NodeAllocate{
		KubeReservedCPU: "1",
		SysReservedCPU:  "1",
		KubeMemory:      "500Mi",
		SysMemory:       "500Mi",
		KubeStorage:     "10Gi",
		SysStorage:      "10Gi",
		EvictionMemory:  "500Mi",
		EvictionNodefs:  "10%",
	}
}
func NewSampleSetupStructure() *Setup {

	Netcomponentsetup := &NetComponentSetup{
		Component: "",
		Version:   "",
	}

	Kubenetwork := &KubeNetwork{
		PodSubnet:     "",
		ServiceSubnet: "",
		DNSdomain:     "",
	}

	AdmissionPlugin := &AdmissionPlugin{
		NamespaceLifecycle:         false,
		LimitRanger:                false,
		ServiceAccount:             false,
		DefaultStorageClass:        false,
		DefaultTolerationSeconds:   false,
		MutatingAdmissionWebhook:   false,
		ValidatingAdmissionWebhook: false,
		ResourceQuota:              false,
	}

	Kubernetessetup := &KubernetesSetup{
		ControllerManagerAddr: "", //0.0.0.0
		SchedulerAddr:         "", //0.0.0.0
		Networking:            *Kubenetwork,
		Version:               "",
		VirtualIP:             "",
		CertSANs:              "", //cluster.k8s.ebupt.com
		NetComponent:          *Netcomponentsetup,
		ControlPlaneEndpoint:  "",
		CertificatesDir:       "",
		ImageRepository:       "",
		AdmissionPlugin:       *AdmissionPlugin,
	}

	EtcdSetup := &EtcdSetup{
		Version: "3.4.12-0",
		DataDir: "/var",
	}

	Dockersetup := &DockerSetup{
		RepositoryName: "string",
		RepositoryIP:   "string",
		Version:        "docker-ce-19.03.15-3.el7",
	}

	return &Setup{
		MasterCount: 0,
		NodeCount:   0,
		AccessCount: 0,

		Masters: []Master{

		},

		Nodes: []Node{

		},

		Docker:     *Dockersetup,
		Etcd:       *EtcdSetup,
		Kubernetes: *Kubernetessetup,
	}
}

func StructureToJSON(s Setup) string {
	jsonByte, _ := json.Marshal(s)
	return string(jsonByte)
}
