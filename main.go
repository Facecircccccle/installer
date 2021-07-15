package main

import (
	"github.com/rivo/tview"
	"installer/pkg/ui"
)

// New return the new Gui structure.
func New() *ui.Gui {
	return &ui.Gui{
		App: tview.NewApplication(),
	}
}

func main() {
	gui := New()
	gui.Welcome()
}
