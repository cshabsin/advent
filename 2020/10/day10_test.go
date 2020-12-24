package main

import (
	"strconv"
	"testing"
)

func TestCalcPerms(t *testing.T) {
	testcases := []struct {
		runLength int
		want      int
	}{
		{
			runLength: 1,
			want:      1,
		},
		{
			runLength: 2,
			want:      2,
		},
		{
			runLength: 3,
			want:      4,
		},
		{
			runLength: 4,
			want:      7,
		},
		{
			runLength: 5,
			want:      13,
		},
		{
			runLength: 6,
			want:      24,
		},
		{
			runLength: 7,
			want:      44,
		},
	}
	for _, tc := range testcases {
		t.Run(strconv.Itoa(tc.runLength), func(t *testing.T) {
			if got := calcPerms(tc.runLength); got != tc.want {
				t.Errorf("calcPerms(%d): got %d, want %d", tc.runLength, got, tc.want)
			}
		})
	}
}

func TestAllPerms(t *testing.T) {
	testcases := []struct {
		desc      string
		inputList []int
		maxVal    int
		want      int
	}{
		{
			desc:      "first",
			inputList: []int{1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22},
			maxVal:    22,
			want:      8,
		},
		{
			desc:      "second",
			inputList: []int{0, 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31, 32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49, 52},
			maxVal:    52,
			want:      19208,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			inputSet := map[int]bool{}
			for _, v := range tc.inputList {
				inputSet[v] = true
			}
			if got := allPermutations(inputSet, tc.maxVal); got != tc.want {
				t.Errorf("allPermutations: got %d, want %d", got, tc.want)
			}
		})
	}
}
