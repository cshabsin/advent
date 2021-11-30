package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day1b()
}

// Solution to first part of day 1 of Advent of Code 2020.
func day1a() {
	ch, err := readinp.Read[int]("input.txt", strconv.Atoi)
	if err != nil {
		log.Fatal(err)
	}
	vals := map[int]bool{}
	for line := range ch {
		val, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		if vals[val] {
			fmt.Printf("%d found, answer is: %d\n", val, val*(2020-val))
			break
		}
		vals[2020-val] = true
	}
}

// Solution to second part of day 1 of Advent of Code 2020.
func day1b() {
	ch, err := readinp.Read[int]("input.txt", strconv.Atoi)
	if err != nil {
		log.Fatal(err)
	}
	vals := map[int]bool{}
	// map of possible third numbers to product of first two
	seeking := map[int]int{}
	for line := range ch {
		val, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		if prod, found := seeking[val]; found {
			fmt.Printf("%d found, answer is: %d\n", val, val*prod)
			break
		}
		for first := range vals {
			seeking[2020-first-val] = first * val
		}
		vals[val] = true
	}
}
