package main

import (
	"github.com/nsf/termbox-go"
)

type Gate struct {
	pos         Position
	game        *Game
	form        Form
	exploded    bool
	switchForm  bool
	switchDelay int
}

func (o *Gate) Init(x, y int, game *Game) {
	o.pos.X = x
	o.pos.Y = y
	o.game = game
	o.switchDelay = 0
	o.game.AddCollideAt(x, y+2)
	o.game.AddCollideAt(x, y+1)
	o.game.AddCollideAt(x, y)
}

func (o *Gate) SetForm(form Form) {
	o.form = form
}

func (o *Gate) Type() GameObjectType {
	if o.exploded {
		return GOGateOpen
	} else {
		return GOGateClosed
	}
}

func (o *Gate) Explode() {
	if o.form == FormDemon {
		o.exploded = true
		o.game.RemoveCollideAt(o.pos.X, o.pos.Y+2)
		o.game.RemoveCollideAt(o.pos.X, o.pos.Y+1)
		o.game.RemoveCollideAt(o.pos.X, o.pos.Y)
	}
}

func (o *Gate) Contains(x, y int) bool {
	if o.pos.X <= x && o.pos.X+3 > x {
		if o.pos.Y <= y && o.pos.Y+3 > y {
			return true
		}
	}
	return false
}

func (o *Gate) Draw() {
	if o.exploded {
		return
	}
	if o.form == FormDemon {
		o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorMagenta)
		// o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorWhite)
		// o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorMagenta)

		// o.game.SetCell(o.pos.X, o.pos.Y+1, termbox.ColorRed)
		o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorRed)
		o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorRed)
		// o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorRed)

		o.game.SetCell(o.pos.X, o.pos.Y+2, termbox.ColorMagenta)
		// o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorWhite)
		// o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorMagenta)

	} else {
		o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorWhite)
		// o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorWhite)
		// o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorWhite)

		o.game.SetCell(o.pos.X, o.pos.Y+1, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorYellow)
		o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorYellow)
		o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorWhite)

		o.game.SetCell(o.pos.X, o.pos.Y+2, termbox.ColorWhite)
		// o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorWhite)
		// o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorWhite)
	}
}

func (o *Gate) Update() {

}
