package main

import (
	"fmt"
	"log"
	"sort"
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
	unvisited := distanceBoard{
		ds:        make(map[board.Coord]intS),
		nexts:     make(map[intS][]board.Coord),
		unvisited: make(map[board.Coord]bool),
	}
	for _, co := range brd.AllCoords() {
		unvisited.setRaw(co, 99999999999)
	}
	current := board.MakeCoord(0, 0)
	target := board.MakeCoord(len(brd)-1, len(brd[0])-1)
	unvisited.setRaw(current, 0)
	for {
		if current == target {
			fmt.Println(unvisited.get(current))
			return
		}
		for _, neigh := range brd.Neighbors4(current) {
			if !unvisited.has(neigh) {
				continue
			}
			unvisited.set(neigh, unvisited.get(current)+brd.GetCoord(neigh))
		}
		unvisited.remove(current)
		current = unvisited.next()
	}
}

type distanceBoard struct {
	ds        map[board.Coord]intS
	unvisited map[board.Coord]bool

	// map from distance to list of coordinates with that size
	nexts map[intS][]board.Coord

	// sorted list (smallest to largest) of distances in the next map
	nextDistances []intS
}

func (d *distanceBoard) set(co board.Coord, dist intS) {
	// fmt.Println("setting", co, "to", dist)
	d.ds[co] = dist
	if d.nexts[dist] == nil {
		d.nextDistances = append(d.nextDistances, dist)
		sort.Slice(d.nextDistances, func(i, j int) bool { return d.nextDistances[i] < d.nextDistances[j] })
	}
	d.nexts[dist] = append(d.nexts[dist], co)
}

func (d *distanceBoard) setRaw(co board.Coord, val intS) {
	d.ds[co] = val
	d.unvisited[co] = true
	// don't put this distance into the nodes to consider for next current node.
}

func (d *distanceBoard) has(co board.Coord) bool {
	return d.unvisited[co]
}

func (d *distanceBoard) get(co board.Coord) intS {
	return d.ds[co]
}

func (d *distanceBoard) remove(co board.Coord) {
	delete(d.unvisited, co)
}

func (d *distanceBoard) next() board.Coord {
	dist := d.nextDistances[0]
	next := d.nexts[dist][0]
	d.nexts[dist] = d.nexts[dist][1:len(d.nexts[dist])]
	if len(d.nexts[dist]) == 0 {
		d.nextDistances = d.nextDistances[1:len(d.nextDistances)]
		delete(d.nexts, dist)
	}
	return next
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
