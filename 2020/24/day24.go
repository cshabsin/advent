package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	b := board{}
	b.Toggle(coord(""))
	b.Toggle(coord("e"))
	fmt.Println(b)
	b.Day()
	fmt.Println(b)

	b = board{}
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		b.Toggle(coord(line.Value()))
	}
	fmt.Println(b.values)
	for i := 0; i < 100; i++ {
		b.Day()
	}
	count := 0
	for _, row := range b.values {
		for _, on := range row {
			if on {
				count++
			}
		}
	}
	fmt.Println(count)
}

func pop(s string) (byte, string) {
	return s[0], s[1:len(s)]
}

func coord(line string) (int, int) {
	var x, y int
	for {
		if line == "" {
			break
		}
		var dir byte
		dir, line = pop(line)
		if dir == 'n' {
			y--
			dir, line = pop(line)
			if dir == 'e' {
				x++
			}
		} else if dir == 's' {
			y++
			dir, line = pop(line)
			if dir == 'w' {
				x--
			}
		} else if dir == 'e' {
			x++
		} else {
			x-- // must be w
		}
	}
	return x, y
}

type board struct {
	minX, minY, maxX, maxY int
	values                 map[int]map[int]bool
}

func (b *board) Toggle(x, y int) {
	if b.values == nil {
		b.values = map[int]map[int]bool{}
	}
	if b.values[x] == nil {
		b.values[x] = map[int]bool{}
	}
	b.values[x][y] = !b.values[x][y]
	if b.minX > x {
		b.minX = x
	}
	if b.minY > y {
		b.minY = y
	}
	if b.maxX < x {
		b.maxX = x
	}
	if b.maxY < y {
		b.maxY = y
	}
}

func (b board) Get(x, y int) bool {
	if b.values == nil {
		return false
	}
	if b.values[x] == nil {
		return false
	}
	return b.values[x][y]
}

func (b *board) Day() {
	newBoard := board{}
	for y := b.minY - 1; y <= b.maxY+1; y++ {
		for x := b.minX - 1; x <= b.maxX+1; x++ {
			n := b.countNeighbors(x, y)
			if b.Get(x, y) {
				if n == 1 || n == 2 {
					newBoard.Toggle(x, y)
				}
			} else {
				if n == 2 {
					newBoard.Toggle(x, y)
				}
			}
		}
	}
	b.minX = newBoard.minX
	b.minY = newBoard.minY
	b.maxX = newBoard.maxX
	b.maxY = newBoard.maxY
	b.values = newBoard.values
}

type d struct {
	dx, dy int
}

func makeD(dir string) d {
	x, y := coord(dir)
	return d{x, y}
}

func (b board) countNeighbors(x, y int) int {
	n := 0
	vals := []d{
		makeD("nw"),
		makeD("ne"),
		makeD("e"),
		makeD("se"),
		makeD("sw"),
		makeD("w"),
	}
	for _, val := range vals {
		if val.dx == 0 && val.dy == 0 {
			continue
		}
		if b.Get(x+val.dx, y+val.dy) {
			n++
		}
	}
	return n
}
