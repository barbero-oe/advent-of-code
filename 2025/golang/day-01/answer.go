package day01

import (
	"strconv"
)

func ProblemOne(lines []string) int {
	position := 50
	times := 0
	for _, line := range lines {
		sign := 1
		if line[0] == 'L' {
			sign = -1
		}
		num, _ := strconv.Atoi(line[1:])
		position += sign * num
		position %= 100
		if position < 0 {
			position += 100
		}
		if position == 0 {
			times++
		}
	}
	return times
}
