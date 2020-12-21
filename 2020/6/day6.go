package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	c := counter{firstLine: true, vals: map[int]bool{}}
	for line := range ch {
		if line.Error != nil {
			log.Fatal(err)
		}
		cont := strings.TrimSpace(*line.Contents)
		if cont == "" {
			total += c.total
			c = counter{firstLine: true, vals: map[int]bool{}}
			continue
		}
		c.count(cont)
	}
	total += c.total
	fmt.Println(total)
}

type counter struct {
	firstLine bool
	vals      map[int]bool
	total     int
}

func (c *counter) count(cont string) {
	if c.firstLine {
		for _, ch := range cont {
			index := int(ch) - int('a')
			if c.vals[index] {
				continue
			}
			c.vals[index] = true
			c.total++
		}
		c.firstLine = false
		return
	}
	thisLineVals := map[int]bool{}
	for _, ch := range cont {
		index := int(ch) - int('a')
		if thisLineVals[index] {
			continue
		}
		thisLineVals[index] = true
	}
	for current, stillValid := range c.vals {
		if stillValid && !thisLineVals[current] {
			c.vals[current] = false
			c.total--
		}
	}
}
