package constants

const (
	ClusterPort         = string("6443")
	HaPort              = string("16443")
	GateVersion         = 119
	InputWidth          = 70
	LabelNodeRolePrefix = "node-role.kubernetes.io/"
	NodeLabelRole       = "kubernetes.io/role"
)

var (
	KubernetesVersion = []string{"v1.20.1", "v1.20.2", "v1.20.3", "v1.20.4", "v1.20.5", "v1.20.6"}

	DockerVersion = []string{"docker-ce-19.03.15"}
	EtcdVersion   = []string{"3.4.13-0"}

	HowToSetup = "The installation section currently provides two modes: ONE master cluster installation and MULTI master (HA) cluster installation."

	OneMasterSetup = "This section installs a single master kubernetes cluster on your remote machine, and you can enter the configuration information for the cluster as needed."

	HAClusterSetup = "This section remotely installs a kubernetes cluster of multiple masters and nodes using your current machine as an operator. You need to make sure that each " +
		"machine's IP has completed an SSH connection with the local machine.You need to enter the role and IP of each machine, the required cluster information, and so on. Some " +
		"configurations are necessary in this part."

	EbKubernetes = "EbK8s is a one-click installer for beginners that can install the Kubernetes cluster. Since it has " +
		"just been developed, the current version only supports some of the basic Features and versions, and will have more Features in the future."

	EbKubernetesInstall = "In the installation section, all you need to do is enter the IP of the cluster and some personalized " +
		"information, make an SSH connection between the local machine and the target IP, then click Start to install."

	EbKubernetesManage = "The management section can acquire privileges through config files to implement cluster-level " +
		"management. Currently, this program only supports the addition and deletion of Nodes."

	SetupListTotalIntro = "The general components required to install a kubernetes cluster are listed in the" +
		" setup list, and you need to configure this simply to start the cluster.\nCluster - Node Information, " +
		"Kubernetes - Cluster Configuration, CNI - Container Runtime Configuration (currently only supports docker, " +
		"other runtimes will be updated later), Etcd - Storage Component, node Allocate - Node Resource Limitation.\n" +
		"When the above configuration is complete, select start to enter automatic installation.\n" +
		"\nKEY: Use '⬆' and '⬇' to scroll, 'Enter' to choose."

	SetupListClusterIntro = "Cluster role part completes the collection of role information.\n" +
		"Input the kind of roles (Master or Node) and IP to add new role.\n" +
		"Please make sure the IP is available and already pass the SSH authentication on this machine.\n" +
		"\nKEY: Use '⬆' and '⬇' to scroll, 'i' to add a new role, 'delete' to remove an existing role, 'b' to back to Setup Menu."

	SetupListKubernetesIntro = "Kubernetes part completes the collection of cluster and network information.\n" +
		"Only Kubernetes Version and NetPlugin must be filled in, others can be left blank if you feel unnecessary or unclear\n" +
		"\nKEY: Use 'TAB' to scroll down, Enter 'next' to the next field. It will return to Setup Menu when the 'next' in last field is chosen."

	SetupListDockerIntro = "Docker part completes the collection of CNI information.\n" +
		"For the time being, we only support docker as a container runtime, and we will provide more options in subsequent versions.\n" +
		"The version of docker is associated with the kubernetes version. When the selected version does not match, there will be a prompt.\n" +
		"\nKEY: Use 'TAB' to scroll down, Enter 'next' to the Setup Menu."

	SetupListEtcdIntro = "Etcd part completes the collection of storage information.\n" +
		"The version of etcd is associated with the kubernetes version. When the selected version does not match, there will be a prompt.\n" +
		"\nKEY: Use 'TAB' to scroll down, Enter 'next' to the Setup Menu."

	SetupListNodeIntro = "Node Allocate part completes the collection of node resource limitation information.\n" +
		"For STORAGE part, please make sure the input should end with 'Gi' OR 'Mi'.\n" +
		"For PERCENT part, please make sure the input should end with '%'.\n" +
		"\nKEY: Use 'TAB' to scroll down, Enter 'next' to the Setup Menu."

	SetupListFeatureIntro = "Feature info"
)
