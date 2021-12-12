package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/cshabsin/advent/commongen/board"
	"github.com/cshabsin/advent/commongen/readinp"
	"github.com/cshabsin/advent/commongen/set"
)

func main() {
	day9a("sample.txt")
	day9b("sample.txt")
	fmt.Println("---")
	day9a("input.txt")
	day9b("input.txt")
}

type intS int

func (i intS) String() string {
	return strconv.Itoa(int(i))
}

func (intS) Delimiter() string {
	return "" // no space between entries
}

func day9a(fn string) {
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
	var rsk intS
	for _, co := range brd.AllCoords() {
		rsk += risk(brd, co)
	}
	fmt.Println(rsk)
}

func day9b(fn string) {
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
	var lows []board.Coord
	for _, co := range brd.AllCoords() {
		if isLow(brd, co) {
			lows = append(lows, co)
		}
	}
	var basins []int
	for _, low := range lows {
		basins = append(basins, basinSize(brd, low))
	}
	sort.Sort(sort.IntSlice(basins))
	fmt.Println(basins)
}

func isLow(b board.Board[intS], co board.Coord) bool {
	for _, dco := range b.Neighbors8(co) {
		if b.GetCoord(dco) <= b.GetCoord(co) {
			return false
		}
	}
	return true
}

func risk(brd board.Board[intS], co board.Coord) intS {
	if !isLow(brd, co) {
		return 0
	}
	return brd.GetCoord(co) + 1
}

func basinSize(brd board.Board[intS], co board.Coord) int {
	return basinSizeInternal(brd, co, set.Set[board.Coord]{})
}

func basinSizeInternal(brd board.Board[intS], co board.Coord, v set.Set[board.Coord]) int {
	v.Add(co)
	if brd.GetCoord(co) == 9 {
		return 0
	}
	basin := 1
	for _, dco := range brd.Neighbors4(co) {
		if v.Contains(dco) {
			continue
		}
		if brd.GetCoord(dco) >= brd.GetCoord(co) {
			basin += basinSizeInternal(brd, dco, v)
		}
	}
	return basin
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
