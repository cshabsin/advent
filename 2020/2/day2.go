package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	lineRegex, err := regexp.Compile(`(\d*)-(\d*) ([[:alpha:]]): ([[:alpha:]]*)`)
	if err != nil {
		log.Fatal(err)
	}
	validPws := 0
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		match := lineRegex.FindStringSubmatch(line)
		if match == nil {
			log.Fatal("line had no matches: ", line)
		}
		min, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}
		char := match[3]
		password := match[4]
		num := count(password, char)
		if num >= min && num <= max {
			validPws++
		}
	}
	fmt.Println(validPws)
}

// count counts the incidence of char in password
func count(password, char string) int {
	count := 0
	for _, c := range password {
		if c == rune(char[0]) {
			count++
		}
	}
	return count
}
