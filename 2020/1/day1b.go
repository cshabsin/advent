package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func day1b() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	vals := map[int]bool{}
	// map of possible third numbers to product of first two
	seeking := map[int]int{}
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		val, err := strconv.Atoi(strings.TrimSpace(line))
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
