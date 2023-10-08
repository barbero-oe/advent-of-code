package day03

import (
	"testing"

	utils "github.com/barbero-oe/advent-of-code/tree/main/2022/golang"
)

func TestPart1(t *testing.T) {
	lines := utils.ReadLines("test", t)
	utils.AssertEquals(157, PartOne(lines), t)
}

func TestPart2(t *testing.T) {
	lines := utils.ReadLines("test", t)
	utils.AssertEquals(70, PartTwo(lines), t)
}

func TestRealPart1(t *testing.T) {
	lines := utils.ReadLines("input", t)
	utils.AssertEquals(7875, PartOne(lines), t)
}

func TestRealPart2(t *testing.T) {
	lines := utils.ReadLines("input", t)
	utils.AssertEquals(2479, PartTwo(lines), t)
}
