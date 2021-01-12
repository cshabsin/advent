package main

import "testing"

var (
	b1    = boardType{[]boardEntry{c(1)}}
	b13   = boardType{[]boardEntry{c(1), c(3)}}
	b1to3 = boardType{[]boardEntry{
		boardEntry{rangeBegin: 1, rangeEnd: 3},
	}}
	b1235 = boardType{[]boardEntry{
		boardEntry{rangeBegin: 1, rangeEnd: 3},
		c(5),
	}}
)

func TestGet(t *testing.T) {
	testcases := []struct {
		desc  string
		board boardType
		loc   int
		want  int
	}{
		{
			desc:  "basic",
			board: b1,
			loc:   0,
			want:  1,
		},
		{
			desc:  "basic2",
			board: b13,
			loc:   1,
			want:  3,
		},
		{
			desc:  "range",
			board: b1to3,
			loc:   0,
			want:  1,
		},
		{
			desc:  "range2",
			board: b1to3,
			loc:   2,
			want:  3,
		},
		{
			desc:  "after range",
			board: b1235,
			loc:   3,
			want:  5,
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
			board:      b13,
			index:      0,
			offset:     0,
			wantIndex:  1,
			wantOffset: 0,
		},
		{
			desc:       "wraparound",
			board:      b13,
			index:      1,
			offset:     0,
			wantIndex:  0,
			wantOffset: 0,
		},
		{
			desc:       "range",
			board:      b1to3,
			index:      0,
			offset:     0,
			wantIndex:  0,
			wantOffset: 1,
		},
		{
			desc:       "range middle",
			board:      b1to3,
			index:      0,
			offset:     1,
			wantIndex:  0,
			wantOffset: 2,
		},
		{
			desc:       "range wraparound",
			board:      b1to3,
			index:      0,
			offset:     2,
			wantIndex:  0,
			wantOffset: 0,
		},
		{
			desc:       "after range",
			board:      b1235,
			index:      0,
			offset:     2,
			wantIndex:  1,
			wantOffset: 0,
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
func TestFind(t *testing.T) {
	testcases := []struct {
		desc  string
		board boardType
		value int
		want  int
	}{
		{
			desc:  "basic",
			board: b13,
			value: 1,
			want:  0,
		},
		{
			desc:  "basic2",
			board: b13,
			value: 3,
			want:  1,
		},
		{
			desc:  "range0",
			board: b1to3,
			value: 1,
			want:  0,
		},
		{
			desc:  "range1",
			board: b1to3,
			value: 2,
			want:  1,
		},
		{
			desc:  "range2",
			board: b1to3,
			value: 3,
			want:  2,
		},
		{
			desc:  "after range",
			board: b1235,
			value: 1,
			want:  0,
		},
		{
			desc:  "after range",
			board: b1235,
			value: 5,
			want:  3,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			if got := tc.board.Find(tc.value); got != tc.want {
				t.Errorf("board.Find(%d) got %d, want %d", tc.value, got, tc.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	testcases := []struct {
		desc  string
		board boardType
		loc   int
		val   []int
		want  []int
	}{
		{
			desc:  "basic",
			board: b13,
			loc:   0,
			val:   []int{5},
			want:  []int{5, 1, 3},
		},
		{
			desc:  "basic2",
			board: b13,
			loc:   1,
			val:   []int{5},
			want:  []int{1, 3, 5},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.board.Insert(tc.loc, tc.val)
			for i, v := range tc.want {
				if got := tc.board.Get(i); got != v {
					t.Errorf("tc.board.Insert value at %d got %d, want %d (board %v)", i, got, v, tc.board)
				}
			}
		})
	}
}
