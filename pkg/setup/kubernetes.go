package setup

import (
	"github.com/rivo/tview"
	"installer/pkg/constants"
	"installer/pkg/util"
	"strings"
)

type Cluster struct {
	ClusterName                  string
	Version                      string
	ControllermanagerBindAddress string
	SchedulerAddress             string
	CertSANs                     string
	ControlPlaneEndpoint         string
	certificatesDir              string
	ImageRepository              string
	useHyperKubeImage            bool
}

type Networking struct {
	PodSubnet     string
	ServiceSubnet string
	DNSdomain     string
}

type NetPlugin struct {
	Plugin  string
	Version string
}

type Clusters struct {
	*tview.Form
}

type NetWorkings struct {
	*tview.Form
}
type NetPlugins struct {
	*tview.Form
}
type AdmissionPlugins struct {
	*tview.Form
}

func (a AdmissionPlugins) SetEntires(s *Setup) {

	s.Kubernetes.AdmissionPlugin.NamespaceLifecycle = a.GetFormItemByLabel("NamespaceLifecycle").(*tview.Checkbox).IsChecked()
	s.Kubernetes.AdmissionPlugin.LimitRanger = a.GetFormItemByLabel("LimitRanger").(*tview.Checkbox).IsChecked()
	s.Kubernetes.AdmissionPlugin.ServiceAccount = a.GetFormItemByLabel("ServiceAccount").(*tview.Checkbox).IsChecked()
	s.Kubernetes.AdmissionPlugin.DefaultStorageClass = a.GetFormItemByLabel("DefaultStorageClass").(*tview.Checkbox).IsChecked()
	s.Kubernetes.AdmissionPlugin.DefaultTolerationSeconds = a.GetFormItemByLabel("DefaultTolerationSeconds").(*tview.Checkbox).IsChecked()
	s.Kubernetes.AdmissionPlugin.MutatingAdmissionWebhook = a.GetFormItemByLabel("MutatingAdmissionWebhook").(*tview.Checkbox).IsChecked()
	s.Kubernetes.AdmissionPlugin.ValidatingAdmissionWebhook = a.GetFormItemByLabel("ValidatingAdmissionWebhook").(*tview.Checkbox).IsChecked()
	s.Kubernetes.AdmissionPlugin.ResourceQuota = a.GetFormItemByLabel("ResourceQuota").(*tview.Checkbox).IsChecked()

}

func NewCluster() *Clusters {

	clusters := &Clusters{
		Form: tview.NewForm().
			AddInputField("ClusterName", "KubernetesCluster", 0, nil, nil).
			AddDropDown("Version", constants.KubernetesVersion, 0, nil).
			AddInputField("ControllerManager_bind_address", "", 0, nil, nil).
			AddInputField("Scheduler_address", "", 0, nil, nil).
			AddInputField("Virtual_IP", "", 0, nil, nil).
			AddInputField("CertSANs", "", 0, nil, nil).
			AddInputField("certificatesDir", "/etc/kubernetes/pki", 0, nil, nil).
			AddInputField("ImageRepository", "", 0, nil, nil),
	}

	clusters.SetBorder(true).SetTitle("Cluster Info").SetTitleAlign(tview.AlignCenter)
	clusters.SetItemPadding(0).SetBorderPadding(0, 0, 0, 1)

	return clusters

}
func (c Clusters) SetEntries(s *Setup) {

	if s.MasterCount > 1 {
		s.Kubernetes.VirtualIP = c.GetFormItemByLabel("Virtual_IP").(*tview.InputField).GetText()
		s.Kubernetes.ControlPlaneEndpoint = c.GetFormItemByLabel("CertSANs").(*tview.InputField).GetText() + ":" + constants.HaPort
	} else if c.GetFormItemByLabel("CertSANs").(*tview.InputField).GetText() != "" {
		s.Kubernetes.ControlPlaneEndpoint = c.GetFormItemByLabel("CertSANs").(*tview.InputField).GetText() + ":" + constants.ClusterPort
	}

	s.Kubernetes.ClusterName = c.GetFormItemByLabel("ClusterName").(*tview.InputField).GetText()
	_, s.Kubernetes.Version = c.GetFormItemByLabel("Version").(*tview.DropDown).GetCurrentOption()
	s.Kubernetes.ControllerManagerAddr = c.GetFormItemByLabel("ControllerManager_bind_address").(*tview.InputField).GetText()
	s.Kubernetes.SchedulerAddr = c.GetFormItemByLabel("Scheduler_address").(*tview.InputField).GetText()
	s.Kubernetes.CertSANs = c.GetFormItemByLabel("CertSANs").(*tview.InputField).GetText()
	s.Kubernetes.CertificatesDir = c.GetFormItemByLabel("certificatesDir").(*tview.InputField).GetText()
	s.Kubernetes.ImageRepository = c.GetFormItemByLabel("ImageRepository").(*tview.InputField).GetText()

	//i.SetText("Clusters  " + StructureToJSON(*s))
}

func NewNetwork() *NetWorkings {

	netWorkings := &NetWorkings{
		Form: tview.NewForm().
			AddInputField("PodSubnet", "10.244.0.0/16", 0, nil, nil).
			AddInputField("ServiceSubnet", "10.96.0.0/12", 0, nil, nil).
			AddInputField("DNSdomain", "", 0, nil, nil),
	}
	netWorkings.SetBorder(true).SetTitle("Network Info").SetTitleAlign(tview.AlignCenter)
	netWorkings.SetItemPadding(0).SetBorderPadding(0, 0, 0, 1)
	return netWorkings
}
func (n NetWorkings) SetEntries(s *Setup) {

	s.Kubernetes.Networking.PodSubnet = n.GetFormItemByLabel("PodSubnet").(*tview.InputField).GetText()
	s.Kubernetes.Networking.ServiceSubnet = n.GetFormItemByLabel("ServiceSubnet").(*tview.InputField).GetText()
	s.Kubernetes.Networking.DNSdomain = n.GetFormItemByLabel("DNSdomain").(*tview.InputField).GetText()

	//i.SetText("network  " + StructureToJSON(*s))
}

func NewNetPlugin() *NetPlugins {

	netPlugins := &NetPlugins{
		Form: tview.NewForm().
			AddDropDown("Plugin", []string{"calico", "flannel"}, 0, nil),
	}
	netPlugins.SetBorder(true).SetTitle("Net Plugin Info").SetTitleAlign(tview.AlignCenter)
	netPlugins.SetItemPadding(0).SetBorderPadding(0, 0, 0, 1)

	return netPlugins
}

func (n NetPlugins) SetEntries(s *Setup) {
	_, s.Kubernetes.NetComponent.Component = n.GetFormItemByLabel("Plugin").(*tview.DropDown).GetCurrentOption()

	//i.SetText("net plugins  " + StructureToJSON(*s))
}

func NewAdmission() *AdmissionPlugins {
	admissionPlugins := &AdmissionPlugins{
		Form: tview.NewForm().
			AddCheckbox("NamespaceLifecycle", false, nil).
			AddCheckbox("LimitRanger", false, nil).
			AddCheckbox("ServiceAccount", false, nil).
			AddCheckbox("DefaultStorageClass", false, nil).
			AddCheckbox("DefaultTolerationSeconds", false, nil).
			AddCheckbox("MutatingAdmissionWebhook", false, nil).
			AddCheckbox("ValidatingAdmissionWebhook", false, nil).
			AddCheckbox("ResourceQuota", false, nil),
	}

	admissionPlugins.SetBorder(true).SetTitle("Admission Plugin Info").SetTitleAlign(tview.AlignCenter)
	admissionPlugins.SetItemPadding(0).SetBorderPadding(0, 0, 0, 1)

	return admissionPlugins
}

func GetAdmissionPlugins(plugin AdmissionPlugin) string {
	result := ""
	if plugin.DefaultStorageClass == true {
		result = util.StringAppend(result, "DefaultStorageClass,")
	}
	if plugin.DefaultTolerationSeconds == true {
		result = util.StringAppend(result, "DefaultTolerationSeconds,")
	}
	if plugin.LimitRanger == true {
		result = util.StringAppend(result, "LimitRanger,")
	}
	if plugin.MutatingAdmissionWebhook == true {
		result = util.StringAppend(result, "MutatingAdmissionWebhook,")
	}
	if plugin.NamespaceLifecycle == true {
		result = util.StringAppend(result, "NamespaceLifecycle,")
	}
	if plugin.ResourceQuota == true {
		result = util.StringAppend(result, "ResourceQuota,")
	}
	if plugin.ServiceAccount == true {
		result = util.StringAppend(result, "ServiceAccount,")
	}
	if plugin.ValidatingAdmissionWebhook == true {
		result = util.StringAppend(result, "ValidatingAdmissionWebhook,")
	}

	return strings.TrimRight(result, ",")
}
