package main

import (
	"fmt"
	"testing"
)

func TestSampleRun(t *testing.T) {
	sampleMoves := []struct {
		pod Pod
		loc Location
	}{
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
	}
	st := sample.initFromPods()
	for _, mv := range sampleMoves {
		if !st.canMove(mv.pod, mv.loc) {
			t.Errorf("canMove returned false moving %d (%d) to %d:\n%v", mv.pod, st.podLocations[mv.pod], mv.loc, st)
			break
		}
		st = st.move(mv.pod, mv.loc, 1)
		fmt.Println(st)
	}
}
