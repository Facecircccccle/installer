module installer

go 1.13

require (
	github.com/gdamore/tcell/v2 v2.2.1
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/rivo/tview v0.0.0-20210514202809-22dbf8415b04
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/metrics v0.21.3
	k8s.io/utils v0.0.0-20210527160623-6fdb442a123b // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.21.1
