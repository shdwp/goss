package ui

import (
	"golang.org/x/image/colornames"
	"goss/game"
	_map "goss/map"
)

type TestWindow struct {
	X, Y uint
	M _map.Map
}

func (self *TestWindow) Render(canvas game.RuneCanvas) {
	canvas.Put('@', self.X, self.Y, colornames.Red)

	for x, row := range self.M.Tiles {
		for y, tile := range row {
			if tile != nil {
				canvas.Put(tile.Rune, uint(x)+self.X, uint(y) + self.Y, tile.Color)
			}
		}
	}
}

func (self *TestWindow) Process(event game.Event, ch chan game.ProcessResult) {
    switch e := event.(type) {
	case game.InputEvent:
		switch e.Input {
		case 'j':
			self.Y++
		case 'k':
			self.Y--
		case 'h':
			self.X--
		case 'l':
			self.X++
		default:
			break
		}
	}

	ch <- game.DoNothing
}


