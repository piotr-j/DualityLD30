package main

import (
	"github.com/nsf/termbox-go"
)

type Wall struct {
	pos  Position
	form Form
	game *Game
}

func (o *Wall) Init(x, y int, game *Game) {
	o.pos.X = x
	o.pos.Y = y
	o.game = game
	o.form = FormAngel
}

func (o *Wall) SetForm(form Form) {
	o.form = form
}

func (o *Wall) Type() GameObjectType {
	return GOWall
}

func (o *Wall) Explode() {

}

func (o *Wall) Contains(x, y int) bool {
	if o.pos.X <= x && o.pos.X+3 > x {
		if o.pos.Y <= y && o.pos.Y+3 > y {
			return true
		}
	}
	return false
}

func (o *Wall) Draw() {
	if o.form == FormDemon {
		//o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorRed)
		o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorRed)
		o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorRed)
		//o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorMagenta)

		o.game.SetCell(o.pos.X, o.pos.Y+1, termbox.ColorRed)
		o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorRed)

		//o.game.SetCell(o.pos.X, o.pos.Y+2, termbox.ColorRed)
		o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorRed)
		o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorRed)
		//o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorMagenta)
	} else {
		o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorYellow)
		o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorYellow)

		o.game.SetCell(o.pos.X, o.pos.Y+1, termbox.ColorCyan)
		o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorCyan)

		o.game.SetCell(o.pos.X, o.pos.Y+2, termbox.ColorYellow)
		o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorYellow)
	}
}

func (o *Wall) Update() {

}
