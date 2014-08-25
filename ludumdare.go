package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"os/signal"
	"time"
)

const SLEEP_TIME = 16 * time.Millisecond // 60ish fps

func print(a ...interface{}) {
	fmt.Println(a...)
}

var isRunning = true

func mainLoop(eventQueue chan termbox.Event) {
	for isRunning {
		update(eventQueue)
		draw()
		// wait for next update
		time.Sleep(SLEEP_TIME)
	}
}

func update(eventQueue chan termbox.Event) {
	select {
	case event := <-eventQueue:
		switch event.Type {
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyEsc:
				isRunning = false
			case termbox.KeyArrowLeft:
				game.Left()
			case termbox.KeyArrowRight:
				game.Right()
			case termbox.KeyArrowUp:
				game.Up()
			case termbox.KeyArrowDown:
				game.Down()
			case termbox.KeySpace:
				game.Attack()
			default: // ignore other
			}
			switch event.Ch {
			case 'a':
				game.Left()
			case 'd':
				game.Right()
			case 'w':
				game.Up()
			case 's':
				game.Down()
			default: // ignore other
			}
		case termbox.EventResize:
			print("resize")
			game.Resize(event.Width, event.Height)
		case termbox.EventError:
			panic(event.Err)
		}
	default:
		// do nothing
	}
	game.Update()
}

func draw() {
	// draw state
	//print("draw...")
	// clear the screen
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	// reset size as event is not getting called for some reason
	game.Resize(termbox.Size())
	game.Draw()
	// draw to screen
	termbox.Flush()
}

var game *Game

func initGame(eventQueue chan termbox.Event) {
	game = new(Game)
	game.Init("level_1.png")
	game.Resize(termbox.Size())
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	go func() {
		for {
			select {
			case signal := <-sigChan:
				isRunning = false
				print("Got", signal, ", quitting")
			}
		}
	}()
}

func main() {
	// init termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	// create channel for nonblocking keyboard event handling
	eventQueue := make(chan termbox.Event)
	initGame(eventQueue)
	mainLoop(eventQueue)
	print("Thanks for playing! Check out my other stuff at piotrjastrzebski.io!")
}
