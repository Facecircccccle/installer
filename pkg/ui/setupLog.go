package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"installer/pkg/menu"
	"installer/pkg/setup"
)

type SetupLog struct {
	app   *tview.Application
	pages *tview.Pages
}

func (g *Gui) setupLog(s *setup.Setup) {

	setupInfo := newSetupInfo(g)
	setupMenu := newSetupMenu(g)
	setupAnsibleHostAndName := newAnsibleHostAndName(g, setupMenu, s, setupInfo)

	gridSetupLog := tview.NewGrid().SetRows(10, -1).SetColumns(-1, -5).
		AddItem(setupInfo, 0, 0, 1, 2, 0, 0, false).
		AddItem(setupMenu, 1, 0, 1, 1, 0, 0, true).
		AddItem(setupAnsibleHostAndName, 1, 1, 1, 1, 0, 0, true)

	setupMenu.SetSelectedFunc(func(row int, column int) {
		switch row {
		case 0:
			g.App.SetFocus(setupAnsibleHostAndName)
			go ProcessNewCluster(g, s, setupAnsibleHostAndName)
			g.App.SetFocus(setupAnsibleHostAndName)

		case 1:
			g.Pages.RemovePage("SetupLog")
			g.SetupChoice()
		}
	})

	g.Pages = tview.NewPages().
		AddAndSwitchToPage("SetupLog", gridSetupLog, true)
	//g.pages.AddAndSwitchToPage("main", gridlist, true)
	_ = g.App.SetRoot(g.Pages, true).Run()
}

func newSetupInfo(g *Gui) *Infos {
	infos := &Infos{
		TextView: tview.NewTextView().SetText("").SetWordWrap(true).SetWrap(true),
	}
	infos.SetTitle("Setup Info").SetTitleAlign(tview.AlignCenter)
	infos.SetBorder(true)
	return infos
}

func newSetupMenu(g *Gui) *menu.Menus {
	menus := &menu.Menus{
		Table: tview.NewTable().SetSelectable(true, true).SetFixed(1, 1),
	}
	menus.SetTitle("Setup Menu").SetTitleAlign(tview.AlignCenter)
	menus.SetBorder(true)
	Menu := []string{"Start", "Done"}
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

func (ansibleName *myText) setKeybinding(g *Gui, setupMenu *menu.Menus) {
	ansibleName.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'b':
			g.App.SetFocus(setupMenu)
		}
		return event
	})
}
func newAnsibleHostAndName(g *Gui, setupMenu *menu.Menus, setup *setup.Setup, setupInfo *Infos) *myText {
	ansibleHost := &myText{
		TextView: tview.NewTextView().SetWordWrap(true).SetWrap(true),
	}

	setupInfo.SetText("new ansible Host")

	ansibleHost.SetTitle("Ansible Log").SetTitleAlign(tview.AlignCenter)
	ansibleHost.SetBorder(true)
	ansibleHost.SetScrollable(true)
	ansibleHost.setKeybinding(g, setupMenu)

	return ansibleHost
}
