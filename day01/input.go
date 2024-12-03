package day01

import (
	"bufio"
	"bytes"
	_ "embed"
	"sort"
	"strings"

	"github.com/ryankennedy/advent-of-code-2024/parse"
)

//go:embed input.txt
var input []byte

func parsedList() ([]int, []int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	left := make([]int, 0, 1)
	right := make([]int, 0, 1)
	for scanner.Scan() {
		ids, err := parse.ParseInts(strings.SplitN(scanner.Text(), "   ", 2))
		if err != nil {
			return []int{}, []int{}, err
		}
		left = append(left, ids[0])
		right = append(right, ids[1])
	}

	return left, right, nil
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
