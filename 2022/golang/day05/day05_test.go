package day05

import (
	"strings"
	"testing"

	. "github.com/barbero-oe/advent-of-code/tree/main/2022/golang"
)

func TestDay05(t *testing.T) {
	RunAoCTest(t, []TestCase[AoC[string]]{
		{Name: "01-Test", Tc: AoC[string]{Filename: "test", Expected: "CMZ", Fn: PartOne}},
		{Name: "01-Real", Tc: AoC[string]{Filename: "input", Expected: "VPCDMSLWJ", Fn: PartOne}},
		{Name: "02-Test", Tc: AoC[string]{Filename: "test", Expected: "MCD", Fn: PartTwo}},
		{Name: "02-Real", Tc: AoC[string]{Filename: "input", Expected: "TPWCGNCCG", Fn: PartTwo}},
	})
}

func TestCratesParser(t *testing.T) {
	crate_str := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `
	crate := ParseCrate(strings.Split(crate_str, "\n"))
	AssertEqualContents([]byte{'N', 'Z'}, crate.stacks[0], t)
	AssertEqualContents([]byte{'D', 'C', 'M'}, crate.stacks[1], t)
	AssertEqualContents([]byte{'P'}, crate.stacks[2], t)
}
