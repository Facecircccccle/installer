package ui

import (
	"bufio"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type myGrid struct {
	*tview.Grid
}

type myTable struct {
	*tview.Table
}

type myText struct {
	*tview.TextView
}








// Command processes shell command and output result in UI.
func (g *Gui) Command(cmd string, log *myText) error {
	c := exec.Command("bash", "-c", cmd)
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			if readString == "\n" {
				continue
			}
			log.SetText(log.GetText(false) + readString)
			log.ScrollToEnd()
			g.App.ForceDraw()
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}

// NewTableFromFile create table UI from the formatted text.
func (g *Gui) NewTableFromFile(s string) {
	table := tview.NewTable()
	f, err := os.Open(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	buf := bufio.NewReader(f)
	var count = 0
	for {
		b, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err.Error())
		}
		if count == 0 {
			headers := strings.Fields(string(b))
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
		} else {
			contents := strings.Fields(string(b))
			for i, content := range contents {
				table.SetCell(count+1, i, tview.NewTableCell(content).
					SetTextColor(tcell.ColorLightYellow).SetMaxWidth(1).SetExpansion(1))
			}
		}
		count++
	}
}
