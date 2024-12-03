package main

import (
	"log"

	"github.com/ryankennedy/advent-of-code-2024/day01"
)

func main() {
	if result, err := day01.Puzzle01(); err != nil {
		log.Fatalf("day01 puzzle01 failed: %s", err)
	} else {
		log.Println("day01 puzzle01 =", result)
	}

	if result, err := day01.Puzzle02(); err != nil {
		log.Fatalf("day01 puzzle02 failed: %s", err)
	} else {
		log.Println("day01 puzzle02 =", result)
	}
}
