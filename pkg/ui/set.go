package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"installer/pkg/constants"
)

// SetupChoice provides ways to setup cluster.
func (g *Gui) SetupChoice() {

	listOneMaster := newlist("One Master Setup")
	listBack := newlist("Back")
	listHA := newlist("Multi Master(HA) Setup")

	gridSetup := tview.NewGrid().SetRows(-4, -4, -1, -4, -1, -4, -1, -4, -4).SetColumns(-6, -1, -3, -6, -6).
		AddItem(tview.NewTextView().SetText(constants.HowToSetup), 1, 1, 1, 3, 1, 1, false).
		AddItem(listOneMaster, 2, 1, 1, 2, 0, 0, true).
		AddItem(tview.NewTextView().SetText(constants.OneMasterSetup), 3, 2, 1, 2, 0, 0, false).
		AddItem(listHA, 4, 1, 1, 2, 2, 1, true).
		AddItem(tview.NewTextView().SetText(constants.HAClusterSetup), 5, 2, 1, 2, 0, 0, false).
		AddItem(listBack, 6, 1, 1, 2, 3, 1, true)

	g.App.SetFocus(listOneMaster)
	listOneMaster.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Key(); key {
		case tcell.KeyEnter:
			g.Pages.RemovePage("Setup")
			g.initGUI(false)
		case tcell.KeyTAB, tcell.KeyDown:
			g.App.SetFocus(listHA)
		case tcell.KeyUp:
			g.App.SetFocus(listBack)
		}
		return event
	})
	listBack.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Key(); key {
		case tcell.KeyEnter:
			g.Pages.RemovePage("Setup")
			g.Welcome()
		case tcell.KeyTAB, tcell.KeyDown:
			g.App.SetFocus(listOneMaster)
		case tcell.KeyUp:
			g.App.SetFocus(listHA)
		}
		return event
	})
	listHA.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Key(); key {
		case tcell.KeyEnter:
			g.Pages.RemovePage("Setup")
			g.initGUI(true)
		case tcell.KeyTAB, tcell.KeyDown:
			g.App.SetFocus(listBack)
		case tcell.KeyUp:
			g.App.SetFocus(listOneMaster)
		}
		return event
	})

	g.Pages.AddAndSwitchToPage("Setup", gridSetup, true)
}
