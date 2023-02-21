package screen

import (
	"fmt"

	"github.com/YuNaga224/goTetris/pkg/block"
)

// Screen represents the game screen.
type Screen struct {
	width  int
	height int
	grid   [][]int
}

// NewScreen returns a new game screen with the given width and height.
func NewScreen(width, height int) *Screen {
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	return &Screen{
		width:  width,
		height: height,
		grid:   grid,
	}
}

// DrawBlock draws the given block on the screen.
func (s *Screen) DrawBlock(b *block.Block) {
	for i := range b.Shape {
		for j := range b.Shape[i] {
			if b.Shape[i][j] == 1 {
				s.grid[b.Y+i][b.X+j] = 1
			}
		}
	}
}


// ClearLines clears any completed lines on the screen.
func (s *Screen) ClearLines() {
	var newGrid [][]int
	for i := s.height - 1; i >= 0; i-- {
		var lineComplete bool
		for j := 0; j < s.width; j++ {
			if s.grid[i][j] == 0 {
				lineComplete = false
				break
			}
			lineComplete = true
		}
		if !lineComplete {
			newGrid = append(newGrid, s.grid[i])
		}
	}

	for i := 0; i < len(newGrid)/2; i++ {
		j := len(newGrid) - i - 1
		newGrid[i], newGrid[j] = newGrid[j], newGrid[i]
	}

	s.grid = newGrid
}

// Display displays the current state of the screen on the command line.
func (s *Screen) Display() {
	for i := range s.grid {
		for j := range s.grid[i] {
			if s.grid[i][j] == 1 {
				fmt.Print("[]")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
