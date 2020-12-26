package main

import (
	"testing"
)

func TestSieve(t *testing.T) {
	testcases := []struct {
		desc  string
		buses map[int]int
		want  int
	}{
		{
			desc: "first",
			buses: map[int]int{
				17: 0, 13: 2, 19: 3,
			},
			want: 3417,
		},
		{
			desc: "second",
			buses: map[int]int{
				67: 0, 7: 1, 59: 2, 61: 3,
			},
			want: 754018,
		},
		{
			desc: "third",
			buses: map[int]int{
				67: 0, 7: 2, 59: 3, 61: 4,
			},
			want: 779210,
		},
		{
			desc: "fourth",
			buses: map[int]int{
				67: 0, 7: 1, 59: 3, 61: 4,
			},
			want: 1261476,
		},
		{
			desc: "fifth",
			buses: map[int]int{
				1789: 0, 37: 1, 47: 2, 1889: 3,
			},
			want: 1202161486,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			if got := sieveBuses(tc.buses,0, 0); got != tc.want {
				t.Errorf("sieveBuses: got %d, want %d", got, tc.want)
			}
		})
	}
}
