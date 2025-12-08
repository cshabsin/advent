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

// rotate calculates the new dial position after a rotation.
// It takes the current dial position, a cumulative counter `c`, a direction multiplier `mul` (1 for right, -1 for left),
// and a magnitude `mag`.
// It returns the new dial position, the updated cumulative counter, and the number of times the dial passed 0 during this rotation.
func rotate(dial, c, mul, mag int) (int, int, int) {
	start := dial
	// The total rotation amount.
	rotation := mul * mag

	// Calculate the number of times we pass 0.
	// We use start and rotation to figure this out before adjusting the dial.
	var passed int
	if rotation > 0 {
		// When moving right (positive rotation), we pass 0 every 100 units.
		// The number of passes is the total travel from start divided by 100.
		passed = (start + rotation) / 100
	} else if rotation < 0 {
		// When moving left (negative rotation), we also pass 0.
		// We can think of the dial wrapping from 0 to 99.
		// The number of passes is the total travel from start (in the negative direction) divided by 100.
		passed = (start + rotation - 99) / 100
	}
	newDial := (start + rotation) % 100
	return newDial, c + passed, passed
}
