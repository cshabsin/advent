package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func LookAndSay(s string) string {
	i := 0
	chunks := make([]string, 0, len(s)*2)
	for i < len(s) {
		v := s[i]
		vstr := s[i:i+1]
		count := 0
		for i<len(s) && s[i] == v {
			count++
			i++
		}
		chunks = append(chunks, strconv.Itoa(count))
		chunks = append(chunks, vstr)
	}
	return strings.Join(chunks, "")
}

func main() {
	var input = flag.String("input", "3113322113", "Input string")
	flag.Parse()

	s := *input
	for i := 0; i<50; i++ {
		s = LookAndSay(s)
	}
	fmt.Print(*input, " : ", s, " (", len(s), ")\n")
}
