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
	for line := range ch {
		if line.Error != nil {
			log.Fatal(err)
		}
		fmt.Println(strings.TrimSpace(*line.Contents))
	}
}
