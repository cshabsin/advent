package board

import (
	"fmt"
)

type Coord struct {
	r, c int
}

func MakeCoord(r, c int) Coord {
	return Coord{r, c}
}

func (c Coord) R() int {
	return c.r
}

func (c Coord) C() int {
	return c.c
}

func (c Coord) Apply(diff [2]int) Coord {
	return MakeCoord(c.R()+diff[0], c.C()+diff[1])
}

type delimiterHaver interface {
	Delimiter() string
}

func delim(str fmt.Stringer) string {
	d, ok := str.(delimiterHaver)
	if !ok {
		return " " // default delimiter
	}
	return d.Delimiter()
}

type Board[T fmt.Stringer] [][]T

// String renders the board in string.
//
// This implementation assumes that the String() function
// for the type returns values of the same width, and there
// is no need for any delimiter
// (designed for the one-digit boards from 2021 day 9 and 11).
func (b Board[T]) String() string {
	var s string
	for r := range b {
		for c := range b[r] {
			s += b[r][c].String() + delim(b[r][c])
		}
		s += "\n"
	}
	return s
}

func (b Board[T]) Get(r, c int) T {
	var zero T
	if r < 0 || c < 0 {
		return zero
	}
	if r >= len(b) {
		return zero
	}
	if c >= len(b[r]) {
		return zero
	}
	return b[r][c]
}

func (b *Board[T]) Set(r, c int, val T) {
	if r < 0 || c < 0 {
		return
	}
	if r >= len(*b) {
		return
	}
	if c >= len((*b)[r]) {
		return
	}
	(*b)[r][c] = val
}

func (b Board[T]) GetCoord(co Coord) T {
	return b.Get(co.R(), co.C())
}

func (b *Board[T]) SetCoord(co Coord, val T) {
	b.Set(co.R(), co.C(), val)
}

func (b Board[T]) AllCoords() []Coord {
	var coords []Coord
	for r := range b {
		for c := range b[r] {
			coords = append(coords, MakeCoord(r, c))
		}
	}
	return coords
}

func (b Board[T]) Height() int {
	return len(b)
}

func (b Board[T]) Width() int {
	return len(b[0])
}

func (b Board[T]) InBounds(co Coord) bool {
	if co.R() < 0 || co.C() < 0 {
		return false
	}
	if co.R() >= len(b) {
		return false
	}
	if co.C() >= len(b[co.R()]) {
		return false
	}
	return true
}

func (b Board[T]) Neighbors4(co Coord) []Coord {
	var coords []Coord
	for _, diff := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		dco := co.Apply(diff)
		if !b.InBounds(dco) {
			continue
		}
		coords = append(coords, dco)
	}
	return coords
}

func (b Board[T]) Neighbors8(co Coord) []Coord {
	var coords []Coord
	for _, diff := range [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
		dco := co.Apply(diff)
		if !b.InBounds(dco) {
			continue
		}
		coords = append(coords, dco)
	}
	return coords
}
