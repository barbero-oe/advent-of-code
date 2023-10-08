package day04

import (
	"testing"

	. "github.com/barbero-oe/advent-of-code/tree/main/2022/golang"
)

func TestDay04(t *testing.T) {
	RunAoCTest(t, []TestCase[AoC[int]]{
		{Name: "01-Test", Tc: AoC[int]{Filename: "test", Expected: 2, Fn: PartOne}},
		{Name: "01-Real", Tc: AoC[int]{Filename: "input", Expected: 580, Fn: PartOne}},
		{Name: "02-Test", Tc: AoC[int]{Filename: "test", Expected: 4, Fn: PartTwo}},
		{Name: "02-Real", Tc: AoC[int]{Filename: "input", Expected: 895, Fn: PartTwo}},
	})
}
