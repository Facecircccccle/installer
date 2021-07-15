package setup

import (
	"github.com/rivo/tview"
)

type Allocates struct {
	*tview.Form
}

func NewAllocate() *Allocates {

	allocates := &Allocates{
		Form: tview.NewForm().
			AddInputField("KubeReservedCPU", "1", 0, nil, nil).
			AddInputField("SysReservedCPU", "1", 0, nil, nil).
			AddInputField("KubeMemory", "500Mi", 0, nil, nil).
			AddInputField("SysMemory", "500Mi", 0, nil, nil).
			AddInputField("KubeStorage", "10Gi", 0, nil, nil).
			AddInputField("SysStorage", "10Gi", 0, nil, nil).
			AddInputField("EvictionMemory", "500Mi", 0, nil, nil).
			AddInputField("EvictionNodefs", "10%", 0, nil, nil),
	}
	allocates.SetBorder(true).SetTitle("Allocates info").SetTitleAlign(tview.AlignLeft)
	allocates.SetItemPadding(1).SetBorderPadding(0, 0, 0, 1)

	return allocates
}

func (a Allocates) SetEntries(s *Setup) {

	allocateTmp := NodeAllocate{
		KubeReservedCPU: a.GetFormItemByLabel("KubeReservedCPU").(*tview.InputField).GetText(),
		SysReservedCPU:  a.GetFormItemByLabel("SysReservedCPU").(*tview.InputField).GetText(),
		KubeMemory:      a.GetFormItemByLabel("KubeMemory").(*tview.InputField).GetText(),
		SysMemory:       a.GetFormItemByLabel("SysMemory").(*tview.InputField).GetText(),
		KubeStorage:     a.GetFormItemByLabel("KubeStorage").(*tview.InputField).GetText(),
		SysStorage:      a.GetFormItemByLabel("SysStorage").(*tview.InputField).GetText(),
		EvictionMemory:  a.GetFormItemByLabel("EvictionMemory").(*tview.InputField).GetText(),
		EvictionNodefs:  a.GetFormItemByLabel("EvictionNodefs").(*tview.InputField).GetText(),
	}
	for i := 0; i < s.NodeCount; i++ {
		s.Nodes[i].Allocate = allocateTmp
	}
}
