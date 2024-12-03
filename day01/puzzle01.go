package day01

func Puzzle01() (int, error) {
	left, right, err := sortedList()
	if err != nil {
		return -1, err
	}

	diffSum := 0
	for i := 0; i < len(left); i++ {
		diffSum += intAbs(left[i] - right[i])
	}

	return diffSum, nil
}

func intAbs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
