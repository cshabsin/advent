package main

import (
	"log"
	"strconv"
	"strings"
)

type data struct {
	x1, y1 int
	x2, y2 int
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func parse(line string) (*data, error) {
	pairs := strings.Split(line, " -> ")
	first := strings.Split(pairs[0], ",")
	second := strings.Split(pairs[1], ",")
	return &data{
		x1: atoi(first[0]),
		y1: atoi(first[1]),
		x2: atoi(second[0]),
		y2: atoi(second[1]),
	}, nil
}
