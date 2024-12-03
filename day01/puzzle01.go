package day01

import "github.com/ryankennedy/advent-of-code-2024/math"

func Puzzle01() (int, error) {
	left, right, err := sortedList()
	if err != nil {
		return -1, err
	}

	diffSum := 0
	for i := 0; i < len(left); i++ {
		diffSum += math.IntAbs(left[i] - right[i])
	}

	return diffSum, nil
}
