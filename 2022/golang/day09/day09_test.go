package day09

import (
	"testing"

	. "github.com/barbero-oe/advent-of-code/tree/main/2022/golang"
)

func TestDay09(t *testing.T) {
	RunAoCTest(t, []TestCase[AoC[int]]{
		{Name: "01-Test", Tc: AoC[int]{Filename: "test", Expected: 13, Fn: PartOne}},
		{Name: "01-Real", Tc: AoC[int]{Filename: "input", Expected: 5874, Fn: PartOne}},
		{Name: "02-Test", Tc: AoC[int]{Filename: "test", Expected: 1, Fn: PartTwo}},
		{Name: "02-Test_2", Tc: AoC[int]{Filename: "test_2", Expected: 36, Fn: PartTwo}},
		{Name: "02-Test", Tc: AoC[int]{Filename: "input", Expected: 2467, Fn: PartTwo}},
	})
}
