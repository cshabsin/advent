package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	// lf, _ := parse("[[[[[9,8],1],2],3],4]")
	// fmt.Println(lf)
	// fmt.Println(reduce(lf, 0, 0))

	// left, _ := parse("[[[[4,3],4],4],[7,[[8,4],9]]]")
	// right, _ := parse("[1,1]")
	// fmt.Printf("%v + %v = %v", left, right, left.add(right))
	part1("sample0.txt")
	part1("sample.txt")
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
	regular     int
	left, right *snailfish
}

func (s snailfish) String() string {
	if !s.isRegular() {
		return fmt.Sprintf("[%v,%v]", s.left, s.right)
	}
	return strconv.Itoa(s.regular)
}

func makeRegular(regular int) *snailfish {
	return &snailfish{regular: regular}
}

func (s snailfish) isRegular() bool {
	return s.left == nil
}

func (s *snailfish) add(t *snailfish) *snailfish {
	rc := &snailfish{left: s, right: t}
	// fmt.Println("after addition:", rc)
	for {
		var changed bool
		rc, _, _, changed = explode(rc, 0, 0, true)
		// fmt.Println("after explode:", rc)
		if changed {
			continue
		}
		rc, changed = split(rc)
		// fmt.Println("after split:", rc)
		if !changed {
			break
		}
	}
	return rc
}

func (s *snailfish) addFromLeft(val int) {
	if s.isRegular() {
		s.regular += val
		return
	}
	s.right.addFromLeft(val)
}

func explode(s *snailfish, depth, rightIn int, doExplode bool) (rc *snailfish, left, right int, changed bool) {
	if s.isRegular() {
		return makeRegular(s.regular + rightIn), 0, 0, rightIn != 0
	}
	if depth == 4 && doExplode {
		// your head a splode
		return makeRegular(0), s.left.regular, s.right.regular, true
	}
	rcleft, left, right, ch := explode(s.left, depth+1, rightIn, doExplode)
	changed = changed || ch
	if ch {
		doExplode = false
	}
	rcRight, rtLeft, right, ch := explode(s.right, depth+1, right, doExplode)
	changed = changed || ch
	rcleft.addFromLeft(rtLeft)

	return &snailfish{left: rcleft, right: rcRight}, left, right, changed
}

func split(s *snailfish) (*snailfish, bool) {
	if s.isRegular() {
		if s.regular >= 10 {
			left := s.regular / 2
			return &snailfish{
				left:  makeRegular(left),
				right: makeRegular(s.regular - left),
			}, true
		}
		return s, false
	}
	left, changed := split(s.left)
	right, ch := split(s.right)
	return &snailfish{
		left:  left,
		right: right,
	}, changed || ch
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
		return &snailfish{
			left:  first,
			right: second,
		}, rest[1:], nil
	}
	regular, err := strconv.Atoi(string(line[0]))
	if err != nil {
		return nil, "", err
	}
	return &snailfish{regular: regular}, line[1:], nil
}
