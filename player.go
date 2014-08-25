package main

import (
	"github.com/nsf/termbox-go"
)

const PLAYER_NAME = "ld30"

type Direction int

const (
	DirLeft Direction = iota
	DirRight
)

const (
	GRAVITY_DEMON        = 3
	GRAVITY_ANGEL        = 10
	JUMP_TIME            = 3
	PLAYER_HEIGHT_OFFSET = 1
	PLAYER_WIDTH_OFFSET  = 1
	DEMON_MAX_JUMPS      = 1
	ANGEL_MAX_JUMPS      = 2
)

type Player struct {
	pos         Position
	nextPos     Position
	canMove     bool
	gravDelay   int
	gravTime    int
	grounded    bool
	moveUp      bool
	moveLeft    bool
	moveRight   bool
	isJumping   bool
	jumpTime    int
	jumpCount   int
	maxJumps    int
	moveCount   int
	shootDelay  int
	shoot       bool
	goType      GameObjectType
	dir         Direction
	projectiles []Projectile
	game        *Game
	form        Form
	lives       int
	gravity     int
}

func (o *Player) Init(x, y int, game *Game) {
	o.pos.X = x
	o.pos.Y = y
	o.nextPos.X = x
	o.nextPos.Y = y
	o.canMove = true
	o.grounded = true
	o.form = FormAngel
	// updates for gravity update
	o.gravity = GRAVITY_ANGEL
	o.gravDelay = o.gravity
	o.gravTime = o.gravity
	o.game = game
	o.goType = GOPlayer
	o.projectiles = make([]Projectile, 8)
	o.lives = 3
	o.maxJumps = ANGEL_MAX_JUMPS
}

func (o *Player) ReSpawn(x, y int) {
	o.pos.X = x
	o.pos.Y = y
	o.nextPos.X = x
	o.nextPos.Y = y
	o.canMove = true
	o.grounded = true
	o.form = FormAngel
	o.gravity = GRAVITY_ANGEL
	o.gravDelay = o.gravity
	o.gravTime = o.gravity
	o.lives = 3
	o.maxJumps = ANGEL_MAX_JUMPS
}

func (o *Player) SetForm(form Form) {
	o.form = form
	if o.form == FormDemon {
		o.maxJumps = DEMON_MAX_JUMPS
		o.gravity = GRAVITY_DEMON
	} else {
		o.maxJumps = ANGEL_MAX_JUMPS
		o.gravity = GRAVITY_ANGEL
	}
}

func (o *Player) Type() GameObjectType {
	return o.goType
}

func (o *Player) GetLives() int {
	return o.lives
}

func (o *Player) IsAlive() bool {
	return o.lives > 0
}

func (o *Player) IsDead() bool {
	return o.lives <= 0
}

func (o *Player) Draw() {
	if o.IsDead() {
		return
	}
	// for xx := -1; xx < 2; xx++ {
	// for yy := -1; yy < 2; yy++ {
	// o.game.SetCell(o.pos.X+xx, o.pos.Y+yy, termbox.ColorCyan)
	// }
	// }
	if o.form == FormDemon {
		if o.dir == DirRight {
			o.game.SetCell(o.pos.X-2, o.pos.Y-1, termbox.ColorRed)
			o.game.SetCell(o.pos.X-1, o.pos.Y-1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X,   o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X+1, o.pos.Y-1, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y-1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorRed)

			o.game.SetCell(o.pos.X-2, o.pos.Y, termbox.ColorMagenta)
			o.game.SetCell(o.pos.X-1, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorYellow)
			o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorMagenta)

			// o.game.SetCell(o.pos.X-2, o.pos.Y+2, termbox.ColorMagenta)
			o.game.SetCell(o.pos.X-1, o.pos.Y+1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X,   o.pos.Y+2, termbox.ColorRed)
			o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorRed)
		} else {
			// o.game.SetCell(o.pos.X-2, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X-1, o.pos.Y-1, termbox.ColorRed)
			o.game.SetCell(o.pos.X, o.pos.Y-1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y-1, termbox.ColorRed)
			o.game.SetCell(o.pos.X+3, o.pos.Y-1, termbox.ColorRed)

			o.game.SetCell(o.pos.X-2, o.pos.Y, termbox.ColorMagenta)
			o.game.SetCell(o.pos.X-1, o.pos.Y, termbox.ColorYellow)
			o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorRed)
			o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorMagenta)

			// o.game.SetCell(o.pos.X-2, o.pos.Y+2, termbox.ColorMagenta)
			// o.game.SetCell(o.pos.X-1, o.pos.Y+2, termbox.ColorRed)
			o.game.SetCell(o.pos.X, o.pos.Y+1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+3, o.pos.Y+2, termbox.ColorRed)
		}
	} else {
		if o.dir == DirRight {
			// o.game.SetCell(o.pos.X-1, o.pos.Y-1, termbox.ColorRed)
			o.game.SetCell(o.pos.X, o.pos.Y-2, termbox.ColorYellow)
			// o.game.SetCell(o.pos.X+1, o.pos.Y-1, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y-2, termbox.ColorYellow)
			o.game.SetCell(o.pos.X+3, o.pos.Y-2, termbox.ColorCyan)

			o.game.SetCell(o.pos.X-1, o.pos.Y-1, termbox.ColorYellow)
			o.game.SetCell(o.pos.X, o.pos.Y-1, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+1, o.pos.Y-1, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+2, o.pos.Y-1, termbox.ColorWhite)
			// o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorMagenta)

			o.game.SetCell(o.pos.X-1, o.pos.Y, termbox.ColorWhite)
			o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorWhite)
			// o.game.SetCell(o.pos.X+1, o.pos.Y+1, termbox.ColorMagenta)
			o.game.SetCell(o.pos.X+2, o.pos.Y, termbox.ColorWhite)
			// o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorRed)

			// o.game.SetCell(o.pos.X-1, o.pos.Y+1, termbox.ColorWhite)
			// o.game.SetCell(o.pos.X,   o.pos.Y+2, termbox.ColorRed)
			// o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorMagenta)
			o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorYellow)
		} else {
			o.game.SetCell(o.pos.X-1, o.pos.Y-2, termbox.ColorCyan)
			o.game.SetCell(o.pos.X, o.pos.Y-2, termbox.ColorYellow)
			// o.game.SetCell(o.pos.X+1, o.pos.Y-1, termbox.ColorRed)
			o.game.SetCell(o.pos.X+2, o.pos.Y-2, termbox.ColorYellow)
			// o.game.SetCell(o.pos.X+3, o.pos.Y-1, termbox.ColorCyan)

			// o.game.SetCell(o.pos.X-1, o.pos.Y, termbox.ColorYellow)
			o.game.SetCell(o.pos.X, o.pos.Y-1, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+1, o.pos.Y-1, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+2, o.pos.Y-1, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+3, o.pos.Y-1, termbox.ColorYellow)

			// o.game.SetCell(o.pos.X-1, o.pos.Y+1, termbox.ColorWhite)
			o.game.SetCell(o.pos.X, o.pos.Y, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+1, o.pos.Y, termbox.ColorWhite)
			// o.game.SetCell(o.pos.X+2, o.pos.Y+1, termbox.ColorWhite)
			o.game.SetCell(o.pos.X+3, o.pos.Y, termbox.ColorWhite)

			o.game.SetCell(o.pos.X-1, o.pos.Y+1, termbox.ColorYellow)
			o.game.SetCell(o.pos.X, o.pos.Y+1, termbox.ColorWhite)
			// o.game.SetCell(o.pos.X+1, o.pos.Y+2, termbox.ColorMagenta)
			// o.game.SetCell(o.pos.X+2, o.pos.Y+2, termbox.ColorWhite)
			// o.game.SetCell(o.pos.X+3, o.pos.Y+1, termbox.ColorWhite)
		}
	}
}

func (o *Player) Contains(x, y int) bool {
	if o.pos.X-1 <= x && o.pos.X+1 >= x {
		if o.pos.Y-1 <= y && o.pos.Y+1 >= y {
			return true
		}
	}
	return false
}

func (o *Player) Left() {
	o.moveLeft = true
	o.dir = DirLeft
	if o.moveRight {
		o.moveCount = 0
	}
}

func (o *Player) Right() {
	o.moveRight = true
	o.dir = DirRight
	if o.moveLeft {
		o.moveCount = 0
	}
}

func (o *Player) Up() {
	o.moveUp = true
}

func (o *Player) Down() {
	// if o.canMove {
	// o.nextPos.Y += 1
	// o.canMove = false
	// }
}

func (o *Player) Attack() {
	if o.shootDelay == 0 && o.form == FormDemon {
		o.shootDelay = 15
		proj := Projectile{}
		proj.InitProjectile(o.pos.X, o.pos.Y, o.game, o.dir, GOProjectile)
		o.game.AddGO(&proj)
	}
}

func (o *Player) isClearLeft() bool {
	return o.game.NoCollideAt(o.pos.X-PLAYER_WIDTH_OFFSET-1, o.pos.Y-1) &&
		o.game.NoCollideAt(o.pos.X-PLAYER_WIDTH_OFFSET-1, o.pos.Y) &&
		o.game.NoCollideAt(o.pos.X-PLAYER_WIDTH_OFFSET-1, o.pos.Y+1)
}

func (o *Player) isClearRight() bool {
	return o.game.NoCollideAt(o.pos.X+PLAYER_WIDTH_OFFSET+1, o.pos.Y-1) &&
		o.game.NoCollideAt(o.pos.X+PLAYER_WIDTH_OFFSET+1, o.pos.Y) &&
		o.game.NoCollideAt(o.pos.X+PLAYER_WIDTH_OFFSET+1, o.pos.Y+1)
}

func (o *Player) isClearUp() bool {
	return o.game.NoCollideAt(o.pos.X-1, o.pos.Y-PLAYER_HEIGHT_OFFSET-1) &&
		o.game.NoCollideAt(o.pos.X, o.pos.Y-PLAYER_HEIGHT_OFFSET-1) &&
		o.game.NoCollideAt(o.pos.X+1, o.pos.Y-PLAYER_HEIGHT_OFFSET-1)
}

func (o *Player) isClearDown() bool {
	return o.game.NoCollideAt(o.pos.X-1, o.pos.Y+PLAYER_HEIGHT_OFFSET+1) &&
		o.game.NoCollideAt(o.pos.X, o.pos.Y+PLAYER_HEIGHT_OFFSET+1) &&
		o.game.NoCollideAt(o.pos.X+1, o.pos.Y+PLAYER_HEIGHT_OFFSET+1)
}

func (o *Player) Explode() {
	if o.lives > 0 {
		o.lives--
	}
}

func (o *Player) Die() {
	o.lives = 0
}

func (o *Player) Update() {
	if o.IsDead() {
		return
	}
	if o.shootDelay > 0 {
		o.shootDelay--
	}
	if o.moveLeft {
		if o.isClearLeft() {
			o.pos.X -= 1
			o.moveCount++
			// move one unit furthe if flying
			if !o.grounded && o.moveCount == 1 {
				if o.isClearLeft() {
					o.pos.X -= 1
				}
			}
			if o.moveCount > 3 {
				o.moveCount = 0
				o.moveLeft = false
			}
		} else {
			o.moveCount = 0
			o.moveLeft = false
		}
	} else if o.moveRight {
		if o.isClearRight() {
			o.pos.X += 1
			o.moveCount++
			// move one unit furthe if flying
			if !o.grounded && o.moveCount == 1 {
				if o.isClearRight() {
					o.pos.X += 1
				}
			}
			if o.moveCount > 3 {
				o.moveCount = 0
				o.moveRight = false
			}
		} else {
			o.moveCount = 0
			o.moveRight = false
		}

	}
	if o.moveUp && o.jumpCount < o.maxJumps {
		if o.isClearUp() {
			o.pos.Y -= 1
			o.jumpCount++
			o.grounded = false
			o.jumpTime = 0
			o.isJumping = true
		}
		o.moveUp = false
	} else if o.isJumping {
		if o.jumpTime < JUMP_TIME {
			if o.isClearUp() {
				o.pos.Y -= 1
			}
			o.jumpTime++
		} else {
			o.jumpTime = 0
			o.gravDelay = o.gravity
			o.isJumping = false
		}
	} else if o.gravDelay == 0 {
		o.gravDelay = o.gravTime
		if o.isClearDown() {
			o.pos.Y += 1
			if o.gravTime > 1 {
				o.gravTime -= 1
			}
		} else {
			// reset time
			o.gravTime = o.gravity
			o.grounded = true
			o.jumpCount = 0
			o.isJumping = false
			o.jumpTime = 0
			o.moveUp = false
		}
	} else {
		o.gravDelay -= 1
	}
}
