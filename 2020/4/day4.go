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
	for passport := range passportChan {
		fmt.Println(passport.values)
	}
}

type passport struct {
	values map[string]string
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
