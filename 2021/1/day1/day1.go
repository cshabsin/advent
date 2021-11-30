package day1

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cshabsin/advent/commongen/readinp"
)

func Day1a() {
	ch, err := readinp.Read("input.txt", strconv.Atoi)
	if err != nil {
		log.Fatal(err)
	}
	var increases int
	prev := 193 // first one from input.txt
	for line := range ch {
		val, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		if val > prev {
			increases++
		}
		prev = val
	}
	fmt.Println(increases)
}

func Day1b() {
	ch, err := readinp.Read("input.txt", strconv.Atoi)
	if err != nil {
		log.Fatal(err)
	}
	var increases int
	var a []int
	for i := 0; i < 3; i++ {
		line := <-ch
		val, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, val)
	}
	prev := add(a)
	var index int
	for line := range ch {
		val, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		a[index] = val
		index = (index + 1) % 3
		cur := add(a)
		if cur > prev {
			increases++
		}
		prev = cur
	}
	fmt.Println(increases)
}

func add(a []int) int {
	fmt.Println(a)
	return a[0] + a[1] + a[2]
}
