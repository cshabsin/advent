package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	part1("sample.txt")
	part2("sample.txt")
	fmt.Println("---")
	part1("input.txt")
	part2("input.txt")
}

func part1(fn string) {
	in, err := load(fn)
	if err != nil {
		log.Fatal(err)
	}
	s := in.template
	for i := 0; i < 10; i++ {
		s = in.apply(s)
	}
	counts := map[rune]int{}
	for _, b := range s {
		counts[b] = counts[b] + 1
	}
	min := 9999999
	max := 0
	var minR, maxR rune
	for r, cnt := range counts {
		if cnt > max {
			maxR = r
			max = cnt
		}
		if cnt < min {
			minR = r
			min = cnt
		}
	}
	fmt.Println(max, maxR, min, minR, max-min)
}

func part2(fn string) {

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
