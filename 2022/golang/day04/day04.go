package day04

import (
	"strconv"
	"strings"
)

func PartOne(lines []string) int {
	total := 0
	for _, line := range lines {
		first, second := parse(line)
		if isContained(first, second) || isContained(second, first) {
			total++
		}
	}
	return total
}

func PartTwo(lines []string) int {
	total := 0
	for _, line := range lines {
		first, second := parse(line)
		if overlaps(first, second) || overlaps(second, first) {
			total++
		}
	}
	return total
}

type Range struct {
	Start int
	End   int
}

func parse(line string) (Range, Range) {
	assignments := strings.Split(line, ",")
	first := createRange(assignments[0])
	second := createRange(assignments[1])
	return first, second
}

func createRange(assignment string) Range {
	r := strings.Split(assignment, "-")
	start, err := strconv.Atoi(r[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(r[1])
	if err != nil {
		panic(err)
	}
	return Range{start, end}
}

func isContained(r1, r2 Range) bool {
	return r1.Start <= r2.Start && r2.End <= r1.End
}

func overlaps(r1, r2 Range) bool {
	return r1.Start <= r2.Start && r2.Start <= r1.End ||
		r1.Start <= r2.End && r2.End <= r1.End
}
