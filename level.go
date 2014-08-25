package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const COLLIDE = 1
const NO_COLLIDE = 0
const LEVEL_SCALE_X = 4
const LEVEL_SCALE_Y = 3

var (
	colorEmpty              = color.RGBA{0, 0, 0, 0}         // transparent
	colorEmpty2             = color.RGBA{255, 255, 255, 255} // transparent
	colorFloor              = color.RGBA{0, 0, 0, 255}       // black
	colorCeiling            = color.RGBA{196, 196, 196, 255} // light gray
	colorWall               = color.RGBA{128, 128, 128, 255} // gray
	colorEnemyLeft          = color.RGBA{255, 0, 0, 255}     // red
	colorEnemyRight         = color.RGBA{128, 0, 0, 255}     // red
	colorGate               = color.RGBA{0, 255, 0, 255}     // green
	colorPortalClosed       = color.RGBA{0, 0, 255, 255}     // blue
	colorPortalOpen         = color.RGBA{0, 255, 255, 255}   // cyan
	colorPlayerSpawnInitial = color.RGBA{255, 0, 255, 255}   // magenta
	colorPlayerSpawn        = color.RGBA{196, 0, 255, 255}   // magenta
	colorPlayerSpawnFinish  = color.RGBA{128, 0, 255, 255}   // magenta
	colorTrap               = color.RGBA{255, 255, 0, 255}   // yellow
)

type Level struct {
	path          string
	player        *Player
	width, height int
	game          *Game
	collisions    [][]uint8
}

func (lvl *Level) LoadLevel(path string) {
	lvl.path = path
	imageData, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer imageData.Close()

	image, err := png.Decode(imageData)
	if err != nil {
		panic(err)
	}

	parseImageData(image, lvl)
}

func (lvl *Level) SetGame(game *Game) {
	lvl.game = game
}

func (lvl *Level) Width() int {
	return lvl.width
}

func (lvl *Level) Height() int {
	return lvl.height
}

func (lvl *Level) Player() *Player {
	return lvl.player
}

func (lvl *Level) addInitialPlayerSpawn(x, y int) {
	//print("player spawn at:", x, y)
	player := Player{}
	player.Init(x, y, lvl.game)
	lvl.player = &player
	lvl.game.AddGO(lvl.player)
	w := Spawn{}
	w.InitSpawn(x, y, lvl.game, GOSpawnInitial)
	lvl.game.AddGO(&w)
	game.SetSpawnPoint(&w)
	//addCollisionAt(x, y, lvl)
}

func (lvl *Level) addTrap(x, y int) {
	w := Trap{}
	w.Init(x, y, lvl.game)
	lvl.game.AddGO(&w)
	//addCollisionAt(x, y, lvl)
}

func (lvl *Level) addPlayerSpawn(x, y int) {
	w := Spawn{}
	w.Init(x, y, lvl.game)
	lvl.game.AddGO(&w)
	//addCollisionAt(x, y, lvl)
}

func (lvl *Level) addFinishPlayerSpawn(x, y int) {
	w := Spawn{}
	w.InitSpawn(x, y, lvl.game, GOSpawnFinish)
	lvl.game.AddGO(&w)
	//addCollisionAt(x, y, lvl)
}

func (lvl *Level) addWall(x, y int) {
	//print("wall at:", x, y)
	w := Wall{}
	w.Init(x, y, lvl.game)
	lvl.game.AddGO(&w)
	addCollisionAt(x, y, lvl)
}

func (lvl *Level) addCeiling(x, y int) {
	//print("ceiling at:", x, y)
	w := Ceiling{}
	w.Init(x, y, lvl.game)
	lvl.game.AddGO(&w)
	addCollisionAt(x, y, lvl)
}

func (lvl *Level) addFloor(x, y int) {
	//print("Floor at:", x, y)
	w := Floor{}
	w.Init(x, y, lvl.game)
	lvl.game.AddGO(&w)
	addCollisionAt(x, y, lvl)
}

func (lvl *Level) addEnemy(x, y int, left bool) {
	//print("enemy at:", x, y)
	w := Enemy{}
	if left {
		w.InitEnemy(x, y, lvl.game, EnemyLeft)
	} else {
		w.InitEnemy(x, y, lvl.game, EnemyRight)
	}
	lvl.game.AddGO(&w)
}

func (lvl *Level) addGate(x, y int) {
	//print("gate at:", x, y)
	w := Gate{}
	w.Init(x, y, lvl.game)
	lvl.game.AddGO(&w)
}

func (lvl *Level) addPortal(x, y int, open bool) {
	//print(open, "portal spawn at:", x, y)
	w := Portal{}
	w.Init(x, y, lvl.game)
	lvl.game.AddGO(&w)
	//addCollisionAt(x, y, lvl)
}

func (lvl *Level) AddCollideAt(x, y int) {
	lvl.collisions[x][y] = COLLIDE
}

func (lvl *Level) RemoveCollideAt(x, y int) {
	lvl.collisions[x][y] = NO_COLLIDE
}

func addCollisionAt(x, y int, lvl *Level) {
	for xx := 0; xx < LEVEL_SCALE_X; xx++ {
		for yy := 0; yy < LEVEL_SCALE_Y; yy++ {
			lvl.collisions[x+xx][y+yy] = COLLIDE
		}
	}
}

func (lvl *Level) IsClearAt(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x > lvl.width-1 || y > lvl.height-1 {
		return false
	}
	return lvl.collisions[x][y] != COLLIDE
}

func parseImageData(image image.Image, level *Level) {
	bounds := image.Bounds()
	level.width = int(bounds.Max.X) * LEVEL_SCALE_X
	level.height = int(bounds.Max.Y) * LEVEL_SCALE_Y
	level.collisions = make([][]uint8, level.width*LEVEL_SCALE_X)
	for x := 0; x < level.width*LEVEL_SCALE_X; x += LEVEL_SCALE_X {
		for scale := 0; scale < LEVEL_SCALE_X; scale++ {
			level.collisions[x+scale] = make([]uint8, level.height*LEVEL_SCALE_Y)
		}
		for y := 0; y < level.height*LEVEL_SCALE_Y; y += LEVEL_SCALE_Y {
			switch image.At(x/LEVEL_SCALE_X, y/LEVEL_SCALE_Y) {
			case colorPlayerSpawnInitial:
				level.addInitialPlayerSpawn(x, y)
			case colorPlayerSpawn:
				level.addPlayerSpawn(x, y)
			case colorPlayerSpawnFinish:
				level.addFinishPlayerSpawn(x, y)
			case colorFloor:
				level.addFloor(x, y)
			case colorCeiling:
				level.addCeiling(x, y)
			case colorWall:
				level.addWall(x, y)
			case colorEnemyLeft:
				level.addEnemy(x, y, true)
			case colorEnemyRight:
				level.addEnemy(x, y, false)
			case colorGate:
				level.addGate(x, y)
			case colorPortalOpen:
				level.addPortal(x, y, true)
			case colorPortalClosed:
				level.addPortal(x, y, false)
			case colorTrap:
				level.addTrap(x, y)
			case colorEmpty, colorEmpty2:
				// do nothing
			default:
				print("Unknown color:", image.At(x/LEVEL_SCALE_X, y/LEVEL_SCALE_Y),
					"at", x/LEVEL_SCALE_X, y/LEVEL_SCALE_Y)
			}
		}
	}
}
