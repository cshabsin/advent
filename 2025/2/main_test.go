package main

import (
	"fmt"
	"testing"
)

func TestIsInvalid(t *testing.T) {
	testcases := []struct {
		val  int
		want bool
	}{
		{
			val:  1,
			want: false,
		},
		{
			val:  2,
			want: false,
		},
		{
			val:  11,
			want: true,
		},
	}
	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%d", tc.val), func(t *testing.T) {
			got := isInvalid(tc.val)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
