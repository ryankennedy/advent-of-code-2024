package parse

import "strconv"

func ParseInts(strs []string) ([]int, error) {
	ints := make([]int, 0, len(strs))
	for _, str := range strs {
		parsed, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, err
		}
		ints = append(ints, parsed)
	}
	return ints, nil
}
