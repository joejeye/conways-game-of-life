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

	fmt.Println("\033[2J") // Clear the terminal

	bLoop := cells.NewBoard(numRows, numCols)

	// Use fast and slow pointers to detect a loop
	for {
		// The fast pointer moves two steps at a time
		c.Next()
		c.Next()

		vis := b.TerminalVisualize()
		moveUp := fmt.Sprintf("\x1b[%dA", numRows+1)
		fmt.Print(moveUp, vis) // Move the cursor up to the top of the terminal and print the visualization

		// Mark the first convergence point
		if b.Equal(c) {
			bLoop = b.Clone()
		}

		// The slow pointer moves one step at a time
		b.Next()

		// The marked convergence point is reached again, implying
		// that the 2D world has gone through at least one cycle
		if b.Equal(bLoop) {
			break
		}

		time.Sleep(125 * time.Millisecond)
	}

	fmt.Println("The world entered a loop!")
}
