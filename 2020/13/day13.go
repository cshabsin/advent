package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	line := <-ch
	if line.Error != nil {
		log.Fatal(err)
	}
	minTime, err := strconv.Atoi(strings.TrimSpace(*line.Contents))
	if err != nil {
		log.Fatal(err)
	}
	line = <-ch
	if line.Error != nil {
		log.Fatal(err)
	}
	buses, err := getBuses(strings.TrimSpace(*line.Contents))
	if err != nil {
		log.Fatal(err)
	}
	day13a(minTime, buses)
}

func day13a(minTime int, buses map[int]int) {
	t := minTime
	id := 0
outer:
	for {
		for bus := range buses {
			if t%bus == 0 {
				id = bus
				break outer
			}
		}
		t++
	}
	fmt.Println("minTime", minTime)
	fmt.Println("t", t)
	fmt.Println("id", id)
	fmt.Println("wait*id", (t-minTime)*id)
}

func getBuses(line string) (map[int]int, error) {
	buses := map[int]int{}
	i := 0
	for _, s := range strings.Split(line, ",") {
		if s == "x" {
			i++
			continue
		}
		bus, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		buses[bus] = i
		i++
	}
	return buses, nil
}
