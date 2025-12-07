package day01

import (
	"bufio"
	"os"
	"testing"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func TestProblemOne(t *testing.T) {
	var cases = []struct {
		day      string
		file     string
		expected int
	}{
		{"Day 01", "./example01", 3},
		{"Day 01", "./input", 3},
	}
	for _, c := range cases {
		lines, err := ReadLines(c.file)
		if err != nil {
			t.Fatal(err)
		}
		t.Run(c.day, func(t *testing.T) {
			result := ProblemOne(lines)
			if result != c.expected {
				t.Errorf("Expected %d, but got %d", c.expected, result)
			}
		})
	}
}
