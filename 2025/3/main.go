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
	var firstDigit int
	var secondDigit int
	for index, c := range line {
		i := int(c - '0')
		if i > firstDigit && index != len(line)-1 {
			firstDigit = i
			secondDigit = 0
			continue
		}
		if i > secondDigit {
			secondDigit = i
		}
	}
	return firstDigit*10 + secondDigit
}

func joltage12(line string) int {
	var digits [12]byte
	lineIndex := 0
	// Load the first 12 digits
	for lineIndex < 12 {
		digits[lineIndex] = line[lineIndex]
		lineIndex++
	}
	lineIndex = 1
	for lineIndex < len(line) {
		// number of characters left in the line
		lineRemaining := len(line) - lineIndex
		searchStart := 0
		if lineRemaining < 12 {
			searchStart = 12 - lineRemaining
		}
		for digitsIndex := searchStart; digitsIndex < 12; digitsIndex++ {
			if line[lineIndex+digitsIndex-searchStart] > digits[digitsIndex] {
				// copy line[lineIndex] to digits for digitsIndex up to 12
				lineOffset := digitsIndex - searchStart
				for digitsIndex < 12 {
					digits[digitsIndex] = line[lineIndex+lineOffset]
					digitsIndex++
					lineOffset++
				}
				break
			}
		}

		lineIndex++
	}
	var rc int
	for _, d := range digits {
		rc = rc*10 + int(d-'0')
	}
	fmt.Println(line, " became ", string(digits[0:12]), " or", rc)
	return rc
}
