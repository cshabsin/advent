package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cshabsin/advent/commongen/board"
	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	part1("sample.txt", false)
	part1("sample.txt", true)
	part1("input.txt", false)
	part1("input.txt", true)
	// part2("sample.txt")
	// fmt.Println("---")
	// part2("input.txt")
}

func part1(fn string, isQuint bool) {
	fmt.Println("---", fn, isQuint, ":")
	brd, err := load(fn)
	if err != nil {
		log.Fatal(err)
	}
	if isQuint {
		brd = quintuple(brd)
	}
	totals := make(board.Board[intS], len(brd), len(brd))
	for i := range totals {
		totals[i] = make([]intS, len(brd[i]))
	}
	fmt.Println(traverse(brd, totals, board.MakeCoord(0, 0), 0))
}

func traverse(brd, totals board.Board[intS], co board.Coord, risk intS) intS {
	if co.R() == len(brd)-1 && co.C() == len(brd[0])-1 {
		return risk
	}
	var minRisk intS = 9999999
	for _, dco := range brd.Neighbors4(co) {
		if totals.GetCoord(dco) == 0 || totals.GetCoord(dco) > risk+brd.GetCoord(dco) {
			totals.SetCoord(dco, risk+brd.GetCoord(dco))
			subRisk := traverse(brd, totals, dco, risk+brd.GetCoord(dco))
			if subRisk < minRisk {
				minRisk = subRisk
			}
		}
	}
	return minRisk
}

func load(fn string) (board.Board[intS], error) {
	ch, err := readinp.Read(fn, func(s string) ([]intS, error) {
		var rc []intS
		for _, c := range s {
			rc = append(rc, intS(c)-'0')
		}
		return rc, nil
	})
	if err != nil {
		return nil, err
	}
	var acc board.Board[intS]
	for line := range ch {
		row, err := line.Get()
		if err != nil {
			return nil, err
		}
		acc = append(acc, row)
	}
	return acc, nil
}

type intS int

func (i intS) String() string {
	return strconv.Itoa(int(i))
}

func quintuple(brd board.Board[intS]) board.Board[intS] {
	var out board.Board[intS]
	for r := 0; r < 5; r++ {
		for rx := 0; rx < len(brd); rx++ {
			var row []intS
			for c := 0; c < 5; c++ {
				for cx := 0; cx < len(brd[0]); cx++ {
					row = append(row, dumbmod(brd.Get(rx, cx)+intS(r+c)))
				}
			}
			out = append(out, row)
		}
	}
	return out
}

func dumbmod(i intS) intS {
	for i > 9 {
		return i - 9
	}
	return i
}
