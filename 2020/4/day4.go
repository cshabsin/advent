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
	passportChan := readPassports(ch)
	valid := 0
	for passport := range passportChan {
		if passport.valid() {
			valid++
		}
	}
	fmt.Println(valid)
}

type passport struct {
	values map[string]string
}

func (p passport) valid() bool {
	for _, key := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if _, ok := p.values[key]; !ok {
			return false
		}
	}
	return true
}

func readPassports(ch chan readinp.Line) chan passport {
	out := make(chan passport)
	go func() {
		p := passport{values: map[string]string{}}
		hasContents := false
		for line := range ch {
			if line.Error != nil {
				log.Fatal(line.Error)
			}
			thisLine := strings.TrimSpace(*line.Contents)
			if thisLine == "" {
				out <- p
				p = passport{values: map[string]string{}}
				hasContents = false
				continue
			}
			hasContents = true
			fields := strings.Fields(thisLine)
			for _, field := range fields {
				splitField := strings.Split(field, ":")
				p.values[splitField[0]] = splitField[1]
			}
		}
		if hasContents {
			out <- p
		}
		close(out)
	}()
	return out
}
