package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	day10a("sample.txt")
	day10b("sample.txt")
	fmt.Println("----")
	day10a("input.txt")
	day10b("input.txt")
}

func day10a(fn string) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	var total int
	for line := range ch {
		score, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		total += score
	}
	fmt.Println(total)
}

func day10b(fn string) {
	ch, err := readinp.Read(fn, parse2)
	if err != nil {
		log.Fatal(err)
	}
	var completions []int
	for line := range ch {
		state, err := line.Get()
		if err != nil {
			continue
		}
		if state == "" {
			continue
		}
		fmt.Println(state, state.completion())
		completions = append(completions, state.completion())
	}
	sort.Sort(sort.IntSlice(completions))
	fmt.Println(completions, len(completions), completions[len(completions)/2])
}
