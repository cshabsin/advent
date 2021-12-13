package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	part1("sample.txt")
	part2("sample.txt")
	fmt.Println("---")
	part1("input.txt")
	part2("input.txt")
}

func part1(fn string) {
	_, err := load(fn)
	if err != nil {
		log.Fatal(err)
	}
}

func part2(fn string) {

}

type pair struct {
	a, b string
}

func (p pair) first() string {
	return p.a
}

func (p pair) second() string {
	return p.b
}

func from(s string) pair {
	fields := strings.Split(s, " -> ")
	return pair{fields[0], fields[1]}
}

type in struct {
	template string
	steps    []pair
}

func load(fn string) (*in, error) {
	ch, err := readinp.Read(fn, readinp.S)
	if err != nil {
		return nil, err
	}
	line := <-ch
	s, err := line.Get()
	if err != nil {
		return nil, err
	}
	rc := &in{template: s}
	<-ch // skip a line
	for line := range ch {
		s, err := line.Get()
		if err != nil {
			return nil, err
		}
		rc.steps = append(rc.steps, from(s))
	}
	return rc, nil
}
