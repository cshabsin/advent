package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	dial := 50
	fmt.Println("The dial starts by pointing at", dial)
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	c := 0
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		val := line.Value()
		mul := 1
		if val[0] == 'L' {
			mul = -1
		}
		mag, err := strconv.Atoi(val[1:])
		if err != nil {
			log.Fatal(err)
		}
		var passed int
		dial, c, passed = rotate(dial, c, mul, mag)

		desc := fmt.Sprintf("The dial is rotated %s to point at %d", val, dial)
		if passed != 0 {
			desc += fmt.Sprintf("; during this rotation, it points at 0 %d times.", passed)
		} else {
			desc += "."
		}
		fmt.Printf("%s (%d)\n", desc, c)
	}
	fmt.Println(c)
}

func rotate(dial, c, mul, mag int) (int, int, int) {
	if mul < 0 && dial == 0 {
		dial = 100
	}
	dial += mul * mag
	var passed int
	for dial < 0 {
		passed++
		dial += 100
	}
	for dial > 99 {
		passed++
		dial -= 100
	}
	if (passed == 0 || mul < 0) && dial == 0 {
		passed++
	}
	return dial, c + passed, passed
}
