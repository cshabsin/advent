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
	values := map[int]map[int]bool{}
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		x, y := coord(line.Value())
		if values[x] == nil {
			values[x] = map[int]bool{}
		}
		values[x][y] = !values[x][y]
	}
	fmt.Println(values)
	count := 0
	for _, row := range values {
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
