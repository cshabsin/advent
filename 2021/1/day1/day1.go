package day1

import (
	"log"
	"strconv"

	"github.com/cshabsin/advent/commongen/readinp"
)

func Day1a() {
	ch, err := readinp.Read("input.txt", strconv.Atoi)
	if err != nil {
		log.Fatal(err)
	}
	for line := range ch {
		_, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
	}
}
