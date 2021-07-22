package ui

import (
	"context"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"installer/pkg/constants"
	"installer/pkg/manage"
	"installer/pkg/menu"
	setup2 "installer/pkg/setup"
	"installer/pkg/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"strings"
)

func setNodeKeybinding(g *Gui, r *myTable, clientset *kubernetes.Clientset, m *menu.Menus, i *Infos, log *myText) {
	r.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		//g.setGlobalKeybinding(event)
		switch event.Key() {
		case tcell.KeyTAB:
			g.App.SetFocus(log)
		}
		switch event.Rune() {
		case 'i':
			importNodeForm(g, r, clientset, m, i, log)
		case 'b':
			g.App.SetFocus(m)
		}
		return event
	})
}

func setNamespaceKeybinding(g *Gui, r *myTable, clientset *kubernetes.Clientset, m *menu.Menus, i *Infos, log *myText) {
	r.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		//g.setGlobalKeybinding(event)
		switch event.Key() {
		case tcell.KeyTAB:
			g.App.SetFocus(log)
		}
		switch event.Rune() {

		case 'b':
			g.App.SetFocus(m)
		}
		return event
	})
}

func setStorageClassKeybinding(g *Gui, r *myTable, clientset *kubernetes.Clientset, m *menu.Menus, i *Infos, log *myText) {
	r.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		//g.setGlobalKeybinding(event)
		switch event.Key() {
		case tcell.KeyTAB:
			g.App.SetFocus(log)
		}
		switch event.Rune() {

		case 'b':
			g.App.SetFocus(m)
		}
		return event
	})
}

func setPVKeybinding(g *Gui, r *myTable, clientset *kubernetes.Clientset, m *menu.Menus, i *Infos, log *myText) {
	r.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		//g.setGlobalKeybinding(event)
		switch event.Key() {
		case tcell.KeyTAB:
			g.App.SetFocus(log)
		}
		switch event.Rune() {

		case 'b':
			g.App.SetFocus(m)
		}
		return event
	})
}

func setNodeEntries(g *Gui, r *myTable, clientset *kubernetes.Clientset, i *Infos, log *myText) {
	table := r.Clear()

	headers := []string{
		"Name",
		"Status",
		"Roles",
		"Time",
		"Version",
	}

	for i, header := range headers {
		table.SetCell(0, i, &tview.TableCell{
			Text:            header,
			NotSelectable:   true,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorWhite,
			BackgroundColor: tcell.ColorDefault,
			Attributes:      tcell.AttrBold,
		})
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var nodeInfo []manage.NodeOutputStructure
	for i := 0; i < len(nodes.Items); i++ {
		node := nodes.Items[i]

		conditionMap := make(map[v1.NodeConditionType]*v1.NodeCondition)
		NodeAllConditions := []v1.NodeConditionType{v1.NodeReady}
		for i := range node.Status.Conditions {
			cond := node.Status.Conditions[i]
			conditionMap[cond.Type] = &cond
		}
		var status []string
		for _, validCondition := range NodeAllConditions {
			if condition, ok := conditionMap[validCondition]; ok {
				if condition.Status == v1.ConditionTrue {
					status = append(status, string(condition.Type))
				} else {
					status = append(status, "Not"+string(condition.Type))
				}
			}
		}
		if len(status) == 0 {
			status = append(status, "Unknown")
		}
		if node.Spec.Unschedulable {
			status = append(status, "SchedulingDisabled")
		}

		roles := strings.Join(manage.FindNodeRoles(&node), ",")
		if len(roles) == 0 {
			roles = "<none>"
		}

		nodeInfo = append(nodeInfo, manage.NodeOutputStructure{
			Name:    node.Name,
			Status:  strings.Join(status, ","),
			Roles:   roles,
			Time:    manage.TranslateTimestampSince(node.CreationTimestamp),
			Version: node.Status.NodeInfo.KubeletVersion,
		})
	}

	for i := 0; i < len(nodeInfo); i++ {
		table.SetCell(i+2, 0, tview.NewTableCell(nodeInfo[i].Name).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 1, tview.NewTableCell(nodeInfo[i].Status).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 2, tview.NewTableCell(nodeInfo[i].Roles).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 3, tview.NewTableCell(nodeInfo[i].Time).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 4, tview.NewTableCell(nodeInfo[i].Version).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))
	}
}

func setStorageClassEntries(g *Gui, r *myTable, clientset *kubernetes.Clientset, i *Infos, log *myText) {
	table := r.Clear()

	headers := []string{
		"Name",
		"Provisioner",
		"ReclaimPolicy",
		"VolumeBindingMode",
		"AllowVolumeExpansion",
	}

	for i, header := range headers {
		table.SetCell(0, i, &tview.TableCell{
			Text:            header,
			NotSelectable:   true,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorWhite,
			BackgroundColor: tcell.ColorDefault,
			Attributes:      tcell.AttrBold,
		})
	}

	sc, err := clientset.StorageV1().StorageClasses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var storageClassInfo []manage.SCOutPutStructure
	for i := 0; i < len(sc.Items); i++ {
		obj := sc.Items[i]

		name := obj.Name
		if IsDefaultAnnotation(obj.ObjectMeta) {
			name += " (default)"
		}
		provtype := obj.Provisioner
		reclaimPolicy := "delete"
		if obj.ReclaimPolicy != nil {
			reclaimPolicy = string(*obj.ReclaimPolicy)
		}

		volumeBindingMode := "Immediate"
		if obj.VolumeBindingMode != nil {
			volumeBindingMode = string(*obj.VolumeBindingMode)
		}

		allowVolumeExpansion := false
		if obj.AllowVolumeExpansion != nil {
			allowVolumeExpansion = *obj.AllowVolumeExpansion
		}

		storageClassInfo = append(storageClassInfo, manage.SCOutPutStructure{
			Name:                 name,
			Provisioner:          provtype,
			ReclaimPolicy:        reclaimPolicy,
			VolumeBindingMode:    volumeBindingMode,
			AllowVolumeExpansion: strconv.FormatBool(allowVolumeExpansion),
		})
	}

	for i := 0; i < len(storageClassInfo); i++ {
		table.SetCell(i+2, 0, tview.NewTableCell(storageClassInfo[i].Name).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 1, tview.NewTableCell(storageClassInfo[i].Provisioner).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 2, tview.NewTableCell(storageClassInfo[i].ReclaimPolicy).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 3, tview.NewTableCell(storageClassInfo[i].VolumeBindingMode).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 4, tview.NewTableCell(storageClassInfo[i].AllowVolumeExpansion).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))
	}
}

func IsDefaultAnnotation(obj metav1.ObjectMeta) bool {
	if obj.Annotations[constants.IsDefaultStorageClassAnnotation] == "true" {
		return true
	}
	if obj.Annotations[constants.BetaIsDefaultStorageClassAnnotation] == "true" {
		return true
	}

	return false
}


func setNodeStatusKeybinding(g *Gui, table *myTable, clientset *kubernetes.Clientset, m *menu.Menus, info *Infos) {

}

func setNodeStatusEntries(g *Gui, r *myTable, clientset *kubernetes.Clientset, info *Infos) {
	table := r.Clear()
	headers := []string{
		"Name",
		"CPU(cores)",
		"CPU%",
		"MEMORY(bytes)",
		"MEMORY%",
	}

	for i, header := range headers {
		table.SetCell(0, i, &tview.TableCell{
			Text:            header,
			NotSelectable:   true,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorWhite,
			BackgroundColor: tcell.ColorDefault,
			Attributes:      tcell.AttrBold,
		})
	}

	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var namespaceInfo []manage.NamespaceOutPutStructure
	for i := 0; i < len(namespaceList.Items); i++ {
		obj := namespaceList.Items[i]

		namespaceInfo = append(namespaceInfo, manage.NamespaceOutPutStructure{
			Name:   obj.Name,
			Status: string(obj.Status.Phase),
			Age:    manage.TranslateTimestampSince(obj.CreationTimestamp),
		})
	}

	for i := 0; i < len(namespaceInfo); i++ {
		table.SetCell(i+2, 0, tview.NewTableCell(namespaceInfo[i].Name).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 1, tview.NewTableCell(namespaceInfo[i].Status).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 2, tview.NewTableCell(namespaceInfo[i].Age).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))
	}
}

func setNamespaceEntries(g *Gui, r *myTable, clientset *kubernetes.Clientset, i *Infos, log *myText) {
	table := r.Clear()
	headers := []string{
		"Name",
		"Status",
		"Age",
	}

	for i, header := range headers {
		table.SetCell(0, i, &tview.TableCell{
			Text:            header,
			NotSelectable:   true,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorWhite,
			BackgroundColor: tcell.ColorDefault,
			Attributes:      tcell.AttrBold,
		})
	}

	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var namespaceInfo []manage.NamespaceOutPutStructure
	for i := 0; i < len(namespaceList.Items); i++ {
		obj := namespaceList.Items[i]

		namespaceInfo = append(namespaceInfo, manage.NamespaceOutPutStructure{
			Name:   obj.Name,
			Status: string(obj.Status.Phase),
			Age:    manage.TranslateTimestampSince(obj.CreationTimestamp),
		})
	}

	for i := 0; i < len(namespaceInfo); i++ {
		table.SetCell(i+2, 0, tview.NewTableCell(namespaceInfo[i].Name).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 1, tview.NewTableCell(namespaceInfo[i].Status).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 2, tview.NewTableCell(namespaceInfo[i].Age).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))
	}
}

func setPVEntries(g *Gui, r *myTable, clientset *kubernetes.Clientset, i *Infos, log *myText) {
	table := r.Clear()

	headers := []string{
		"Name",
		"Capacity",
		"AccessModes",
		"ReclaimPolicy",
		"Status",
		"Claim",
		"StorageClass",
		"Reason",
		"Age",
	}

	for i, header := range headers {
		table.SetCell(0, i, &tview.TableCell{
			Text:            header,
			NotSelectable:   true,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorWhite,
			BackgroundColor: tcell.ColorDefault,
			Attributes:      tcell.AttrBold,
		})
	}

	persistentVolumeList, err := clientset.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var pv []manage.PVOutputStructure
	for i := 0; i < len(persistentVolumeList.Items); i++ {
		obj := persistentVolumeList.Items[i]

		claimRefUID := ""
		if obj.Spec.ClaimRef != nil {
			claimRefUID += obj.Spec.ClaimRef.Namespace
			claimRefUID += "/"
			claimRefUID += obj.Spec.ClaimRef.Name
		}

		modesStr := manage.GetAccessModesAsString(obj.Spec.AccessModes)
		reclaimPolicyStr := string(obj.Spec.PersistentVolumeReclaimPolicy)

		aQty := obj.Spec.Capacity["storage"]
		aSize := aQty.String()

		phase := obj.Status.Phase
		if obj.ObjectMeta.DeletionTimestamp != nil {
			phase = "Terminating"
		}

		pv = append(pv, manage.PVOutputStructure{
			Name:          obj.Name,
			Capacity:      aSize,
			AccessModes:   modesStr,
			ReclaimPolicy: reclaimPolicyStr,
			Status:        string(phase),
			Claim:        claimRefUID,
			StorageClass: manage.GetPersistentVolumeClass(&obj),
			Reason:       obj.Status.Reason,
			Age:          manage.TranslateTimestampSince(obj.CreationTimestamp),
		})
	}

	for i := 0; i < len(pv); i++ {
		table.SetCell(i+2, 0, tview.NewTableCell(pv[i].Name).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 1, tview.NewTableCell(pv[i].Capacity).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 2, tview.NewTableCell(pv[i].AccessModes).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 3, tview.NewTableCell(pv[i].ReclaimPolicy).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 4, tview.NewTableCell(pv[i].Status).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 5, tview.NewTableCell(pv[i].Claim).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 6, tview.NewTableCell(pv[i].StorageClass).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 7, tview.NewTableCell(pv[i].Reason).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 8, tview.NewTableCell(pv[i].Age).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))
	}
}

func importNodeForm(g *Gui, r *myTable, clientset *kubernetes.Clientset, m *menu.Menus, i *Infos, log *myText) {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignCenter)
	form.SetTitle("Import Node")
	form.AddInputField("IP", "", constants.InputWidth, nil, nil).
		AddInputField("docker-registries", "", constants.InputWidth, nil, nil).
		AddButton("Load", func() {
			nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				panic(err)
			}
			masterCount := 0
			isHA := false

			for i := 0; i < len(nodes.Items); i++ {
				if strings.Contains(strings.Join(manage.FindNodeRoles(&nodes.Items[i]), ","), "master") {
					masterCount++
				}
			}
			if masterCount > 1{
				isHA = true
			}

			result, reason := setup2.AddRoleCheck(form, isHA)
			if result == true {
				nodeIP := form.GetFormItemByLabel("IP").(*tview.InputField).GetText()
				dockerRegistries := form.GetFormItemByLabel("docker-registries").(*tview.InputField).GetText()

				//获取k8s docker版本, 仓库, masterip
				kubeVersion, containerRuntimeVersion, masterIP := manage.GetKubeInfo(clientset)
				if dockerRegistries == "" {
					dockerRegistries = "0.0.0.0/0"
				}

				//删除当前的ansible-host, 添加[node] 192.168.48.1
				//_ = g.Command("rm -rf /etc/ansible/hosts", log)
				_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-node "+nodeIP, log)
				_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-master-init "+masterIP, log)
				//_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-master:children "+"k8s-master-init", log)
				//_ = g.Command("sh localScript/add_ansible_host.sh "+"allnodes:children "+"k8s-master", log)
				_ = g.Command("sh localScript/add_ansible_host.sh "+"allnodes:children "+"k8s-node", log)


				//在192.168.48.1上创建满足preset的节点
				_ = g.Command("cp -r k8s-installer-fix k8s-installer", log)

				//k8s
				_ = g.Command("sed -i 's/KUBEADM_VERSION/"+kubeVersion+"/g' k8s-installer/k8s-script/osinit/installk8s.sh", log)
				_ = g.Command("sed -i 's/KUBELET_VERSION/"+kubeVersion+"/g' k8s-installer/k8s-script/osinit/installk8s.sh", log)
				_ = g.Command("sed -i 's/KUBECTL_VERSION/"+kubeVersion+"/g' k8s-installer/k8s-script/osinit/installk8s.sh", log)

				_ = g.Command("sed -i 's/KUBE_APISERVER_VERSION/"+kubeVersion+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", log)
				_ = g.Command("sed -i 's/KUBE_PROXY_VERSION/"+kubeVersion+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", log)
				_ = g.Command("sed -i 's/KUBE_SCHEDULER_VERSION/"+kubeVersion+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", log)
				_ = g.Command("sed -i 's/KUBE_CONTROLLER_VERSION/"+kubeVersion+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", log)
				_ = g.Command("sed -i 's/ETCD_VERSION/"+util.GetEtcdVersion(kubeVersion)+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", log)
				_ = g.Command("sed -i 's/COREDNS_VERSION/"+util.GetCoreDNSVersion(kubeVersion)+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", log)
				_ = g.Command("sed -i 's/PAUSE_VERSION/"+util.GetPauseVersion(kubeVersion)+"/g' k8s-installer/k8s-script/cluster/pull_k8s_image.sh", log)

				//Docker
				_ = g.Command("sed -i 's/DOCKER_VERSION/"+containerRuntimeVersion+"/g' k8s-installer/k8s-script/osinit/installdocker.sh", log)
				_ = g.Command("sed -i 's/DOCKER_REGISTRIES/"+dockerRegistries+"/g' k8s-installer/k8s-script/osinit/installdocker.sh", log)

				//addNode
				_ = g.Command("sed -i 's/MASTER_IP/"+masterIP+"/g' k8s-installer/addNode/addnode.sh", log)

				//权限
				_ = g.Command("chmod -R +x .", log)

				//打包
				_ = g.Command("tar -zcvf k8s-installer.tar.gz k8s-installer", log)

				//开始
				_ = g.Command("sh start.sh ", log)
				_ = g.Command("sh ./k8s-installer/preSet/preSet.sh ", log)
				_ = g.Command("sh ./k8s-installer/addNode/addnode.sh ", log)

				setNodeEntries(g, r, clientset, i, log)

				_ = g.Command("sh localScript/delete_ansible_host.sh "+"k8s-master-init", log)
				_ = g.Command("sh localScript/delete_ansible_host.sh "+"k8s-node", log)
				//_ = g.Command("sh localScript/delete_ansible_host.sh "+"k8s-master:children", log)
				//_ = g.Command("sh localScript/delete_ansible_host.sh "+"allnodes:children", log)
				_ = g.Command("sh localScript/delete_ansible_host.sh "+"allnodes:children", log)
				_ = g.Command("rm -rf k8s-installer.tar.gz", log)
				_ = g.Command("rm -rf k8s-installer", log)

				g.Pages.RemovePage("form")
				g.App.SetFocus(r)
			} else {
				modal := tview.NewModal().
					SetText(reason).
					AddButtons([]string{"ok"})
				modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "ok" {
						g.Pages.SwitchToPage("form")
					}
				})
			}
		}).
		AddButton("Cancel", func() {
			g.Pages.RemovePage("form")
			g.App.SetFocus(r)
		})

	g.Pages.AddAndSwitchToPage("form", g.Modal(form, 80, 16), true).ShowPage("Manage")
}
