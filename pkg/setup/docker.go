package setup

import (
	"github.com/rivo/tview"
	"installer/pkg/constants"
)

// Dockers struct.
type Dockers struct {
	*tview.Form
}

// NewDocker build the Docker Form in UI.
func NewDocker() *Dockers {

	dockers := &Dockers{
		Form: tview.NewForm().
			AddDropDown("Version", constants.DockerVersion, 0, nil).
			AddInputField("Registries", "core.harbor.k8s.ebupt.com", 0, nil, nil),
	}
	dockers.SetBorder(true).SetTitle("docker info").SetTitleAlign(tview.AlignCenter)
	dockers.SetItemPadding(1).SetBorderPadding(0, 0, 0, 1)

	return dockers
}

// SetEntries set entries for setup structure.
func (d Dockers) SetEntries(s *Setup) {

	_, s.Docker.Version = d.GetFormItemByLabel("Version").(*tview.DropDown).GetCurrentOption()
	s.Docker.RepositoryName = d.GetFormItemByLabel("Registries").(*tview.InputField).GetText()
}
