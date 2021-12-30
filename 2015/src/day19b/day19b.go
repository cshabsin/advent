package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/cshabsin/advent/commongen/heapof"
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
	fmt.Println(target)

	for _, s := range []string{
		"e",
		"CRnCa",
		"CRnCaArArMg",
		"CRnCaArArMgP",
		target.String(),
	} {
		fmt.Println(s, ":", formula{
			form:   dict.Parse(s),
			target: target.form,
			dict:   dict,
		}.Cost())
	}
	tgtStr := target.String()
	initial := &formula{
		dict:   dict,
		form:   []int{0},
		target: target.form,
	} // "e"
	formHeap := heapof.Make([]*formula{initial})
	i := 0
	for {
		if i%100 == 0 {
			fmt.Println(i, "states:", formHeap.Len())
		}
		i++
		if formHeap.Len() == 0 {
			fmt.Println("out of states!")
		}
		form := formHeap.PopHeap()
		nexts := transforms.nexts(form)
		for _, next := range nexts {
			if next.String() == tgtStr {
				fmt.Println("win!")
				fmt.Println(next, "steps:", next.steps)
				return
			}
			formHeap.PushHeap(next)
		}
	}
}

type formula struct {
	dict   *dictionary
	form   []int
	target []int
	steps  int // number of steps to get here
}

func (f formula) Cost() int {
	cost := f.steps
	var fi, ti int
	var prevMatches int // prefer streaks
	for ti < len(f.target) {
		if f.form[fi] == f.form[ti] {
			cost -= prevMatches
			prevMatches++
			fi++
			ti++
			if fi >= len(f.form) {
				// we've matched the whole input, now add the length
				// of the target's remainder
				cost += len(f.target) - ti
				break
			}
			continue
		}
		prevMatches = 0
		cost++
		ti++
	}
	return cost
}

func (f formula) String() string {
	return f.dict.ToString(f.form)
}

type transforms map[int][]transform
type transform []int

func (t transforms) nexts(from *formula) []*formula {
	var out []*formula
	for i, elem := range from.form {
		for _, tr := range t[elem] {
			var nextVals []int
			nextVals = append(nextVals, from.form[0:i]...)
			nextVals = append(nextVals, tr...)
			nextVals = append(nextVals, from.form[i+1:len(from.form)]...)
			out = append(out, &formula{dict: from.dict, form: nextVals})
		}
	}
	return out
}

func ReadInput(fn string) (*dictionary, transforms, *formula, error) {
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
	return dict, transforms, &formula{
		dict:   dict,
		form:   target,
		target: target,
	}, nil
}
