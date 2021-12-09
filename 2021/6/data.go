package main

import (
	"strings"

	"github.com/cshabsin/advent/commongen/readinp"
)

type data []int

func (d data) nextgen() data {
	d2 := make(data, 9, 9)
	for i := 1; i < 9; i++ {
		d2[i-1] = d[i]
	}
	d2[6] += d[0]
	d2[8] = d[0]
	return d2
}

func (d data) len() int {
	var len int
	for _, cnt := range d {
		len += cnt
	}
	return len
}

func parse(line string) (data, error) {
	d := make(data, 9, 9)
	for _, val := range strings.Split(line, ",") {
		d[readinp.Atoi(val)]++
	}
	return d, nil
}
