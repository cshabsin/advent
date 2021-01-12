package main

import "testing"

func TestGet(t *testing.T) {
	testcases := []struct {
		desc  string
		board boardType
		loc   int
		want  int
	}{
		{
			desc:  "basic",
			board: boardType{[]boardEntry{c(1)}},
			loc:   0,
			want:  1,
		},
		{
			desc:  "basic2",
			board: boardType{[]boardEntry{c(1), c(3)}},
			loc:   1,
			want:  3,
		},
		{
			desc: "range",
			board: boardType{[]boardEntry{
				boardEntry{rangeBegin: 1, rangeEnd: 3},
			}},
			loc:  0,
			want: 1,
		},
		{
			desc: "range2",
			board: boardType{[]boardEntry{
				boardEntry{rangeBegin: 1, rangeEnd: 3},
			}},
			loc:  2,
			want: 3,
		},
		{
			desc: "after range",
			board: boardType{[]boardEntry{
				boardEntry{rangeBegin: 1, rangeEnd: 3},
				c(5),
			}},
			loc:  3,
			want: 5,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			if got := tc.board.Get(tc.loc); got != tc.want {
				t.Errorf("board.Get(%d) got %d, want %d", tc.loc, got, tc.want)
			}
		})
	}
}

func TestNextIndexAndOffset(t *testing.T) {
	testcases := []struct {
		desc       string
		board      boardType
		index      int
		offset     int
		wantIndex  int
		wantOffset int
	}{
		{
			desc:       "basic",
			board:      boardType{[]boardEntry{c(1), c(3)}},
			index:      0,
			offset:     0,
			wantIndex:  1,
			wantOffset: 0,
		},
		{
			desc:       "wraparound",
			board:      boardType{[]boardEntry{c(1), c(3)}},
			index:      1,
			offset:     0,
			wantIndex:  0,
			wantOffset: 0,
		},
		{
			desc: "range",
			board: boardType{[]boardEntry{
				boardEntry{rangeBegin: 1, rangeEnd: 3},
			}},
			index:      0,
			offset:     0,
			wantIndex:  0,
			wantOffset: 1,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			gotIndex, gotOffset := tc.board.nextIndexAndOffset(tc.index, tc.offset)
			if gotIndex != tc.wantIndex {
				t.Errorf("board.nextIndexAndOffset(%d, %d) got index %d, want index %d",
					tc.index, tc.offset, gotIndex, tc.wantIndex)
			}
			if gotOffset != tc.wantOffset {
				t.Errorf("board.nextIndexAndOffset(%d, %d) got offset %d, want offset %d",
					tc.index, tc.offset, gotOffset, tc.wantOffset)
			}
		})
	}
}
