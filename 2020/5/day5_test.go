package main

import "testing"

func TestGetVals(t *testing.T) {
	testcases := []struct {
		input      string
		wantRow    int
		wantCol    int
		wantSeatID int
	}{
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}
	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			if got := getRow(tc.input); got != tc.wantRow {
				t.Errorf("getRow(%q) got %d, want %d", tc.input, got, tc.wantRow)
			}
			if got := getCol(tc.input); got != tc.wantCol {
				t.Errorf("getCol(%q) got %d, want %d", tc.input, got, tc.wantCol)
			}
			if got := getSeatID(tc.input); got != tc.wantSeatID {
				t.Errorf("getSeatID(%q) got %d, want %d", tc.input, got, tc.wantSeatID)
			}
		})
	}
}
