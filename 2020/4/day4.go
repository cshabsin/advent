package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	day4b()
}

func day4b() {
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

var (
	yrRegex    = regexp.MustCompile(`^\d{4}$`)
	hclRegex   = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	pidRegex   = regexp.MustCompile(`^\d{9}$`)
	hgtCmRegex = regexp.MustCompile(`^(\d*)cm$`)
	hgtInRegex = regexp.MustCompile(`^(\d*)in$`)
)

func inRange(value string, min, max int) bool {
	if !yrRegex.MatchString(value) {
		return false
	}
	if val, err := strconv.Atoi(value); err != nil {
		return false
	} else if val < min || val > max {
		return false
	}
	return true
}

func isValidHeight(hgt string) bool {
	if matches := hgtCmRegex.FindStringSubmatch(hgt); matches != nil {
		hCm, err := strconv.Atoi(matches[1])
		if err != nil {
			return false
		}
		if hCm < 150 || hCm > 193 {
			return false
		}
	} else if matches := hgtInRegex.FindStringSubmatch(hgt); matches != nil {
		hIn, err := strconv.Atoi(matches[1])
		if err != nil {
			return false
		}
		if hIn < 59 || hIn > 76 {
			return false
		}
	} else {
		return false
	}
	return true
}

func (p passport) valid() bool {
	for _, key := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if _, ok := p.values[key]; !ok {
			return false
		}
	}

	if !inRange(p.values["byr"], 1920, 2002) {
		return false
	}

	if !inRange(p.values["iyr"], 2010, 2020) {
		return false
	}

	if !inRange(p.values["eyr"], 2020, 2030) {
		return false
	}

	if !isValidHeight(p.values["hgt"]) {
		return false
	}

	if !hclRegex.MatchString(p.values["hcl"]) {
		return false
	}

	if !map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}[p.values["ecl"]] {
		return false
	}

	if !pidRegex.MatchString(p.values["pid"]) {
		return false
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
