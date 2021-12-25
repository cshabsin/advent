package main

import (
	"fmt"
	"strings"
)

// #############
// #..x.x.x.x..#   01 2 3 4 56
// ###A#D#A#C###   7, 11, 15, 19
//   #D#C#B#A#     8  12  16  20
//   #D#B#A#C#     9  13  17  21
//   #C#D#B#B#     10 14  18  22
//   #########

// Cost 2 connections:
// 7-1, 7-2, 1-2, 11-2, 11-3, 2-3, 15-3, 15-4, 3-4, 19-4, 19-5, 4-5

func main() {
	initial := state{
		locations: [16]int{
			7, 15, 17, 20, // A
			13, 16, 18, 22, // B
			10, 12, 19, 21, // C
			8, 9, 11, 14, // D
		},
	}
	fmt.Println(initial)
	p2, _ := initial.move(10, 5, 2)
	fmt.Println(p2, p2.prevMover, p2.prevDir)
	fmt.Println(p2.canMove(10, 4))
	possibilities := map[state]int{initial: 0}
	for {
		newPoss := map[state]int{}
		for state, cost := range possibilities {
			next := state.possibleNext(cost)
			for newState, newCost := range next {
				if _, ok := newPoss[newState]; !ok {
					newPoss[newState] = newCost
				}
				if newPoss[newState] > newCost {
					newPoss[newState] = newCost
				}
			}
		}
		for st, cost := range newPoss {
			if st.win() {
				fmt.Println(st, cost)
			}
		}
		possibilities = newPoss
		fmt.Println("=====")
		fmt.Println(len(possibilities))
		// fmt.Println(possibilities)
		i := 0
		for p, cost := range possibilities {
			if i > 3 {
				break
			}
			i++
			fmt.Println(cost, p)
			fmt.Println("---")
		}
	}
}

type state struct {
	// 0-3 a, 4-7 b, 8-11 c, 12-15 d
	locations [16]int
	prevMover int
	prevDir   int // -1 for leftward, 1 for rightward
}

func replaceAtIndex(s string, i int, r rune) string {
	out := []rune(s)
	out[i] = r
	return string(out)
}

func (s state) String() string {
	board := []string{
		"#############",
		"#...........#",
		"###.#.#.#.###",
		"  #.#.#.#.#  ",
		"  #.#.#.#.#  ",
		"  #.#.#.#.#  ",
		"  #########  ",
	}
	locations := [][2]int{
		{1, 1},
		{1, 2},
		{1, 4},
		{1, 6},
		{1, 8},
		{1, 10},
		{1, 11},

		{2, 3},
		{3, 3},
		{4, 3},
		{5, 3},

		{2, 5},
		{3, 5},
		{4, 5},
		{5, 5},

		{2, 7},
		{3, 7},
		{4, 7},
		{5, 7},

		{2, 9},
		{3, 9},
		{4, 9},
		{5, 9},
	}
	for i, loc := range s.locations {
		var r rune
		if isA(i) {
			r = 'A'
		} else if isB(i) {
			r = 'B'
		} else if isC(i) {
			r = 'C'
		} else if isD(i) {
			r = 'D'
		} else {
			r = '?'
		}
		board[locations[loc][0]] = replaceAtIndex(board[locations[loc][0]], locations[loc][1], r)
	}
	return fmt.Sprintf("%d:\n%s", s.value(), strings.Join(board, "\n"))
}

func (s state) win() bool {
	for i, loc := range s.locations {
		if !locMatchesPod(i, loc) {
			return false
		}
	}
	return true
}

// neighbor links of distance 1
var neighbors1 = [][]int{
	{1}, // hall: 0
	{},  // 1
	{},  // 2
	{},  // 3
	{},  // 4
	{6}, // 5
	{5}, // 6

	{8},     // a: 7
	{7, 9},  // 8
	{8, 10}, // 9
	{9},     // 10

	{12},     // b: 11
	{11, 13}, // 12
	{12, 14}, // 13
	{13},     // 14

	{16},     // c: 15
	{15, 17}, // 16
	{16, 18}, // 17
	{17},     // 18

	{20},     // d: 19
	{19, 21}, // 20
	{20, 22}, // 21
	{21},     // 22
}

// neighbor links of distance 2
var neighbors2 = [][]int{
	{},             // hall: 0
	{2, 7},         // 1
	{1, 3, 7, 11},  // 2
	{2, 4, 11, 15}, // 3
	{3, 5, 15, 19}, // 4
	{4, 19},        // 5
	{},             // 6

	{1, 2}, // a: 7
	{},     // 8
	{},     // 9
	{},     // 10

	{2, 3}, // b: 11
	{},     // 12
	{},     // 13
	{},     // 14

	{3, 4}, // c: 15
	{},     // 16
	{},     // 17
	{},     // 18

	{4, 5}, // d: 19
	{},     // 20
	{},     // 21
	{},     // 22
}

var aVal = []int{
	6, 5, 5, 7, 9, 11, 12, // hall
	3, 2, 1, 0, // a
	7, 8, 9, 10, // b
	9, 10, 11, 12, // c
	11, 12, 13, 14, // d
}

var bVal = []int{
	8, 7, 5, 5, 7, 9, 10, // hall
	7, 8, 9, 10, // a
	3, 2, 1, 0, // b
	7, 8, 9, 10, // c
	9, 10, 11, 12, // d
}

var cVal = []int{
	10, 9, 7, 5, 5, 7, 8, // hall
	9, 10, 11, 12, // a
	7, 8, 9, 10, // b
	3, 2, 1, 0, // c
	7, 8, 9, 10, // d
}

var dVal = []int{
	12, 11, 9, 7, 5, 5, 6, // hall
	11, 12, 13, 14, // a
	9, 10, 11, 12, // b
	7, 8, 9, 10, // c
	3, 2, 1, 0, // d
}

func (s state) value() int {
	var value int
	for i := 0; i < 4; i++ {
		value += aVal[s.locations[i]]
	}
	for i := 4; i < 8; i++ {
		value += bVal[s.locations[i]]
	}
	for i := 8; i < 12; i++ {
		value += cVal[s.locations[i]]
	}
	for i := 12; i < 16; i++ {
		value += dVal[s.locations[i]]
	}
	return value
}

// cost for pod index i to move one square
func cost(i int) int {
	if i < 4 {
		return 1
	}
	if i < 8 {
		return 10
	}
	if i < 12 {
		return 100
	}
	return 1000
}

func isA(i int) bool {
	return i >= 0 && i < 4
}

func isB(i int) bool {
	return i >= 4 && i < 8
}

func isC(i int) bool {
	return i >= 8 && i < 12
}

func isD(i int) bool {
	return i >= 12 && i < 16
}

func isLocA(loc int) bool {
	return loc >= 7 && loc <= 10
}

func isLocB(loc int) bool {
	return loc >= 11 && loc <= 14
}

func isLocC(loc int) bool {
	return loc >= 15 && loc <= 18
}

func isLocD(loc int) bool {
	return loc >= 19 && loc <= 22
}

func isHall(loc int) bool {
	return loc < 7
}

func locMatchesPod(i int, loc int) bool {
	if isA(i) && isLocA(loc) {
		return true
	}
	if isB(i) && isLocB(loc) {
		return true
	}
	if isC(i) && isLocC(loc) {
		return true
	}
	if isD(i) && isLocD(loc) {
		return true
	}
	return false
}

func (s state) isADone() bool {
	for j := 7; j < 11; j++ {
		for j := range s.locations {
			if isA(j) {
				continue
			}
			if isLocA(s.locations[j]) {
				return false
			}
		}
	}
	return true
}

func (s state) isBDone() bool {
	for j := 11; j < 15; j++ {
		for j := range s.locations {
			if isB(j) {
				continue
			}
			if isLocB(s.locations[j]) {
				return false
			}
		}
	}
	return true
}

func (s state) isCDone() bool {
	for j := range s.locations {
		if isC(j) {
			continue
		}
		if isLocC(s.locations[j]) {
			return false
		}
	}
	return true
}

func (s state) isDDone() bool {
	for j := range s.locations {
		if isD(j) {
			continue
		}
		if isLocD(s.locations[j]) {
			return false
		}
	}
	return true
}

func (s state) canMove(i int, to int) bool {
	if !isHall(s.locations[i]) && (to < s.locations[i]) {
		return true
	}
	if s.prevMover != i {
		if isA(i) && !s.isADone() {
			return false
		}
		if isB(i) && !s.isBDone() {
			return false
		}
		if isC(i) && !s.isCDone() {
			return false
		}
		if isD(i) && !s.isDDone() {
			return false
		}
	}
	if isLocA(to) {
		return locMatchesPod(i, to) && s.isADone()
	}
	if isLocB(to) {
		return locMatchesPod(i, to) && s.isBDone()
	}
	if isLocC(to) {
		return locMatchesPod(i, to) && s.isCDone()
	}
	if isLocD(to) {
		return locMatchesPod(i, to) && s.isDDone()
	}
	if s.prevMover == i {
		if s.locations[i] > to {
			// moving to the left is only possible if previous direction was left
			return s.prevDir < 0
		}
		return s.prevDir >= 0
	}
	return false
}

func (s state) direction(i int, to int) int {
	if !isHall(to) {
		return 0
	}
	from := s.locations[i]
	if isHall(from) {
		if from < to {
			return -1
		} else {
			return 1
		}
	}
	if isLocA(from) {
		if to == 1 {
			return -1
		}
		return 1
	}
	if isLocB(from) {
		if to == 2 {
			return -1
		}
		return 1
	}
	if isLocC(from) {
		if to == 3 {
			return -1
		}
		return 1
	}
	if isLocD(from) {
		if to == 4 {
			return -1
		}
		return 1
	}
	return 0 // should never get here
}

func (s state) move(i int, loc int, dist int) (*state, int) {
	if !s.canMove(i, loc) {
		return nil, 0
	}
	for j, pod := range s.locations {
		if i != j && pod == loc {
			return nil, 0 // someone else is already there!
		}
	}
	s2 := &state{
		locations: cp(s.locations),
		prevMover: i,
		prevDir:   s.direction(i, loc),
	}
	s2.locations[i] = loc
	return s2, dist * cost(i) // return state where given pod moves to location
}

func cp(is [16]int) [16]int {
	var js [16]int
	for i := range is {
		js[i] = is[i]
	}
	return js
}

func (s state) possibleNext(prevCost int) map[state]int {
	next := map[state]int{}
	for i, podLoc := range s.locations {
		for _, neigh := range neighbors1[podLoc] {
			mv, cost := s.move(i, neigh, 1)
			cost += prevCost
			if mv != nil {
				if _, ok := next[*mv]; !ok {
					next[*mv] = cost
				}
				if next[*mv] > cost {
					next[*mv] = cost
				}
			}
		}
		for _, neigh := range neighbors2[podLoc] {
			mv, cost := s.move(i, neigh, 2)
			cost += prevCost
			if mv != nil {
				if _, ok := next[*mv]; !ok {
					next[*mv] = cost
				}
				if next[*mv] > cost {
					next[*mv] = cost
				}
			}
		}
	}
	return next
}
