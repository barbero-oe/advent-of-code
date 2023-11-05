package day09

import (
	"math"
	"strconv"
	"strings"
)

func PartOne(lines []string) int {
	head := [2]int{0, 0}
	tail := [2]int{0, 0}
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
			x, y := head[0]-tail[0], head[1]-tail[1]
			switch {
			case x != 0 && y != 0:
				if max(math.Abs(float64(x)), math.Abs(float64(y))) > 1 {
					tail[0] += moveBoth(x)
					tail[1] += moveBoth(y)
				}
			case x != 0:
				tail[0] += move(x)
			case y != 0:
				tail[1] += move(y)
			}
			visited[tail] = struct{}{}
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
	return 0
}
