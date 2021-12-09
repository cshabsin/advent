package main

import (
	"strings"

	"github.com/cshabsin/advent/commongen/readinp"
)

type data struct {
	x1, y1 int
	x2, y2 int
}

func parse(line string) (*data, error) {
	pairs := strings.Split(line, " -> ")
	first := strings.Split(pairs[0], ",")
	second := strings.Split(pairs[1], ",")
	return &data{
		x1: readinp.Atoi(first[0]),
		y1: readinp.Atoi(first[1]),
		x2: readinp.Atoi(second[0]),
		y2: readinp.Atoi(second[1]),
	}, nil
}
