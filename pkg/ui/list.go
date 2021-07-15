package ui

import (
	"github.com/rivo/tview"
	"installer/pkg/constants"
	menu2 "installer/pkg/menu"
	setup2 "installer/pkg/setup"
)

// Gui struct.
type Gui struct {
	App   *tview.Application
	Pages *tview.Pages
}

// Infos struct.
type Infos struct {
	*tview.TextView
}

func newInfo(g *Gui) *Infos {
	infos := &Infos{
		TextView: tview.NewTextView().SetText("").SetWordWrap(true).SetWrap(true),
	}
	infos.SetTitle("Cluster Setup").SetTitleAlign(tview.AlignCenter)
	infos.SetBorder(true)
	return infos
}

func (g *Gui) initGUI(isHA bool) {

	setup := setup2.NewSampleSetupStructure()

	info := newInfo(g)
	menu := menu2.NewMenu()
	role := newRoleGrid(g, menu, setup, isHA)

	cni := newCnis(g, menu, setup)
	storage := newStorages(g, menu, setup)
	nodeAllocate := newNodeAllocates(g, menu, setup)
	feature := newFeature(g, menu)

	kubernetes := newKubernetes(g, menu, setup)

	info.SetText(constants.SetupListTotalIntro)

	gridList := tview.NewGrid().SetRows(10, -1).SetColumns(-1, -5).
		AddItem(info, 0, 0, 1, 2, 0, 0, false).
		AddItem(menu, 1, 0, 1, 1, 0, 0, true).
		AddItem(role, 1, 1, 1, 1, 0, 0, false)

	var c = role
	menu.SetSelectedFunc(func(row int, column int) {
		switch row {
		case 0:
			info.SetText(constants.SetupListClusterIntro)
			gridList.RemoveItem(c)
			c = role
			gridList.AddItem(role, 1, 1, 1, 1, 0, 0, true)
			g.App.SetFocus(role)
		case 1:
			info.SetText(constants.SetupListKubernetesIntro)
			gridList.RemoveItem(c)
			c = kubernetes
			gridList.AddItem(kubernetes, 1, 1, 1, 1, 0, 0, true)
			g.App.SetFocus(kubernetes)
		case 2:
			info.SetText(constants.SetupListDockerIntro)
			gridList.RemoveItem(c)
			c = cni
			gridList.AddItem(cni, 1, 1, 1, 1, 0, 0, true)
			g.App.SetFocus(cni)
		case 3:
			info.SetText(constants.SetupListEtcdIntro)
			gridList.RemoveItem(c)
			c = storage
			gridList.AddItem(storage, 1, 1, 1, 1, 0, 0, true)
			g.App.SetFocus(storage)
		case 4:
			info.SetText(constants.SetupListNodeIntro)
			gridList.RemoveItem(c)
			c = nodeAllocate
			gridList.AddItem(nodeAllocate, 1, 1, 1, 1, 0, 0, true)
			g.App.SetFocus(nodeAllocate)
		case 5:

			info.SetText(constants.SetupListFeatureIntro)

			gridList.RemoveItem(c)
			c = feature
			gridList.AddItem(feature, 1, 1, 1, 1, 0, 0, true)
			g.App.SetFocus(feature)

		case 6:
			gridList.RemoveItem(c)
			g.Pages.RemovePage("main")
			g.setupLog(setup)

		case 7:
			gridList.RemoveItem(c)
			g.Pages.RemovePage("main")
			g.SetupChoice()
		}

	})
	g.Pages = tview.NewPages().
		AddAndSwitchToPage("main", gridList, true)
	g.App.SetRoot(g.Pages, true).Run()
}

// Modal creates small window in UI.
func (g *Gui) Modal(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(0, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}
