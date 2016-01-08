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

func LineMatches(target map[string]int, line string) bool {
	skip_name := line[strings.Index(line, ": ")+2:]
	entries := strings.Split(strings.Trim(skip_name, "\n"), ", ")
	for _, entry := range entries {
		colon_index := strings.Index(entry, ": ")
		key := entry[:colon_index]
		value, err := strconv.Atoi(entry[colon_index+2:])
		if err != nil {
			log.Fatal(err)
		}
		if target_value, found := target[key]; found {
			if key == "cats" || key == "trees" {
				if target_value >= value {
					return false
				}
			} else if key == "pomeranians" || key == "goldfish" {
				if target_value <= value {
					return false
				}
			} else if target_value != value {
				return false
			}
		}
	}
	return true
}

func main() {
	var infile = flag.String("infile", "input16.txt", "Input file")
	flag.Parse()

	target := map[string]int{"children": 3,
		"cats": 7,
		"samoyeds": 2,
		"pomeranians": 3,
		"akitas": 0,
		"vizslas": 0,
		"goldfish": 5,
		"trees": 3,
		"cars": 2,
		"perfumes": 1}

	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if LineMatches(target, line) {
			fmt.Print(line)
		}
	}
}
