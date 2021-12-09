package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day9a("sample.txt")
	// day9b("sample.txt")
	fmt.Println("---")
	day9a("input.txt")
	// day9b("input.txt")
}

func day9a(fn string) {
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

func day9b(fn string) {
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
