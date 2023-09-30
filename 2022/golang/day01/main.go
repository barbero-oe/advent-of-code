package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Test Input:")
	run("test")
	fmt.Println()
	fmt.Println("Real Input:")
	run("input")
}

func run(input string) {
	top := rankTop(3, input)
	first := top[0]
	sum := Sum(top...)
	fmt.Println("Max Calories:", first)
	fmt.Println("Top Sum Calories:", sum)
}

func rankTop(top int, input string) []int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	ranking := NewRanking(3)
	parser := NewCaloriesParser(file)
	for parser.HasNext() {
		calories := parser.Calories()
		calorie := Sum(calories...)
		ranking.Add(calorie)
	}
	return ranking.Top()
}

type CaloriesParser struct {
	scanner *bufio.Scanner
	hasNext bool
}

func NewCaloriesParser(reader io.Reader) *CaloriesParser {
	return &CaloriesParser{
		scanner: bufio.NewScanner(reader),
		hasNext: true,
	}
}

func (p *CaloriesParser) HasNext() bool {
	return p.hasNext
}

func (p *CaloriesParser) Calories() []int {
	calories := make([]int, 0)
	for p.scanner.Scan() {
		line := p.scanner.Text()
		if line == "" {
			return calories
		} else {
			calorie, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			calories = append(calories, calorie)
		}
	}
	p.hasNext = false
	return calories
}

func Sum(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

type Ranking struct {
	top     int
	ranking []int
}

func NewRanking(top int) *Ranking {
	return &Ranking{
		top:     top,
		ranking: make([]int, 0, top+1),
	}
}

func (r *Ranking) Add(item int) {
	if len(r.ranking) < r.top {
		r.ranking = append(r.ranking, item)
		slices.Sort(r.ranking)
		slices.Reverse(r.ranking)
	} else {
		for i, current := range r.ranking {
			if current < item {
				r.ranking = append(r.ranking[:i+1], r.ranking[i:]...)
				r.ranking[i] = item
				r.ranking = r.ranking[:r.top]
				break
			}
		}
	}
}

func (r *Ranking) Top() []int {
	return r.ranking[:]
}
