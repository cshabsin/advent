package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day11a("sample.txt")
	// day11b("sample.txt")
	fmt.Println("---")
	day11a("input.txt")
	// day11b("input.txt")
}

func day11a(fn string) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	var board board
	for line := range ch {
		row, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		board = append(board, row)
	}
	var flashes100 int
	var i int
	for {
		i++
		var f int
		f, board = board.next()
		if i < 100 {
			flashes100 += f
		}
		if f == 100 {
			break
		}
	}
	fmt.Println(flashes100, i)
}

type board [][]int

func (b board) String() string {
	var s string
	for r := range b {
		for c := range b[r] {
			s += strconv.Itoa(b[r][c])
		}
		s += "\n"
	}
	return s
}

func (b board) get(r, c int) int {
	if r < 0 || c < 0 {
		return 0 // the outer edge has no energy
	}
	if r >= len(b) {
		return 0
	}
	if c >= len(b[r]) {
		return 0
	}
	return b[r][c]
}

func (b *board) set(r, c, val int) {
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

func (b board) getCoord(co coord) int {
	return b.get(co.r(), co.c())
}

type coord int

func makeCoord(r, c int) coord {
	return coord(r*1000 + c)
}

func (c coord) r() int {
	return int(c / 1000)
}

func (c coord) c() int {
	return int(c % 1000)
}

func (b board) next() (int, board) {
	var next board
	var flashes int
	for r := range b {
		next = append(next, make([]int, len(b[r]), len(b[r])))
		for c := range b[r] {
			next[r][c] = b[r][c] + 1
		}
	}
	for {
		var changed bool
		for r := range b {
			for c := range b[r] {
				if next.get(r, c) < 10 {
					continue
				}
				next.set(r, c, 0)
				flashes++
				changed = true
				for _, diff := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
					dr, dc := diff[0], diff[1]
					if next.get(r+dr, c+dc) == 0 {
						continue // already exploded
					}
					next.set(r+dr, c+dc, next.get(r+dr, c+dc)+1)
				}
			}
		}
		if !changed {
			break
		}
	}
	return flashes, next
}

func parse(line string) ([]int, error) {
	var d []int
	for _, c := range line {
		if c < '0' || c > '9' {
			return nil, fmt.Errorf("in line %q, illegal character %c", line, c)
		}
		d = append(d, int(c-'0'))
	}
	return d, nil
}
