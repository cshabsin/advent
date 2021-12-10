package main

import "fmt"

var (
	pairs = map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
		'<': '>',
	}
	scores = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	completions = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

type state string

func (s state) push(r rune) state {
	return state(string(s) + string(r))
}

type badClosing int

func (b *badClosing) Error() string {
	return fmt.Sprintf("bad closing for score %d", b)
}

func (s state) pop(r rune) (state, error) {
	last := rune(s[len(s)-1])
	if pairs[last] != r {
		bc := badClosing(scores[r])
		return "", &bc
	}
	return s[0 : len(s)-1], nil
}

func (s state) next(r rune) (state, error) {
	if _, ok := pairs[r]; ok {
		return s.push(r), nil
	}
	return s.pop(r)
}

func (s state) completion() int {
	var total int
	for i := len(s) - 1; i >= 0; i-- {
		total = 5*total + completions[pairs[rune(s[i])]]
	}
	return total
}

func parse(line string) (int, error) {
	var s state
	for _, r := range line {
		var err error
		s, err = s.next(rune(r))
		if err != nil {
			bc, ok := err.(*badClosing)
			if !ok {
				return 0, err
			}
			return int(*bc), nil
		}
	}
	return 0, nil
}

func parse2(line string) (state, error) {
	var s state
	for _, r := range line {
		var err error
		s, err = s.next(rune(r))
		if err != nil {
			return "", nil
		}
	}
	return s, nil
}
