package main

import (
	"github.com/nsf/termbox-go"
)

type Portal struct {
	pos         Position
	game        *Game
	form        Form
	switchForm  bool
	switchDelay int
}

func (o *Portal) Init(x, y int, game *Game) {
	o.pos.X = x
	o.pos.Y = y
	o.game = game
	o.switchDelay = 0
}
func (o *Portal) SetForm(form Form) {
	o.form = form
}
func (o *Portal) Type() GameObjectType {
	return GOPortalOpen
}

func (o *Portal) Explode() {

}

func (o *Portal) Contains(x, y int) bool {
	if o.pos.X <= x && o.pos.X+3 > x {
		if o.pos.Y <= y && o.pos.Y+3 > y {
			return true
		}
	}
	return false
}

func (o *Portal) Draw() {
	if o.form == FormAngel {
		o.game.SetCell(o.pos.X-1, o.pos.Y-1, termbox.ColorRed)
		// o.game.SetCell(o.pos.X,   o.pos.Y-1, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X+1, o.pos.Y-1, termbox.ColorRed)
		// o.game.SetCell(o.pos.X+2, o.pos.Y-1, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+3, o.pos.Y-1, termbox.ColorRed)

		// o.game.SetCell(o.pos.X-1, o.pos.Y, termbox.ColorRed)
		o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorRed)
		o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorRed)
		// o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorMagenta)

		o.game.SetCell(o.pos.X-1, o.pos.Y+1, termbox.ColorRed)
		o.game.SetCell(o.pos.X, o.pos.Y+1, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorRed)

		// o.game.SetCell(o.pos.X-1, o.pos.Y+2, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X, o.pos.Y+2, termbox.ColorRed)
		o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorRed)
		// o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorMagenta)
	} else {
		o.game.SetCell(o.pos.X-1, o.pos.Y-1, termbox.ColorCyan)
		// o.game.SetCell(o.pos.X,   o.pos.Y-1, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X+1, o.pos.Y-1, termbox.ColorCyan)
		// o.game.SetCell(o.pos.X+2, o.pos.Y-1, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+3, o.pos.Y-1, termbox.ColorCyan)

		// o.game.SetCell(o.pos.X-1, o.pos.Y, termbox.ColorRed)
		o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorYellow)
		o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorYellow)
		// o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorMagenta)

		o.game.SetCell(o.pos.X-1, o.pos.Y+1, termbox.ColorYellow)
		o.game.SetCell(o.pos.X, o.pos.Y+1, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorYellow)

		// o.game.SetCell(o.pos.X-1, o.pos.Y+2, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X, o.pos.Y+2, termbox.ColorYellow)
		o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorYellow)
		// o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorMagenta)
	}
}

func (o *Portal) Update() {
	if g := o.game.GameObjectAtWithType(o.pos.X+1, o.pos.Y+1, GOPlayer); g != nil {
		gameObject := *g
		switch gameObject.Type() {
		case GOPlayer:
			if !o.switchForm {
				o.switchForm = true
				if o.form == FormAngel {
					o.game.SetForm(FormDemon)
				} else {
					o.game.SetForm(FormAngel)
				}
				// 1 second
				o.switchDelay = 60
			}

		default:
			// do nothing
		}
	} else {
		if o.switchDelay > 0 {
			o.switchDelay--
			if o.switchDelay == 0 {
				o.switchForm = false
			}
		}
	}
}
