package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	Day3a("input.txt")
	fmt.Println("---")
	// Day2b("input.txt")
}

type foo [12]bool

func parseFoo(s string) (foo, error) {
	var rc foo
	for i, c := range s {
		if c == '1' {
			rc[i] = true
		}
	}
	return rc, nil
}

// Day3a solves part 1 of day 3
func Day3a(fn string) {
	ch, err := readinp.Read(fn, parseFoo)
	if err != nil {
		log.Fatal(err)
	}
	var count int
	var set [12]int
	for line := range ch {
		l, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		count++
		for i := 0; i < 12; i++ {
			if l[i] {
				set[i]++
			}
		}
	}
	var gamma, epsilon int
	for i := 0; i < 12; i++ {
		fmt.Println(set[i], count, count/2)
		if set[i] >= count/2 {
			gamma += 1 << (11 - i)
		} else {
			epsilon += 1 << (11 - i)
		}
	}
	fmt.Println(gamma, epsilon, gamma*epsilon)
}
