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
	tokenIndices map[string]byte
	tokens       []string
}

func (d *dictionary) Token(tok string) byte {
	if i, ok := d.tokenIndices[tok]; ok {
		return i
	}
	i := byte(len(d.tokens))
	d.tokenIndices[tok] = i
	d.tokens = append(d.tokens, tok)

	return i
}

func (d dictionary) ToString(tokens []byte) string {
	var out string
	for _, tok := range tokens {
		out += d.tokens[tok]
	}
	return out
}

func (d *dictionary) Parse(in string) []byte {
	var tokens []byte
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

	seen := map[string]bool{}
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
	tgtStr := "e"
	// initial := &formula{
	// 	dict:   dict,
	// 	form:   []byte{0},
	// 	target: target.form,
	// } // "e"
	formHeap := heapof.Make([]*formula{target})
	i := 0
	for {
		if i%100 == 0 {
			fmt.Println(i, "states:", formHeap.Len(), "(visited", len(seen), ")")
		}
		i++
		if formHeap.Len() == 0 {
			fmt.Println("out of states!")
		}
		form := formHeap.PopHeap()
		nexts := transforms.prevs(form)
		for _, next := range nexts {
			if len(next.form) > len(target.form) {
				continue
			}
			nextStr := next.String()
			if seen[nextStr] {
				continue
			}
			seen[nextStr] = true
			if nextStr == tgtStr {
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
	form   []byte
	target []byte
	steps  int // number of steps to get here
}

func (f formula) Cost() int {
	cost := f.steps + len(f.form)
	// var fi, ti int
	// var prevMatches int // prefer streaks
	// for ti < len(f.target) {
	// 	if f.form[fi] == f.form[ti] {
	// 		cost -= prevMatches
	// 		prevMatches++
	// 		fi++
	// 		ti++
	// 		if fi >= len(f.form) {
	// 			// we've matched the whole input, now add the length
	// 			// of the target's remainder
	// 			cost += (len(f.target) - ti) / 4
	// 			break
	// 		}
	// 		continue
	// 	}
	// 	prevMatches = 0
	// 	cost++
	// 	ti++
	// }
	return cost
}

func (f formula) String() string {
	return f.dict.ToString(f.form)
}

func (f formula) Replace(i, l int, tr []byte) *formula {
	var nextVals []byte
	nextVals = append(nextVals, f.form[0:i]...)
	nextVals = append(nextVals, tr...)
	nextVals = append(nextVals, f.form[i+l:len(f.form)]...)
	return &formula{
		dict:   f.dict,
		form:   nextVals,
		target: f.target,
		steps:  f.steps + 1,
	}
}

type transforms struct {
	forward  map[byte][]transform
	reverse8 map[[8]byte]byte
	reverse6 map[[6]byte]byte
	reverse4 map[[4]byte]byte
	reverse3 map[[3]byte]byte
	reverse2 map[[2]byte]byte
	reverse1 map[byte]byte
}

type transform []byte

func (t transforms) nexts(from *formula) []*formula {
	var out []*formula
	for i, elem := range from.form {
		for _, tr := range t.forward[elem] {
			out = append(out, from.Replace(i, 1, tr))
		}
	}
	return out
}

func (t transforms) prevs(from *formula) []*formula {
	var out []*formula
	for i := range from.form {
		for f, t := range t.reverse1 {
			if f == from.form[i] {
				out = append(out, from.Replace(i, 1, []byte{t}))
			}
		}
		if i == len(from.form)-1 {
			continue
		}
		next2 := [2]byte{from.form[i], from.form[i+1]}
		for f, t := range t.reverse2 {
			if f == next2 {
				out = append(out, from.Replace(i, 2, []byte{t}))
			}
		}
		if i == len(from.form)-2 {
			continue
		}
		next3 := [3]byte{from.form[i], from.form[i+1], from.form[i+2]}
		for f, t := range t.reverse3 {
			if f == next3 {
				out = append(out, from.Replace(i, 3, []byte{t}))
			}
		}
		if i == len(from.form)-3 {
			continue
		}
		next4 := [4]byte{from.form[i], from.form[i+1], from.form[i+2], from.form[i+3]}
		for f, t := range t.reverse4 {
			if f == next4 {
				out = append(out, from.Replace(i, 4, []byte{t}))
			}
		}
		if i >= len(from.form)-5 {
			continue
		}
		next6 := [6]byte{from.form[i], from.form[i+1], from.form[i+2], from.form[i+3], from.form[i+4], from.form[i+5]}
		for f, t := range t.reverse6 {
			if f == next6 {
				out = append(out, from.Replace(i, 6, []byte{t}))
			}
		}
		if i >= len(from.form)-7 {
			continue
		}
		next8 := [8]byte{from.form[i], from.form[i+1], from.form[i+2], from.form[i+3], from.form[i+4], from.form[i+5], from.form[i+6], from.form[i+7]}
		for f, t := range t.reverse8 {
			if f == next8 {
				out = append(out, from.Replace(i, 8, []byte{t}))
			}
		}
	}
	return out
}

func ReadInput(fn string) (*dictionary, *transforms, *formula, error) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	dict := &dictionary{
		tokenIndices: map[string]byte{"e": 0},
		tokens:       []string{"e"},
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	transforms := &transforms{
		forward:  make(map[byte][]transform),
		reverse8: make(map[[8]byte]byte),
		reverse6: make(map[[6]byte]byte),
		reverse4: make(map[[4]byte]byte),
		reverse3: make(map[[3]byte]byte),
		reverse2: make(map[[2]byte]byte),
		reverse1: make(map[byte]byte),
	}
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
		transforms.forward[from[0]] = append(transforms.forward[from[0]], to)
		switch len(to) {
		case 1:
			transforms.reverse1[to[0]] = from[0]
		case 2:
			transforms.reverse2[[2]byte{to[0], to[1]}] = from[0]
		case 3:
			transforms.reverse3[[3]byte{to[0], to[1], to[2]}] = from[0]
		case 4:
			transforms.reverse4[[4]byte{to[0], to[1], to[2], to[3]}] = from[0]
		case 6:
			transforms.reverse6[[6]byte{to[0], to[1], to[2], to[3], to[4], to[5]}] = from[0]
		case 8:
			transforms.reverse8[[8]byte{to[0], to[1], to[2], to[3], to[4], to[5], to[6], to[7]}] = from[0]
		default:
			return nil, nil, nil, fmt.Errorf("unexpected length for to (%q): %d", tokens[1], len(to))
		}
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
