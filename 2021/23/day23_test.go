package main

import (
	"fmt"
	"testing"
)

func TestSampleRun(t *testing.T) {
	sampleMoves := []move{
		{15, 5}, {15, 6},
		{2, 19}, {2, 4}, {2, 3}, {2, 2}, {2, 1}, {2, 0},
		{6, 4}, {6, 5},
		{7, 15}, {7, 4},
		{1, 16}, {1, 15}, {1, 3}, {1, 2}, {1, 1},
		{8, 3}, {8, 15}, {8, 16}, {8, 17},
		{9, 11}, {9, 3}, {9, 15}, {9, 16},
		{5, 12}, {5, 11}, {5, 3},
		{14, 13}, {14, 12}, {14, 11}, {14, 2},
		{5, 11}, {5, 12}, {5, 13}, {5, 14},
		{7, 3}, {7, 11}, {7, 12}, {7, 13},
		{6, 4}, {6, 3}, {6, 11}, {6, 12},
		{11, 20}, {11, 19}, {11, 4}, {11, 15},
		{3, 21}, {3, 20}, {3, 19}, {3, 5},
		{14, 3}, {14, 4}, {14, 19}, {14, 20}, {14, 21}, {14, 22},
		{4, 2}, {4, 11},
		{12, 7}, {12, 2}, {12, 3}, {12, 4}, {12, 19}, {12, 20}, {12, 21},
		{13, 8}, {13, 7}, {13, 2}, {13, 3}, {13, 4}, {13, 19}, {13, 20},
		{1, 7}, {1, 8}, {1, 9},
		{2, 1}, {2, 7}, {2, 8},
		{3, 4}, {3, 3}, {3, 2}, {3, 7},
		{15, 5}, {15, 19},
	}
	st := sample.initFromPods()
	for _, mv := range sampleMoves {
		fmt.Println(st)
		fmt.Println(st.possibleNexts())
		if !st.canMove(mv.pod, mv.loc) {
			t.Errorf("canMove returned false moving %d (%d) to %d:\n%v", mv.pod, st.podLocations[mv.pod], mv.loc, st)
			break
		}
		st = st.move(mv.pod, mv.loc)
	}
	fmt.Println(st)
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
