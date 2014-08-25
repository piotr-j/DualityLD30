package main

import (
	"github.com/nsf/termbox-go"
)

const ANIM_TIME = 5

type Trap struct {
	pos      Position
	game     *Game
	form     Form
	offset   int
	animTime int
	active   bool
	trapType GameObjectType
}

func (o *Trap) Init(x, y int, game *Game) {
	o.pos.X = x
	o.pos.Y = y
	o.game = game
	o.trapType = GOTrap
	o.SetForm(FormAngel)
	o.animTime = ANIM_TIME
}

func (o *Trap) SetForm(form Form) {
	o.form = form
	// x := o.pos.X
	// y := o.pos.Y
	// if form == FormAngel {
	// o.game.AddCollideAt(x, y)
	// o.game.AddCollideAt(x+1, y)
	// o.game.AddCollideAt(x+2, y)
	// o.game.AddCollideAt(x+3, y)
	// } else {

	// o.game.RemoveCollideAt(x, y)
	// o.game.RemoveCollideAt(x+1, y)
	// o.game.RemoveCollideAt(x+2, y)
	// o.game.RemoveCollideAt(x+3, y)
	// }
}

func (o *Trap) Type() GameObjectType {
	return o.trapType
}

func (o *Trap) Explode() {

}

func (o *Trap) Pos() Position {
	return o.pos
}

func (o *Trap) Contains(x, y int) bool {
	if o.pos.X <= x && o.pos.X+3 > x {
		if o.pos.Y <= y && o.pos.Y+3 > y {
			return true
		}
	}
	return false
}

func (o *Trap) Draw() {
	if o.form == FormAngel {
		for xx := 0; xx < 4; xx++ {
			for yy := 2; yy < 3; yy++ {
				o.game.SetCell(o.pos.X+xx, o.pos.Y+yy, termbox.ColorMagenta)
			}
		}
		o.game.SetCell(o.pos.X+o.offset, o.pos.Y+1, termbox.ColorRed)
		o.game.SetCell(o.pos.X+2+o.offset, o.pos.Y+2, termbox.ColorRed)
	} else {
		for xx := 0; xx < 4; xx++ {
			for yy := 2; yy < 3; yy++ {
				o.game.SetCell(o.pos.X+xx, o.pos.Y+yy, termbox.ColorWhite)
			}
		}
		o.game.SetCell(o.pos.X, o.pos.Y+1+o.offset, termbox.ColorYellow)
		o.game.SetCell(o.pos.X+2, o.pos.Y+2+o.offset, termbox.ColorYellow)
	}
}

func (o *Trap) Update() {
	if o.animTime == 0 {
		if o.offset == 0 {
			o.offset = 1
		} else {
			o.offset = 0
		}
		o.animTime = ANIM_TIME
	} else {
		o.animTime--
	}
	for i := 0; i < 4; i++ {
		if g := o.game.GameObjectAtWithType(o.pos.X+i, o.pos.Y+1, GOPlayer); g != nil {
			gameObject := *g
			switch gameObject.Type() {
			case GOPlayer:
				o.game.KillPlayer()
				return
			default:
				// do nothing
			}
		}
	}

}
