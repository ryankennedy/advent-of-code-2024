package day01

func Puzzle02() (int, error) {
	left, right, err := parsedList()
	if err != nil {
		return -1, err
	}

	similarity := 0
	for i := 0; i < len(left); i++ {
		similarity += left[i] * occurrences(left[i], right)
	}

	return similarity, nil
}

func occurrences(value int, values []int) int {
	count := 0
	for _, v := range values {
		if value == v {
			count++
		}
	}
	return count
}
