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
	c := counter{vals: map[int]bool{}}
	for line := range ch {
		if line.Error != nil {
			log.Fatal(err)
		}
		cont := strings.TrimSpace(*line.Contents)
		if cont == "" {
			total += c.total
			c = counter{vals: map[int]bool{}}
			continue
		}
		c.count(cont)
	}
	total += c.total
	fmt.Println(total)
}

type counter struct {
	vals  map[int]bool
	total int
}

func (c *counter) count(cont string) {
	for _, ch := range cont {
		index := int(ch) - int('a')
		if c.vals[index] {
			continue
		}
		c.vals[index] = true
		c.total++
	}
}
