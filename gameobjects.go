package main

type Position struct {
	X, Y int
}

type GameObjectType int
type Form int

const (
	GOWall GameObjectType = iota // 0...
	GOCeiling
	GOFloor
	GOPlayer
	GOEnemy
	GOEnemyExploded
	GOGateOpen
	GOGateClosed
	GOPortalOpen
	GOPortalClosed
	GOProjectile
	GOProjectileEnemy
	GOSpawn
	GOSpawnInitial
	GOSpawnFinish
	GOTrap
)
const (
	FormAngel = iota
	FormDemon
)

type GameObject interface {
	Init(x, y int, game *Game)
	Draw()
	Update()
	Contains(x, y int) bool
	Type() GameObjectType
	Explode()
	SetForm(form Form)
}
