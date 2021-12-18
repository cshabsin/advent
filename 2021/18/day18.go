package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"strconv"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	part1("sample0.txt")
	part1("sample.txt")
	part1("input.txt")
}

func part1(fn string) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	v := <-ch
	tot, err := v.Get()
	if err != nil {
		log.Fatal(err)
	}
	for v := range ch {
		s, err := v.Get()
		if err != nil {
			log.Fatal(err)
		}
		tot = tot.add(s)
	}
	fmt.Println(fn, ":", tot)
}

type snailfish struct {
	regular       int
	parent        *snailfish
	first, second *snailfish
	isFirst       bool
}

func (s snailfish) String() string {
	if !s.isRegular() {
		return fmt.Sprintf("[%v,%v]", s.first, s.second)
	}
	return strconv.Itoa(s.regular)
}

func makeRegular(regular int) *snailfish {
	return &snailfish{regular: regular}
}

func (s snailfish) isRegular() bool {
	return s.first == nil
}

var doDebug = false

func dbg(x ...interface{}) {
	if doDebug {
		fmt.Println(x...)
	}
}

func (s *snailfish) add(t *snailfish) *snailfish {
	if doDebug {
		debug.PrintStack()
	}
	dbg("---", s, "+", t, ":")
	rc := &snailfish{
		first:   s,
		second:  t,
		parent:  s.parent, // do these matter? we only ever "add" top-level values
		isFirst: s.isFirst,
	}
	dbg("after addition:", rc)
	for {
		var changed bool
		rc, changed = explode(rc, 0, true)
		if changed {
			dbg("after explode:", rc)
			continue
		}
		rc, changed = split(rc)
		dbg("after split:", rc)
		if !changed {
			break
		}
	}
	return rc
}

func (s *snailfish) addToLeft(val int) {
	if s == nil || s.parent == nil {
		return
	}
	if s.isFirst {
		s.parent.addToLeft(val)
		return
	}
	// start with the left sibling
	current := s.parent.first
	for {
		if current == nil {
			return
		}
		if current.isRegular() {
			current.regular += val
			return
		}
		// go down the right-side path
		current = current.second
	}
}

func (s *snailfish) addToRight(val int) {
	if s == nil || s.parent == nil {
		return
	}
	if !s.isFirst {
		s.parent.addToRight(val)
		return
	}
	current := s.parent.second
	for {
		if current == nil {
			return
		}
		if current.isRegular() {
			current.regular += val
			return
		}
		// go down the left-side path
		current = current.first
	}
}

func (s *snailfish) top() *snailfish {
	if s.parent == nil {
		return s
	}
	return s.parent.top()
}

func explode(s *snailfish, depth int, doExplode bool) (rc *snailfish, changed bool) {
	if s.isRegular() {
		return s, false
	}
	if depth == 4 && doExplode {
		// your head a splode
		// fmt.Println("before explode", s.top())
		s.addToLeft(s.first.regular)
		// fmt.Println("after addToLeft", s.top())
		s.addToRight(s.second.regular)
		// fmt.Println("after addToRight", s.top())
		rc := makeRegular(0)
		rc.parent = s.parent
		rc.isFirst = s.isFirst
		return rc, true
	}
	rcFirst, changed := explode(s.first, depth+1, doExplode)
	if changed {
		rc := &snailfish{
			first:   rcFirst,
			second:  s.second,
			parent:  s.parent,
			isFirst: s.isFirst,
		}
		rc.first.parent = rc
		rc.first.isFirst = true
		rc.second.parent = rc
		return rc, true
	}

	rcSecond, changed := explode(s.second, depth+1, doExplode)

	rc = &snailfish{
		first:   rcFirst,
		second:  rcSecond,
		parent:  s.parent,
		isFirst: s.isFirst,
	}
	rc.first.parent = rc
	rc.first.isFirst = true
	rc.second.parent = rc

	return rc, changed
}

func split(s *snailfish) (*snailfish, bool) {
	if s.isRegular() {
		if s.regular >= 10 {
			left := s.regular / 2
			rc := &snailfish{
				first:   makeRegular(left),
				second:  makeRegular(s.regular - left),
				parent:  s.parent,
				isFirst: s.isFirst,
			}
			rc.first.parent = rc
			rc.first.isFirst = true
			rc.second.parent = rc
			return rc, true
		}
		return s, false
	}
	left, changed := split(s.first)
	if changed {
		rc := &snailfish{
			first:   left,
			second:  s.second,
			parent:  s.parent,
			isFirst: s.isFirst,
		}
		rc.first.parent = rc
		rc.first.isFirst = true
		rc.second.parent = rc
		return rc, true
	}
	right, ch := split(s.second)
	rc := &snailfish{
		first:   left,
		second:  right,
		parent:  s.parent,
		isFirst: s.isFirst,
	}
	rc.first.parent = rc
	rc.first.isFirst = true
	rc.second.parent = rc
	return rc, changed || ch
}

func parse(line string) (*snailfish, error) {
	s, rest, err := parsePartial(line)
	if err != nil {
		return nil, err
	}
	if rest != "" {
		return nil, fmt.Errorf("line %q had remainder %q", line, rest)
	}
	return s, nil
}

func parsePartial(line string) (*snailfish, string, error) {
	if line[0] == '[' {
		first, rest, err := parsePartial(line[1:])
		if err != nil {
			return nil, "", err
		}
		if rest[0] != ',' {
			return nil, "", fmt.Errorf("parsePartial on line %q expected comma after parsing first, got %q", line, rest)
		}
		second, rest, err := parsePartial(rest[1:])
		if err != nil {
			return nil, "", err
		}
		if rest[0] != ']' {
			return nil, "", fmt.Errorf("parsePartial on line %q expected close bracket after parsing second, got %q", line, rest)
		}
		s := &snailfish{
			first:  first,
			second: second,
		}
		first.parent = s
		first.isFirst = true
		second.parent = s
		return s, rest[1:], nil
	}
	regular, err := strconv.Atoi(string(line[0]))
	if err != nil {
		return nil, "", err
	}
	return &snailfish{regular: regular}, line[1:], nil
}
