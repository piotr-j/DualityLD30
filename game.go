package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

type Cell struct {
	x, y   int
	color  termbox.Attribute
	symbol rune
}

type Game struct {
	lvl           *Level
	player        *Player
	viewPort      ViewPort
	displayBuffer [][]Cell
	gameObjects   []GameObject
	projectiles   []Projectile
	spawn         *Spawn
	form          Form
	moveShown     bool
	shootShown    bool
	gameWon       bool
}

type ViewPort struct {
	width, height int
}

func (g *Game) Init(level string) {
	g.gameObjects = make([]GameObject, 128)
	g.projectiles = make([]Projectile, 128)
	g.lvl = new(Level)
	g.lvl.SetGame(g)
	g.lvl.LoadLevel(level)
	g.player = g.lvl.Player()
	g.viewPort = ViewPort{80, 43}
	g.displayBuffer = make([][]Cell, g.lvl.Width())
	for x := 0; x < g.lvl.Width(); x++ {
		g.displayBuffer[x] = make([]Cell, g.lvl.Height())
		for y := 0; y < g.lvl.Height(); y++ {
			g.displayBuffer[x][y].symbol = ' '
		}
	}
}

func (g *Game) AddGO(gameObject GameObject) {
	g.gameObjects = append(g.gameObjects, gameObject)
}

func (g *Game) Draw() {
	for _, gameObject := range g.gameObjects {
		if gameObject != nil {
			gameObject.Draw()
		}
	}
	if g.form == FormAngel && !g.moveShown {
		g.DrawText("WSAD to move", 8, 8, termbox.ColorDefault)
	} else if g.form == FormDemon && !g.shootShown {
		g.DrawText("SPACE to shoot -->", 16, 8, termbox.ColorDefault)
	}
	if g.gameWon {
		g.DrawText("YOU WON!", 12, 40, termbox.ColorDefault)
		g.DrawText("Press ESC to quit", 12, 42, termbox.ColorDefault)
	}
	halfWidth := g.viewPort.width / 2
	halfHeight := g.viewPort.height / 2
	//print(halfWidth, halfHeight)
	offsetX := max(0, min(g.player.pos.X-halfWidth-1, g.lvl.Width()-g.viewPort.width))
	offsetY := max(0, min(g.player.pos.Y-halfHeight-1, g.lvl.Height()-g.viewPort.height))
	for x := 0; x < g.lvl.Width(); x++ {
		for y := 0; y < g.lvl.Height(); y++ {
			termbox.SetCell(
				g.displayBuffer[x][y].x-offsetX,
				g.displayBuffer[x][y].y-offsetY,
				g.displayBuffer[x][y].symbol,
				termbox.ColorDefault,
				g.displayBuffer[x][y].color)
		}
	}

	for x := 0; x < g.lvl.Width(); x++ {
		for y := 0; y < g.lvl.Height(); y++ {
			g.displayBuffer[x][y].x = 0
			g.displayBuffer[x][y].y = 0
			g.displayBuffer[x][y].color = termbox.ColorDefault
			g.displayBuffer[x][y].symbol = ' '
		}
	}

	// draw ui
	if g.form == FormDemon {
		text := fmt.Sprintf("HP: %d", g.player.GetLives())
		g.DrawGuiText(text, 1, 1, termbox.ColorBlack)
	}
}

func (g *Game) DrawText(text string, x, y int, bg termbox.Attribute) {
	for i := 0; i < len(text); i++ {
		g.SetCellSymbol(
			x+i, y,
			bg,
			rune(text[i]))
	}
}

func (g *Game) DrawGuiText(text string, x, y int, bg termbox.Attribute) {
	for i := 0; i < len(text); i++ {
		termbox.SetCell(
			x+i, y,
			rune(text[i]),
			termbox.ColorDefault,
			bg)
	}
}

func (g *Game) Update() {
	for _, gameObject := range g.gameObjects {
		if gameObject != nil {
			gameObject.Update()
		}
	}
	if g.player.IsDead() {
		g.player.ReSpawn(g.spawn.Pos().X, g.spawn.Pos().Y)
		g.SetForm(FormAngel)
	}
}

func (g *Game) Win() {
	print("You won!")
	g.gameWon = true
}

func (g *Game) KillPlayer() {
	g.player.Die()
}

func (g *Game) SetForm(form Form) {
	g.form = form
	if g.form == FormDemon {
		g.moveShown = true
	}
	for _, gameObject := range g.gameObjects {
		if gameObject != nil {
			gameObject.SetForm(form)
		}
	}
}

func (g *Game) Resize(width, height int) {
	g.viewPort.width = width
	g.viewPort.height = height
}

func (g *Game) SetCell(x, y int, color termbox.Attribute) {
	g.displayBuffer[x][y].color = color
	g.displayBuffer[x][y].x = x
	g.displayBuffer[x][y].y = y
	g.displayBuffer[x][y].symbol = ' '
}

func (g *Game) SetCellSymbol(x, y int, color termbox.Attribute, symbol rune) {
	g.displayBuffer[x][y].color = color
	g.displayBuffer[x][y].x = x
	g.displayBuffer[x][y].y = y
	g.displayBuffer[x][y].symbol = symbol
}

func (g *Game) CollideAt(x, y int) bool {
	return !g.lvl.IsClearAt(x, y)
}

func (g *Game) AddCollideAt(x, y int) {
	g.lvl.AddCollideAt(x, y)
}

func (g *Game) RemoveCollideAt(x, y int) {
	g.lvl.RemoveCollideAt(x, y)
}

func (g *Game) NoCollideAt(x, y int) bool {
	return g.lvl.IsClearAt(x, y)
}

func (g *Game) GameObjectAt(x, y int) *GameObject {
	for _, gameObject := range g.gameObjects {
		if gameObject != nil {
			if gameObject.Contains(x, y) {
				return &gameObject
			}
		}
	}
	return nil
}

func (g *Game) GameObjectAtWithType(x, y int, goType GameObjectType) *GameObject {
	for _, gameObject := range g.gameObjects {
		if gameObject != nil {
			if gameObject.Contains(x, y) && gameObject.Type() == goType {
				return &gameObject
			}
		}
	}
	return nil
}

func (g *Game) GameObjectAtWithoutType(x, y int, goType GameObjectType) *GameObject {
	for _, gameObject := range g.gameObjects {
		if gameObject != nil {
			if gameObject.Contains(x, y) && gameObject.Type() != goType {
				return &gameObject
			}
		}
	}
	return nil
}

func (g *Game) SetSpawnPoint(spawn *Spawn) {
	if g.spawn != nil {
		g.spawn.SetActive(false)
	}
	spawn.SetActive(true)
	g.spawn = spawn
}

func (g *Game) Up() {
	g.player.Up()
}

func (g *Game) Down() {
	g.player.Down()
}

func (g *Game) Left() {
	g.player.Left()
}

func (g *Game) Right() {
	g.player.Right()
}

func (g *Game) Attack() {
	if g.form == FormDemon {
		g.shootShown = true
	}
	g.player.Attack()
}
