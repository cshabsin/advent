package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	Day2a("input.txt")
}

type foo struct {
	a int
	b string
}

func parseFoo(s string) (*foo, error) {
	arr := strings.Split(s, ",")
	a, err := strconv.Atoi(arr[0])
	if err != nil {
		return nil, err
	}
	return &foo{
		a: a,
		b: arr[1],
	}, nil
}

// Day2a solves part 1 of day 2
func Day2a(fn string) {
	ch, err := readinp.Read(fn, parseFoo)
	if err != nil {
		log.Fatal(err)
	}
	for line := range ch {
		line.Get()
	}
}

// Day2b solves part 2 of day 2
func Day2b(fn string) {
	ch, err := readinp.Read(fn, parseFoo)
	if err != nil {
		log.Fatal(err)
	}
	for line := range ch {
		line.Get()
	}
}
