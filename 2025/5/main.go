// Package main solves an Advent of Code style puzzle involving ranges and ingredients.
// It reads a list of allowed ranges, merges them, and then checks a list of ingredients
// to see how many fall within the allowed ranges.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Range represents an inclusive interval [Start, End].
type Range struct {
	Start, End int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input.txt: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ranges := make(map[Range]bool)

	// Part 1: Parse and merge ranges.
	// The input format expects ranges (e.g., "5-10") until an empty line.
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			log.Fatalf("malformed range line: %q", line)
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("invalid start value in line %q: %v", line, err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("invalid end value in line %q: %v", line, err)
		}

		addRange(ranges, Range{Start: start, End: end})
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading ranges: %v", err)
	}

	// Part 2: Count available ingredients.
	// The rest of the file contains one ingredient (integer) per line.
	var availableCount int
	for scanner.Scan() {
		line := scanner.Text()
		ingredient, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("invalid ingredient line %q: %v", line, err)
		}

		for r := range ranges {
			if ingredient >= r.Start && ingredient <= r.End {
				availableCount++
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading ingredients: %v", err)
	}

	fmt.Println(availableCount)
	fmt.Printf("%d : %v\n", countCovered(ranges), ranges)
}

// addRange adds a new range to the set, merging it with any existing overlapping ranges.
func addRange(ranges map[Range]bool, newRange Range) {
	for {
		var overlapped Range
		found := false
		for r := range ranges {
			// Check for overlap: [Start, End] overlaps with [r.Start, r.End]
			// if Start <= r.End && r.Start <= End.
			if newRange.Start <= r.End && r.Start <= newRange.End {
				overlapped = r
				found = true
				break
			}
		}

		if !found {
			break
		}

		// Merge the found overlapping range into newRange.
		if overlapped.Start < newRange.Start {
			newRange.Start = overlapped.Start
		}
		if overlapped.End > newRange.End {
			newRange.End = overlapped.End
		}
		// Remove the old unmerged range.
		delete(ranges, overlapped)
	}
	ranges[newRange] = true
}

// countCovered calculates the total number of integers covered by the ranges.
func countCovered(ranges map[Range]bool) int {
	var count int
	for r := range ranges {
		count += r.End - r.Start + 1
	}
	return count
}
