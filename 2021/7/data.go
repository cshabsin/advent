package main

import (
	"log"
	"strconv"
	"strings"
)

type data struct {
	m   map[int]int
	max int
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func parse(line string) (data, error) {
	d := data{
		m: map[int]int{},
	}
	for _, val := range strings.Split(line, ",") {
		v := atoi(val)
		d.m[v]++
		if v > d.max {
			d.max = v
		}
	}
	return d, nil
}

func (d *data) fuel(tgt int) int {
	fuel := 0
	for loc, cnt := range d.m {
		dst := loc - tgt
		if dst < 0 {
			dst = -dst
		}
		fuel += (dst * (dst + 1)) * cnt / 2
	}
	return fuel
}
