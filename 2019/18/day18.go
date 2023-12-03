package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	keys := map[rune][2]int{}
	doors := map[rune][2]int{}
	var vault [][]rune
	var posR, posC int
	row := 0
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		vault = append(vault, []rune(line))
		for col, r := range line {
			if unicode.IsLetter(r) && unicode.IsLower(r) {
				keys[unicode.ToUpper(r)] = [2]int{row, col}
			}
			if unicode.IsLetter(r) && unicode.IsUpper(r) {
				doors[r] = [2]int{row, col}
			}
			if r == '@' {
				posR = row
				posC = col
			}
		}
		row++
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	b := board{vault, posR, posC, keys, doors, map[rune]bool{}}
	fmt.Println(b)
}

type board struct {
	vault      [][]rune
	posR, posC int
	keys       map[rune][2]int
	doors      map[rune][2]int
	opened     map[rune]bool
}
