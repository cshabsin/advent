package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/cshabsin/advent/2020/console"
	"github.com/cshabsin/advent/common/readinp"
)

var bagRuleRegex = regexp.MustCompile(`^([a-z ]*) bags contain (.*)$`)
var subbagRegex = regexp.MustCompile(`(\d) ([a-z ]*) bag`)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	cons := console.New()
	for line := range ch {
		if line.Error != nil {
			log.Fatal(err)
		}
		if err := cons.ReadInstruction(strings.TrimSpace(*line.Contents)); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(cons.Run(0, nil))
}
