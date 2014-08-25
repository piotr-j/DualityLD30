package main

import (
	"github.com/nsf/termbox-go"
)

const (
	SHOOT_DELAY = 45
	MOVE_DELAY  = 10
)

const (
	EnemyLeft  int = -1
	EnemyRight int = 1
)

type Enemy struct {
	pos         Position
	game        *Game
	form        Form
	goType      GameObjectType
	exploded    bool
	switchForm  bool
	switchDelay int
	shootDelay  int
	moveDelay   int
	moveDir     int
	shootDir    int
}

func (o *Enemy) Init(x, y int, game *Game) {
	o.pos.X = x
	o.pos.Y = y
	o.game = game
	o.switchDelay = 0
	o.goType = GOEnemy
	o.shootDir = EnemyLeft
	o.shootDelay = SHOOT_DELAY
	o.moveDir = -1
}

func (o *Enemy) InitEnemy(x, y int, game *Game, shootDir int) {
	o.pos.X = x
	o.pos.Y = y
	o.game = game
	o.switchDelay = 0
	o.goType = GOEnemy
	o.shootDir = shootDir
	o.shootDelay = SHOOT_DELAY
	o.moveDir = -1
}

func (o *Enemy) SetForm(form Form) {
	o.form = form
}

func (o *Enemy) Type() GameObjectType {
	return o.goType

}

func (o *Enemy) Explode() {
	if o.form == FormDemon {
		o.exploded = true
		o.goType = GOEnemyExploded
	}
}

func (o *Enemy) Contains(x, y int) bool {
	if o.form != FormDemon {
		return false
	}
	if o.pos.X-2 <= x && o.pos.X+3 > x {
		if o.pos.Y-2 <= y && o.pos.Y+3 > y {
			return true
		}
	}
	return false
}

func (o *Enemy) Draw() {
	if o.form == FormDemon && !o.exploded {
		if o.shootDir == EnemyRight {
			o.game.SetCell(o.pos.X-1, o.pos.Y-2, termbox.ColorRed)
			// o.game.SetCell(o.pos.X,   o.pos.Y-2, termbox.ColorMagenta)
			// o.game.SetCell(o.pos.X+1, o.pos.Y-2, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+2, o.pos.Y-2, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+3, o.pos.Y-2, termbox.ColorRed)

			// o.game.SetCell(o.pos.X-1, o.pos.Y-1, termbox.ColorRed)
			o.game.SetCell(o.pos.X, o.pos.Y-1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+1, o.pos.Y-1, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y-1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+3, o.pos.Y-1, termbox.ColorRed)

			o.game.SetCell(o.pos.X-1, o.pos.Y, termbox.ColorMagenta)
			o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorYellow)
			o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorMagenta)

			// o.game.SetCell(o.pos.X-1, o.pos.Y+1, termbox.ColorRed)
			o.game.SetCell(o.pos.X, o.pos.Y+1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorRed)

			o.game.SetCell(o.pos.X-1, o.pos.Y+2, termbox.ColorRed)
			// o.game.SetCell(o.pos.X,   o.pos.Y+2, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorMagenta)
			// o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorRed)
			o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorRed)
		} else {
			o.game.SetCell(o.pos.X-1, o.pos.Y-2, termbox.ColorRed)
			// o.game.SetCell(o.pos.X,   o.pos.Y-2, termbox.ColorMagenta)
			// o.game.SetCell(o.pos.X+1, o.pos.Y-2, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+2, o.pos.Y-2, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+3, o.pos.Y-2, termbox.ColorRed)

			// o.game.SetCell(o.pos.X-1, o.pos.Y-1, termbox.ColorRed)
			o.game.SetCell(o.pos.X, o.pos.Y-1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+1, o.pos.Y-1, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y-1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+3, o.pos.Y-1, termbox.ColorRed)

			o.game.SetCell(o.pos.X-1, o.pos.Y, termbox.ColorMagenta)
			o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorYellow)
			o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorMagenta)

			// o.game.SetCell(o.pos.X-1, o.pos.Y+1, termbox.ColorRed)
			o.game.SetCell(o.pos.X, o.pos.Y+1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorRed)

			o.game.SetCell(o.pos.X-1, o.pos.Y+2, termbox.ColorRed)
			// o.game.SetCell(o.pos.X,   o.pos.Y+2, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorMagenta)
			// o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorRed)
			o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorRed)
		}
	}
}

func (o *Enemy) Update() {
	if o.form != FormDemon || o.exploded {
		return
	}
	if o.shootDelay == 0 {
		proj := Projectile{}
		if o.shootDir == EnemyLeft {
			proj.InitProjectile(o.pos.X, o.pos.Y, o.game, DirLeft, GOProjectileEnemy)
		} else {
			proj.InitProjectile(o.pos.X, o.pos.Y, o.game, DirRight, GOProjectileEnemy)
		}
		o.game.AddGO(&proj)
		o.shootDelay = SHOOT_DELAY
	} else {
		o.shootDelay--
	}

	if o.moveDelay == 0 {
		if o.moveDir == -1 {
			if o.game.NoCollideAt(o.pos.X, o.pos.Y-4) {
				o.pos.Y += o.moveDir
			} else {
				o.moveDir *= -1
			}
		} else {
			if o.game.NoCollideAt(o.pos.X, o.pos.Y+4) {
				o.pos.Y += o.moveDir
			} else {
				o.moveDir *= -1
			}
		}

		o.moveDelay = MOVE_DELAY
	} else {
		o.moveDelay--
	}

}
