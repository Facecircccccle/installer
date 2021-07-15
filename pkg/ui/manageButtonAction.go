package ui

import (
	"context"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"installer/pkg/manage"
	"installer/pkg/menu"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)

func SetNodeKeybinding(g *Gui, r *MyTable, clientset *kubernetes.Clientset, m *menu.Menus, i *Infos, log *MyText) {
	r.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		//g.setGlobalKeybinding(event)
		switch event.Key() {
		case tcell.KeyTAB:
			g.App.SetFocus(log)
		}
		switch event.Rune() {
		case 'i':
			ImportNodeForm(g, r, clientset, m, i, log)
		case 'b':
			g.App.SetFocus(m)
		}
		return event
	})

}

func SetNodeEntries(g *Gui, r *MyTable, clientset *kubernetes.Clientset, i *Infos, log *MyText) {
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
