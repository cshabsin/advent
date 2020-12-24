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
	joltages := map[int]bool{}
	maxJ := 0
	for line := range ch {
		if line.Error != nil {
			log.Fatal(err)
		}
		j, err := strconv.Atoi(strings.TrimSpace(*line.Contents))
		if err != nil {
			log.Fatal(err)
		}
		joltages[j] = true
		if j > maxJ {
			maxJ = j
		}
	}
	joltages[maxJ+3] = true
	diff := 0
	diffCounts := [3]int{}
	for i := 1; i <= maxJ+3; i++ {
		if joltages[i] {
			diffCounts[diff]++
			diff = 0
		} else {
			diff++
		}
	}
	fmt.Println(diffCounts)
	fmt.Println(diffCounts[0] * diffCounts[2])
}
