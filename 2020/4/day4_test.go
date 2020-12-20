package main

import "testing"

func TestIsValidHeight(t *testing.T) {
	testcases := []struct {
		hgt  string
		want bool
	}{
		{
			hgt:  "60in",
			want: true,
		},
		{
			hgt:  "190cm",
			want: true,
		},
		{
			hgt:  "190in",
			want: false,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.hgt, func(t *testing.T) {
			if got := isValidHeight(tc.hgt); got != tc.want {
				t.Errorf("isValidHeight(%q): got %t, want %t", tc.hgt, got, tc.want)
			}
		})
	}
}

func TestDay4B(t *testing.T) {
	day4b()
}
