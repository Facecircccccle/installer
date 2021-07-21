package ui

import (
	"context"
	"flag"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"installer/pkg/constants"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func newlist(s string) *tview.List {
	list := tview.NewList().AddItem(s, "", '>', nil)
	list.SetShortcutColor(tcell.ColorYellow)
	//list.SetMainTextColor(tcell.ColorBlack)
	//list.SetSelectedBackgroundColor(tcell.ColorBlack)
	list.SetSelectedFocusOnly(true)
	return list
}

// Welcome build the main window UI.
func (g *Gui) Welcome() {
	listSetup := newlist("cluster Setup")
	listManage := newlist("cluster management")
	listExit := newlist("exit")

	grid := tview.NewGrid().SetRows(-5, -5, -1, -5, -1, -5, -1, -5).SetColumns(-6, -1, -3, -6, -6).
		AddItem(tview.NewTextView().SetText(constants.EbKubernetes), 1, 1, 1, 3, 1, 1, false).
		AddItem(listSetup, 2, 1, 1, 2, 0, 0, true).
		AddItem(tview.NewTextView().SetText(constants.EbKubernetesInstall), 3, 2, 1, 2, 0, 0, false).
		AddItem(listManage, 4, 1, 1, 2, 2, 1, true).
		AddItem(tview.NewTextView().SetText(constants.EbKubernetesManage), 5, 2, 1, 2, 0, 0, false).
		AddItem(listExit, 6, 1, 1, 2, 3, 1, true)

	listSetup.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Key(); key {
		case tcell.KeyEnter:
			g.Pages.RemovePage("Welcome")
			g.SetupChoice()
		case tcell.KeyTAB, tcell.KeyDown:
			g.App.SetFocus(listManage)
		case tcell.KeyUp:
			g.App.SetFocus(listExit)
		}
		return event
	})
	listManage.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Key(); key {
		case tcell.KeyEnter:
			g.importClusterForm(listManage)
		case tcell.KeyTAB, tcell.KeyDown:
			g.App.SetFocus(listExit)
		case tcell.KeyUp:
			g.App.SetFocus(listSetup)
		}
		return event
	})
	listExit.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Key(); key {
		case tcell.KeyEnter:
			g.App.Stop()
		case tcell.KeyTAB, tcell.KeyDown:
			g.App.SetFocus(listSetup)
		case tcell.KeyUp:
			g.App.SetFocus(listManage)
		}
		return event
	})

	g.Pages = tview.NewPages().
		AddAndSwitchToPage("Welcome", grid, true)

	_ = g.App.SetRoot(g.Pages, true).Run()
}

func (g *Gui) importClusterForm(l *tview.List) {

	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignLeft)
	form.SetTitle("Import Cluster")
	form.
		//AddInputField("IP", "", inputWidth, nil, nil).
		//AddInputField("Code", "", inputWidth, nil, nil).
		AddCheckbox("config file is already in ~/.kube/config", true, nil).
		//	AddInputField("Config addr", "", inputWidth, nil, nil).
		AddButton("Load", func() {
			if form.GetFormItemByLabel("config file is already in ~/.kube/config").(*tview.Checkbox).IsChecked() {
				result, reason, clientset := configCheck()
				if result {
					Manage(g, clientset)

					g.Pages.RemovePage("newClusterForm")
					g.App.SetFocus(l)
				} else {
					modal := tview.NewModal().
						SetText(reason).
						AddButtons([]string{"ok"})
					modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonLabel == "ok" {
							g.Pages.RemovePage("Modal")
							g.App.SetFocus(form)
							g.Pages.AddAndSwitchToPage("newClusterForm", g.Modal(form, 70, 10), true).ShowPage("Welcome")
						}
					})
					g.Pages.AddAndSwitchToPage("Modal", g.Modal(modal, 40, 10), true).ShowPage("Welcome")
				}
			} else {
				modal := tview.NewModal().
					SetText("Please copy the config file in ~/.kube/config, then tick the box and Load.").
					AddButtons([]string{"ok"})
				modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "ok" {
						g.Pages.RemovePage("Modal")
						g.App.SetFocus(form)
						g.Pages.AddAndSwitchToPage("newClusterForm", g.Modal(form, 70, 10), true).ShowPage("Welcome")
					}
				})
				g.Pages.AddAndSwitchToPage("Modal", g.Modal(modal, 40, 10), true).ShowPage("Welcome")
			}

		}).SetButtonsAlign(tview.AlignRight).
		AddButton("Cancel", func() {
			g.Pages.RemovePage("newClusterForm")
			g.App.SetFocus(l)
		}).SetButtonsAlign(tview.AlignRight)

	g.Pages.AddAndSwitchToPage("newClusterForm", g.Modal(form, 80, 10), true).ShowPage("Welcome")
}

var kubeconfig *string

func configCheck() (bool, string, *kubernetes.Clientset) {
	var kubeconfig *string
	var clientset *kubernetes.Clientset

	if kubeconfig == nil {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "config")
		}
		flag.Parse()
		kubeconfig = kubeconfig
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return false, "build config error, check input file path", nil
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return false, "create clientset error, check input file available", nil
	}

	_, err = clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, "can not connect to the cluster, please check the config file.", nil
	}

	return true, "", clientset
}
