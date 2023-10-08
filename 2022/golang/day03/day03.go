package day03

import (
	"strings"
)

var priority string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func PartOne(lines []string) int {
	total := 0
	for _, line := range lines {
		item := GetWrongItem(line)
		priority := GetPriority(item)
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

func GetPriority(item rune) int {
	return strings.IndexRune(priority, item) + 1
}

func PartTwo(lines []string) int {
	total := 0
	group := make([]string, 0, 3)
	for _, line := range lines {
		group = append(group, line)
		if len(group) < 3 {
			continue
		}
		badge := GetGroupBadge(group)
		total += GetPriority(badge)
		group = group[:0]
	}
	return total
}

func GetGroupBadge(group []string) rune {
	ruckstacks := make([]map[rune]int, 0, len(group))
	for _, ruckstack := range group {
		count := make(map[rune]int)
		for _, item := range ruckstack {
			count[item] += 1
		}
		ruckstacks = append(ruckstacks, count)
	}

	intersection := Fold(SumIntersect, ruckstacks, nil)
	for badge := range intersection {
		return badge
	}
	panic("No badge")
}

func SumIntersect(a, b map[rune]int) map[rune]int {
	if a == nil {
		return b
	}
	intersection := make(map[rune]int)
	for key, av := range a {
		bv, exists := b[key]
		if exists {
			intersection[key] = av + bv
		}
	}
	return intersection
}

func Fold[T any, R any](f func(R, T) R, sequence []T, first R) R {
	acc := first
	for _, item := range sequence {
		acc = f(acc, item)
	}
	return acc
}
