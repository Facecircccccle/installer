package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"installer/pkg/constants"
	"installer/pkg/menu"
	"installer/pkg/setup"
)

func NewCnis(g *Gui, m *menu.Menus, s *setup.Setup) *MyGrid {
	docker := setup.NewDocker()
	SetCnisConnection(docker, g, m, s)

	cnis := &MyGrid{
		Grid: tview.NewGrid().SetRows(-1, -2).SetBorders(false).
			AddItem(docker, 0, 0, 1, 2, 0, 0, true),
	}

	cnis.SetTitle("docker setting").SetTitleAlign(tview.AlignLeft)
	cnis.SetBorder(true)
	return cnis
}

func NewStorages(g *Gui, m *menu.Menus, s *setup.Setup) *MyGrid {
	etcd := setup.NewEtcd()
	SetStoragesConnection(etcd, g, m, s)

	storages := &MyGrid{
		Grid: tview.NewGrid().SetRows(-1, -2).SetBorders(false).
			AddItem(etcd, 0, 0, 1, 2, 0, 0, true),
	}

	storages.SetTitle("etcd setting").SetTitleAlign(tview.AlignLeft)
	storages.SetBorder(true)
	return storages
}

func NewFeature(g *Gui, m *menu.Menus) *MyGrid {
	feature := setup.NewFeatures()

	features := &MyGrid{
		Grid: tview.NewGrid().SetBorders(false).
			AddItem(feature, 0, 0, 1, 1, 0, 0, true),
	}


	feature.AddButton("done", func() {
		g.App.SetFocus(m)
	}).SetButtonsAlign(tview.AlignRight)

	features.SetTitle("Feature gate setting").SetTitleAlign(tview.AlignCenter)
	features.SetBorder(true)
	return features
}

func NewKubernetes(g *Gui, m *menu.Menus, s *setup.Setup) *MyGrid {

	cluster := setup.NewCluster()
	networking := setup.NewNetwork()
	netplugin := setup.NewNetPlugin()
	admissionplugin := setup.NewAdmission()

	SetKubernetesConnection(cluster, networking, netplugin, admissionplugin, g, m, s)

	kubernetes := &MyGrid{
		Grid: tview.NewGrid().SetRows(-2, -1, -1, -2).SetBorders(false).
			AddItem(cluster, 0, 0, 1, 2, 0, 0, true).
			AddItem(networking, 1, 0, 1, 2, 0, 0, true).
			AddItem(netplugin, 2, 0, 1, 2, 0, 0, true).
			AddItem(admissionplugin, 3, 0, 1, 2, 0, 0, true),
	}

	kubernetes.SetTitle("Kubernetes Setting").SetTitleAlign(tview.AlignCenter)
	kubernetes.SetBorder(true)
	return kubernetes
}

func NewNodeAllocates(g *Gui, m *menu.Menus, s *setup.Setup) *MyGrid {
	allocate := setup.NewAllocate()
	SetAllocatesConnection(allocate, g, m, s)

	nodeAllocates := &MyGrid{
		Grid: tview.NewGrid().SetRows(-1, -1).SetBorders(false).
			AddItem(allocate, 0, 0, 1, 2, 0, 0, true),
	}

	nodeAllocates.SetTitle("nodeAllocates setting").SetTitleAlign(tview.AlignLeft)
	nodeAllocates.SetBorder(true)
	return nodeAllocates
}

func NewRoleGrid(g *Gui, m *menu.Menus, s *setup.Setup, isHA bool) *MyGrid {
	role := NewRole(g, m, s, isHA)

	roleGrid := &MyGrid{
		Grid: tview.NewGrid().SetBorders(false).
			AddItem(role, 0, 0, 1, 1, 0, 0, true),
	}

	roleGrid.SetTitle("").SetTitleAlign(tview.AlignLeft)
	roleGrid.SetBorder(false)
	return roleGrid
}

func NewRole(g *Gui, m *menu.Menus, s *setup.Setup, isHA bool) *MyTable {
	roles := &MyTable{
		Table: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
	}

	roles.SetTitle("Role List").SetTitleAlign(tview.AlignCenter)
	roles.SetBorder(true)
	SetRoleEntries(roles, s)
	SetKeybinding(g, roles, m, s, isHA)
	return roles
}


func ImportRoleForm(g *Gui, r *MyTable, s *setup.Setup, isHA bool) {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignCenter)
	form.SetTitle("Import New Role")
	form.AddDropDown("Role", []string{"Master", "Node", "Access"}, 0, nil).
		AddInputField("IP", "", constants.InputWidth, nil, nil)
	//AddInputField("Name", "", inputWidth, nil, nil).
	//AddInputField("Username", "", inputWidth, nil, nil).
	//AddInputField("Code", "", inputWidth, nil, nil).
	if isHA{
		form.AddInputField("Network Card", "", constants.InputWidth, nil, nil)
	}
	form.AddButton("Load", func() {
		result, reason := setup.AddRoleCheck(form, isHA)
		if result == true {
			_, role := form.GetFormItemByLabel("Role").(*tview.DropDown).GetCurrentOption()
			ip := form.GetFormItemByLabel("IP").(*tview.InputField).GetText()
			//name := form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
			//username := form.GetFormItemByLabel("Username").(*tview.InputField).GetText()
			//code := form.GetFormItemByLabel("Code").(*tview.InputField).GetText()

			// write to structure
			// do something
			if role == "Master" {
				var master setup.Master
				if isHA{
					master = setup.Master{
						IPAddr: ip,
						NetCard: form.GetFormItemByLabel("Network Card").(*tview.InputField).GetText(),
					}
				}else {
					master = setup.Master{
						IPAddr: ip,
					}
				}
				s.Masters = append(s.Masters, master)
				s.MasterCount++
			}else if role == "Node" {
				node := setup.Node{
					IPAddr:   ip,
					Allocate: setup.NewSampleNodeAllocate(),
				}
				s.Nodes = append(s.Nodes, node)
				s.NodeCount++
			}else {
				access := setup.Access{
					IPAddr:   ip,
				}
				s.Accesses = append(s.Accesses, access)
				s.AccessCount++
			}

			SetRoleEntries(r, s)

			g.Pages.RemovePage("form")
			g.App.SetFocus(r)
		}else {
			modal := tview.NewModal().
				SetText(reason).
				AddButtons([]string{"ok"})
			modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "ok" {
					g.Pages.RemovePage("Modal")
					g.App.SetFocus(form)
					g.Pages.AddAndSwitchToPage("form", g.Modal(form, 70, 10), true).ShowPage("main")
				}
			})
			g.Pages.AddAndSwitchToPage("Modal", g.Modal(modal, 40, 16), true).ShowPage("main")

		}
	}).
		AddButton("Cancel", func() {
			g.Pages.RemovePage("form")
			g.App.SetFocus(r)
		})

	g.Pages.AddAndSwitchToPage("form", g.Modal(form, 70, 11), true).ShowPage("main")
}

func DeleteRoleForm(r *MyTable, s *setup.Setup) {
	row, _ := r.GetSelection()

	if r.GetCell(row, 0).Text == "Access" {
		s.Accesses = append(s.Accesses[:row-2], s.Accesses[row-2+1:]...)
		s.AccessCount--
	}
	if r.GetCell(row, 0).Text == "Master" {
		s.Masters = append(s.Masters[:row-s.AccessCount-2], s.Masters[row-s.AccessCount-2+1:]...)
		s.MasterCount--
	}
	if r.GetCell(row, 0).Text == "Node" {
		s.Nodes = append(s.Nodes[:row-s.AccessCount-s.MasterCount-2], s.Nodes[row-s.AccessCount-s.MasterCount-2+1:]...)
		s.NodeCount--
	}
	SetRoleEntries(r, s)
}

func SetRoleEntries(r *MyTable, s *setup.Setup) {
	table := r.Clear()

	headers := []string{
		"Role",
		"IP",
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

	var Roles []setup.Role
	for i:= 0; i < s.AccessCount; i++ {
		Roles = append(Roles, setup.Role{
			Role: "Access"},
			//Name: s.Accesses[i].Name,
			//User: s.Accesses[i].UserName,
			//Code: s.Accesses[i].code})
		)
	}
	for i:= 0; i < s.MasterCount; i++ {
		Roles = append(Roles, setup.Role{
			Role: "Master",
			IP:   s.Masters[i].IPAddr,
			//Name: s.Masters[i].Name,
			//User: s.Masters[i].UserName,
			//Code: s.Masters[i].code
		})
	}
	for i:= 0; i < s.NodeCount; i++ {
		Roles = append(Roles, setup.Role{
			Role: "Node",
			IP:   s.Nodes[i].IPAddr,
			//Name: s.Nodes[i].Name,
			//User: s.Nodes[i].UserName,
			//Code: s.Nodes[i].code
		})
	}

	for i := 0; i < len(Roles); i++{
		table.SetCell(i+2, 0, tview.NewTableCell(Roles[i].Role).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))

		table.SetCell(i+2, 1, tview.NewTableCell(Roles[i].IP).
			SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))
	}
}

