package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	c := 0
	c12 := 0
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		c += joltage(line.Value())
		c12 += joltage12(line.Value())
	}
	fmt.Println(c)
	fmt.Println(c12)
}

func joltage(line string) int {
	var firstDigit, firstIdx int
	// Find the largest digit in the string, excluding the last character.
	for i := 0; i < len(line)-1; i++ {
		val := int(line[i] - '0')
		if val > firstDigit {
			firstDigit = val
			firstIdx = i
		}
	}

	var secondDigit int
	// Find the largest digit after the first digit.
	for i := firstIdx + 1; i < len(line); i++ {
		val := int(line[i] - '0')
		if val > secondDigit {
			secondDigit = val
		}
	}

	return firstDigit*10 + secondDigit
}

func joltage12(line string) int {
	// Find the lexicographically largest substring of length 12.
	maxSub := line[:12]
	for i := 1; i <= len(line)-12; i++ {
		if sub := line[i : i+12]; sub > maxSub {
			maxSub = sub
		}
	}

	digits := []byte(maxSub)

	// Check if any suffix of the line (length < 12) can improve the end of our number.
	for i := len(line) - 11; i < len(line); i++ {
		suffix := line[i:]
		offset := 12 - len(suffix)
		for j := 0; j < len(suffix); j++ {
			if suffix[j] > digits[offset+j] {
				copy(digits[offset+j:], suffix[j:])
				break
			} else if suffix[j] < digits[offset+j] {
				break
			}
		}
	}
	var rc int
	for _, d := range digits {
		rc = rc*10 + int(d-'0')
	}
	return rc
}
