package day08

import (
	"math"
)

func PartOne(lines []string) int {
	m := NewMatrix(lines)
	fns := []func(*Matrix) [][2]int{knownVisible, LeftToRight, RightToLeft, TopToBottom, BottomToTop}
	coords := make([][2]int, len(fns))
	for _, fn := range fns {
		coord := fn(&m)
		coords = append(coords, coord...)
	}
	size := combine(coords)
	return size
}

func combine(coords [][2]int) int {
	set := make(map[[2]int]struct{})
	for _, coord := range coords {
		set[coord] = struct{}{}
	}
	return len(set)
}

func knownVisible(matrix *Matrix) [][2]int {
	known := make([][2]int, 0)
	height := matrix.H() - 1
	width := matrix.W() - 1
	for i := 0; i <= height; i++ {
		known = append(known, [2]int{0, i})
		known = append(known, [2]int{width, i})
	}
	for j := 0; j <= width; j++ {
		known = append(known, [2]int{j, 0})
		known = append(known, [2]int{j, height})
	}
	return known
}

func PartTwo(lines []string) int {
	m := NewMatrix(lines)
	maxScore := 0
	for x := 0; x < m.W(); x++ {
		for y := 0; y < m.H(); y++ {
			current := m.MeasureScenicScore(x, y)
			if current > maxScore {
				maxScore = current
			}
		}
	}
	return maxScore
}


func (m *Matrix) MeasureScenicScore(x, y int) int {
	left := m.leftScore(x, y)
	right := m.rightScore(x, y)
	up := m.upScore(x, y)
	down := m.downScore(x, y)
	return left * right * up * down
}

func (m *Matrix) leftScore(x, y int) int {
	if x == 0 {
		return 0
	}
	height := m.At(x, y)
	score := 0
	for i := x - 1; i >= 0; i-- {
		score++
		current := m.At(i, y)
		if current >= height {
			break
		}
	}
	return score
}

func (m *Matrix) rightScore(x, y int) int {
	if x >= m.W() - 1 {
		return 0
	}
	height := m.At(x, y)
	score := 0
	for i := x + 1; i < m.W(); i++ {
		score++
		current := m.At(i, y)
		if current >= height {
			break
		}
	}
	return score
}

func (m *Matrix) upScore(x, y int) int {
	if y == 0 {
		return 0
	}
	height := m.At(x, y)
	score := 0
	for j := y - 1; j >= 0; j-- {
		score++
		current := m.At(x, j)
		if current >= height {
			break
		}
	}
	return score
}

func (m *Matrix) downScore(x, y int) int {
	if y >= m.H() {
		return 0
	}
	height := m.At(x, y)
	score := 0
	for j := y + 1; j < m.H(); j++ {
		score++
		current := m.At(x, j)
		if current >= height {
			break
		}
	}
	return score
}

type Matrix struct {
	m  [][]byte
	W  func() int
	H  func() int
	At func(int, int) byte
}

func NewMatrix(lines []string) Matrix {
	matrix := make([][]byte, len(lines))
	for i := 0; i < len(lines); i++ {
		vect := make([]byte, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			vect[j] = byte(lines[i][j]) - 48
		}
		matrix[i] = vect
	}
	return Matrix{
		m:  matrix,
		W:  func() int { return len(matrix[0]) },
		H:  func() int { return len(matrix) },
		At: func(x, y int) byte { return matrix[y][x] },
	}
}

func LeftToRight(matrix *Matrix) [][2]int {
	return VisibleTrees(matrix)
}

func RightToLeft(matrix *Matrix) [][2]int {
	reverse := ReverseCoords(matrix)
	m := Matrix{
		m:  matrix.m,
		W:  matrix.W,
		H:  matrix.H,
		At: func(x, y int) byte { return matrix.At(reverse(x, y)) }}
	coords := VisibleTrees(&m)
	for i := range coords {
		coords[i][0], coords[i][1] = reverse(coords[i][0], coords[i][1])
	}
	return coords
}

func ReverseCoords(matrix *Matrix) func(int, int) (int, int) {
	width := matrix.W() - 1
	height := matrix.H() - 1
	return func(x, y int) (int, int) {
		i := int(math.Abs(float64(x - width)))
		j := int(math.Abs(float64(y - height)))
		return i, j
	}
}

func TopToBottom(matrix *Matrix) [][2]int {
	m := Matrix{
		m:  matrix.m,
		W:  matrix.H,
		H:  matrix.W,
		At: func(x, y int) byte { return matrix.At(y, x) }}
	coords := VisibleTrees(&m)
	for i := range coords {
		coords[i][0], coords[i][1] = coords[i][1], coords[i][0]
	}
	return coords
}

func BottomToTop(matrix *Matrix) [][2]int {
	reverse := ReverseCoords(matrix)
	transposeReverse := func(x, y int) (int, int) {
		y, x = x, y
		return reverse(x, y)
	}
	m := Matrix{
		m:  matrix.m,
		H:  matrix.W,
		W:  matrix.H,
		At: func(x, y int) byte { return matrix.At(transposeReverse(x, y)) },
	}
	coords := VisibleTrees(&m)
	for i := range coords {
		coords[i][0], coords[i][1] = transposeReverse(coords[i][0], coords[i][1])
	}
	return coords
}

func VisibleTrees(matrix *Matrix) [][2]int {
	visible := make([][2]int, 0)
	for y := 1; y < matrix.H()-1; y++ {
		tall := matrix.At(0, y)
		for x := 1; x < matrix.W()-1; x++ {
			if tall == 9 {
				break
			}
			if item := matrix.At(x, y); item > tall {
				visible = append(visible, [2]int{x, y})
				tall = item
			}
		}
	}
	return visible
}
