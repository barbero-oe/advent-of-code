package utils

import (
	"bufio"
	"os"
	"testing"
)

func ReadLines(input string) []string {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func AssertEquals[T comparable](expected, actual T, t *testing.T) {
	if expected != actual {
		t.Fatal("Expected", expected, "but got", actual)
	}
}

func AssertEqualContents[T comparable](expected, actual []T, t *testing.T) {
	if len(expected) != len(actual) {
		t.Fatal("Expected", expected, "but got", actual)
	} else {
		for i := 0; i < len(expected); i++ {
			if expected[i] != actual[i] {
				t.Fatal("Expected", expected, "but got", actual)
			}
		}
	}
}

type TestCase[TC any] struct {
	Name string
	Tc   TC
}

type AoC[T comparable] struct {
	Input    string
	Filename string
	Expected T
	Fn       func([]string) T
}

func Parametrize[T any](t *testing.T, parameters []TestCase[T], fn func(T, *testing.T)) {
	for _, test := range parameters {
		t.Run(test.Name, func(t *testing.T) { fn(test.Tc, t) })
	}
}

func RunAoCTest[T comparable](t *testing.T, parameters []TestCase[AoC[T]]) {
	Parametrize(t, parameters, func(tc AoC[T], t *testing.T) {
		var lines []string
		if tc.Input == "" {
			lines = ReadLines(tc.Filename)
		} else {
			lines = []string{tc.Input}
		}
		actual := tc.Fn(lines)
		AssertEquals(tc.Expected, actual, t)
	})
}
