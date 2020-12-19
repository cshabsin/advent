package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Matches(containers []int, included []bool, desired int) bool {
	total := 0
	for i, v := range(containers) {
		if included[i] {
			total += v
			if total > desired {
				return false
			}
		}
	}
	return total == desired
}

func IncrementIncluded(included []bool) error {
	for i, v := range included {
		included[i] = !v
		if !v {
			return nil
		}
	}
	return io.EOF
}

func main() {
	var infile = flag.String("infile", "input17.txt", "Input file")
	var liters = flag.Int("liters", 150, "Liters of eggnog")
	flag.Parse()

	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	containers := make([]int, 0, 21)
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		size, err := strconv.Atoi(strings.Trim(line, "\n"))
		if err != nil {
			log.Fatal(err)
		}
		containers = append(containers, size)
	}
	included := make([]bool, len(containers))
	var matching int
	matching_by_len := make(map[int]int)
	for {
		if Matches(containers, included, *liters) {
			used := make([]int, 0, len(containers))
			for i, v := range included {
				if v {
					used = append(used, containers[i])
				}
			}
			fmt.Print(used, "\n")
			matching++
			matching_by_len[len(used)] = matching_by_len[len(used)] + 1
		}
		err := IncrementIncluded(included)
		if err == io.EOF {
			break
		}
	}
	fmt.Print(matching_by_len, "\n")
}
