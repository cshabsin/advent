package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type stateDataList []stateData

func (sdl stateDataList) String() string {
	var locstrs []string
	for i := range sdl {
		locstrs = append(locstrs, "[16]Location"+sdl[i].String())
	}
	return "{" + strings.Join(locstrs, ", ") + "}"
}

type stateData [16]Location

func (sd stateData) String() string {
	var locstrs []string
	for i := range sd {
		locstrs = append(locstrs, strconv.Itoa(int(sd[i])))
	}
	return "{" + strings.Join(locstrs, ", ") + "}"
}

func TestSampleRun(t *testing.T) {
	sampleMoves := []struct {
		pod  Pod
		from Location
		to   Location
	}{
		{15, 19, 5}, {15, 5, 6},
		{2, 20, 19}, {2, 19, 4}, {2, 4, 3}, {2, 3, 2}, {2, 2, 1}, {2, 1, 0},
		{6, 15, 4}, {6, 4, 5},
		{7, 16, 15}, {7, 15, 4},
		{1, 17, 16}, {1, 16, 15}, {1, 15, 3}, {1, 3, 2}, {1, 2, 1},
		{8, 11, 3}, {8, 3, 15}, {8, 15, 16}, {8, 16, 17},
		{9, 12, 11}, {9, 11, 3}, {9, 3, 15}, {9, 15, 16},
		{5, 13, 12}, {5, 12, 11}, {5, 11, 3},
		{14, 14, 13}, {14, 13, 12}, {14, 12, 11}, {14, 11, 2},
		{5, 3, 11}, {5, 11, 12}, {5, 12, 13}, {5, 13, 14},
		{7, 4, 3}, {7, 3, 11}, {7, 11, 12}, {7, 12, 13},
		{6, 5, 4}, {6, 4, 3}, {6, 3, 11}, {6, 11, 12},
		{11, 21, 20}, {11, 20, 19}, {11, 19, 4}, {11, 4, 15},
		{3, 22, 21}, {3, 21, 20}, {3, 20, 19}, {3, 19, 5},
		{14, 2, 3}, {14, 3, 4}, {14, 4, 19}, {14, 19, 20}, {14, 20, 21}, {14, 21, 22},
		{4, 7, 2}, {4, 2, 11},
		{12, 8, 7}, {12, 7, 2}, {12, 2, 3}, {12, 3, 4}, {12, 4, 19}, {12, 19, 20}, {12, 20, 21},
		{13, 9, 8}, {13, 8, 7}, {13, 7, 2}, {13, 2, 3}, {13, 3, 4}, {13, 4, 19}, {13, 19, 20},
		{1, 1, 7}, {1, 7, 8}, {1, 8, 9},
		{2, 0, 1}, {2, 1, 7}, {2, 7, 8},
		{3, 5, 4}, {3, 4, 3}, {3, 3, 2}, {3, 2, 7},
		{15, 6, 5}, {15, 5, 19},
	}
	st := sample.initFromPods()
	var states stateDataList
	for i, mv := range sampleMoves {
		// fmt.Println(st)
		nexts := st.possibleNexts()
		if i < len(sampleMoves)-1 {
			nextMove := sampleMoves[i+1]
			var found bool
			for _, next := range nexts {
				if next.podLocations[nextMove.pod] == nextMove.from {
					found = true
					break
				}
			}
			if !found {
				t.Error("couldn't find nextMove", nextMove, "among nexts", nexts)
			}
		}
		if mv.from != st.podLocations[mv.pod] {
			t.Error("data entry problem, 'from' wrong on move", mv)
		}
		if !st.canMove(mv.pod, mv.to) {
			t.Errorf("canMove returned false moving %d (%d) to %d:\n%v", mv.pod, st.podLocations[mv.pod], mv.to, st)
			break
		}
		st = st.move(mv.pod, mv.to)
		states = append(states, st.podLocations)
	}
	fmt.Println(states)
}

// [{2 7} {2 2} {1 8} {2 3} {2 4} {2 5} {1 7} {3 9} {3 8} {3 7} {3 2} {3 3} {3 4} {13 19} {12 20} {14 21} {13 4} {12 19} {13 3} {12 4} {13 2} {13 7} {14 20} {12 3} {14 19} {12 2} {14 4} {15 22} {15 21} {15 20} {13 8} {4 11} {6 12} {7 13} {12 7} {4 2} {14 3} {15 19} {15 4} {3 5} {3 19} {11 15} {3 20} {9 16} {8 17} {11 4} {3 21} {11 19} {11 20} {2 6} {2 5} {2 19} {15 3} {15 4} {5 14} {6 11} {7 12} {6 3} {5 13} {6 4} {7 11} {7 3} {5 12} {5 11} {14 2} {14 11} {5 3} {9 15} {14 12} {14 13} {5 11} {5 12} {9 3} {9 11} {8 16} {8 15} {8 3} {1 1} {1 2} {1 3} {1 15} {1 16} {7 4} {7 15} {6 5} {6 4}]
func TestWeirdWinningRun(t *testing.T) {
	output := [][]int{{2, 7}, {2, 2}, {1, 8}, {2, 3}, {2, 4}, {2, 5}, {1, 7}, {3, 9}, {3, 8}, {3, 7}, {3, 2}, {3, 3}, {3, 4}, {13, 19}, {12, 20}, {14, 21}, {13, 4}, {12, 19}, {13, 3}, {12, 4}, {13, 2}, {13, 7}, {14, 20}, {12, 3}, {14, 19}, {12, 2}, {14, 4}, {15, 22}, {15, 21}, {15, 20}, {13, 8}, {4, 11}, {6, 12}, {7, 13}, {12, 7}, {4, 2}, {14, 3}, {15, 19}, {15, 4}, {3, 5}, {3, 19}, {11, 15}, {3, 20}, {9, 16}, {8, 17}, {11, 4}, {3, 21}, {11, 19}, {11, 20}, {2, 6}, {2, 5}, {2, 19}, {15, 3}, {15, 4}, {5, 14}, {6, 11}, {7, 12}, {6, 3}, {5, 13}, {6, 4}, {7, 11}, {7, 3}, {5, 12}, {5, 11}, {14, 2}, {14, 11}, {5, 3}, {9, 15}, {14, 12}, {14, 13}, {5, 11}, {5, 12}, {9, 3}, {9, 11}, {8, 16}, {8, 15}, {8, 3}, {1, 1}, {1, 2}, {1, 3}, {1, 15}, {1, 16}, {7, 4}, {7, 15}, {6, 5}, {6, 4}}
	var sampleMoves []move
	for i := len(output) - 1; i >= 0; i-- {
		sampleMoves = append(sampleMoves, move{
			pod: Pod(output[i][0]),
			loc: Location(output[i][1]),
		})
	}
	fmt.Println(sampleMoves)
	st := sample.initFromPods()
	for _, mv := range sampleMoves {
		if !st.canMove(mv.pod, mv.loc) {
			t.Errorf("canMove returned false moving %d (%d) to %d:\n%v", mv.pod, st.podLocations[mv.pod], mv.loc, st)
			break
		}
		st = st.move(mv.pod, mv.loc)
		fmt.Println(st)
	}
}
