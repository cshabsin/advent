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
		'}': 1197,
		']': 57,
		'>': 25137,
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

func (s state) pop(r rune) (state, *badClosing) {
	last := rune(s[len(s)-1])
	if pairs[last] != r {
		bc := badClosing(scores[r])
		return "", &bc
	}
	return s[0 : len(s)-1], nil
}

func (s state) next(r rune) (state, *badClosing) {
	if _, ok := pairs[r]; ok {
		return s.push(r), nil
	}
	return s.pop(r)
}

func parse(line string) (int, error) {
	var s state
	for _, r := range line {
		var err *badClosing
		s, err = s.next(rune(r))
		if err != nil {
			return int(*err), nil
		}
	}
	return 0, nil
}
