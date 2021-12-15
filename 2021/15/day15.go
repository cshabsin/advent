package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	part1("sample.txt")
	part1("input.txt")
	// part2("sample.txt")
	// fmt.Println("---")
	// part2("input.txt")
}

func part1(fn string) {
	fmt.Println("---", fn, ":")
	_, err := load(fn)
	if err != nil {
		log.Fatal(err)
	}
}

func load(fn string) (string, error) {
	ch, err := readinp.Read(fn, readinp.S)
	if err != nil {
		return "", err
	}
	var acc string
	for line := range ch {
		s, err := line.Get()
		if err != nil {
			return "", err
		}
		acc += s
	}
	return acc, nil
}
