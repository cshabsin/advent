package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day5a("sample.txt")
	fmt.Println("---")
	day5a("input.txt")
}

type board struct {
	data [][]int
}

func (b *board) incr(x, y int) {
	for len(b.data) <= y {
		b.data = append(b.data, nil)
	}
	for len(b.data[y]) <= x {
		b.data[y] = append(b.data[y], 0)
	}
	b.data[y][x]++
}

func day5a(fn string) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	var b board
	for line := range ch {
		d, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		if d.x1 != d.x2 && d.y1 != d.y2 {
			// only consider horizontal/vertical lines
			continue
		}
		if d.x1 == d.x2 {
			var first, second int
			if d.y1 < d.y2 {
				first, second = d.y1, d.y2
			} else {
				first, second = d.y2, d.y1
			}
			for y := first; y <= second; y++ {
				b.incr(d.x1, y)
			}
		} else {
			var first, second int
			if d.x1 < d.x2 {
				first, second = d.x1, d.x2
			} else {
				first, second = d.x2, d.x1
			}
			for x := first; x <= second; x++ {
				b.incr(x, d.y1)
			}
		}
	}
	var cnt int
	for _, l := range b.data {
		for _, n := range l {
			if n > 1 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
