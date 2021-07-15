package setup

import (
	"github.com/rivo/tview"
	"installer/pkg/constants"
)

// Etcds struct.
type Etcds struct {
	*tview.Form
}

// NewEtcd build the Etcd Form in UI.
func NewEtcd() *Etcds {

	etcds := &Etcds{
		Form: tview.NewForm().
			AddDropDown("Version", constants.EtcdVersion, 0, nil).
			AddInputField("DataDir", "/var/lib/etcd", 0, nil, nil),
	}
	etcds.SetBorder(true).SetTitle("etcd info").SetTitleAlign(tview.AlignLeft)
	etcds.SetItemPadding(1).SetBorderPadding(0, 0, 0, 1)

	return etcds
}

// SetEntries set entries for setup structure.
func (e Etcds) SetEntries(s *Setup) {

	_, s.Etcd.Version = e.GetFormItemByLabel("Version").(*tview.DropDown).GetCurrentOption()
	s.Etcd.DataDir = e.GetFormItemByLabel("DataDir").(*tview.InputField).GetText()
}
