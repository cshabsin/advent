package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	last25 := [25]int{}
	last25set := map[int]bool{}
	i := 0
	initial := true
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		val, err := strconv.Atoi(strings.TrimSpace(*line.Contents))
		if err != nil {
			log.Fatal(err)
		}
		if initial {
			last25[i] = val
			last25set[val] = true
			i++
			if i == 25 {
				initial = false
				i = 0
			}
		} else {
			valid := false
			for v := range last25set {
				if valid {
					break
				}
				for v2 := range last25set {
					if v == v2 {
						continue
					}
					if v+v2 == val {
						valid = true
						break
					}
				}
			}
			if !valid {
				fmt.Println(val)
				return
			}
			last25set[last25[i]] = false
			last25set[val] = true
			last25[i] = val
			i = (i + 1) % 25
		}
	}
	fmt.Println("none?")
}
