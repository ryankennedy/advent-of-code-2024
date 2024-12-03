package day02

import (
	"bufio"
	"bytes"
	_ "embed"
	"strings"

	"github.com/ryankennedy/advent-of-code-2024/parse"
)

//go:embed input.txt
var input []byte

type Report struct {
	Levels []int
}

func reports() ([]Report, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	reports := make([]Report, 0, 1)
	for scanner.Scan() {
		levels, err := parse.ParseInts(strings.Split(scanner.Text(), " "))
		if err != nil {
			return []Report{}, err
		}
		reports = append(reports, Report{levels})
	}

	return reports, nil
}
