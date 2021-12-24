package main

import (
	"fmt"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	part1("sample.txt")
	part1("input.txt")
}

func part1(fn string) {
	b := readBoard(fn)
	changed := true
	var steps int
	for changed {
		b, changed = next(b)
		steps++
	}
	fmt.Println(b)
	fmt.Println(steps)
}

type board [][]space

func (b board) String() string {
	var out string
	for _, row := range b {
		for _, c := range row {
			switch c {
			case blank:
				out += "."
			case east:
				out += ">"
			case south:
				out += "v"
			}
		}
		out += "\n"
	}
	return out
}

// returns next board and whether or not anything changed
func next(in board) (board, bool) {
	var out board
	for _, r := range in {
		out = append(out, make([]space, len(r), len(r)))
	}
	var changed bool
	// place eastbound cucumbers in out
	for r, rowSpaces := range in {
		for c, sp := range rowSpaces {
			if sp == east {
				eastCol := (c + 1) % len(rowSpaces)
				if in[r][eastCol] == blank {
					changed = true
					out[r][eastCol] = east
				} else {
					out[r][c] = east
				}
			}
		}
	}
	// place southbound cucumbers in out
	for r, rowSpaces := range in {
		for c, sp := range rowSpaces {
			if sp == south {
				southRow := (r + 1) % len(in)
				// if the input has a south there, we're blocked.
				// if the input has an east there, make sure we look at the output instead, in case it's moved
				if in[southRow][c] != south && out[southRow][c] == blank {
					changed = true
					out[southRow][c] = south
				} else {
					out[r][c] = south
				}
			}
		}
	}
	return out, changed
}

type space int

const (
	blank space = iota
	east
	south
)

func readBoard(fn string) board {
	ch := readinp.MustRead(fn, parse)
	var b board
	for line := range ch {
		b = append(b, line.MustGet())
	}
	return b
}

func parse(line string) ([]space, error) {
	var rc []space
	for _, c := range line {
		switch c {
		case '.':
			rc = append(rc, blank)
		case '>':
			rc = append(rc, east)
		case 'v':
			rc = append(rc, south)
		}
	}
	return rc, nil
}
