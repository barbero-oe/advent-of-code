package day06

import (
	"testing"

	. "github.com/barbero-oe/advent-of-code/tree/main/2022/golang"
)

func TestDay06(t *testing.T) {
	RunAoCTest(t, []TestCase[AoC[int]]{
		{Name: "01/1-Test", Tc: AoC[int]{Input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", Expected: 7, Fn: PartOne}},
		{Name: "01/2-Test", Tc: AoC[int]{Input: "bvwbjplbgvbhsrlpgdmjqwftvncz", Expected: 5, Fn: PartOne}},
		{Name: "01/3-Test", Tc: AoC[int]{Input: "nppdvjthqldpwncqszvftbrmjlhg", Expected: 6, Fn: PartOne}},
		{Name: "01/4-Test", Tc: AoC[int]{Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", Expected: 10, Fn: PartOne}},
		{Name: "01/5-Test", Tc: AoC[int]{Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", Expected: 11, Fn: PartOne}},
		{Name: "01-Real", Tc: AoC[int]{Filename: "input", Expected: 1912, Fn: PartOne}},
		{Name: "02/1-Test", Tc: AoC[int]{Input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", Expected: 19, Fn: PartTwo}},
		{Name: "02/2-Test", Tc: AoC[int]{Input: "bvwbjplbgvbhsrlpgdmjqwftvncz", Expected: 23, Fn: PartTwo}},
		{Name: "02/3-Test", Tc: AoC[int]{Input: "nppdvjthqldpwncqszvftbrmjlhg", Expected: 23, Fn: PartTwo}},
		{Name: "02/4-Test", Tc: AoC[int]{Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", Expected: 29, Fn: PartTwo}},
		{Name: "02/5-Test", Tc: AoC[int]{Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", Expected: 26, Fn: PartTwo}},
		{Name: "02-Real", Tc: AoC[int]{Filename: "input", Expected: 2122, Fn: PartTwo}},
	})
}
