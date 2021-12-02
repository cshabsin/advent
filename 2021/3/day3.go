package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	Day3a("input.txt")
	fmt.Println("---")
	// Day2b("input.txt")
}

type foo struct {
	direction string
	amount    int
}

func parseFoo(s string) (*foo, error) {
	arr := strings.Split(s, " ")
	a, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, err
	}
	return &foo{
		direction: arr[0],
		amount:    a,
	}, nil
}

// Day3a solves part 1 of day 3
func Day3a(fn string) {
	ch, err := readinp.Read(fn, parseFoo)
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
