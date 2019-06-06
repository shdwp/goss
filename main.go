package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"goss/game"
	_map "goss/map"
	"goss/ui"
	"time"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	m := _map.NewMap(_map.Gen())

	runecanvas := game.NewRuneCanvas(win)
	loop := game.NewLoop(runecanvas)
	loop.AddComponent(&ui.TestWindow{0,0,m})

	sync := time.Tick(time.Second / 60)
	for !win.Closed() {
		for _, r := range win.Typed() {
			loop.Input(r)
		}

		win.Clear(colornames.Black)
		runecanvas.Clear()

		loop.Render()
        runecanvas.Draw()

		win.Update()
		<- sync
	}
}

func main() {
	pixelgl.Run(run)
}