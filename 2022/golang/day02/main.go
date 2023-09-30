package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	rock = iota + 1
	paper
	scissors
)
const (
	needsToLose = iota + 1
	needsToDraw
	needsToWin
)
const (
	loss = 0
	draw = 3
	win  = 6
)

var wins_over = []int{-1, paper, scissors, rock}
var loses_against = []int{-1, scissors, rock, paper}

func main() {
	fmt.Println("My Score 1:", partOne("input"))
	fmt.Println("My Score 2:", partTwo("input"))
}

func partOne(input string) int {
	s, cleanup := scan(input)
	defer cleanup()
	total := 0
	for s.Scan() {
		line := s.Text()
		them, me := parse(line)
		total += score(them, me)
	}
	return total
}

func partTwo(input string) int {
	s, cleanup := scan(input)
	defer cleanup()
	total := 0
	for s.Scan() {
		line := s.Text()
		them, me := parse(line)
		me = chooseMine(them, me)
		total += score(them, me)
	}
	return total
}

func chooseMine(them, me int) int {
	switch me {
	case needsToDraw:
		return them
	case needsToWin:
		return wins_over[them]
	default:
		return loses_against[them]
	}
}

func score(them, me int) int {
	if me == them {
		return me + draw
	} else if me == wins_over[them] {
		return me + win
	} else {
		return me + loss
	}
}

func scan(path string) (*bufio.Scanner, func() error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	return scanner, file.Close
}

func parse(line string) (int, int) {
	round := strings.Split(line, " ")
	them := int(round[0][0] - 'A' + 1)
	me := int(round[1][0] - 'X' + 1)
	return them, me
}
