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

func getBuses(line string) (map[int]bool, error) {
	buses := map[int]bool{}
	for _, s := range strings.Split(line, ",") {
		if s == "x" {
			continue
		}
		bus, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		buses[bus] = true
	}
	return buses, nil
}
