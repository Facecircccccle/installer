package ui

import (
	"installer/pkg/setup"
	"installer/pkg/util"
	"strconv"
)

// ProcessNewCluster sends commands to the remote host to complete the installation process.
func ProcessNewCluster(g *Gui, s *setup.Setup, ansibleLog *myText) {
	processAnsibleHosts(g, s, ansibleLog)

	copyScript(g, ansibleLog)
	processKubeadmConfig(s)
	processGeneralScript(g, s, ansibleLog)
	processHAScript(g, s, ansibleLog)

	processPermission(g, ansibleLog)
	processTARScript(g, ansibleLog)

	processSendPackage(g, ansibleLog)
	processInstall(g, s, ansibleLog)

	deleteTmpScript(g, s, ansibleLog)
}

func processInstall(g *Gui, s *setup.Setup, ansibleLog *myText) {
	if s.MasterCount > 1 {
		for i := 0; i < s.MasterCount; i++ {
			_ = g.Command("sh localScript/add_ansible_host.sh "+"netCardChange "+s.Masters[i].IPAddr, ansibleLog)
			_ = g.Command("ansible-playbook localScript/netCardChange.yaml -e netCard="+s.Masters[i].NetCard, ansibleLog)
			_ = g.Command("sh localScript/delete_ansible_host.sh "+"netCardChange", ansibleLog)
		}
		_ = g.Command("sh startHaCluster.sh ", ansibleLog)
	} else {
		_ = g.Command("sh startOneMasterCluster.sh ", ansibleLog)
	}
}

func processSendPackage(g *Gui, ansibleLog *myText) {
	_ = g.Command("sh start.sh ", ansibleLog)
}

func processTARScript(g *Gui, ansibleLog *myText) {
	_ = g.Command("tar -zcvf k8s-installer.tar.gz k8s-installer", ansibleLog)
}

func processPermission(g *Gui, ansibleLog *myText) {
	_ = g.Command("chmod -R +x .", ansibleLog)
}

func processHAScript(g *Gui, s *setup.Setup, ansibleLog *myText) {
	if s.MasterCount > 1 {
		_ = g.Command("sed -i 's/VIRTUAL_IP/"+s.Kubernetes.VirtualIP+"/g' k8s-installer/addFirstMaster/HaMasterPreSet.yaml", ansibleLog)
		_ = g.Command("sed -i 's/CERT_SANS/"+s.Kubernetes.CertSANs+"/g' k8s-installer/addFirstMaster/HaMasterPreSet.yaml", ansibleLog)

		_ = g.Command("sed -i 's/VIRTUAL_IP/"+s.Kubernetes.VirtualIP+"/g' k8s-installer/addFirstMaster/oneMasterPreSet.yaml", ansibleLog)
		_ = g.Command("sed -i 's/CERT_SANS/"+s.Kubernetes.CertSANs+"/g' k8s-installer/addFirstMaster/oneMasterPreSet.yaml", ansibleLog)

		for i := 0; i < s.MasterCount; i++ {
			s := "    server  app" + strconv.Itoa(i) + " " + s.Masters[i].IPAddr + ":6443 check"
			_ = g.Command("sed -i '/#SET_MASTER_SERVER_HERE/a"+s+"' k8s-installer/config/haproxy.cfg", ansibleLog)
		}

		//	_ = g.Command("sed -i 's/NET_CARD/" + Setup.Kubernetes.NetCard + "/g' k8s-installer/config/keepalived.conf", ansibleLog)
		_ = g.Command("sed -i 's/VIRTUAL_IP/"+s.Kubernetes.VirtualIP+"/g' k8s-installer/config/keepalived.conf", ansibleLog)
	}
}

func processGeneralScript(g *Gui, s *setup.Setup, ansibleLog *MyText) {
	_ = g.Command("sed -i 's/kube_apiserver_version/"+s.Kubernetes.Version+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/kube_proxy_version/"+s.Kubernetes.Version+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/kube_scheduler_version/"+s.Kubernetes.Version+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/kube_controller_version/"+s.Kubernetes.Version+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/etcd_version/"+s.Etcd.Version+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/coredns_version/"+util.GetCoreDNSVersion(s.Kubernetes.Version)+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/pause_version/"+util.GetPauseVersion(s.Kubernetes.Version)+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)

	//k8s
	_ = g.Command("sed -i 's/KUBEADM_VERSION/"+s.Kubernetes.Version[1:len(s.Kubernetes.Version)]+"/g' k8s-installer/k8s-script/osinit/installk8s.sh", ansibleLog)
	_ = g.Command("sed -i 's/KUBELET_VERSION/"+s.Kubernetes.Version[1:len(s.Kubernetes.Version)]+"/g' k8s-installer/k8s-script/osinit/installk8s.sh", ansibleLog)
	_ = g.Command("sed -i 's/KUBECTL_VERSION/"+s.Kubernetes.Version[1:len(s.Kubernetes.Version)]+"/g' k8s-installer/k8s-script/osinit/installk8s.sh", ansibleLog)

	//Docker
	_ = g.Command("sed -i 's/DOCKER_VERSION/"+s.Docker.Version+"/g' k8s-installer/k8s-script/osinit/installdocker.sh", ansibleLog)
	_ = g.Command("sed -i 's/DOCKER_REGISTRIES/"+s.Docker.RepositoryName+"/g' k8s-installer/k8s-script/osinit/installdocker.sh", ansibleLog)

	//网络插件
	_ = g.Command("sed -i 's/NETWORK/"+s.Kubernetes.NetComponent.Component+"/g' k8s-installer/k8s-script/cluster/init_first_master.sh", ansibleLog)
}

func processKubeadmConfig(s *setup.Setup) {
	s1 := "apiVersion: kubeadm.k8s.io/v1beta1\nkind: ClusterConfiguration\n"

	inputControllerManager(&s1, s)

	inputScheduler(&s1, s)

	inputClusterInfo(&s1, s)

	inputNetworking(&s1, s)

	inputAPIServer(&s1, s)

	inputOtherInfo(&s1, s)

	util.WriteToNewFile("k8s-installer/k8s-script/cluster/kubeadmin_init.yaml", s1)
}

func inputOtherInfo(s1 *string, s *setup.Setup) {
	if s.Kubernetes.ControlPlaneEndpoint != "" {
		*s1 = util.StringAppend(*s1, "controlPlaneEndpoint: \""+s.Kubernetes.ControlPlaneEndpoint+"\"\n")
	}

	if s.Kubernetes.CertificatesDir != "" {
		*s1 = util.StringAppend(*s1, "certificatesDir: \""+s.Kubernetes.CertificatesDir+"\"\n")
	}

	if s.Kubernetes.ImageRepository != "" {
		*s1 = util.StringAppend(*s1, "imageRepository: \""+s.Kubernetes.ImageRepository+"\"\n")
	}
}

func inputScheduler(s1 *string, s *setup.Setup) {
	if s.Kubernetes.SchedulerAddr == "" && setup.Changed == nil {
	} else {
		*s1 = util.StringAppend(*s1, "scheduler:\n  extraArgs:\n")
		if s.Kubernetes.SchedulerAddr != "" {
			*s1 = util.StringAppend(*s1, "    address: "+s.Kubernetes.SchedulerAddr+"\n")
		}
		if setup.Changed != nil {
			*s1 = util.StringAppend(*s1, "    Feature-gates: \""+setup.GetFeatureGates(setup.Changed)+"\"\n")
		}
	}
}

func inputClusterInfo(s1 *string, s *setup.Setup) {
	if s.Kubernetes.ClusterName != "" {
		*s1 = util.StringAppend(*s1, "clusterName: "+s.Kubernetes.ClusterName+"\n")
	}

	*s1 = util.StringAppend(*s1, "kubernetesVersion: "+s.Kubernetes.Version+"\n")
}

func inputNetworking(s1 *string, s *setup.Setup) {
	if s.Kubernetes.Networking.PodSubnet == "" && s.Kubernetes.Networking.ServiceSubnet == "" && s.Kubernetes.Networking.DNSdomain == "" {
	} else {
		*s1 = util.StringAppend(*s1, "networking:\n")
		if s.Kubernetes.Networking.PodSubnet != "" {
			*s1 = util.StringAppend(*s1, "  podSubnet: "+s.Kubernetes.Networking.PodSubnet+"\n")
		}
		if s.Kubernetes.Networking.ServiceSubnet != "" {
			*s1 = util.StringAppend(*s1, "  serviceSubnet: "+s.Kubernetes.Networking.ServiceSubnet+"\n")
		}
		if s.Kubernetes.Networking.DNSdomain != "" {
			*s1 = util.StringAppend(*s1, "  dnsDomain: \""+s.Kubernetes.Networking.DNSdomain+"\"\n")
		}
	}
}

func inputAPIServer(s1 *string, s *setup.Setup) {
	if s.Kubernetes.CertSANs == "" && setup.GetAdmissionPlugins(s.Kubernetes.AdmissionPlugin) == "" && setup.Changed == nil {
	} else {
		*s1 = util.StringAppend(*s1, "apiServer:\n")
		if s.Kubernetes.CertSANs != "" {
			*s1 = util.StringAppend(*s1, "  certSANs:\n  - \""+s.Kubernetes.CertSANs+"\"\n")
		}
		if setup.GetAdmissionPlugins(s.Kubernetes.AdmissionPlugin) == "" && setup.Changed == nil {
		} else {
			*s1 = util.StringAppend(*s1, "  extraArgs:\n")
			if setup.GetAdmissionPlugins(s.Kubernetes.AdmissionPlugin) != "" {
				*s1 = util.StringAppend(*s1, "    enable-admission-plugins: "+setup.GetAdmissionPlugins(s.Kubernetes.AdmissionPlugin)+"\n")
			}
			if setup.Changed != nil {
				*s1 = util.StringAppend(*s1, "    Feature-gates: \""+setup.GetFeatureGates(setup.Changed)+"\"\n")
			}
		}
	}
}

func inputControllerManager(s1 *string, s *setup.Setup) {
	if s.Kubernetes.ControllerManagerAddr == "" && setup.Changed == nil {
	} else {
		*s1 = util.StringAppend(*s1, "controllerManager:\n  extraArgs:\n")
		if s.Kubernetes.ControllerManagerAddr != "" {
			*s1 = util.StringAppend(*s1, "    bind-address: "+s.Kubernetes.ControllerManagerAddr+"\n")
		}
		if setup.Changed != nil {
			*s1 = util.StringAppend(*s1, "    Feature-gates: \""+setup.GetFeatureGates(setup.Changed)+"\"\n")
		}
	}
}

func copyScript(g *Gui, ansibleLog *myText) {
	_ = g.Command("cp -r k8s-installer-fix k8s-installer", ansibleLog)
	ansibleLog.SetText(ansibleLog.GetText(false) + "Making new script...").ScrollToEnd()
	ansibleLog.SetText(ansibleLog.GetText(false) + "Copy from template...").ScrollToEnd()
	g.App.ForceDraw()
}

func processAnsibleHosts(g *Gui, s *setup.Setup, ansibleLog *myText) {
	//	_ = g.Command("rm -rf /etc/ansible/hosts", ansibleLog)
	for i := 0; i < s.MasterCount; i++ {
		//	_ = g.Command("sh localScript/add_ansible_host.sh "+"hostname "+s.Masters[i].IPAddr, ansibleLog)
		_ = g.Command("sed -i 's/"+s.Masters[i].IPAddr+"//g' /etc/ansible/hosts", ansibleLog)
		if i == 0 {
			_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-master-init "+"master1 ansible_host="+s.Masters[i].IPAddr, ansibleLog)
		} else {
			_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-master-other "+s.Masters[i].IPAddr, ansibleLog)
		}
	}

	for i := 0; i < s.NodeCount; i++ {
		//	_ = g.Command("sh localScript/add_ansible_host.sh "+"hostname "+s.Nodes[i].IPAddr, ansibleLog)
		_ = g.Command("sed -i 's/"+s.Nodes[i].IPAddr+"//g' /etc/ansible/hosts", ansibleLog)
		_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-node "+s.Nodes[i].IPAddr, ansibleLog)
	}

	_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-master:children "+"k8s-master-init", ansibleLog)
	if s.MasterCount > 1 {
		_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-master:children "+"k8s-master-other", ansibleLog)
	}
	_ = g.Command("sh localScript/add_ansible_host.sh "+"allnodes:children "+"k8s-master", ansibleLog)
	_ = g.Command("sh localScript/add_ansible_host.sh "+"allnodes:children "+"k8s-node", ansibleLog)
}

func deleteTmpScript(g *Gui, s *setup.Setup, ansibleLog *myText) {
	_ = g.Command("sh localScript/delete_ansible_host.sh "+"k8s-master-init", ansibleLog)
	for i := 0; i < s.NodeCount; i++ {
		_ = g.Command("sh localScript/delete_ansible_host.sh "+"k8s-node", ansibleLog)
	}

	_ = g.Command("sh localScript/delete_ansible_host.sh "+"k8s-master:children", ansibleLog)
	if s.MasterCount > 1 {
		for i := 0; i < s.MasterCount-1; i++ {
			_ = g.Command("sh localScript/delete_ansible_host.sh "+"k8s-master-other", ansibleLog)
		}
		_ = g.Command("sh localScript/delete_ansible_host.sh "+"k8s-master:children", ansibleLog)
	}

	_ = g.Command("sh localScript/delete_ansible_host.sh "+"allnodes:children", ansibleLog)
	_ = g.Command("sh localScript/delete_ansible_host.sh "+"allnodes:children", ansibleLog)

	_ = g.Command("rm -rf k8s-installer.tar.gz", ansibleLog)
	_ = g.Command("rm -rf k8s-installer", ansibleLog)
}
