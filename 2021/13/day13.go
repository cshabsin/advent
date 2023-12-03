package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cshabsin/advent/commongen/board"
	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day13a("sample.txt")
	day13b("sample.txt")
	fmt.Println("---")
	day13a("input.txt")
	day13b("input.txt")
}

type fold struct {
	isX bool
	val int
}

func day13a(fn string) {
	ch, err := readinp.Read(fn, readinp.S)
	if err != nil {
		log.Fatal(err)
	}
	b, folds := readall(ch)
	b = perform(b, folds[0])
	fmt.Println(len(b))
}

func day13b(fn string) {
	ch, err := readinp.Read(fn, readinp.S)
	if err != nil {
		log.Fatal(err)
	}
	b, folds := readall(ch)
	for _, f := range folds {
		b = perform(b, f)
	}
	print(b)
}

func perform(b map[board.Coord]bool, f fold) map[board.Coord]bool {
	if f.isX {
		// fold along x=val, fold the right side left.
		newB := map[board.Coord]bool{}
		for co := range b {
			c := co.C()
			if c > f.val {
				c = f.val*2 - c // val-(c-val)
				if c < 0 {
					continue
				}
			}
			newB[board.MakeCoord(co.R(), c)] = true
		}
		return newB
	}
	// fold along y=val, fold the bottom up.
	newB := map[board.Coord]bool{}
	for co := range b {
		r := co.R()
		if r > f.val {
			r = f.val*2 - r // val-(r-val)
			if r < 0 {
				continue
			}
		}
		newB[board.MakeCoord(r, co.C())] = true
	}
	return newB
}

func print(b map[board.Coord]bool) {
	var maxR, maxC int
	for co := range b {
		if co.R() > maxR {
			maxR = co.R()
		}
		if co.C() > maxC {
			maxC = co.C()
		}
	}
	var brd [][]bool
	for r := 0; r <= maxR; r++ {
		brd = append(brd, make([]bool, maxC+1, maxC+1))
	}
	for co := range b {
		brd[co.R()][co.C()] = true
	}
	for _, r := range brd {
		var rowS string
		for _, c := range r {
			if c {
				rowS += "#"
			} else {
				rowS += "."
			}
		}
		fmt.Println(rowS)
	}
}

func readall(ch chan readinp.Line[string]) (map[board.Coord]bool, []fold) {
	b := map[board.Coord]bool{}
	for line := range ch {
		s, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		if s == "" {
			break
		}
		coS := strings.Split(s, ",")
		x, y := readinp.Atoi(coS[0]), readinp.Atoi(coS[1])
		b[board.MakeCoord(y, x)] = true
	}
	var folds []fold
	for line := range ch {
		s, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		s = strings.TrimPrefix(s, "fold along ")
		fields := strings.Split(s, "=")
		folds = append(folds, fold{
			isX: fields[0] == "x",
			val: readinp.Atoi(fields[1]),
		})
	}
	return b, folds
}
