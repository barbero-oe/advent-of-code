package main

import (
	"bufio"
	"os"
	"testing"
)

func readLines(input string, t *testing.T) []string {
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

func TestDay(t *testing.T) {
	lines := readLines("test", t)
	AssertEquals(157, PartOne(lines), t)
}

func TestRealPart1(t *testing.T) {
	lines := readLines("input", t)
	AssertEquals(7875, PartOne(lines), t)
}

func AssertEquals[T comparable](expected, actual T, t *testing.T) {
	if expected != actual {
		t.Fatal("Expected", expected, "but got", actual)
	}
}
