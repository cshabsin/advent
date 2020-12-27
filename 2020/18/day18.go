package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	total := 0
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		val, _, err := calc(line.Value())
		if err != nil {
			fmt.Printf("error parsing line %q: %v\n", line.Value(), err)
		}
		total += val
	}
	fmt.Println(total)
}

func calc(line string) (int, int, error) {
	fmt.Printf("calc(%q)\n", line)
	var off, val int
	operator := '+'
	for {
		var subval, l int
		if line[off] == '(' {
			var err error
			subval, l, err = calc(line[off+1 : len(line)])
			fmt.Printf("line %q: subval %d, len %d\n", line[off+1:len(line)], subval, l)
			if err != nil {
				return 0, 0, err
			}
			l += 2 // add one for the '(' and skip the ')'
		} else {
			var err error
			subval, err = strconv.Atoi(line[off : off+1])
			if err != nil {
				return 0, 0, fmt.Errorf("parsing digit from %q: %v", line, err)
			}
			l = 1
		}
		switch operator {
		case '+':
			fmt.Println("applying + to", val, subval)
			val += subval
		case '-':
			fmt.Println("applying - to", val, subval)
			val -= subval
		case '*':
			fmt.Println("applying * to", val, subval)
			val *= subval
		}
		off += l
		if off == len(line) {
			return val, off, nil
		}
		if line[off] == ')' {
			return val, off, nil
		}
		off++ // skip the space
		operator = rune(line[off])
		off += 2
	}
}
