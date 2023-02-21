package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/YuNaga224/goTetris/pkg/block"
	"github.com/YuNaga224/goTetris/pkg/game"
	"github.com/YuNaga224/goTetris/pkg/screen"
)

func main() {
	s := screen.NewScreen(10, 20)
	b := block.NewBlock([][]int{{1, 1}, {1, 1}}, 3, 0)
	g := game.NewGame(s, b)

	go func() {
		for {
			g.Start()
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		text, _ := reader.ReadString('\n')
		switch text {
		case "left\n":
			g.MoveBlock(-1, 0)
		case "right\n":
			g.MoveBlock(1, 0)
		case "down\n":
			g.MoveBlock(0, 1)
		case "rotate\n":
			g.RotateBlock()
		case "quit\n":
			os.Exit(0)
		}
	}
}

