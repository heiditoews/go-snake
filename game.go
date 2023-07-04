package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen tcell.Screen
	Speed int
	Snakes SnakeBody
	Food SnakeSeg
}

func (g *Game) run() {
 
    defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
    g.Screen.SetStyle(defStyle)
	width, height := g.Screen.Size()

	g.Snakes.reset()
	g.UpdateFood(width, height)

    for {
        g.Screen.Clear()

		foundFood := false

		if g.checkCollision(g.Snakes.Segments, g.Food) {
			foundFood = true
			g.UpdateFood(width, height)
		}
		if g.checkCollision(g.Snakes.Segments[:len(g.Snakes.Segments)-1], g.Snakes.Segments[len(g.Snakes.Segments)-1]) {
			break
		}

		g.Snakes.update(width, height, foundFood)
		g.showSnake(g.Screen, g.Snakes.Segments, snakeStyle)
		g.showFood(g.Screen, g.Food, defStyle)

        g.Screen.Show()
        time.Sleep(40 * time.Millisecond)
    }

}

func (g *Game) showSnake(screen tcell.Screen, segs []SnakeSeg, style tcell.Style) {
	// g.Screen.SetContent(g.Snakes.X, g.Snakes.Y, g.Snakes.display(), nil, defStyle)
	for _, s := range segs {
		g.Screen.SetContent(s.X, s.Y, s.display(), nil, style)
	}
}

func (g *Game) showFood(screen tcell.Screen, food SnakeSeg, style tcell.Style) {
	g.Screen.SetContent(food.X, food.Y, '\u25CF', nil, style)
}

func (g *Game) checkCollision(list []SnakeSeg, other SnakeSeg) bool {
	for _, seg := range list {
		if seg.X == other.X && seg.Y == other.Y {
			return true
		}
	}
	return false
}

func (g *Game) UpdateFood(width int, height int) {
	g.Food.X = rand.Intn(width)
	g.Food.Y = rand.Intn(height)
}