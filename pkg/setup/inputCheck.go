package setup

import (
	"github.com/rivo/tview"
	"installer/pkg/util"
	"installer/pkg/version"
	"strings"
)

// AddRoleCheck verifies that you entered the correct IP and network card when adding a new role.
func AddRoleCheck(form *tview.Form, isHA bool) (b bool, s string) {
	ip := form.GetFormItemByLabel("IP").(*tview.InputField).GetText()
	if ip == "" || !util.IsIp(ip) {
		return false, "Illegal IP address"
	}

	//	util.ExecShell("rm -rf /etc/ansible/hosts")
	util.ExecShell("sh localScript/add_ansible_host.sh " + "hostnameCheck " + ip)
	if !util.SshSuccess(ip) {
		return false, "IP can not be reached, check ssh connect"
	}
	util.ExecShell("sh localScript/delete_ansible_host.sh " + "hostnameCheck")
	if isHA {
		netCard := form.GetFormItemByLabel("Network Card").(*tview.InputField).GetText()
		if !util.CheckNetCard(ip, netCard) {
			return false, "Can not find target net card named " + "" + " on remote server, please use 'ip addr' to check."
		}
	}
	return true, ""
}

// InputRoleBackCheck verifies that the number of masters matches the current pattern.
func InputRoleBackCheck(setup *Setup, isHA bool) (bool, string) {
	if setup.MasterCount < 1 {
		return false, "The cluster must have one master at least."
	}
	if setup.MasterCount > 1 && !isHA {
		return false, "This is single master cluster part, make sure the master number is 1. If you still need more masters, choose BACK and use HA setup."
	}
	if setup.MasterCount == 1 && isHA {
		return false, "Multiple masters are typically used in the HA cluster. If you need only one master, choose BACK and use Basic setup."
	}
	if (setup.MasterCount%2) == 0 && isHA {
		return false, "According to ETCD fault tolerance requirements, an odd number of master machines are required."
	}
	return true, ""
}

// EtcdAndKubernetesVersionCheck verifies that the ETCD version matches the current Kubernetes version.
func EtcdAndKubernetesVersionCheck(e *Etcds, s *Setup) (bool, string) {
	_, etcdVersion := e.GetFormItemByLabel("Version").(*tview.DropDown).GetCurrentOption()

	if util.StringIsInArray(etcdVersion, version.GetComponentVersion()[s.Kubernetes.Version].EtcdVersion) {
		return true, ""
	}
	return false, "In kubernetes " + s.Kubernetes.Version + ", etcd version should be " + strings.Join(version.GetComponentVersion()[s.Kubernetes.Version].EtcdVersion, ", ")
}

// DockerAndKubernetesVersionCheck verifies that the Docker version matches the current Kubernetes version.
func DockerAndKubernetesVersionCheck(d *Dockers, s *Setup) (bool, string) {
	_, dockerVersion := d.GetFormItemByLabel("Version").(*tview.DropDown).GetCurrentOption()

	if util.StringIsInArray(dockerVersion, version.GetComponentVersion()[s.Kubernetes.Version].DockerVersion) {
		return true, ""
	}
	return false, "In kubernetes " + s.Kubernetes.Version + ", docker version should be " + strings.Join(version.GetComponentVersion()[s.Kubernetes.Version].DockerVersion, ", ")
}

// NetworkingCheck verifies that you entered the correct PodSubnet and ServiceSubnet.
func NetworkingCheck(networking *NetWorkings) (b bool, s string) {
	var reason = ""
	if !util.IsIpWithSubnet(networking.GetFormItemByLabel("PodSubnet").(*tview.InputField).GetText()) {
		return false, "Illegal Pod Subnet address"
	}
	if !util.IsIpWithSubnet(networking.GetFormItemByLabel("ServiceSubnet").(*tview.InputField).GetText()) {
		return false, "Illegal Service Subnet address"
	}
	return true, reason
}

// ClusterInfoCheck verifies that you entered the correct Controller manager and scheduler address, virtual IP and certSANs for HA as well.
func ClusterInfoCheck(c *Clusters, setup *Setup) (b bool, s string) {
	var reason = ""
	if !util.IsIp(c.GetFormItemByLabel("ControllerManager_bind_address").(*tview.InputField).GetText()) {
		return false, "Illegal controller manager bind address"
	}
	if !util.IsIp(c.GetFormItemByLabel("Scheduler_address").(*tview.InputField).GetText()) {
		return false, "Illegal Scheduler address"
	}

	if setup.MasterCount > 1 {
		if c.GetFormItemByLabel("Virtual_IP").(*tview.InputField).GetText() == "" {
			return false, "If the number of masters is greater than 1, Virtual_IP is needed."
		}
		if !util.IsIp(c.GetFormItemByLabel("Virtual_IP").(*tview.InputField).GetText()) {
			return false, "Illegal Virtual IP address"
		}
		//检查vip 和master处于同一子网
		if !util.IsVipInSameSubnet(c.GetFormItemByLabel("Virtual_IP").(*tview.InputField).GetText(), setup.Masters[0].IPAddr, util.GetNetmask(setup.Masters[0].NetCard)) {
			return false, "Virtual IP address needs to be in the same subnet with Maters"
		}
		if c.GetFormItemByLabel("CertSANs").(*tview.InputField).GetText() == "" {
			return false, "If the number of masters is greater than 1, CertSANs (e.g. cluster.<name>.com etc..) is needed."
		}
	}
	return true, reason
}

// AllocateCheck verifies that you entered the correct number in node allocate part.
func AllocateCheck(a *Allocates) (b bool, s string) {
	if !util.IsDigit(a.GetFormItemByLabel("KubeReservedCPU").(*tview.InputField).GetText()) {
		return false, "Illegal Kube Reserved CPU"
	}
	if !util.IsDigit(a.GetFormItemByLabel("SysReservedCPU").(*tview.InputField).GetText()) {
		return false, "Illegal Sys Reserved CPU"
	}
	if !util.IsDigitWithStorage(a.GetFormItemByLabel("KubeMemory").(*tview.InputField).GetText()) {
		return false, "Illegal Kube Memory"
	}
	if !util.IsDigitWithStorage(a.GetFormItemByLabel("SysMemory").(*tview.InputField).GetText()) {
		return false, "Illegal Sys Memory"
	}
	if !util.IsDigitWithStorage(a.GetFormItemByLabel("KubeStorage").(*tview.InputField).GetText()) {
		return false, "Illegal Kube Storage"
	}
	if !util.IsDigitWithStorage(a.GetFormItemByLabel("SysStorage").(*tview.InputField).GetText()) {
		return false, "Illegal Sys Storage"
	}
	if !util.IsDigitWithStorage(a.GetFormItemByLabel("EvictionMemory").(*tview.InputField).GetText()) {
		return false, "Illegal Eviction Memory"
	}
	if !util.IsDigitWithPercent(a.GetFormItemByLabel("EvictionNodefs").(*tview.InputField).GetText()) {
		return false, "Illegal Eviction Node"
	}

	return true, ""
}
