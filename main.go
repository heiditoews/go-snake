package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

// go run main.go game.go snake.go

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	snakeParts := []SnakeSeg{
		{
			X: 5,
			Y: 10,
		},
		{
			X: 6,
			Y: 10,
		},
		{
			X: 7,
			Y: 10,
		},
	}

	sn := SnakeBody {
		Segments: snakeParts,
		direction: 1,
	}

	game := Game{
		Screen: screen,
		Speed: 2,
		Snakes: sn,
	}

	go game.run()

	for {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				game.Snakes.changeDirection(0)
			} else if event.Key() == tcell.KeyRight {
				game.Snakes.changeDirection(1)
			} else if event.Key() == tcell.KeyDown {
				game.Snakes.changeDirection(2)
			} else if event.Key() == tcell.KeyLeft {
				game.Snakes.changeDirection(3)
			}
		}
    }

	
}