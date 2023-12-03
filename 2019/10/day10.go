package main

import (
	"constraints"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/cshabsin/advent/commongen/ansi"
	"github.com/cshabsin/advent/commongen/board"
	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	b := load("input.txt")
	part1(b)
	part2(b)
}

func part1(b board.Board[boolS]) {
	var max int
	var station board.Coord
	for _, s := range b.AllCoords() {
		c := canSee(b, s)
		cnt := len(c)
		if b.GetCoord(station) {
			cnt++ // station can see the asteroid it's on
		}
		if cnt > max {
			max = cnt
			station = s
		}
	}
	fmt.Println(station, max)
}

func part2(b board.Board[boolS]) {
	station := board.MakeCoord(19, 22)
	canSee := canSee(b, station)
	var tgts fraclist
	for f := range canSee {
		tgts = append(tgts, f)
	}
	sort.Sort(tgts)
	// fmt.Println(tgts)
	ansi.Clear()
	fmt.Println(b)
	ansi.Loc(station.R(), station.C())
	ansi.Color(35, "!")
	for i, tgt := range tgts {
		time.Sleep(time.Millisecond * 5)
		coord := canSee[tgt]
		ansi.Loc(coord.R(), coord.C())
		c := 34
		if i == 199 {
			c = 36
		}
		ansi.Color(c, "#")
		// ansi.Loc(coord.R(), coord.C()+40)
		// fmt.Print(tgt, coord)
	}
	ansi.Loc(28, 0)
	fmt.Println(tgts[199], canSee[tgts[199]])
}

type fraction struct {
	num, denom int
}

func abs[T constraints.Signed](a T) T {
	if a < T(0) {
		return T(-a)
	}
	return a
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func makeFrac(num, denom int) fraction {
	g := gcd(num, denom)
	if num > 0 {
		num = abs(num / g)
	} else {
		num = -abs(num / g)
	}
	if denom > 0 {
		denom = abs(denom / g)
	} else {
		denom = -abs(denom / g)
	}
	return fraction{num, denom}
}

type fraclist []fraction

func (f fraclist) Len() int {
	return len(f)
}

func (f fraclist) Swap(i, j int) {
	t := f[i]
	f[i] = f[j]
	f[j] = t
}

func (f fraclist) Less(i, j int) bool {
	iAtan := math.Atan2(-float64(f[i].denom), float64(f[i].num))
	jAtan := math.Atan2(-float64(f[j].denom), float64(f[j].num))
	return iAtan < jAtan
}

func canSee(b board.Board[boolS], station board.Coord) map[fraction]board.Coord {
	canSee := map[fraction]board.Coord{}
	for _, target := range b.AllCoords() {
		if target == station {
			continue
		}
		if !b.GetCoord(target) {
			continue
		}
		f := makeFrac(target.R()-station.R(), target.C()-station.C())
		if alreadySeen, ok := canSee[f]; ok {
			if abs(station.R()-alreadySeen.R()) < abs(station.R()-target.R()) {
				continue // existing one is closer
			}
		}
		canSee[f] = target
	}
	return canSee
}

func load(fn string) board.Board[boolS] {
	var b board.Board[boolS]
	for line := range readinp.MustRead(fn, parse) {
		b = append(b, line.MustGet())
	}
	return b
}

func parse(line string) ([]boolS, error) {
	var b []boolS
	for _, c := range line {
		b = append(b, c == '#')
	}
	return b, nil
}

type boolS bool

func (b boolS) String() string {
	if b {
		return "#"
	}
	return "."
}

func (b boolS) Delimiter() string {
	return ""
}
