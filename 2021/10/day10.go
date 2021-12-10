package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day10a("sample.txt")
	fmt.Println("----")
	day10a("input.txt")
}

func day10a(fn string) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	for line := range ch {
		_, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
	}
}
