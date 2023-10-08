package day03

import (
	"testing"

	. "github.com/barbero-oe/advent-of-code/tree/main/2022/golang"
)

func TestDay03(t *testing.T) {
	RunAoCTest(t, []TestCase[AoC[int]]{
		{Name: "01-Test", Tc: AoC[int]{Filename: "test", Expected: 157, Fn: PartOne}},
		{Name: "01-Real", Tc: AoC[int]{Filename: "input", Expected: 7875, Fn: PartOne}},
		{Name: "02-Test", Tc: AoC[int]{Filename: "test", Expected: 70, Fn: PartTwo}},
		{Name: "02-Real", Tc: AoC[int]{Filename: "input", Expected: 2479, Fn: PartTwo}},
	})
}
