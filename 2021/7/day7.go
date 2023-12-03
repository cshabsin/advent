package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day7a("sample.txt")
	// day6b("sample.txt")
	fmt.Println("---")
	day7a("input.txt")
	// day6b("input.txt")
}

func day7a(fn string) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	line := <-ch
	dat, err := line.Get()
	if err != nil {
		log.Fatal(err)
	}
	var best int
	bestFuel := 99999999
	for i := 0; i < dat.max; i++ {
		f := dat.fuel(i)
		if f < bestFuel {
			best = i
			bestFuel = f
		}
	}
	fmt.Println(best, bestFuel)
}
