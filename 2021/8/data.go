package main

import (
	"log"
	"sort"
	"strings"
)

type data struct {
	line     string
	patterns []string
	output   []string
}

func parse(line string) (data, error) {
	var d data
	d.line = line
	parts := strings.Split(line, "|")
	for _, out := range strings.Split(parts[0], " ") {
		d.patterns = append(d.patterns, strings.TrimSpace(out))
	}
	for _, out := range strings.Split(parts[1], " ") {
		d.output = append(d.output, strings.TrimSpace(out))
	}
	return d, nil
}

var normalMap = [][]int{
	{0, 1, 2, 4, 5, 6},
	{2, 5},
	{0, 2, 3, 4, 6},
	{0, 2, 3, 5, 6},
	{1, 2, 3, 5},
	{0, 1, 3, 5, 6},
	{0, 1, 3, 4, 5, 6},
	{0, 2, 5},
	{0, 1, 2, 3, 4, 5, 6},
	{0, 1, 2, 3, 5, 6},
}

// segment -> a-g
type mapping map[rune]int

func (m mapping) translate(s string) int {
	set := map[int]bool{}
	for _, r := range s {
		set[m[r]] = true
	}
	var sortedSegments []int
	for seg := range set {
		sortedSegments = append(sortedSegments, seg)
	}
	sort.Ints(sortedSegments)
	for i, iSegs := range normalMap {
		if sliceEq(sortedSegments, iSegs) {
			return i
		}
	}
	log.Fatalf("whoa: %q, %v", s, m)
	return 0
}

func sliceEq[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (d data) allEntries() []string {
	return append(d.output, d.patterns...)
}

func (d data) getMapping() mapping {
	// map of signal to possible target segments
	possibilities := map[rune]map[int]bool{}
	for i := 'a'; i < 'h'; i++ {
		possibilities[i] = map[int]bool{}
		for j := 0; j < 7; j++ {
			possibilities[i][j] = true
		}
	}
	for _, ent := range d.allEntries() {
		switch len(ent) {
		case 2:
			// only 1
			for i := 0; i < len(ent); i++ {
				possibilities[rune(ent[i])] = setOnly(possibilities[rune(ent[i])], makeMap(normalMap[1]...))
			}
		case 3:
			// only 7
			for i := 0; i < len(ent); i++ {
				possibilities[rune(ent[i])] = setOnly(possibilities[rune(ent[i])], makeMap(normalMap[7]...))
			}
		case 4:
			// only 4
			for i := 0; i < len(ent); i++ {
				possibilities[rune(ent[i])] = setOnly(possibilities[rune(ent[i])], makeMap(normalMap[4]...))
			}
		case 5:
			// 2, 3, 5
			unset := getUnsetSignals(ent)
			// the unset ones can only be segments 1 (in rendering 2 or 3), 2 (5), 4 (3, 5), or 5 (2)
			delete(possibilities[unset[0]], 0)
			delete(possibilities[unset[0]], 3)
			delete(possibilities[unset[0]], 6)
			delete(possibilities[unset[1]], 0)
			delete(possibilities[unset[1]], 3)
			delete(possibilities[unset[1]], 6)
		case 6:
			// 0, 6, 9
			unset := getUnsetSignals(ent)
			// the unset one can only be segment 2 (in rendering 6), 3 (0), or 4 (9)
			delete(possibilities[unset[0]], 0)
			delete(possibilities[unset[0]], 1)
			delete(possibilities[unset[0]], 5)
			delete(possibilities[unset[0]], 6)
		case 7:
			// only 8
			for i := 0; i < len(ent); i++ {
				possibilities[rune(ent[i])] = setOnly(possibilities[rune(ent[i])], makeMap(normalMap[8]...))
			}
		}
	}
	// set of found segments
	found := map[int]rune{}
	for {
		for r, vals := range possibilities {
			if len(vals) == 1 {
				found[getTheVal(vals)] = r
				continue
			}
			newVals := map[int]bool{}
			for v := range vals {
				// only include values that haven't already been found
				if _, ok := found[v]; !ok {
					newVals[v] = true
				}
			}
			possibilities[r] = newVals
		}
		if len(found) == 7 {
			break
		}
	}
	rc := map[rune]int{}
	for r, seg := range possibilities {
		rc[r] = getTheVal(seg)
	}
	return rc
}

func getTheVal(m map[int]bool) int {
	for k := range m {
		return k
	}
	log.Fatal("no val?")
	return 0
}

func makeMap[T comparable](vs ...T) map[T]bool {
	m := map[T]bool{}
	for _, v := range vs {
		m[v] = true
	}
	return m
}

func setOnly(m map[int]bool, only map[int]bool) map[int]bool {
	r := map[int]bool{}
	for v := range m {
		if only[v] {
			r[v] = true
		}
	}
	return r
}

func getUnsetSignals(ent string) []rune {
	signals := map[rune]bool{}
	for _, e := range ent {
		signals[e] = true
	}
	var unset []rune
	for r := 'a'; r <= 'g'; r++ {
		if !signals[r] {
			unset = append(unset, r)
		}
	}
	return unset
}
