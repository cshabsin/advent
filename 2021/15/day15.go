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
	"github.com/cshabsin/advent/commongen/set"
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
		board:     brd,
		distances: make(map[board.Coord]intS),
		paths:     make(map[board.Coord]set.Set[board.Coord]),
		unvisited: make(map[board.Coord]bool),
		nexts:     make(map[intS][]board.Coord),
	}
	distBrd.paths[board.MakeCoord(0, 0)] = make(set.Set[board.Coord])
	distBrd.paths[board.MakeCoord(0, 0)][board.MakeCoord(0, 0)] = true
	for _, co := range brd.AllCoords() {
		distBrd.initialize(co, 99999999999)
	}
	current := board.MakeCoord(0, 0)
	target := board.MakeCoord(len(brd)-1, len(brd[0])-1)
	distBrd.initialize(current, 0)
	outGIF := &gif.GIF{}
	var iter int
	for {
		if fn == "sample.txt" || (!isQuint && iter%100 == 0) || iter%2000 == 0 {
			outGIF.Image = append(outGIF.Image, distBrd.visualize(brd, current))
			outGIF.Delay = append(outGIF.Delay, 0)
		}
		iter++
		if current == target {
			fmt.Println(distBrd.get(current))
			outGIF.Image = append(outGIF.Image, distBrd.visualize(brd, current))
			outGIF.Delay = append(outGIF.Delay, 1000)
			break
		}
		for _, neigh := range brd.Neighbors4(current) {
			if !distBrd.isUnvisited(neigh) {
				continue
			}
			newDist := distBrd.get(current) + brd.GetCoord(neigh)
			if newDist < distBrd.get(neigh) {
				distBrd.set(current, neigh)
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
	board     board.Board[intS]
	distances map[board.Coord]intS
	paths     map[board.Coord]set.Set[board.Coord]
	unvisited map[board.Coord]bool

	// map from distance to list of coordinates with that size
	nexts map[intS][]board.Coord

	// sorted list (smallest to largest) of distances in the next map
	nextDistances []intS
}

func (d *distanceBoard) set(current, co board.Coord) {
	dist := d.get(current) + d.board.GetCoord(co)
	// fmt.Println("setting", co, "to", dist)
	d.distances[co] = dist
	d.paths[co] = make(set.Set[board.Coord])
	d.paths[co][co] = true
	for c := range d.paths[current] {
		d.paths[co].Add(c)
	}
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

var palette21 = color.Palette{
	// 0-9 unvisited
	color.RGBA{
		R: 0, G: 0, B: 0x40, A: 0xff,
	},
	color.RGBA{
		R: 0, G: 0x18, B: 0x28, A: 0xff,
	},
	color.RGBA{
		R: 0, G: 0x30, B: 0x10, A: 0xff,
	},
	color.RGBA{
		R: 0, G: 0x40, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0x20, G: 0x40, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0x40, G: 0x40, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0x40, G: 0x30, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0x40, G: 0x20, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0x40, G: 0x10, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0x40, G: 0, B: 0, A: 0xff,
	},
	// 10-19 visited
	color.RGBA{
		R: 0, G: 0, B: 0xff, A: 0xff,
	},
	color.RGBA{
		R: 0, G: 0x60, B: 0xa0, A: 0xff,
	},
	color.RGBA{
		R: 0, G: 0xb0, B: 0x40, A: 0xff,
	},
	color.RGBA{
		R: 0, G: 0xff, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0x80, G: 0xff, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0xff, G: 0xff, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0xff, G: 0xb0, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0xff, G: 0x80, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0xff, G: 0x40, B: 0, A: 0xff,
	},
	color.RGBA{
		R: 0xff, G: 0, B: 0, A: 0xff,
	},
	// 20 = path
	color.RGBA{
		R: 0xff, G: 0xff, B: 0xff, A: 0xff,
	},
}

func (d *distanceBoard) visCoord(brd board.Board[intS], co, current board.Coord) intS {
	if d.paths[current][co] {
		return 20
	}
	val := brd.GetCoord(co)
	if !d.isUnvisited(co) {
		val += 10
	}
	return val
}

func (d *distanceBoard) visDist(brd board.Board[intS], co, current board.Coord) uint8 {
	if d.paths[current][co] {
		return 201
	}
	if d.isUnvisited(co) {
		return uint8(d.visCoord(brd, co, current))
	}
	curDist := d.get(current)
	coDist := d.get(co)

	// funky!
	return uint8(((200 + coDist - curDist) * 200 / curDist))
	// // scale it into a range from 0-199
	// return uint8(199 - ((curDist - coDist) * 200 / curDist))
}

func (d *distanceBoard) visualize(brd board.Board[intS], current board.Coord) *image.Paletted {
	mul := 500 / brd.Width()
	var pix []uint8
	for r := 0; r < brd.Height(); r++ {
		var pixRow []uint8
		for c := 0; c < brd.Width(); c++ {
			co := board.MakeCoord(r, c)
			// val := d.visCoord(brd, co, current)
			val := d.visDist(brd, co, current)
			for i := 0; i < mul; i++ {
				pixRow = append(pixRow, uint8(val))
			}
		}
		for i := 0; i < mul; i++ {
			pix = append(pix, pixRow...)
		}
	}
	var palette201 []color.Color
	for i := uint8(0); i < 200; i++ {
		palette201 = append(palette201, color.RGBA{i, i / 2, i, 0xff})
	}
	palette201 = append(palette201, color.RGBA{0, 0xff, 0, 0xff})
	palette201 = append(palette201, color.White)
	return &image.Paletted{
		Pix:    pix,
		Stride: brd.Width() * mul, // 1 byte per entry
		Rect: image.Rectangle{
			Min: image.Pt(0, 0),
			Max: image.Pt((mul*brd.Width())-1, (mul*brd.Height())-1),
		},
		Palette: palette201,
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
