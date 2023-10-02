package main

import "strings"

var priority string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func PartOne(lines []string) int {
	total := 0
	for _, line := range lines {
		item := GetWrongItem(line)
		priority := strings.IndexRune(priority, item) + 1
		total += priority
	}
	return total
}

func GetWrongItem(line string) rune {
	set := make(map[rune]struct{})
	midPoint := len(line) / 2
	first_compartment, second_compartment := line[:midPoint], line[midPoint:]
	for _, item := range first_compartment {
		set[item] = struct{}{}
	}
	for _, item := range second_compartment {
		_, exists := set[item]
		if exists {
			return item
		}
	}
	panic("There should be a misplaced item")
}
