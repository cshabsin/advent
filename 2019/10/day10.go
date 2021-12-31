package main

import (
	"constraints"
	"fmt"

	"github.com/cshabsin/advent/commongen/board"
	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	b := load("input.txt")
	part1(b)
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
	if num < 0 {
		return fraction{-num / g, -denom / g}
	}
	return fraction{num / g, denom / g}
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
