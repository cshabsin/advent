package main

import "testing"

func TestRotate(t *testing.T) {
	testcases := []struct {
		name string
		dial int
		mul  int
		mag  int

		wantDial   int
		wantC      int
		wantPassed int
	}{
		{
			name: "basic",
			dial: 50,
			mul:  1,
			mag:  10,

			wantDial:   60,
			wantC:      0,
			wantPassed: 0,
		},
		{
			name: "passes 100",
			dial: 50,
			mul:  1,
			mag:  100,

			wantDial:   50,
			wantC:      1,
			wantPassed: 1,
		},
		{
			name: "passes 0",
			dial: 50,
			mul:  -1,
			mag:  100,

			wantDial:   50,
			wantC:      1,
			wantPassed: 1,
		},
		{
			name: "multiple pos",
			dial: 50,
			mul:  1,
			mag:  150,

			wantDial:   0,
			wantC:      2,
			wantPassed: 2,
		},
		{
			name: "multiple pos from 0",
			dial: 0,
			mul:  1,
			mag:  200,

			wantDial:   0,
			wantC:      2,
			wantPassed: 2,
		},
		{
			name: "pos from 0",
			dial: 0,
			mul:  1,
			mag:  199,

			wantDial:   99,
			wantC:      1,
			wantPassed: 1,
		},
		{
			name: "multiple neg",
			dial: 50,
			mul:  -1,
			mag:  150,

			wantDial:   0,
			wantC:      2,
			wantPassed: 2,
		},
		{
			name: "multiple neg from 0",
			dial: 0,
			mul:  -1,
			mag:  200,

			wantDial:   0,
			wantC:      2,
			wantPassed: 2,
		},
		{
			name: "neg from 0",
			dial: 0,
			mul:  -1,
			mag:  199,

			wantDial:   1,
			wantC:      1,
			wantPassed: 1,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			gotDial, gotC, gotPassed := rotate(tc.dial, 0, tc.mul, tc.mag)
			if gotDial != tc.wantDial {
				t.Errorf("dial got %d, want %d", gotDial, tc.wantDial)
			}
			if gotC != tc.wantC {
				t.Errorf("c got %d, want %d", gotC, tc.wantC)
			}
			if gotPassed != tc.wantPassed {
				t.Errorf("passed got %d, want %d", gotPassed, tc.wantPassed)
			}
		})
	}
}
