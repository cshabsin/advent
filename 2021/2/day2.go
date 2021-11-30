package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	Day2a("input.txt")
	fmt.Println("---")
	Day2b("input.txt")
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

// Day2a solves part 1 of day 2
func Day2a(fn string) {
	ch, err := readinp.Read(fn, parseFoo)
	if err != nil {
		log.Fatal(err)
	}
	var pos, depth int
	for line := range ch {
		dir, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		switch dir.direction {
		case "forward":
			pos += dir.amount
		case "down":
			depth += dir.amount
		case "up":
			depth -= dir.amount
		default:
			log.Fatal("unknown dir", dir.direction)
		}
	}
	fmt.Println(pos * depth)
}

// Day2b solves part 2 of day 2
func Day2b(fn string) {
	ch, err := readinp.Read(fn, parseFoo)
	if err != nil {
		log.Fatal(err)
	}
	var pos, depth, aim int
	for line := range ch {
		dir, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		switch dir.direction {
		case "forward":
			pos += dir.amount
			depth += aim * dir.amount
		case "down":
			aim += dir.amount
		case "up":
			aim -= dir.amount
		default:
			log.Fatal("unknown dir", dir.direction)
		}
	}
	fmt.Println(pos * depth)
}
