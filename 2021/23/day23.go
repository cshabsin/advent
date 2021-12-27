package main

import (
	"container/heap"
	"fmt"
	"log"
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

var (
	input = &state{
		podLocations: [16]Location{
			7, 15, 17, 20, // A
			13, 16, 18, 22, // B
			10, 12, 19, 21, // C
			8, 9, 11, 14, // D
		},
	}
	sample = &state{
		podLocations: [16]Location{
			10, 17, 20, 22,
			7, 13, 15, 16,
			11, 12, 18, 21,
			8, 9, 14, 19,
		},
	}
)

type Location byte
type Pod byte

func main() {
	sh := &stateHeap{[]*state{sample.initFromPods()}}
	heap.Init(sh)
	i := 0
	visitedStates := map[[16]Location]bool{}
	for {
		if len(sh.states) == 0 {
			fmt.Println("out of states!")
			break
		}
		nextState := heap.Pop(sh).(*state)
		if i == 0 {
			fmt.Println("====== Processing:", nextState, "(", len(sh.states), ")")
		}
		i++
		if i == 500000 {
			i = 0
		}
		if visitedStates[nextState.podLocations] {
			continue
		}
		visitedStates[nextState.podLocations] = true
		next := nextState.possibleNexts()
		for _, s := range next {
			if s.win() {
				fmt.Println("win!")
				fmt.Println(s, "cost:", s.costSoFar)
				return
			}
			if !visitedStates[s.podLocations] {
				heap.Push(sh, s)
			}
		}
	}
}

type state struct {
	// 0-3 a, 4-7 b, 8-11 c, 12-15 d
	podLocations [16]Location
	locContents  [23]*Pod
	prevMover    Pod
	prevDir      int // -1 for leftward, 1 for rightward

	costSoFar int
}

type stateHeap struct {
	states []*state
}

func (sh *stateHeap) Len() int {
	return len(sh.states)
}

func (sh *stateHeap) Less(i, j int) bool {
	return sh.states[i].value()+sh.states[i].costSoFar < sh.states[j].value()+sh.states[j].costSoFar
}

func (sh *stateHeap) Swap(i, j int) {
	tmp := sh.states[j]
	sh.states[j] = sh.states[i]
	sh.states[i] = tmp
}

func (sh *stateHeap) Push(x interface{}) {
	sh.states = append(sh.states, x.(*state))
}

func (sh *stateHeap) Pop() interface{} {
	tmp := sh.states[sh.Len()-1]
	sh.states = sh.states[:sh.Len()-1]
	return tmp
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
	for i, loc := range s.podLocations {
		iPod := Pod(i)
		var r rune
		if iPod.isA() {
			r = 'A'
		} else if iPod.isB() {
			r = 'B'
		} else if iPod.isC() {
			r = 'C'
		} else if iPod.isD() {
			r = 'D'
		} else {
			r = '?'
		}
		board[locations[loc][0]] = replaceAtIndex(board[locations[loc][0]], locations[loc][1], r)
	}
	return fmt.Sprintf("\n%d - %d:\n%s", s.value(), s.costSoFar, strings.Join(board, "\n"))
}

func (s *state) initFromPods() *state {
	for i, loc := range s.podLocations {
		p := Pod(i)
		s.locContents[loc] = &p
	}
	return s
}

func (s state) win() bool {
	for i, loc := range s.podLocations {
		if !locMatchesPod(Pod(i), loc) {
			return false
		}
	}
	return true
}

// cost for pod index i to move one square
func (p Pod) cost() int {
	if p < 4 {
		return 1
	}
	if p < 8 {
		return 10
	}
	if p < 12 {
		return 100
	}
	return 1000
}

func (p Pod) isA() bool {
	return p >= 0 && p < 4
}

func (p Pod) isB() bool {
	return p >= 4 && p < 8
}

func (p Pod) isC() bool {
	return p >= 8 && p < 12
}

func (p Pod) isD() bool {
	return p >= 12 && p < 16
}

func (loc Location) isA() bool {
	return loc >= 7 && loc <= 10
}

func (loc Location) isB() bool {
	return loc >= 11 && loc <= 14
}

func (loc Location) isC() bool {
	return loc >= 15 && loc <= 18
}

func (loc Location) isD() bool {
	return loc >= 19 && loc <= 22
}

func (loc Location) isHall() bool {
	return loc < 7
}

func locMatchesPod(i Pod, loc Location) bool {
	if i.isA() && loc.isA() {
		return true
	}
	if i.isB() && loc.isB() {
		return true
	}
	if i.isC() && loc.isC() {
		return true
	}
	if i.isD() && loc.isD() {
		return true
	}
	return false
}

// neighbor links of distance 1
var neighbors1 = [][]Location{
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
var neighbors2 = [][]Location{
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
	2, 5, 5, 7, 9, 11, 8, // hall
	3, 2, 1, 0, // a
	7, 8, 9, 10, // b
	9, 10, 11, 12, // c
	11, 12, 13, 14, // d
}

var bVal = []int{
	4, 7, 5, 5, 7, 9, 6, // hall
	7, 8, 9, 10, // a
	3, 2, 1, 0, // b
	7, 8, 9, 10, // c
	9, 10, 11, 12, // d
}

var cVal = []int{
	6, 9, 7, 5, 5, 7, 4, // hall
	9, 10, 11, 12, // a
	7, 8, 9, 10, // b
	3, 2, 1, 0, // c
	7, 8, 9, 10, // d
}

var dVal = []int{
	8, 11, 9, 7, 5, 5, 2, // hall
	11, 12, 13, 14, // a
	9, 10, 11, 12, // b
	7, 8, 9, 10, // c
	3, 2, 1, 0, // d
}

func (s state) isDone(loc Location) bool {
	if loc.isA() {
		return s.isADone()
	}
	if loc.isB() {
		return s.isBDone()
	}
	if loc.isC() {
		return s.isCDone()
	}
	if loc.isD() {
		return s.isDDone()
	}
	log.Fatal("isDone called on ", loc)
	return true
}

// return true if the only thing in the A column is A pods, i.e. the column is "done" enough for more A pods to move in
func (s state) isADone() bool {
	for j := range s.podLocations {
		if Pod(j).isA() {
			continue
		}
		if s.podLocations[j].isA() {
			return false
		}
	}
	return true
}

func (s state) isBDone() bool {
	for j := range s.podLocations {
		if Pod(j).isB() {
			continue
		}
		if s.podLocations[j].isB() {
			return false
		}
	}
	return true
}

func (s state) isCDone() bool {
	for j := range s.podLocations {
		if Pod(j).isC() {
			continue
		}
		if s.podLocations[j].isC() {
			return false
		}
	}
	return true
}

func (s state) isDDone() bool {
	for j := range s.podLocations {
		if Pod(j).isD() {
			continue
		}
		if s.podLocations[j].isD() {
			return false
		}
	}
	return true
}

func dbg(x ...interface{}) {
	fmt.Println(x...)
}

func (s state) canMove(i Pod, to Location) bool {
	from := s.podLocations[i]
	if from.isHall() && to.isHall() {
		if s.prevMover == i {
			if to < from {
				// moving to the left is only possible if previous direction was left
				return s.prevDir < 0
			}
			return s.prevDir >= 0
		}
		return true
	}
	if from.isHall() {
		if !locMatchesPod(i, to) {
			return false // mismatched pod can never enter
		}
		return s.isDone(to) // only enter if the location is done.
	}
	// only move out to the hall if leaving a column that's not done
	if to.isHall() {
		return !s.isDone(from)
	}

	// if the column is done, only allow moves further in.
	if s.isDone(to) {
		return to > from
	}
	// column isn't done, only allow moves out.
	return to < from
}

func (s state) direction(i Pod, to Location) int {
	if !to.isHall() {
		return 0
	}
	from := s.podLocations[i]
	if from.isHall() {
		if to < from {
			return -1
		} else {
			return 1
		}
	}
	if from.isA() {
		if to == 1 {
			return -1
		}
		return 1
	}
	if from.isB() {
		if to == 2 {
			return -1
		}
		return 1
	}
	if from.isC() {
		if to == 3 {
			return -1
		}
		return 1
	}
	if from.isD() {
		if to == 4 {
			return -1
		}
		return 1
	}
	return 0 // should never get here
}

func (s state) isAHome(i Pod) bool {
	if !s.podLocations[i].isA() {
		return false
	}
	for loc := s.podLocations[i]; loc <= 10; loc++ {
		if s.locContents[loc] != nil && !s.locContents[loc].isA() {
			return false
		}
	}
	return true
}

func (s state) isBHome(i Pod) bool {
	if !s.podLocations[i].isB() {
		return false
	}
	for loc := s.podLocations[i]; loc <= 14; loc++ {
		if s.locContents[loc] != nil && !s.locContents[loc].isB() {
			return false
		}
	}
	return true
}

func (s state) isCHome(i Pod) bool {
	if !s.podLocations[i].isC() {
		return false
	}
	for loc := s.podLocations[i]; loc <= 18; loc++ {
		if s.locContents[loc] != nil && !s.locContents[loc].isC() {
			return false
		}
	}
	return true
}

func (s state) isDHome(i Pod) bool {
	if !s.podLocations[i].isD() {
		return false
	}
	for loc := s.podLocations[i]; loc <= 22; loc++ {
		if s.locContents[loc] != nil && !s.locContents[loc].isD() {
			return false
		}
	}
	return true
}

func (s state) value() int {
	var value int
	for i := Pod(0); i < 4; i++ {
		value += aVal[s.podLocations[i]]
		if s.isAHome(i) {
			value -= 100
		}
	}
	for i := Pod(4); i < 8; i++ {
		value += 10 * bVal[s.podLocations[i]]
		if s.isBHome(i) {
			value -= 100
		}
	}
	for i := Pod(8); i < 12; i++ {
		value += 100 * cVal[s.podLocations[i]]
		if s.isCHome(i) {
			value -= 100
		}
	}
	for i := Pod(12); i < 16; i++ {
		value += 1000 * dVal[s.podLocations[i]]
		if s.isDHome(i) {
			value -= 100
		}
	}
	return value
}
func (s state) move(i Pod, loc Location, dist int) *state {
	if !s.canMove(i, loc) {
		return nil
	}
	for j, pod := range s.podLocations {
		if i != Pod(j) && pod == loc {
			return nil // someone else is already there!
		}
	}
	s2 := &state{
		podLocations: s.podLocations,
		locContents:  s.locContents,
		prevMover:    i,
		prevDir:      s.direction(i, loc),
		costSoFar:    s.costSoFar + dist*i.cost(),
	}
	from := s.podLocations[i]
	s2.podLocations[i] = loc
	s2.locContents[from] = nil
	s2.locContents[loc] = &i
	return s2 // return state where given pod moves to location
}

func (s state) possibleNexts() []*state {
	var next []*state
	for i, podLoc := range s.podLocations {
		iPod := Pod(i)
		for _, neigh := range neighbors1[podLoc] {
			mv := s.move(iPod, neigh, 1)
			if mv != nil {
				next = append(next, mv)
			}
		}
		for _, neigh := range neighbors2[podLoc] {
			mv := s.move(iPod, neigh, 2)
			if mv != nil {
				next = append(next, mv)
			}
		}
	}
	return next
}
