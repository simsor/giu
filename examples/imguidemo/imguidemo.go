package main

import (
	g "github.com/simsor/giu"
	"github.com/simsor/giu/imgui"
)

func loop() {
	imgui.ShowDemoWindow(nil)
}

func main() {
	wnd := g.NewMasterWindow("Widgets", 1024, 768, 0, nil)
	wnd.Main(loop)
}
