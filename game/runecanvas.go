package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"log"
)

type (
	RuneCanvas struct {
		window *pixelgl.Window
		atlas *text.Atlas
		text *text.Text
	}
)

func NewRuneCanvas(window *pixelgl.Window) RuneCanvas {
    return RuneCanvas{
    	window,
    	text.NewAtlas(basicfont.Face7x13, text.ASCII),
    	text.New(pixel.V(0, 0), text.NewAtlas(basicfont.Face7x13, text.ASCII)),
	}
}

func (self *RuneCanvas) Clear() {
	self.text.Clear()
}

func (self *RuneCanvas) Put(r rune, x uint, y uint, color color.Color) {
	self.text.Dot.X = float64(x * 10)
	self.text.Dot.Y = self.window.Bounds().H() / 10 - float64(y * 10)

	self.text.Color = color
	if _, e := self.text.WriteRune(r); e != nil {
        log.Println(e)
	}
}

func (self *RuneCanvas) Draw() {
	im := pixel.IM
	im = im.Scaled(pixel.V(0, 0), 0.5)
	im = im.Moved(pixel.V(0, self.window.Bounds().H() - 90))
    self.text.Draw(self.window, im)
}

