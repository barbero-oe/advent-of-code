package day05

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne(lines []string) string {
	crates_section, instructions_section := splitInstructions(lines)
	crates := ParseCrate(crates_section)
	for _, line := range instructions_section {
		command := parseCommand(line)
		crates.Move(command)
	}
	return crates.TopCrates()
}

func PartTwo(lines []string) string {
	crates_section, instructions_section := splitInstructions(lines)
	crates := ParseCrate(crates_section)
	for _, line := range instructions_section {
		command := parseCommand(line)
		crates.MoveMultiple(command)
	}
	return crates.TopCrates()
}

func splitInstructions(lines []string) ([]string, []string) {
	for i := range lines {
		if lines[i] == "" {
			return lines[:i], lines[i+1:]
		}
	}
	panic("Wrong format")
}

func ParseCrate(crates_lines []string) Crates {
	locations := stackLocations(crates_lines[len(crates_lines)-1])
	stacks := make([][]byte, len(locations))
	for _, line := range crates_lines[:len(crates_lines)-1] {
		for stack_n, index := range locations {
			if crate := line[index]; crate != ' ' {
				stacks[stack_n] = append(stacks[stack_n], crate)
			}
		}
	}
	return Crates{stacks}
}

func stackLocations(stack_definition string) []int {
	loc_indexes := make([]int, 0)
	for i, item := range stack_definition {
		if item != ' ' {
			loc_indexes = append(loc_indexes, i)
		}
	}
	return loc_indexes
}

type Crates struct {
	stacks [][]byte
}

func (c Crates) Move(command Command) {
	source := c.stacks[command.Source]
	tmp := source[:command.Quantity]
	slices.Reverse(tmp)
	c.stacks[command.Destination] = slices.Insert(c.stacks[command.Destination], 0, tmp...)
	result := source[:0]
	if len(source) >= command.Quantity {
		result = slices.Delete(source, 0, command.Quantity)
	}
	c.stacks[command.Source] = result
}

func (c Crates) MoveMultiple(command Command) {
	source := c.stacks[command.Source]
	tmp := source[:command.Quantity]
	c.stacks[command.Destination] = slices.Insert(c.stacks[command.Destination], 0, tmp...)
	result := source[:0]
	if len(source) >= command.Quantity {
		result = slices.Delete(source, 0, command.Quantity)
	}
	c.stacks[command.Source] = result
}

func (c Crates) TopCrates() string {
	tops := make([]byte, 0)
	for _, stack := range c.stacks {
		if len(stack) > 0 {
			tops = append(tops, stack[0])
		}
	}
	return string(tops)
}

type Command struct {
	Quantity    int
	Source      int
	Destination int
}

func parseCommand(line string) Command {
	parts := strings.Split(line, " ")
	return Command{
		Quantity:    AtoiPanic(parts[1]),
		Source:      AtoiPanic(parts[3]) - 1,
		Destination: AtoiPanic(parts[5]) - 1,
	}
}

func AtoiPanic(value string) int {
	n, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return n
}
