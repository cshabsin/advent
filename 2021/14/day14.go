package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cshabsin/advent/commongen/pair"
	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	in, _ := load("input.txt")
	counts := in.count2("SHBN", 4)
	fmt.Println(counts)
	fmt.Println("---")

	part1("sample.txt")
	part1("input.txt")
	fmt.Println("---")
	part2("sample.txt")
	part2("input.txt")
}

func part1(fn string) {
	in, err := load(fn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("new")
	counts := in.count2(in.template, 10)
	fmt.Println(counts)
	s := in.template
	for i := 0; i < 10; i++ {
		s = in.apply(s)
	}
	fmt.Println("old")
	counts = count(s)
	fmt.Println(counts)
	max, min := maxmin(counts)
	fmt.Println(max - min)
}

func dbg(v ...interface{}) {
	// fmt.Println(v...)
}

func part2(fn string) {
	in, err := load(fn)
	if err != nil {
		log.Fatal(err)
	}
	counts := in.count2(in.template, 40)
	fmt.Println(counts)
	max, min := maxmin(counts)
	fmt.Println(max, min, max-min)
}

type countMap map[rune]int

func (c countMap) add(a countMap) {
	rs := map[rune]bool{}
	for r := range a {
		rs[r] = true
	}
	for r := range c {
		rs[r] = true
	}
	for r := range rs {
		c[r] = a[r] + c[r]
	}
}

func (form in) count2(s string, depth int) countMap {
	cache := map[pair.Pair[string, int]]countMap{}
	counts := countMap{}
	for i, r := range s {
		if i == len(s)-1 {
			counts[r]++
			break
		}
		insert := form.steps[s[i:i+2]]
		counts.add(form.c2h(string(s[i])+insert, depth-1, cache))
		counts.add(form.c2h(insert+string(s[i+1]), depth-1, cache))
	}
	return counts
}

func (form in) c2h(s string, depth int, cache map[pair.Pair[string, int]]countMap) countMap {
	if val := cache[pair.Make(s, depth)]; val != nil {
		return val
	}
	counts := countMap{}
	dbg(s, depth)
	child := string(rune(s[0])) + form.steps[s]
	if depth == 1 {
		dbg("counting", child)
		for _, r := range child {
			counts[r]++
		}
		cache[pair.Make(s, depth)] = counts
		return counts
	}
	counts.add(form.c2h(child, depth-1, cache))
	counts.add(form.c2h(form.steps[s]+string(rune(s[1])), depth-1, cache))
	cache[pair.Make(s, depth)] = counts
	return counts
}

func count(s string) map[rune]int {
	counts := map[rune]int{}
	for _, b := range s {
		counts[b] = counts[b] + 1
	}
	return counts
}

func maxmin(counts map[rune]int) (int, int) {
	min := counts[66]
	max := 0
	for _, cnt := range counts {
		if cnt > max {
			max = cnt
		}
		if cnt < min {
			min = cnt
		}
	}
	return max, min
}

type in struct {
	template string
	steps    map[string]string
}

func (form in) apply(s string) string {
	var t string
	for i := range s {
		if i == len(s)-1 {
			t += string(s[i])
			break
		}
		t += string(s[i])
		if insert, ok := form.steps[s[i:i+2]]; ok {
			t += insert
		}
	}
	return t
}

func load(fn string) (*in, error) {
	ch, err := readinp.Read(fn, readinp.S)
	if err != nil {
		return nil, err
	}
	line := <-ch
	s, err := line.Get()
	if err != nil {
		return nil, err
	}
	rc := &in{template: s, steps: map[string]string{}}
	<-ch // skip a line
	for line := range ch {
		s, err := line.Get()
		if err != nil {
			return nil, err
		}
		fields := strings.Split(s, " -> ")
		rc.steps[fields[0]] = fields[1]
	}
	return rc, nil
}
