package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/tbouasli/game-of-life/game"
)

func main() {
	grid, err := game.FromLife106Format(bufio.NewScanner(os.Stdin))
	if err != nil {
		log.Fatalf("Error parsing Life 1.06 file: %v", err)
	}

	for range 10 {
		grid = grid.NextGeneration()
	}

	fmt.Print(game.ToLife106Format(grid))
}
