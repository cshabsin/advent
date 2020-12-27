package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

func TestParse(t *testing.T) {
	testcases := []struct {
		eq      string
		want    expr
		wantLen int
	}{
		{
			eq: "1",
			want: expr{
				entry{
					operator: '+',
					term:     term{literal: 1},
				},
			},
		},
		{
			eq: "1 + 1",
			want: expr{
				entry{
					operator: '+',
					term:     term{literal: 1},
				},
				entry{
					operator: '+',
					term:     term{literal: 1},
				},
			},
		},
		{
			eq: "(1 + 1)",
			want: expr{
				entry{
					operator: '+',
					term: term{parenthetical: expr{
						entry{
							operator: '+',
							term:     term{literal: 1},
						},
						entry{
							operator: '+',
							term:     term{literal: 1},
						},
					},
					},
				},
			},
		},
		// {
		// 	eq:      "1 + 1)",
		// 	want:    2,
		// 	wantLen: 5,
		// },
		// {
		// 	eq:   "1 + 2 * 3",
		// 	want: 7,
		// },
		// {
		// 	eq:   "2 * 3 + 1",
		// 	want: 7,
		// },
		// {
		// 	eq:   "1 + (2 + 3)",
		// 	want: 6,
		// },
		// {
		// 	eq:   "5 + (6 + 7) + 8",
		// 	want: 26,
		// },
		// {
		// 	eq:   "((9 + 1))",
		// 	want: 10,
		// },
	}
	for _, tc := range testcases {
		t.Run(tc.eq, func(t *testing.T) {
			got, off := parse(tc.eq)
			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(entry{}, term{})); diff != "" {
				t.Errorf("calc unexpected result, -want +got:\n%s", diff)
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

func TestCalc2(t *testing.T) {
	testcases := []struct {
		eq   string
		want int
	}{
		{"4 * 3 + 9", 48},
		{"1 + 1", 2},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
	}
	for _, tc := range testcases {
		t.Run(tc.eq, func(t *testing.T) {
			fmt.Println("---", tc.eq, "---")
			if val := calc2(tc.eq); val != tc.want {
				t.Errorf("calc got %d want %d", val, tc.want)
			}
		})
	}
}
