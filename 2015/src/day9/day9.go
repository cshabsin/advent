package main
/// BUGGY

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

func AddArcs(arcs map[string]map[string]int, locations map[string]bool, line string) {
	tokens := strings.Split(strings.Trim(line, "\n"), " ")
	length, err := strconv.ParseInt(tokens[4], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	if arcs[tokens[0]] == nil {
		arcs[tokens[0]] = make(map[string]int)
	}
	arcs[tokens[0]][tokens[2]] = int(length)
	if arcs[tokens[2]] == nil {
		arcs[tokens[2]] = make(map[string]int)
	}
	arcs[tokens[2]][tokens[0]] = int(length)
	locations[tokens[0]] = true
	locations[tokens[2]] = true
}

func find_shortest_path(arcs map[string]map[string]int, locations map[string]bool) int {
	shortest_length := 0
	for location, _ := range locations {
		new_locations := make(map[string]bool)
		for k, _ := range locations {
			if k != location {
				new_locations[k] = true
			}
		}
		current_length := find_shortest_path_internal(
			location, arcs, new_locations)
		if shortest_length == 0 || current_length > shortest_length {
			shortest_length = current_length
		}
	}
	return shortest_length
}

func find_shortest_path_internal(source string, arcs map[string]map[string]int, locations map[string]bool) int {
	for k, _ := range locations {
		if k == source {
			log.Fatal("Found source ", source, " in locations ", locations)
		}
	}
        if len(locations) == 1 {
		for location := range locations {
			return arcs[source][location]
		}
	}
	shortest_len := 0
	for location, _ := range locations {
		new_locations := make(map[string]bool)
		for k, _ := range locations {
			if k != location {
				new_locations[k] = true
			}
		}
		current_len := find_shortest_path_internal(
			location, arcs, new_locations)
		if shortest_len == 0 || current_len > shortest_len {
			shortest_len = current_len + arcs[source][location]
		}
	}
	return shortest_len
}

func main() {
	var infile = flag.String("infile", "input9.txt", "Input filename")
	flag.Parse()
	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	locations := make(map[string]bool, 0)
	arcs := make(map[string]map[string]int)
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		AddArcs(arcs, locations, line)
	}
	length := find_shortest_path(arcs, locations)
	fmt.Print(locations, "\n")
	fmt.Print(length, "\n")
}
