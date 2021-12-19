package main

import (
	"fmt"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {

}

func part1(fn string) {
	ch := readinp.MustRead(fn, parse)
	for l := range ch {
		rec := l.MustGet()
		fmt.Println(rec.s)
	}
}

type rec struct {
	s string
}

func parse(line string) (rec, error) {
	return rec{line}, nil
}
