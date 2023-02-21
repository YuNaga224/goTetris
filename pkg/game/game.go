package game

import (
	"fmt"
	"time"

	"github.com/YuNaga224/goTetris/pkg/block"
	"github.com/YuNaga224/goTetris/pkg/screen"
)

// Game represents a game of Tetris.
type Game struct {
	screen *screen.Screen
	block  *block.Block
	score  int
	speed  int
}

// NewGame returns a new game of Tetris with the given screen and block.
func NewGame(s *screen.Screen, b *block.Block) *Game {
	return &Game{
		screen: s,
		block:  b,
		score:  0,
		speed:  500,
	}
}

// Start starts the game of Tetris.
func (g *Game) Start() {
	for {
		g.screen.Display()
		g.block.Move(0, 1)
		g.screen.DrawBlock(g.block)
		g.screen.ClearLines()
		time.Sleep(time.Duration(g.speed) * time.Millisecond)
	}
}

// MoveBlock moves the block in the given direction.
func (g *Game) MoveBlock(dx, dy int) {
	g.block.Move(dx, dy)
}

// RotateBlock rotates the block 90 degrees clockwise.
func (g *Game) RotateBlock() {
	g.block.Rotate()
}

// DisplayScore displays the current score of the game.
func (g *Game) DisplayScore() {
	fmt.Println("Score:", g.score)
}
