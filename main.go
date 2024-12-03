package main

import (
	"log"

	"github.com/ryankennedy/advent-of-code-2024/day01"
	"github.com/ryankennedy/advent-of-code-2024/day02"
)

type PuzzleFunc func() (int, error)

type Puzzle struct {
	Day    int
	Puzzle int
	Func   PuzzleFunc
}

func main() {
	puzzles := []Puzzle{
		Puzzle{1, 1, day01.Puzzle01},
		Puzzle{1, 2, day01.Puzzle02},
		Puzzle{2, 1, day02.Puzzle01},
	}

	for _, puzzle := range puzzles {
		result, err := puzzle.Func()
		if err != nil {
			log.Fatalf("Day %d Puzzle %d: error %s", puzzle.Day, puzzle.Puzzle, err)
		}
		log.Printf("Day %d Puzzle %d: result=%d\n", puzzle.Day, puzzle.Puzzle, result)
	}
}
