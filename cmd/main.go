package main

import (
	"fmt"
	"myplay/conwayslifegame/cells"
	"myplay/conwayslifegame/cli"
	"os"
	"time"
)

func main() {
	numRows, numCols, init, err := cli.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	b := cells.NewBoard(numRows, numCols)
	b.RandomInit(init)
	c := b.Clone()

	for {
		c.Next()
		c.Next()
		vis := b.TerminalVisualize()
		fmt.Print("\x0c", vis)
		if b.Equal(c) {
			break
		}
		b.Next()
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Loop detected!")
}
