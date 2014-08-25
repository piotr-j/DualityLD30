package main

import (
	"github.com/nsf/termbox-go"
)

type Projectile struct {
	pos         Position
	dir         Direction
	explode     bool
	exploded    bool
	game        *Game
	explodeTime int
	projType    GameObjectType
	movedDst    int
}

func (o *Projectile) SetForm(form Form) {

}

func (p *Projectile) Init(x, y int, game *Game) {
	p.InitProjectile(x, y, game, DirRight, GOProjectile)
}

func (p *Projectile) InitProjectile(x, y int, game *Game, dir Direction, projType GameObjectType) {
	p.pos.X = x
	p.pos.Y = y
	p.game = game
	p.dir = dir
	p.explode = false
	p.exploded = false
	p.explodeTime = 5
	p.projType = projType
}

func (p *Projectile) Type() GameObjectType {
	return p.projType
}

func (p *Projectile) IsExploded() bool {
	return p.exploded
}

func (p *Projectile) SetDirection(dir Direction) {
	p.dir = dir
}

func (p *Projectile) Contains(x, y int) bool {
	if p.exploded {
		return false
	}
	if p.pos.X == x && p.pos.Y == y {
		return true
	}
	return false
}

func (p *Projectile) Draw() {
	if p.explode && p.explodeTime > 0 {
		p.game.SetCell(p.pos.X+1, p.pos.Y, termbox.ColorYellow)
		p.game.SetCell(p.pos.X-1, p.pos.Y, termbox.ColorYellow)
		p.game.SetCell(p.pos.X, p.pos.Y+1, termbox.ColorYellow)
		p.game.SetCell(p.pos.X, p.pos.Y-1, termbox.ColorYellow)
	} else if !p.exploded {
		diff := 0
		switch p.dir {
		case DirLeft:
			diff = 1
		case DirRight:
			diff = -1
		default:
			// do nothing
		}
		p.game.SetCell(p.pos.X+diff, p.pos.Y, termbox.ColorYellow)
		p.game.SetCell(p.pos.X, p.pos.Y, termbox.ColorRed)
	}
}

func (p *Projectile) Explode() {
	p.explode = true
}

func (p *Projectile) Update() {
	if p.explode {
		p.explodeTime--
		if p.explodeTime == 0 {
			p.exploded = true
			p.explode = false
		}
	} else if !p.exploded {
		diff := 0
		switch p.dir {
		case DirLeft:
			diff = -1
		case DirRight:
			diff = 1
		default:
			// do nothing
		}
		if g := p.game.GameObjectAt(p.pos.X+diff, p.pos.Y); g != nil {
			gameObject := *g
			switch gameObject.Type() {
			case GOCeiling:
				//p.explode = true
			case GOFloor:
				//p.explode = true
			case GOWall:
				p.explode = true
			case GOPlayer:
				if p.projType == GOProjectileEnemy {
					p.explode = true
					gameObject.Explode()
				}
			case GOEnemy:
				if p.projType == GOProjectile {
					p.explode = true
					gameObject.Explode()
				}
			case GOGateClosed:
				p.explode = true
				if p.projType == GOProjectile {
					gameObject.Explode()
				}
			case GOProjectileEnemy:
				p.explode = true
				gameObject.Explode()
			default:
				// do nothing
			}
		}
		if !p.explode {
			if !p.game.NoCollideAt(p.pos.X+diff, p.pos.Y) {
				p.explode = true
			} else if p.projType == GOProjectile && p.movedDst > 42 {
				p.explode = true
			} else {
				if p.projType == GOProjectile {
					p.movedDst++
				}
				p.pos.X += diff
			}
		}
	}
}
