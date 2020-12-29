package main

import "testing"

func TestEdgesMatch(t *testing.T) {
	testcases := []struct {
		a, b int
		want bool
	}{
		{0, 0, true},
		{1, 512, true},
		{512, 1, true},
		{1, 511, false},
		{3, 768, true},
		{1, 1, false},
	}
	for _, tc := range testcases {
		if got := edgesMatch(tc.a, tc.b); got != tc.want {
			t.Errorf("edgesMatch(%d, %d); got %v, want %v", tc.a, tc.b, got, tc.want)
		}
		if got := edgeDual(tc.a); tc.want && got != tc.b {
			t.Errorf("edgeDual(%d): got %d, want %d", tc.a, got, tc.b)
		}
	}
}
