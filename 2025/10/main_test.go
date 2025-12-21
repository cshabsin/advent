package main

import (
	"testing"
)

// func TestMatrixVoltages(t *testing.T) {
// 	machine := ParseLine("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}")
// 	fmt.Println(matrixVoltages(machine))
// }

func TestDepthFirstJoltages(t *testing.T) {
	machine := ParseLine("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}")
	got := depthFirstJoltages(machine)
	want := 10
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
