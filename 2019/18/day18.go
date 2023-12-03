package main

import (
	"constraints"
	"fmt"

	"github.com/cshabsin/advent/commongen/ansi"
	"github.com/cshabsin/advent/commongen/board"
	"github.com/cshabsin/advent/commongen/heapof"
	"github.com/cshabsin/advent/commongen/readinp"
	"github.com/cshabsin/advent/commongen/set"
)

// sample1: 8
// sample2: 86
// sample3: 132 (but current algorithm yields 140?)

func main() {
	b := load("sample3.txt")
	part1(b)
}

func part1(b *Maze) {
	h := heapof.Make([]*Maze{b})
	var i int
	for {
		if h.Len() == 0 {
			fmt.Println("no more states!")
			return
		}
		state := h.PopHeap()
		if i%10000 == 0 {
			ansi.Loc(0, 0)
			fmt.Println(state)
			fmt.Println("traversals:", i, "; outstanding states:", h.Len())
		}
		i++
		for _, next := range state.Nexts() {
			if next.Win() {
				fmt.Println("win!", next.turns)
				fmt.Println(next)
				return
			}
			h.PushHeap(next)
		}
	}
}

type Maze struct {
	b board.Board[runeS] // board is immutable, so we can just copy the slice around

	wantKeys     set.Set[runeS]
	keyLocations map[runeS]board.Coord
	location     board.Coord
	visited      set.Set[board.Coord]

	turns int
}

func (m Maze) String() string {
	return fmt.Sprintf("%s\ncost: %d\nlocation: %v (need keys %d)", m.Highlight(), m.Cost(), m.location, len(m.wantKeys))
}

func (m Maze) Highlight() string {
	var s string
	for r := range m.b {
		for c := range m.b[r] {
			co := board.MakeCoord(r, c)
			var lit bool
			if co == m.location {
				s += "\x1b[34;47m"
				lit = true
			}
			if m.visited.Contains(co) {
				s += "\x1b[32;44m"
				lit = true
			}

			s += m.b[r][c].String()
			if lit {
				s += "\x1b[0m"
			}
		}
		s += "\n"
	}
	return s
}

func (m Maze) Cost() int {
	cost := m.turns
	for r, loc := range m.keyLocations {
		if !m.wantKeys.Contains(r) {
			continue
		}
		cost += abs(loc.R()-m.location.R()) + abs(loc.C()-m.location.C())
	}
	return cost
}

func (m Maze) Win() bool {
	return len(m.wantKeys) == 0
}

func (m Maze) MoveTo(loc board.Coord) *Maze {
	v := m.visited.Clone()
	v.Add(loc)
	return &Maze{
		b:            m.b,
		wantKeys:     m.wantKeys.Clone(),
		keyLocations: m.keyLocations,
		location:     loc,
		visited:      v,
		turns:        m.turns + 1,
	}
}

func (m Maze) Nexts() []*Maze {
	var nexts []*Maze
	for _, nextCoord := range m.b.Neighbors4(m.location) {
		if m.visited.Contains(nextCoord) {
			continue
		}
		switch ch := m.b.GetCoord(nextCoord); ch {
		case '#':
			continue
		case '.':
			nexts = append(nexts, m.MoveTo(nextCoord))
		case '@':
			nexts = append(nexts, m.MoveTo(nextCoord))
		default:
			if 'a' <= ch && ch <= 'z' {
				to := m.MoveTo(nextCoord)
				delete(to.wantKeys, ch)
				to.visited = set.Make(nextCoord)
				nexts = append(nexts, to)
				continue
			}
			if 'A' <= ch && ch <= 'Z' {
				if m.wantKeys.Contains(ch + 32) {
					continue // can't pass through
				}
				nexts = append(nexts, m.MoveTo(nextCoord))
				continue
			}
		}
	}
	return nexts
}

func load(fn string) *Maze {
	var b board.Board[runeS]
	wantKeys := set.Set[runeS]{}
	keyLocations := map[runeS]board.Coord{}
	var location board.Coord
	row := 0
	for line := range readinp.MustRead(fn, parse) {
		l := line.MustGet()
		b = append(b, l)
		for col, r := range l {
			if r >= 'a' && r <= 'z' {
				wantKeys.Add(r)
				keyLocations[r] = board.MakeCoord(row, col)
			}
			if r == '@' {
				location = board.MakeCoord(row, col)
			}
		}
		row++
	}
	return &Maze{
		b:            b,
		wantKeys:     wantKeys,
		keyLocations: keyLocations,
		location:     location,
		visited:      set.Make(location),
	}
}

func parse(line string) ([]runeS, error) {
	return []runeS(line), nil
}

type runeS rune

func (r runeS) String() string {
	return string(r)
}

func (r runeS) Delimiter() string {
	return ""
}

func abs[T constraints.Signed](a T) T {
	if a < T(0) {
		return T(-a)
	}
	return a
}
