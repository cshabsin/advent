package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

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
	outGIF := &gif.GIF{}
	var iter int
	for {
		if fn == "sample.txt" || iter%2000 == 0 {
			outGIF.Image = append(outGIF.Image, distBrd.visualize(brd, current))
			outGIF.Delay = append(outGIF.Delay, 0)
		}
		iter++
		if current == target {
			fmt.Println(distBrd.get(current))
			outGIF.Image = append(outGIF.Image, distBrd.visualize(brd, current))
			outGIF.Delay = append(outGIF.Delay, 0)
			break
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
	outFile := strings.TrimSuffix(fn, ".txt")
	if isQuint {
		outFile += "5"
	}
	outFile += "_anim.gif"
	f, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	gif.EncodeAll(f, outGIF)
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

func (d *distanceBoard) visualize(brd board.Board[intS], current board.Coord) *image.Paletted {
	var pix []uint8
	for r := 0; r < brd.Height(); r++ {
		for c := 0; c < brd.Width(); c++ {
			co := board.MakeCoord(r, c)
			val := brd.GetCoord(co)
			if !d.isUnvisited(co) {
				val += 10
			}
			pix = append(pix, uint8(val))
		}
	}
	var palette color.Palette
	for i := 0; i < 19; i++ {
		clr := color.RGBA{
			R: uint8(0x10 * (i % 10)),
			B: uint8(0xf0 - (0x10 * (i % 10))),
			A: 0xff,
		}
		if i > 9 {
			clr.R += 0x7f
			clr.B += 0x7f
		}
		palette = append(palette, clr)
	}
	return &image.Paletted{
		Pix:    pix,
		Stride: brd.Width(), // 1 byte per entry
		Rect: image.Rectangle{
			Min: image.Pt(0, 0),
			Max: image.Pt(brd.Width()-1, brd.Height()-1),
		},
		Palette: palette,
	}
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
