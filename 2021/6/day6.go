package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day6("sample.txt", 80)
	day6("sample.txt", 256)
	fmt.Println("---")
	day6("input.txt", 80)
	day6("input.txt", 256)
}

func day6(fn string, iter int) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	line := <-ch
	data, err := line.Get()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < iter; i++ {
		data = data.nextgen()
	}
	fmt.Println(data.len())
}
