package block

import "fmt"

// Block represents a single block in the game.
type Block struct {
	Shape [][]int
	X     int
	Y     int
}

// NewBlock returns a new block with the given shape and position.
func NewBlock(shape [][]int, x, y int) *Block {
	return &Block{
		Shape: shape,
		X:     x,
		Y:     y,
	}
}

// Move moves the block in the given direction.
func (b *Block) Move(dx, dy int) {
	b.X += dx
	b.Y += dy
}

// Rotate rotates the block 90 degrees clockwise.
func (b *Block) Rotate() {
	n := len(b.Shape)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := b.Shape[i][j]
			b.Shape[i][j] = b.Shape[n-j-1][i]
			b.Shape[n-j-1][i] = b.Shape[n-i-1][n-j-1]
			b.Shape[n-i-1][n-j-1] = b.Shape[j][n-i-1]
			b.Shape[j][n-i-1] = temp
		}
	}
}

// Display displays the block on the screen.
func (b *Block) Display() {
	for i := range b.Shape {
		for j := range b.Shape[i] {
			if b.Shape[i][j] == 1 {
				fmt.Print("[]")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
