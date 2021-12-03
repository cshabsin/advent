package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	Day3a("sample.txt")
	fmt.Println("---")
	Day3b("sample.txt")
	fmt.Println("---")
	Day3a("input.txt")
	fmt.Println("---")
	Day3b("input.txt")
}

type foo []bool

func parseFoo(s string) (foo, error) {
	var rc foo
	for _, c := range s {
		rc = append(rc, c == '1')
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
	var set []int
	for line := range ch {
		l, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		count++
		if set == nil {
			set = make([]int, len(l), len(l))
		}
		for i := 0; i < len(l); i++ {
			if l[i] {
				set[i]++
			}
		}
	}
	var gamma, epsilon int
	for i := 0; i < len(set); i++ {
		if set[i] >= count/2 {
			gamma += 1 << (len(set) - 1 - i)
		} else {
			epsilon += 1 << (len(set) - 1 - i)
		}
	}
	fmt.Println(gamma, epsilon, gamma*epsilon)
}

// Day3b solves part 2 of day 3
func Day3b(fn string) {
	ch, err := readinp.Read(fn, parseFoo)
	if err != nil {
		log.Fatal(err)
	}
	var ratings []foo
	for line := range ch {
		l, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		ratings = append(ratings, l)
	}
	oxyRating := calcRating(ratings, true)
	co2Rating := calcRating(ratings, false)
	fmt.Println(oxyRating, co2Rating, oxyRating*co2Rating)
}

func calcRating(ratings []foo, oxy bool) int {
	for i := 0; i < len(ratings[0]); i++ {
		ratings = prune(ratings, i, oxy)
		if len(ratings) == 1 {
			break
		}
	}
	var rating int
	for i := 0; i < len(ratings[0]); i++ {
		if ratings[0][i] {
			rating += 1 << (len(ratings[0]) - 1 - i)
		}
	}
	return rating
}

func prune(ratings []foo, bit int, greater bool) []foo {
	var set, unset int
	for _, rating := range ratings {
		if rating[bit] {
			set++
		} else {
			unset++
		}
	}
	var filterBit bool
	if greater {
		filterBit = set >= unset
	} else {
		filterBit = set < unset
	}
	fmt.Println(filterBit)
	var newratings []foo
	for _, rating := range ratings {
		if rating[bit] == filterBit {
			newratings = append(newratings, rating)
		}
	}
	return newratings
}
