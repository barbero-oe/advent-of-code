package day09

import (
	"math"
	"strconv"
	"strings"
)

func PartOne(lines []string) int {
	return Simulate(lines, 2)
}

func Simulate(lines []string, ropeLenght int) int {
	rope := make([][2]int, ropeLenght)
	for i := 0; i < ropeLenght; i++ {
		rope[i] = [2]int{0, 0}
	}
	head := &rope[0]
	tail := &rope[len(rope)-1]
	visited := make(map[[2]int]struct{})
	for _, line := range lines {
		direction, count := parse(line)
		for i := 0; i < count; i++ {
			switch direction {
			case "R":
				head[0]++
			case "L":
				head[0]--
			case "U":
				head[1]++
			case "D":
				head[1]--
			}
			for i := 1; i < len(rope); i++ {
				previous := &rope[i-1]
				current := &rope[i]
				x, y := previous[0]-current[0], previous[1]-current[1]
				switch {
				case x != 0 && y != 0:
					if max(math.Abs(float64(x)), math.Abs(float64(y))) > 1 {
						current[0] += moveBoth(x)
						current[1] += moveBoth(y)
					}
				case x != 0:
					current[0] += move(x)
				case y != 0:
					current[1] += move(y)
				}
			}
			visited[*tail] = struct{}{}
		}

	}
	return len(visited)
}

func move(x int) int {
	if math.Abs(float64(x)) <= 1 {
		return 0
	}
	if x > 0 {
		return 1
	}
	return -1
}

func moveBoth(x int) int {
	if x > 0 {
		return 1
	}
	return -1
}

func parse(line string) (string, int) {
	parts := strings.Split(line, " ")
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return parts[0], n
}

func PartTwo(lines []string) int {
	return Simulate(lines, 10)
}
