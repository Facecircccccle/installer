package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"installer/pkg/menu"
	setup2 "installer/pkg/setup"
)

func setCnisConnection(d *setup2.Dockers, g *Gui, m *menu.Menus, s *setup2.Setup) {
	d.AddButton("done", func() {
		result, reason := setup2.DockerAndKubernetesVersionCheck(d, s)
		if result == true {
			RUNTIME_SET = true
			d.SetEntries(s)
			g.App.SetFocus(m)
		} else {
			modal := tview.NewModal().
				SetText(reason).
				AddButtons([]string{"ok"})
			modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "ok" {
					g.Pages.SwitchToPage("main")
					g.App.SetFocus(d)
				}
			})
			g.Pages.AddAndSwitchToPage("Modal", g.Modal(modal, 40, 16), true).ShowPage("main")
		}
	}).SetButtonsAlign(tview.AlignRight)
}

func setStoragesConnection(e *setup2.Etcds, g *Gui, m *menu.Menus, s *setup2.Setup) {
	e.AddButton("done", func() {
		result, reason := setup2.EtcdAndKubernetesVersionCheck(e, s)
		if result == true {
			STORAGE_SET = true
			e.SetEntries(s)
			g.App.SetFocus(m)
		} else {
			modal := tview.NewModal().
				SetText(reason).
				AddButtons([]string{"ok"})
			modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "ok" {
					e.SetEntries(s)
					g.App.SetFocus(m)
				}
			})
			g.Pages.AddAndSwitchToPage("Modal", g.Modal(modal, 40, 16), true).ShowPage("main")
		}

	}).SetButtonsAlign(tview.AlignRight)
}

func setKubernetesConnection(c *setup2.Clusters, networking *setup2.NetWorkings, netPlugin *setup2.NetPlugins, admissionPlugin *setup2.AdmissionPlugins, g *Gui, m *menu.Menus, s *setup2.Setup) {
	c.AddButton("next", func() {
		result, reason := setup2.ClusterInfoCheck(c, s)
		if result == true {
			c.SetEntries(s)
			g.App.SetFocus(networking.Form)
		} else {
			modal := tview.NewModal().
				SetText(reason).
				AddButtons([]string{"ok"})
			modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "ok" {
					g.Pages.SwitchToPage("main")
					g.App.SetFocus(c)
				}
			})
			g.Pages.AddAndSwitchToPage("Modal", g.Modal(modal, 40, 16), true).ShowPage("main")
		}

	}).SetButtonsAlign(tview.AlignRight)

	networking.AddButton("next", func() {
		result, reason := setup2.NetworkingCheck(networking)
		if result == true {
			networking.SetEntries(s)
			g.App.SetFocus(netPlugin.Form)
		} else {
			modal := tview.NewModal().
				SetText(reason).
				AddButtons([]string{"ok"})
			modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "ok" {
					g.Pages.SwitchToPage("main")
					g.App.SetFocus(networking)
				}
			})
			g.Pages.AddAndSwitchToPage("Modal", g.Modal(modal, 40, 16), true).ShowPage("main")
		}
	}).SetButtonsAlign(tview.AlignRight)

	netPlugin.AddButton("next", func() {
		netPlugin.SetEntries(s)
		g.App.SetFocus(admissionPlugin.Form)
	}).SetButtonsAlign(tview.AlignRight)

	admissionPlugin.AddButton("done", func() {
		KUBE_SET = true
		admissionPlugin.SetEntries(s)
		g.App.SetFocus(m)
	}).SetButtonsAlign(tview.AlignRight)
}

func setAllocatesConnection(a *setup2.Allocates, g *Gui, m *menu.Menus, s *setup2.Setup) {
	a.AddButton("done", func() {
		result, reason := setup2.AllocateCheck(a)
		if result == true {
			a.SetEntries(s)
			g.App.SetFocus(m)
		} else {
			modal := tview.NewModal().
				SetText(reason).
				AddButtons([]string{"ok"})
			modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "ok" {
					g.Pages.SwitchToPage("main")
				}
			})
			g.Pages.AddAndSwitchToPage("Modal", g.Modal(modal, 40, 16), true).ShowPage("main")
		}
	}).SetButtonsAlign(tview.AlignRight)
}

func setKeybinding(g *Gui, r *myTable, m *menu.Menus, s *setup2.Setup, isHA bool) {
	r.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyDelete:
			deleteRoleForm(r, s)
		}
		switch event.Rune() {
		case 'i':
			importRoleForm(g, r, s, isHA)
		case 'b':
			result, reason := setup2.InputRoleBackCheck(s, isHA)
			if result {
				ROLE_SET = true
				g.App.SetFocus(m)
			} else {
				modal := tview.NewModal().
					SetText(reason).
					AddButtons([]string{"ok"})
				modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "ok" {
						g.Pages.RemovePage("Modal")
						g.App.SetFocus(r)
					}
				})
				g.Pages.AddAndSwitchToPage("Modal", g.Modal(modal, 40, 16), true).ShowPage("main")
			}
		}
		return event
	})
}
