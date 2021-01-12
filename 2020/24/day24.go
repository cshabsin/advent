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
	fmt.Println(coord("nwswnese"))
	fmt.Println(coord("nwwswee"))
	b := board{}
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		b.Set(coord(line.Value()))
	}
	fmt.Println(b.values)
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

func (b *board) Set(x, y int) {
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
