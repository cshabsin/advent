package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt", readinp.S)
	if err != nil {
		log.Fatal(err)
	}
	joltages := map[int]bool{0: true}
	maxJ := 0
	for line := range ch {
		contents, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		j, err := strconv.Atoi(strings.TrimSpace(contents))
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

	fmt.Println(allPermutations(joltages, maxJ+3))

	for i := 1; i < 15; i++ {
		fmt.Println(i, ":", calcPermsHard(i))
	}
}

var perms = []int{1, 1, 2, 4}

func calcPerms(runLength int) int {
	if runLength == 0 {
		return 1
	}
	if len(perms) < runLength {
		for i := len(perms); i < runLength; i++ {
			perms = append(perms, perms[i-1]+perms[i-2]+perms[i-3])
		}
	}
	return perms[runLength-1]
}

func calcPermsHard(runLength int) int {
	if runLength <= 1 {
		return 1
	}
	if runLength == 2 {
		return 2
	}
	if runLength == 3 {
		return 4
	}
	return calcPermsHard(runLength-1) + calcPermsHard(runLength-2) + calcPermsHard(runLength-3)
}

func allPermutations(joltages map[int]bool, maxJ int) int {
	runLength := 0
	permutations := 1
	for i := 0; i <= maxJ; i++ {
		if joltages[i] {
			runLength++
		} else {
			fmt.Println("runlength", runLength, "to", i-1, "perms", calcPermsHard(runLength-1))
			permutations *= calcPermsHard(runLength - 1)
			runLength = 0
		}
	}
	return permutations
}
