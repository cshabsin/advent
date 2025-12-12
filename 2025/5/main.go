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

type rang struct {
	first, last int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("input.txt: %v", err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	fresh := map[rang]bool{}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("read: %v", err)
		}
		line = line[:len(line)-1] // trim the delimiter
		if line == "" {
			break
		}
		vals := strings.Split(line, "-")
		first, err := strconv.Atoi(vals[0])
		if err != nil {
			log.Fatalf("atoi(%s): %v", vals[0], err)
		}
		last, err := strconv.Atoi(vals[1])
		if err != nil {
			log.Fatalf("atoi(%s): %v", vals[1], err)
		}
		addRang(fresh, first, last)
	}
	var available int
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("read: %v", err)
		}
		line = line[:len(line)-1] // trim the delimiter
		ingredient, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("atoi(%s): %v", line, err)
		}
		for r := range fresh {
			if ingredient >= r.first && ingredient <= r.last {
				available++
				break
			}
		}
	}
	fmt.Println(available)
	fmt.Println(countIngredients(fresh), ":", fresh)
}

func addRang(fresh map[rang]bool, first, last int) {
	for {
		var found rang
		for r := range fresh {
			if last < r.first || first > r.last {
				continue
			}
			// They overlap.
			found = r
			if r.first < first {
				first = r.first
			}
			if r.last > last {
				last = r.last
			}
			break
		}
		if found.first == 0 || found.last == 0 {
			// no overlap found
			break
		}
		delete(fresh, found)
	}
	fresh[rang{first, last}] = true
}

func countIngredients(fresh map[rang]bool) int {
	var count int
	for r := range fresh {
		count += r.last - r.first + 1
	}
	return count
}
