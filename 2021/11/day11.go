package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cshabsin/advent/commongen/board"
	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day11a("sample.txt")
	// day11b("sample.txt")
	fmt.Println("---")
	day11a("input.txt")
	// day11b("input.txt")
}

type intS int

func (i intS) String() string {
	return strconv.Itoa(int(i))
}

func (intS) Delimiter() string {
	return "" // no space between entries
}

func day11a(fn string) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	var brd board.Board[intS]
	for line := range ch {
		row, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		brd = append(brd, row)
	}
	var flashes100 int
	var i int
	for {
		i++
		var f int
		f, brd = nextBoard(brd)
		if i < 100 {
			flashes100 += f
			fmt.Println(brd)
		}
		if f == 100 {
			break
		}
	}
	fmt.Println(flashes100, i)
}
func nextBoard(b board.Board[intS]) (int, board.Board[intS]) {
	var next board.Board[intS]
	var flashes int
	for r := range b {
		next = append(next, make([]intS, len(b[r]), len(b[r])))
		for c := range b[r] {
			next[r][c] = b[r][c] + 1
		}
	}
	for {
		var changed bool
		for _, co := range b.AllCoords() {
			if next.GetCoord(co) < 10 {
				continue
			}
			next.SetCoord(co, 0)
			flashes++
			changed = true
			for _, dco := range b.Neighbors8(co) {
				if next.GetCoord(dco) == 0 {
					continue // already exploded
				}
				next.SetCoord(dco, next.GetCoord(dco)+1)
			}
		}
		if !changed {
			break
		}
	}
	return flashes, next
}

func parse(line string) ([]intS, error) {
	var d []intS
	for _, c := range line {
		if c < '0' || c > '9' {
			return nil, fmt.Errorf("in line %q, illegal character %c", line, c)
		}
		d = append(d, intS(c-'0'))
	}
	return d, nil
}
