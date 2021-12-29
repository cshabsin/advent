package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	part2("input.txt")
}

// set of known atoms
type dictionary struct {
	tokenIndices map[string]int
	tokens       []string
}

func (d *dictionary) Token(tok string) int {
	if i, ok := d.tokenIndices[tok]; ok {
		return i
	}
	i := len(d.tokens)
	d.tokenIndices[tok] = i
	d.tokens = append(d.tokens, tok)

	return i
}

func (d dictionary) ToString(tokens []int) string {
	var out string
	for _, tok := range tokens {
		out += d.tokens[tok]
	}
	return out
}

func (d *dictionary) Parse(in string) []int {
	var tokens []int
	var i int
	for i < len(in) {
		tok := string(in[i])
		i++
		if len(in) > i && unicode.IsLower(rune(in[i])) {
			tok += string(in[i])
			i++
		}
		tokens = append(tokens, d.Token(tok))
	}
	return tokens
}

func part2(fn string) {
	dict, transforms, target, err := ReadInput(fn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dict)
	fmt.Println(transforms)
	fmt.Println(dict.ToString(target))
}

type transform []int

func ReadInput(fn string) (*dictionary, map[int][]transform, []int, error) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	dict := &dictionary{
		tokenIndices: map[string]int{"e": 0},
		tokens:       []string{"e"},
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	transforms := make(map[int][]transform)
	for {
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return nil, nil, nil, err
			}
			return nil, nil, nil, errors.New("file ended during phase 1 of parsing")
		}
		line := scanner.Text()
		if line == "" {
			break
		}
		tokens := strings.Split(line, " => ")
		from := dict.Parse(tokens[0])
		if len(from) != 1 {
			return nil, nil, nil, fmt.Errorf("Couldn't parse from %v", tokens[0])
		}
		to := dict.Parse(tokens[1])
		transforms[from[0]] = append(transforms[from[0]], to)
	}
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, nil, nil, err
		}
		return nil, nil, nil, errors.New("file ended during phase 2 of parsing")
	}
	target := dict.Parse(scanner.Text())
	return dict, transforms, target, nil
}
