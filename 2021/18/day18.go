package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	part1("sample.txt")
}

func part1(fn string) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	for v := range ch {
		s, err := v.Get()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)
	}
}

type snailfish struct {
	regular     int
	left, right *snailfish
}

func (s snailfish) String() string {
	if s.left != nil {
		return fmt.Sprintf("[%v,%v]", s.left, s.right)
	}
	return strconv.Itoa(s.regular)
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
