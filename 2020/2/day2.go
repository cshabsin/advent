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
	validTobogganPws := 0
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
		matches := 0
		if password[min-1] == char[0] {
			matches++
		}
		if password[max-1] == char[0] {
			matches++
		}
		if matches == 1 {
			validTobogganPws++
		}
	}
	fmt.Println(validPws)
	fmt.Println(validTobogganPws)
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
