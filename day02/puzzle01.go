package day02

import (
	"slices"

	"github.com/ryankennedy/advent-of-code-2024/math"
)

func Puzzle01() (int, error) {
	reports, err := reports()
	if err != nil {
		return -1, err
	}

	safeCount := 0
	for _, report := range reports {
		reverseLevels := slices.Clone(report.Levels)
		slices.Reverse(reverseLevels)
		if !slices.IsSorted(report.Levels) && !slices.IsSorted(reverseLevels) {
			// Levels must either be all ascending or all descending to be safe
			continue
		}

		isSafe := true
		for i := 0; i < len(report.Levels)-1; i++ {
			first := report.Levels[i]
			second := report.Levels[i+1]
			diff := math.IntAbs(second - first)
			if diff < 1 || diff > 3 {
				// Step up/down must be no less than 1 and no greater than 3
				isSafe = false
				break
			}
		}

		if isSafe {
			safeCount++
		}
	}

	return safeCount, nil
}
