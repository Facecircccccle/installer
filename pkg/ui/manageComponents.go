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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)

// Manage implements cluster management operations.
func Manage(g *Gui, clientset *kubernetes.Clientset) {

	manageInfo := newManageInfo()
	manageMenu := newManageMenu()
	manageNode := newNodeManageGrid(g, clientset, manageInfo, manageMenu)

	managePV := newPVManageGrid(g, clientset, manageInfo, manageMenu)
	manageSC := newSCManageGrid(g, clientset, manageInfo, manageMenu)
	manageNamespace := newNamespaceManageGrid(g, clientset, manageInfo, manageMenu)

	gridSetupLog := tview.NewGrid().SetRows(10, -1).SetColumns(-1, -5).
		AddItem(manageInfo, 0, 0, 1, 2, 0, 0, false).
		AddItem(manageMenu, 1, 0, 1, 1, 0, 0, true).
		AddItem(manageNode, 1, 1, 1, 1, 0, 0, false)

	var c = manageNode
	manageMenu.SetSelectedFunc(func(row int, column int) {
		switch row {
		//Cluster cpu...
		case 0:

		//Node
		case 1:
			gridSetupLog.RemoveItem(c)
			c = manageNode
			gridSetupLog.AddItem(manageNode, 1, 1, 1, 1, 0, 0, true)
			g.App.SetFocus(manageNode)
		//Namespace
		case 2:
			gridSetupLog.RemoveItem(c)
			c = manageNamespace
			gridSetupLog.AddItem(manageNamespace, 1, 1, 1, 1, 0, 0, true)
			g.App.SetFocus(manageNamespace)
		//PV
		case 3:
			gridSetupLog.RemoveItem(c)
			c = managePV
			gridSetupLog.AddItem(managePV, 1, 1, 1, 1, 0, 0, true)
			g.App.SetFocus(managePV)
		//sc
		case 4:
			gridSetupLog.RemoveItem(c)
			c = manageSC
			gridSetupLog.AddItem(manageSC, 1, 1, 1, 1, 0, 0, true)
			g.App.SetFocus(manageSC)
		//back
		case 5:
			g.Pages.RemovePage("Manage")
			g.Welcome()
		}
	})

	g.Pages = tview.NewPages().
		AddAndSwitchToPage("Manage", gridSetupLog, true)
	_ = g.App.SetRoot(g.Pages, true).Run()
}

func newNamespaceManageGrid(g *Gui, clientset *kubernetes.Clientset, info *Infos, m *menu.Menus) *myGrid {
	table, log := newNamespaceManage(g, clientset, info, m)

	grid := &myGrid{
		Grid: tview.NewGrid().SetBorders(false).SetRows(-1, -1).
			AddItem(table, 0, 0, 1, 1, 0, 0, true).
			AddItem(log, 1, 0, 1, 1, 0, 0, true),
	}
	grid.SetTitle("").SetTitleAlign(tview.AlignCenter)
	grid.SetBorder(false)
	return grid
}

func newNamespaceManage(g *Gui, clientset *kubernetes.Clientset, info *Infos, m *menu.Menus) (*myTable, *myText) {
	table := &myTable{
		Table: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
	}
	log := &myText{
		TextView: tview.NewTextView().SetWordWrap(true).SetWrap(true),
	}

	info.SetText("node Manage part")

	table.SetTitle("Namespace List").SetTitleAlign(tview.AlignCenter)
	table.SetBorder(true)
	setNamespaceEntries(g, table, clientset, info, log)
	setNamespaceKeybinding(g, table, clientset, m, info, log)

	log.SetTitle("Namespace Log").SetTitleAlign(tview.AlignCenter)
	log.SetBorder(true)
	log.SetScrollable(true)

	return table, log
}

func newSCManageGrid(g *Gui, clientset *kubernetes.Clientset, info *Infos, m *menu.Menus) *myGrid {
	table, log := newSCManage(g, clientset, info, m)

	grid := &myGrid{
		Grid: tview.NewGrid().SetBorders(false).SetRows(-1, -1).
			AddItem(table, 0, 0, 1, 1, 0, 0, true).
			AddItem(log, 1, 0, 1, 1, 0, 0, true),
	}
	grid.SetTitle("").SetTitleAlign(tview.AlignCenter)
	grid.SetBorder(false)
	return grid
}

func newSCManage(g *Gui, clientset *kubernetes.Clientset, info *Infos, m *menu.Menus) (*myTable, *myText) {
	table := &myTable{
		Table: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
	}
	log := &myText{
		TextView: tview.NewTextView().SetWordWrap(true).SetWrap(true),
	}

	info.SetText("node Manage part")

	table.SetTitle("StorageClass List").SetTitleAlign(tview.AlignCenter)
	table.SetBorder(true)
	setStorageClassEntries(g, table, clientset, info, log)
	setStorageClassKeybinding(g, table, clientset, m, info, log)

	log.SetTitle("StorageClass Log").SetTitleAlign(tview.AlignCenter)
	log.SetBorder(true)
	log.SetScrollable(true)

	return table, log
}

func newPVManageGrid(g *Gui, clientset *kubernetes.Clientset, info *Infos, m *menu.Menus) *myGrid {
	table, log := newPVManage(g, clientset, info, m)

	grid := &myGrid{
		Grid: tview.NewGrid().SetBorders(false).SetRows(-1, -1).
			AddItem(table, 0, 0, 1, 1, 0, 0, true).
			AddItem(log, 1, 0, 1, 1, 0, 0, true),
	}
	grid.SetTitle("").SetTitleAlign(tview.AlignCenter)
	grid.SetBorder(false)
	return grid
}

func newPVManage(g *Gui, clientset *kubernetes.Clientset, info *Infos, m *menu.Menus) (*myTable, *myText) {
	table := &myTable{
		Table: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
	}
	log := &myText{
		TextView: tview.NewTextView().SetWordWrap(true).SetWrap(true),
	}

	info.SetText("node Manage part")

	table.SetTitle("PV List").SetTitleAlign(tview.AlignCenter)
	table.SetBorder(true)
	setPVEntries(g, table, clientset, info, log)
	setPVKeybinding(g, table, clientset, m, info, log)

	log.SetTitle("PV Log").SetTitleAlign(tview.AlignCenter)
	log.SetBorder(true)
	log.SetScrollable(true)

	return table, log
}

func newManageInfo() *Infos {
	infos := &Infos{
		TextView: tview.NewTextView().SetText("here log").SetWordWrap(true).SetWrap(true),
	}
	infos.SetTitle("info").SetTitleAlign(tview.AlignCenter)
	infos.SetBorder(true)

	return infos
}

func newManageMenu() *menu.Menus {
	menus := &menu.Menus{
		Table: tview.NewTable().SetSelectable(true, true).SetFixed(1, 1),
	}
	menus.SetTitle("Management list").SetTitleAlign(tview.AlignCenter)
	menus.SetBorder(true)
	Menu := []string{"Cluster", "Node List", "Namespace", "PV", "SC", "Back"}
	table := menus.Clear()
	for i := 0; i < len(Menu); i++ {
		cell := &tview.TableCell{
			Text:            Menu[i],
			Align:           tview.AlignLeft,
			Color:           tcell.ColorWhite,
			BackgroundColor: tcell.ColorDefault,
			Attributes:      tcell.AttrBold,
		}
		table.SetCell(i, 0, cell.SetMaxWidth(1).SetExpansion(1))
	}

	return menus
}







func newNodeManageGrid(g *Gui, clientset *kubernetes.Clientset, info *Infos, m *menu.Menus) *myGrid {
	table, log := newNodeManage(g, clientset, info, m)

	grid := &myGrid{
		Grid: tview.NewGrid().SetBorders(false).SetRows(-1, -1).
			AddItem(table, 0, 0, 1, 1, 0, 0, true).
			AddItem(log, 1, 0, 1, 1, 0, 0, true),
	}
	grid.SetTitle("").SetTitleAlign(tview.AlignCenter)
	grid.SetBorder(false)
	return grid
}

func newNodeManage(g *Gui, clientset *kubernetes.Clientset, info *Infos, m *menu.Menus) (*myTable, *myText) {
	table := &myTable{
		Table: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
	}
	log := &myText{
		TextView: tview.NewTextView().SetWordWrap(true).SetWrap(true),
	}

	info.SetText("node Manage part")

	table.SetTitle("Node List").SetTitleAlign(tview.AlignCenter)
	table.SetBorder(true)
	setNodeEntries(g, table, clientset, info, log)
	setNodeKeybinding(g, table, clientset, m, info, log)

	log.SetTitle("Node Log").SetTitleAlign(tview.AlignCenter)
	log.SetBorder(true)
	log.SetScrollable(true)

	return table, log
}

func importNodeForm(g *Gui, r *myTable, clientset *kubernetes.Clientset, m *menu.Menus, i *Infos, log *myText) {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignCenter)
	form.SetTitle("Import Node")
	form.AddDropDown("Role", []string{"Node"}, 0, nil).
		AddInputField("IP", "", constants.InputWidth, nil, nil).
		AddInputField("Name", "", constants.InputWidth, nil, nil).
		AddInputField("Username", "", constants.InputWidth, nil, nil).
		AddInputField("Code", "", constants.InputWidth, nil, nil).
		AddInputField("docker-registries", "", constants.InputWidth, nil, nil).
		AddButton("Load", func() {
			nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				panic(err)
			}
			setup := setup2.NewSampleSetupStructure()
			isHA := false
			for i := 0; i < len(nodes.Items); i++ {
				if strings.Contains(strings.Join(manage.FindNodeRoles(&nodes.Items[i]), ","), "master") {
					setup.MasterCount++
					isHA = true
				}
			}

			result, reason := setup2.AddRoleCheck(form, isHA)
			if result == true {
				_, role := form.GetFormItemByLabel("Role").(*tview.DropDown).GetCurrentOption()
				nodeIP := form.GetFormItemByLabel("IP").(*tview.InputField).GetText()
				name := form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
				username := form.GetFormItemByLabel("Username").(*tview.InputField).GetText()
				code := form.GetFormItemByLabel("Code").(*tview.InputField).GetText()
				dockerRegistries := form.GetFormItemByLabel("docker-registries").(*tview.InputField).GetText()

				log.SetText(role + name + username + code)

				//获取k8s docker版本, 仓库, masterip
				kubeVersion, containerRuntimeVersion, masterIP := manage.GetKubeInfo(clientset)
				if dockerRegistries == "" {
					dockerRegistries = "0.0.0.0/0"
				}

				//删除当前的ansible-host, 添加[node] 192.168.48.1
				_ = g.Command("rm -rf /etc/ansible/hosts", log)
				_ = g.Command("sh localScript/add_ansible_host.sh "+"allnodes "+nodeIP, log)
				_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-node "+nodeIP, log)
				_ = g.Command("sh localScript/add_ansible_host.sh "+"k8s-master-init "+masterIP, log)

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
				_ = g.Command("sh startPreSet.sh ", log)
				_ = g.Command("sh startNode.sh ", log)

				setNodeEntries(g, r, clientset, i, log)

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
