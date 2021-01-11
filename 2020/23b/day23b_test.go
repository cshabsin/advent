package main

import "testing"

func TestGet(t *testing.T) {
	testcases := []struct {
		desc  string
		board boardType
		index int
		want  int
	}{
		{
			desc:  "basic",
			board: boardType{[]boardEntry{c(1)}},
			index: 0,
			want:  1,
		},
		{
			desc:  "range",
			board: boardType{[]boardEntry{boardEntry{rangeBegin: 1, rangeEnd: 1}}},
			index: 0,
			want:  1,
		},
		{
			desc:  "range 5",
			board: boardType{[]boardEntry{boardEntry{rangeBegin: 1, rangeEnd: 5}}},
			index: 4,
			want:  5,
		},
		{
			desc:  "complex range",
			board: boardType{[]boardEntry{c(5), boardEntry{rangeBegin: 1, rangeEnd: 4}}},
			index: 4,
			want:  4,
		},
		{
			desc:  "card after range",
			board: boardType{[]boardEntry{c(5), boardEntry{rangeBegin: 1, rangeEnd: 4}, c(6)}},
			index: 5,
			want:  6,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			if got := tc.board.get(tc.index); got != tc.want {
				t.Errorf("board.get(%d): got %d, want %d", tc.index, got, tc.want)
			}
		})
	}
}
