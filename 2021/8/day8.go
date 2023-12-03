package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	firstSample, err := parse("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	if err != nil {
		log.Fatal(err)
	}
	firstSample.getMapping()

	day8a("sample.txt")
	day8b("sample.txt")
	fmt.Println("---")
	day8a("input.txt")
	day8b("input.txt")
}

func day8a(fn string) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	var cnt int
	for line := range ch {
		out, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		for _, o := range out.output {
			if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

func day8b(fn string) {
	ch, err := readinp.Read(fn, parse)
	if err != nil {
		log.Fatal(err)
	}
	var cnt int
	for line := range ch {
		out, err := line.Get()
		if err != nil {
			log.Fatal(err)
		}
		mapping := out.getMapping()
		var tot int
		for _, outVal := range out.output {
			if outVal == "" {
				continue
			}
			tot = 10*tot + mapping.translate(outVal)
		}
		fmt.Println(out.line, tot)
		cnt += tot
	}
	fmt.Println(cnt)
}
