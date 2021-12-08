package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day8a("sample.txt")
	// day6b("sample.txt")
	fmt.Println("---")
	day8a("input.txt")
	// day6b("input.txt")
}

func day8a(fn string) {
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
