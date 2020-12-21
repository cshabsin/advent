package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var min, max int
	foundIds := map[int]bool{}
	for line := range ch {
		if line.Error != nil {
			log.Fatal(err)
		}

		cont := strings.TrimSpace(*line.Contents)
		seatID := getSeatID(cont)
		if seatID > max {
			max = seatID
		}
		if min == 0 || seatID < min {
			min = seatID
		}
		foundIds[seatID] = true
	}
	for id := min; id < max; id++ {
		if foundIds[id] {
			continue
		}
		if foundIds[id-1] && foundIds[id+1] {
			fmt.Println(id)
		}
	}
}

func getRow(line string) int {
	row := 0
	for i := 0; i < 7; i++ {
		if line[i] == 'B' {
			row = 2*row + 1
		} else {
			row = 2 * row
		}
	}
	return row
}

func getCol(line string) int {
	col := 0
	for i := 7; i < 10; i++ {
		if line[i] == 'R' {
			col = 2*col + 1
		} else {
			col = 2 * col
		}
	}
	return col
}

func getSeatID(line string) int {
	return getRow(line)*8 + getCol(line)
}
