package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day9a("sample.txt")
	day9b("sample.txt")
	fmt.Println("---")
	day9a("input.txt")
	day9b("input.txt")
}

func day9a(fn string) {
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
	var risk int
	for r := range board {
		for c := range board[r] {
			risk += board.risk(r, c)
		}
	}
	fmt.Println(risk)
}

func day9b(fn string) {
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
	var lows [][2]int
	for r := range board {
		for c := range board[r] {
			if board.isLow(r, c) {
				lows = append(lows, [2]int{r, c})
			}
		}
	}
	var basins []int
	for _, low := range lows {
		r, c := low[0], low[1]
		basins = append(basins, board.basinSize(r, c))
	}
	sort.Sort(sort.IntSlice(basins))
	fmt.Println(basins)
}

type board [][]int

func (b board) get(r, c int) int {
	if r < 0 || c < 0 {
		return 9 // the outer edge is "higher" than everything
	}
	if r >= len(b) {
		return 9
	}
	if c >= len(b[r]) {
		return 9
	}
	return b[r][c]
}

func (b board) isLow(r, c int) bool {
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			if b.get(r+dr, c+dc) <= b.get(r, c) {
				return false
			}
		}
	}
	return true
}

func (b board) risk(r, c int) int {
	if !b.isLow(r, c) {
		return 0
	}
	return b.get(r, c) + 1
}

func (b board) basinSize(r, c int) int {
	return b.basinSizeInternal(r, c, visited{})
}

type visited map[int]bool

func (v visited) isVisited(r, c int) bool {
	return v[r*1000+c]
}

func (v visited) visit(r, c int) {
	v[r*1000+c] = true
}

func (b board) basinSizeInternal(r, c int, v visited) int {
	v.visit(r, c)
	if b.get(r, c) == 9 {
		return 0
	}
	basin := 1
	for _, diff := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		dr := diff[0]
		dc := diff[1]
		if v.isVisited(r+dr, c+dc) {
			continue
		}
		if b.get(r+dr, c+dc) >= b.get(r, c) {
			basin += b.basinSizeInternal(r+dr, c+dc, v)
		}
	}
	return basin
}
