package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	highestCalories("input")
	topCalories("input")
}

func topCalories(input string) {
	f := readLines(input)
	top := make([]int, 0, 3)
	currentCalories := 0
	for f.Scan() {
		line := f.Text()
		if line == "" {
			if l := len(top); l < cap(top) {
				top = append(top, currentCalories)
				slices.Sort(top)
				slices.Reverse(top)
			} else {
				for i, v := range top {
					if v < currentCalories {
						top = append(top[:i+1], top[i:]...)
						top[i] = currentCalories
						top = top[:3]
						break
					}
				}
			}
			currentCalories = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			currentCalories += calories
		}
	}
	for i, v := range top {
		if v < currentCalories {
			top = append(top[:i+1], top[i:]...)
			top[i] = currentCalories
			top = top[:3]
			break
		}
	}
	sum := 0
	for _, v := range top {
		sum += v
	}
	fmt.Println("Top Sum Calories:", sum)
}

func highestCalories(input string) {
	f := readLines(input)
	maxCalories := 0
	currentCalories := 0
	for f.Scan() {
		line := f.Text()
		if line == "" {
			if maxCalories < currentCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			currentCalories += calories
		}
	}
	if maxCalories < currentCalories {
		maxCalories = currentCalories
	}
	fmt.Println("Max Calories:", maxCalories)
}

func readLines(input string) *bufio.Scanner {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}
