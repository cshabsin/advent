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
	distBrd := distanceBoard{
		distances: make(map[board.Coord]intS),
		unvisited: make(map[board.Coord]bool),
		nexts:     make(map[intS][]board.Coord),
	}
	for _, co := range brd.AllCoords() {
		distBrd.initialize(co, 99999999999)
	}
	current := board.MakeCoord(0, 0)
	target := board.MakeCoord(len(brd)-1, len(brd[0])-1)
	distBrd.initialize(current, 0)
	for {
		// distBrd.visualize(brd, current)
		// time.Sleep(time.Millisecond * 50)
		if current == target {
			fmt.Println(distBrd.get(current))
			return
		}
		for _, neigh := range brd.Neighbors4(current) {
			if !distBrd.isUnvisited(neigh) {
				continue
			}
			newDist := distBrd.get(current) + brd.GetCoord(neigh)
			if newDist < distBrd.get(neigh) {
				distBrd.set(neigh, newDist)
			}
		}
		distBrd.remove(current)
		current = distBrd.next()
	}
}

type distanceBoard struct {
	distances map[board.Coord]intS
	unvisited map[board.Coord]bool

	// map from distance to list of coordinates with that size
	nexts map[intS][]board.Coord

	// sorted list (smallest to largest) of distances in the next map
	nextDistances []intS
}

func (d *distanceBoard) set(co board.Coord, dist intS) {
	// fmt.Println("setting", co, "to", dist)
	d.distances[co] = dist
	if d.nexts[dist] == nil {
		d.nextDistances = append(d.nextDistances, dist)
		sort.Slice(d.nextDistances, func(i, j int) bool { return d.nextDistances[i] < d.nextDistances[j] })
	}
	d.nexts[dist] = append(d.nexts[dist], co)
}

func (d *distanceBoard) initialize(co board.Coord, val intS) {
	d.distances[co] = val
	d.unvisited[co] = true
	// don't put this distance into the nodes to consider for next current node.
}

func (d *distanceBoard) isUnvisited(co board.Coord) bool {
	return d.unvisited[co]
}

func (d *distanceBoard) get(co board.Coord) intS {
	return d.distances[co]
}

func (d *distanceBoard) remove(co board.Coord) {
	delete(d.unvisited, co)
}

func (d *distanceBoard) next() board.Coord {
	var next board.Coord
	for {
		dist := d.nextDistances[0]
		next = d.nexts[dist][0]
		d.nexts[dist] = d.nexts[dist][1:len(d.nexts[dist])]
		if len(d.nexts[dist]) == 0 {
			d.nextDistances = d.nextDistances[1:len(d.nextDistances)]
			delete(d.nexts, dist)
		}
		if d.isUnvisited(next) {
			break
		}
	}
	return next
}

func (d *distanceBoard) visualize(brd board.Board[intS], current board.Coord) {
	for r := 0; r < brd.Height(); r++ {
		for c := 0; c < brd.Width(); c++ {
			co := board.MakeCoord(r, c)
			var format string
			if co == current {
				format = "\033[1;33m%s\033[0m"
			} else if d.isUnvisited(co) {
				format = "\033[1;32m%s\033[0m"
			} else {
				format = "\033[0;36m%s\033[0m"
			}
			fmt.Printf(format, strconv.Itoa(int(brd.GetCoord(co))))
		}
		fmt.Print("\n")
	}
	fmt.Print("\033[1;1H")
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
