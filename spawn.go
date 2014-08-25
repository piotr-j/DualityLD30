package main

import (
	"github.com/nsf/termbox-go"
)

type Spawn struct {
	pos       Position
	game      *Game
	form      Form
	active    bool
	spawnType GameObjectType
}

func (o *Spawn) Init(x, y int, game *Game) {
	o.pos.X = x
	o.pos.Y = y
	o.game = game
	o.spawnType = GOSpawn
	o.active = false
}

func (o *Spawn) InitSpawn(x, y int, game *Game, spawnType GameObjectType) {
	o.pos.X = x
	o.pos.Y = y
	o.game = game
	o.spawnType = spawnType
	if spawnType == GOSpawnInitial {
		o.active = true
	}
}

func (o *Spawn) SetActive(active bool) {
	o.active = active
}

func (o *Spawn) SetForm(form Form) {
	o.form = form
}
func (o *Spawn) Type() GameObjectType {
	return o.spawnType
}

func (o *Spawn) Explode() {

}

func (o *Spawn) Pos() Position {
	return o.pos
}

func (o *Spawn) Contains(x, y int) bool {
	if o.pos.X <= x && o.pos.X+3 > x {
		if o.pos.Y <= y && o.pos.Y+3 > y {
			return true
		}
	}
	return false
}

func (o *Spawn) Draw() {
	if o.form == FormDemon {
		if o.active {
			o.game.SetCell(o.pos.X+1, o.pos.Y-1, termbox.ColorMagenta)
			o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorMagenta)
			o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorMagenta)
			o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorMagenta)
			o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorMagenta)
		}

		o.game.SetCell(o.pos.X-1, o.pos.Y+1, termbox.ColorRed)
		o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorRed)

		// o.game.SetCell(o.pos.X-1, o.pos.Y+2, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X, o.pos.Y+2, termbox.ColorRed)
		o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorRed)
		// o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorMagenta)
	} else {
		if o.active {
			o.game.SetCell(o.pos.X+1, o.pos.Y-1, termbox.ColorWhite)
			o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorWhite)
		}
		o.game.SetCell(o.pos.X-1, o.pos.Y+1, termbox.ColorYellow)
		// o.game.SetCell(o.pos.X,   o.pos.Y+1, termbox.ColorWhite)
		// o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorWhite)
		// o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorWhite)
		o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorYellow)

		// o.game.SetCell(o.pos.X-1, o.pos.Y+2, termbox.ColorMagenta)
		o.game.SetCell(o.pos.X, o.pos.Y+2, termbox.ColorYellow)
		o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorCyan)
		o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorYellow)
		// o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorMagenta)
	}
}

func (o *Spawn) Update() {
	if g := o.game.GameObjectAtWithType(o.pos.X+1, o.pos.Y+1, GOPlayer); g != nil {
		gameObject := *g
		switch gameObject.Type() {
		case GOPlayer:
			o.game.SetSpawnPoint(o)
			if o.spawnType == GOSpawnFinish {
				o.game.Win()
			}
		default:
			// do nothing
		}
	}
}
