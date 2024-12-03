package day01

import (
	"bufio"
	"bytes"
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func parsedList() ([]int, []int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	left := make([]int, 0, 1)
	right := make([]int, 0, 1)
	for scanner.Scan() {
		ids, err := parseInts(strings.SplitN(scanner.Text(), "   ", 2))
		if err != nil {
			return []int{}, []int{}, err
		}
		left = append(left, ids[0])
		right = append(right, ids[1])
	}

	return left, right, nil
}

func parseInts(strs []string) ([]int, error) {
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

func sortedList() ([]int, []int, error) {
	left, right, err := parsedList()
	if err != nil {
		return []int{}, []int{}, err
	}
	sort.Ints(left)
	sort.Ints(right)
	return left, right, nil
}
