package day01

import (
	"testing"

	utils "github.com/barbero-oe/advent-of-code/tree/main/2022/golang"
)

func TestDay01(t *testing.T) {
	testCases := []struct {
		part     int
		filename string
		expected int
	}{
		{1, "test", 24000},
		{1, "input", 66616},
		{2, "test", 45000},
		{2, "input", 199172},
	}

	for _, tc := range testCases {
		switch tc.part {
		case 1:
			t.Run("Part1", func(t *testing.T) { Part1(tc.filename, tc.expected, t) })
		case 2:
			t.Run("Part2", func(t *testing.T) { Part2(tc.filename, tc.expected, t) })
		default:
			t.FailNow()
		}
	}

}

func Part1(file string, expected int, t *testing.T) {
	max, _ := run(file)
	utils.AssertEquals(expected, max, t)
}

func Part2(file string, expected int, t *testing.T) {
	_, sum := run(file)
	utils.AssertEquals(expected, sum, t)
}
