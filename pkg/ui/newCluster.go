package ui

import (
	"installer/pkg/setup"
	"installer/pkg/util"
	"strconv"
)

func ProcessNewCluster(g *Gui, s *setup.Setup, ansibleLog *MyText) {

	_ = g.Command("rm -rf /etc/ansible/hosts", ansibleLog)
	for i := 0; i < s.MasterCount; i++ {
		_ = g.Command("sh localScript/add_ansible_host.sh "+"hostname "+s.Masters[i].IPAddr, ansibleLog)
		_ = g.Command("sed -i 's/"+s.Masters[i].IPAddr+"//g' /etc/ansible/hosts", ansibleLog)
		if i == 0 {
			_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-master-init "+"master1 ansible_host="+s.Masters[i].IPAddr, ansibleLog)
		} else {
			_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-master-other "+s.Masters[i].IPAddr, ansibleLog)
		}
	}

	for i := 0; i < s.NodeCount; i++ {
		_ = g.Command("sh localScript/add_ansible_host.sh "+"hostname "+s.Nodes[i].IPAddr, ansibleLog)
		_ = g.Command("sed -i 's/"+s.Nodes[i].IPAddr+"//g' /etc/ansible/hosts", ansibleLog)
		_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-node "+s.Nodes[i].IPAddr, ansibleLog)
	}

	_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-master:children "+"k8s-master-init", ansibleLog)
	if s.MasterCount > 1 {
		_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-master:children "+"k8s-master-other", ansibleLog)
	}
	_ = g.Command("sh localScript/add_ansible_host.sh "+"allnodes:children "+"k8s-master", ansibleLog)
	_ = g.Command("sh localScript/add_ansible_host.sh "+"allnodes:children "+"k8s-node", ansibleLog)

	ansibleLog.SetText(ansibleLog.GetText(false) + "Making new script...").ScrollToEnd()
	ansibleLog.SetText(ansibleLog.GetText(false) + "Copy from template...").ScrollToEnd()
	g.App.ForceDraw()

	_ = g.Command("cp -r k8s-installer-fix k8s-installer", ansibleLog)

	s1 := "apiVersion: kubeadm.k8s.io/v1beta1\nkind: ClusterConfiguration\n"

	if s.Kubernetes.ControllerManagerAddr == "" && setup.Changed == nil {
	} else {
		s1 = util.StringAppend(s1, "controllerManager:\n  extraArgs:\n")
		if s.Kubernetes.ControllerManagerAddr != "" {
			s1 = util.StringAppend(s1, "    bind-address: "+s.Kubernetes.ControllerManagerAddr+"\n")
		}
		if setup.Changed != nil {
			s1 = util.StringAppend(s1, "    Feature-gates: \""+setup.GetFeatureGates(setup.Changed)+"\"\n")
		}
	}

	if s.Kubernetes.SchedulerAddr == "" && setup.Changed == nil {
	} else {
		s1 = util.StringAppend(s1, "scheduler:\n  extraArgs:\n")
		if s.Kubernetes.SchedulerAddr != "" {
			s1 = util.StringAppend(s1, "    address: "+s.Kubernetes.SchedulerAddr+"\n")
		}
		if setup.Changed != nil {
			s1 = util.StringAppend(s1, "    Feature-gates: \""+setup.GetFeatureGates(setup.Changed)+"\"\n")
		}
	}

	if s.Kubernetes.ClusterName != "" {
		s1 = util.StringAppend(s1, "clusterName: "+s.Kubernetes.ClusterName+"\n")
	}

	s1 = util.StringAppend(s1, "kubernetesVersion: "+s.Kubernetes.Version+"\n")

	if s.Kubernetes.Networking.PodSubnet == "" && s.Kubernetes.Networking.ServiceSubnet == "" && s.Kubernetes.Networking.DNSdomain == "" {
	} else {
		s1 = util.StringAppend(s1, "networking:\n")
		if s.Kubernetes.Networking.PodSubnet != "" {
			s1 = util.StringAppend(s1, "  podSubnet: "+s.Kubernetes.Networking.PodSubnet+"\n")
		}
		if s.Kubernetes.Networking.ServiceSubnet != "" {
			s1 = util.StringAppend(s1, "  serviceSubnet: "+s.Kubernetes.Networking.ServiceSubnet+"\n")
		}
		if s.Kubernetes.Networking.DNSdomain != "" {
			s1 = util.StringAppend(s1, "  dnsDomain: \""+s.Kubernetes.Networking.DNSdomain+"\"\n")
		}
	}

	if s.Kubernetes.CertSANs == "" && setup.GetAdmissionPlugins(s.Kubernetes.AdmissionPlugin) == "" && setup.Changed == nil {
	} else {
		s1 = util.StringAppend(s1, "apiServer:\n")
		if s.Kubernetes.CertSANs != "" {
			s1 = util.StringAppend(s1, "  certSANs:\n  - \""+s.Kubernetes.CertSANs+"\"\n")
		}
		if setup.GetAdmissionPlugins(s.Kubernetes.AdmissionPlugin) == "" && setup.Changed == nil {
		} else {
			s1 = util.StringAppend(s1, "  extraArgs:\n")
			if setup.GetAdmissionPlugins(s.Kubernetes.AdmissionPlugin) != "" {
				s1 = util.StringAppend(s1, "    enable-admission-plugins: "+setup.GetAdmissionPlugins(s.Kubernetes.AdmissionPlugin)+"\n")
			}
			if setup.Changed != nil {
				s1 = util.StringAppend(s1, "    Feature-gates: \""+setup.GetFeatureGates(setup.Changed)+"\"\n")
			}
		}
	}

	if s.Kubernetes.ControlPlaneEndpoint != "" {
		s1 = util.StringAppend(s1, "controlPlaneEndpoint: \""+s.Kubernetes.ControlPlaneEndpoint+"\"\n")
	}

	if s.Kubernetes.CertificatesDir != "" {
		s1 = util.StringAppend(s1, "certificatesDir: \""+s.Kubernetes.CertificatesDir+"\"\n")
	}

	if s.Kubernetes.ImageRepository != "" {
		s1 = util.StringAppend(s1, "imageRepository: \""+s.Kubernetes.ImageRepository+"\"\n")
	}

	//	s1 = StringAppend(s1, "useHyperKubeImage: " + Setup.Kubernetes.UseHyperKubeImage + "\n")

	util.WriteToNewFile("k8s-installer/k8s-script/cluster/kubeadmin_init.yaml", s1)

	_ = g.Command("sed -i 's/KUBE_APISERVER_VERSION/"+s.Kubernetes.Version+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/KUBE_PROXY_VERSION/"+s.Kubernetes.Version+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/KUBE_SCHEDULER_VERSION/"+s.Kubernetes.Version+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/KUBE_CONTROLLER_VERSION/"+s.Kubernetes.Version+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/ETCD_VERSION/"+s.Etcd.Version+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/COREDNS_VERSION/"+util.GetCoreDNSVersion(s.Kubernetes.Version)+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)
	_ = g.Command("sed -i 's/PAUSE_VERSION/"+util.GetPauseVersion(s.Kubernetes.Version)+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", ansibleLog)

	//k8s
	_ = g.Command("sed -i 's/KUBEADM_VERSION/"+s.Kubernetes.Version[1:len(s.Kubernetes.Version)]+"/g' k8s-installer/k8s-script/osinit/installk8s.sh", ansibleLog)
	_ = g.Command("sed -i 's/KUBELET_VERSION/"+s.Kubernetes.Version[1:len(s.Kubernetes.Version)]+"/g' k8s-installer/k8s-script/osinit/installk8s.sh", ansibleLog)
	_ = g.Command("sed -i 's/KUBECTL_VERSION/"+s.Kubernetes.Version[1:len(s.Kubernetes.Version)]+"/g' k8s-installer/k8s-script/osinit/installk8s.sh", ansibleLog)

	//Docker
	_ = g.Command("sed -i 's/DOCKER_VERSION/"+s.Docker.Version+"/g' k8s-installer/k8s-script/osinit/installdocker.sh", ansibleLog)
	_ = g.Command("sed -i 's/DOCKER_REGISTRIES/"+s.Docker.RepositoryName+"/g' k8s-installer/k8s-script/osinit/installdocker.sh", ansibleLog)

	//网络插件
	_ = g.Command("sed -i 's/NETWORK/"+s.Kubernetes.NetComponent.Component+"/g' k8s-installer/k8s-script/cluster/init_first_master.sh", ansibleLog)

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

	//权限
	_ = g.Command("chmod -R +x .", ansibleLog)

	//tar script
	_ = g.Command("tar -zcvf k8s-installer.tar.gz k8s-installer", ansibleLog)
	//	log.SetText(log.GetText(false) + outTarScript)

	_ = g.Command("sh start.sh ", ansibleLog)

	if s.MasterCount > 1 {
		for i := 0; i < s.MasterCount; i++ {
			//添加到hosts
			_ = g.Command("sh localScript/add_ansible_host.sh "+"netCardChange "+s.Masters[i].IPAddr, ansibleLog)
			//发一个命令 让20
			_ = g.Command("ansible-playbook localScript/netCardChange.yaml -e netCard="+s.Masters[i].NetCard, ansibleLog)
			//	_ = g.Command("ansible-playbook localScript/netCardChange.yaml -e netCard=ens33", ansibleLog)

			_ = g.Command("sh localScript/delete_ansible_host.sh "+"netCardChange", ansibleLog)
			ansibleLog.SetText(ansibleLog.GetText(false) + s.Masters[i].IPAddr + " " + s.Masters[i].NetCard).ScrollToEnd()
		}
		_ = g.Command("sh startHaCluster.sh ", ansibleLog)
	} else {
		_ = g.Command("sh startOneMasterCluster.sh ", ansibleLog)
		//	//运行 preset.sh
		//	_ = g.Command("sh startPreSet.sh ", ansibleLog)
		//
		//	ansibleLog.SetText(ansibleLog.GetText(false) + "preSet success, docker k8s etcd already installed\n").ScrollToEnd()
		//	g.App.ForceDraw()
		//
		//	ansibleLog.SetText(ansibleLog.GetText(false) + "init master...\n").ScrollToEnd()
		//	g.App.ForceDraw()
		//
		/////运行 master.sh
		//	_ = g.Command("sh startMaster.sh ", ansibleLog)
		//
		//	//	ansibleLog.SetText(ansibleLog.GetText(false) + "master already joined...\n").ScrollToEnd()
		//	g.App.ForceDraw()
		//
		//	ansibleLog.SetText(ansibleLog.GetText(false) + "init node...\n").ScrollToEnd()
		//	//运行 node.sh
		//	_ = g.Command("sh startNode.sh ", ansibleLog)

		//	ansibleLog.SetText(ansibleLog.GetText(false) + "node already joined...\n").ScrollToEnd()
	}

	_ = g.Command("rm -rf k8s-installer.tar.gz", ansibleLog)
	_ = g.Command("rm -rf k8s-installer", ansibleLog)
}
