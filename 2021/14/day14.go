package main

import (
	"fmt"
	"log"

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

func load(fn string) (string, error) {
	ch, err := readinp.Read(fn, readinp.S)
	if err != nil {
		return "", err
	}
	var tot string
	for line := range ch {
		s, err := line.Get()
		if err != nil {
			return "", err
		}
		tot += s + "\n"
	}
	return tot, nil
}
