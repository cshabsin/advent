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

func day5a(fn string) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	for d := range ch {
		fmt.Println(d)
	}
}
