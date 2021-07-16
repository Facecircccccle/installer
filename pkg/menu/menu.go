package menu

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Menus struct.
type Menus struct {
	*tview.Table
}

func (m Menus) setEntries() {
	Menu := []string{"Cluster", "Kubernetes", "Docker", "Etcd", "Node Allocate", "Feature Gate", "start", "Back"}
	table := m.Clear()
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
}

// NewMenu build the Menu table in UI.
func NewMenu() *Menus {
	menus := &Menus{
		Table: tview.NewTable().SetSelectable(true, true).SetFixed(1, 1),
	}

	menus.SetTitle("Setup List").SetTitleAlign(tview.AlignCenter)
	menus.SetBorder(true)
	menus.setEntries()
	return menus
}
