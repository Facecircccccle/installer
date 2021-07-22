package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"installer/pkg/menu"
	"k8s.io/client-go/kubernetes"
)

// Manage implements cluster management operations.
func Manage(g *Gui, clientset *kubernetes.Clientset) {

	manageInfo := newManageInfo()
	manageMenu := newManageMenu()
	manageNode := newNodeManageGrid(g, clientset, manageInfo, manageMenu)

	manageNodeStatus := newNodeStatusManageGrid(g, clientset, manageInfo, manageMenu)
	managePV := newPVManageGrid(g, clientset, manageInfo, manageMenu)
	manageSC := newSCManageGrid(g, clientset, manageInfo, manageMenu)
	manageNamespace := newNamespaceManageGrid(g, clientset, manageInfo, manageMenu)

	gridSetupLog := tview.NewGrid().SetRows(10, -1).SetColumns(-1, -5).
		AddItem(manageInfo, 0, 0, 1, 2, 0, 0, false).
		AddItem(manageMenu, 1, 0, 1, 1, 0, 0, true).
		AddItem(manageNodeStatus, 1, 1, 1, 1, 0, 0, false)

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

func newNodeStatusManageGrid(g *Gui, clientset *kubernetes.Clientset, info *Infos, m *menu.Menus) *myGrid {
	table := newNodeStatusManage(g, clientset, info, m)

	grid := &myGrid{
		Grid: tview.NewGrid().SetBorders(false).SetRows(-1, -1).
			AddItem(table, 0, 0, 1, 1, 0, 0, true),
	}
	grid.SetTitle("").SetTitleAlign(tview.AlignCenter)
	grid.SetBorder(false)
	return grid
}

func newNodeStatusManage(g *Gui, clientset *kubernetes.Clientset, info *Infos, m *menu.Menus) *myTable {
	table := &myTable{
		Table: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
	}

	info.SetText("node Manage part")

	table.SetTitle("Namespace List").SetTitleAlign(tview.AlignCenter)
	table.SetBorder(true)
	setNodeStatusEntries(g, table, clientset, info)
	setNodeStatusKeybinding(g, table, clientset, m, info)

	return table
}

func newNamespaceManageGrid(g *Gui, clientset *kubernetes.Clientset, info *Infos, m *menu.Menus) *myGrid {
	table, log := newNamespaceManage(g, clientset, info, m)

	grid := &myGrid{
		Grid: tview.NewGrid().SetBorders(false).SetRows(-1, -1).
			AddItem(table, 0, 0, 1, 1, 0, 0, true).
			AddItem(log, 1, 0, 1, 1, 0, 0, false),
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
			AddItem(log, 1, 0, 1, 1, 0, 0, false),
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
			AddItem(log, 1, 0, 1, 1, 0, 0, false),
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
			AddItem(log, 1, 0, 1, 1, 0, 0, false),
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
