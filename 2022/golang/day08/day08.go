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
	return 0
}

type MatrixView interface {
	H() int
	W() int
	At(int, int) byte
}

type Matrix struct {
	m          [][]byte
	coord_fn   func(int, int) (int, int)
	transposed bool
}

func (m *Matrix) H() int {
	if m.transposed {
		return len(m.m[0])
	} else {
		return len(m.m)
	}
}

func (m *Matrix) W() int {
	if m.transposed {
		return len(m.m[0])
	} else {
		return len(m.m)
	}
}

func (m *Matrix) At(x, y int) byte {
	i, j := m.coord_fn(x, y)
	return m.m[j][i]
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
	return Matrix{m: matrix, coord_fn: func(x, y int) (int, int) { return x, y }}
}

func LeftToRight(matrix *Matrix) [][2]int {
	return VisibleTrees(matrix)
}

func RightToLeft(matrix *Matrix) [][2]int {
	reverse := ReverseCoords(matrix)
	m := Matrix{m: matrix.m, coord_fn: reverse}
	coords := VisibleTrees(&m)
	for i := range coords {
		coords[i][0], coords[i][1] = reverse(coords[i][0], coords[i][1])
	}
	return coords
}

func ReverseCoords(matrix MatrixView) func(int, int) (int, int) {
	width := matrix.W() - 1
	height := matrix.H() - 1
	return func(x, y int) (int, int) {
		i := int(math.Abs(float64(x - width)))
		j := int(math.Abs(float64(y - height)))
		return i, j
	}
}

func TopToBottom(matrix *Matrix) [][2]int {
	transpose := func(x, y int) (int, int) { return y, x }
	m := Matrix{m: matrix.m, coord_fn: transpose, transposed: true}
	coords := VisibleTrees(&m)
	for i := range coords {
		coords[i][0], coords[i][1] = transpose(coords[i][0], coords[i][1])
	}
	return coords
}

func BottomToTop(matrix *Matrix) [][2]int {
	reverse := ReverseCoords(matrix)
	transposeReverse := func(x, y int) (int, int) {
		y, x = x, y
		return reverse(x, y)
	}
	m := Matrix{m: matrix.m, coord_fn: transposeReverse, transposed: true}
	coords := VisibleTrees(&m)
	for i := range coords {
		coords[i][0], coords[i][1] = transposeReverse(coords[i][0], coords[i][1])
	}
	return coords
}

func VisibleTrees(matrix MatrixView) [][2]int {
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
