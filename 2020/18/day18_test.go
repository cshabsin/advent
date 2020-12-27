package main

import "testing"

func TestCalc(t *testing.T) {
	testcases := []struct {
		eq      string
		want    int
		wantLen int
	}{
		{
			eq:   "1",
			want: 1,
		},
		{
			eq:   "1 + 1",
			want: 2,
		},
		{
			eq:   "(1 + 1)",
			want: 2,
		},
		{
			eq:      "1 + 1)",
			want:    2,
			wantLen: 5,
		},
		{
			eq:   "1 + 2 * 3",
			want: 9,
		},
		{
			eq:   "1 + (2 + 3)",
			want: 6,
		},
		{
			eq:   "5 + (6 + 7) + 8",
			want: 26,
		},
		{
			eq:   "((9 + 1))",
			want: 10,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.eq, func(t *testing.T) {
			val, off, err := calc(tc.eq)
			if err != nil {
				t.Error(err)
			}
			if val != tc.want {
				t.Errorf("calc: want %d got %d", tc.want, val)
			}
			wantLen := tc.wantLen
			if wantLen == 0 {
				wantLen = len(tc.eq)
			}
			if off != wantLen {
				t.Errorf("calc len: want %d got %d", wantLen, off)
			}

		})
	}
}
