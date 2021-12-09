package main

import "fmt"

func parse(line string) ([]int, error) {
	var d []int
	for _, c := range line {
		if c < '0' || c > '9' {
			return nil, fmt.Errorf("in line %q, illegal character %c", line, c)
		}
		d = append(d, int(c-'0'))
	}
	return d, nil
}
