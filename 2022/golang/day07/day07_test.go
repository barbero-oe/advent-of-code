package day07

import (
	"testing"

	. "github.com/barbero-oe/advent-of-code/tree/main/2022/golang"
)

func TestDay07(t *testing.T) {
	RunAoCTest(t, []TestCase[AoC[int]]{
		{Name: "01-Test", Tc: AoC[int]{Filename: "test", Expected: 95437, Fn: PartOne}},
		{Name: "01-Real", Tc: AoC[int]{Filename: "input", Expected: 1477771, Fn: PartOne}},
		{Name: "02-Test", Tc: AoC[int]{Filename: "test", Expected: 24933642, Fn: PartTwo}},
		{Name: "02-Test", Tc: AoC[int]{Filename: "input", Expected: 3579501, Fn: PartTwo}},
	})
}
