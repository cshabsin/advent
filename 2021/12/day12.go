package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cshabsin/advent/commongen/readinp"
	"github.com/cshabsin/advent/commongen/set"
)

func main() {
	day12("sample.txt", false)
	day12("sample.txt", true)
	fmt.Println("---")
	day12("input.txt", false)
	day12("input.txt", true)
}

func day12(fn string, allowSecondVisit bool) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	caves := map[string][]string{} // map from node to node's children
	for line := range ch {
		np, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		caves[np[0]] = append(caves[np[0]], np[1])
		caves[np[1]] = append(caves[np[1]], np[0])
	}
	fmt.Println(paths(caves, set.Set[string]{}, allowSecondVisit, []string{"start"}))
}

func paths(caves map[string][]string, visited set.Set[string], allowSecondVisit bool, current []string) int {
	curNode := current[len(current)-1]
	var tot int
	for _, child := range caves[curNode] {
		if child == "end" {
			tot++
			continue
		}
		newAllowSecondVisit := allowSecondVisit
		if isSmall(child) && visited.Contains(child) {
			if allowSecondVisit && child != "start" {
				newAllowSecondVisit = false
			} else {
				continue
			}
		}
		vsub := visited.Clone()
		vsub[curNode] = true
		tot += paths(caves, vsub, newAllowSecondVisit, append(current, child))
	}
	return tot
}

func isSmall(cave string) bool {
	return cave[0] >= 'a' && cave[0] <= 'z'
}

type nodePair [2]string

func parse(line string) (nodePair, error) {
	fields := strings.Split(line, "-")
	return nodePair{fields[0], fields[1]}, nil
}
