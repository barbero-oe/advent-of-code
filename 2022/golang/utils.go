package utils

import (
	"bufio"
	"os"
	"testing"
)

func ReadLines(input string, t *testing.T) []string {
	f, err := os.Open(input)
	if err != nil {
		t.Fatal(err)
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
