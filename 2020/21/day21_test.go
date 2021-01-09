package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseLine(t *testing.T) {
	testcases := []struct {
		input           string
		wantIngredients []string
		wantAllergens   []string
	}{
		{
			input:           "a e (contains b)",
			wantIngredients: []string{"a", "e"},
			wantAllergens:   []string{"b"},
		},
		{
			input:           "a (contains q, r, s)",
			wantIngredients: []string{"a"},
			wantAllergens:   []string{"q", "r", "s"},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			gotI, gotA := parseLine(tc.input)
			if diff := cmp.Diff(gotI, tc.wantIngredients); diff != "" {
				t.Errorf("parseLine ingredients -got +want:\n%s", diff)
			}
			if diff := cmp.Diff(gotA, tc.wantAllergens); diff != "" {
				t.Errorf("parseLine allergens -got +want:\n%s", diff)
			}
		})
	}
}
