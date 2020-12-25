package main

import (
	"testing"
)

func TestDay12(t *testing.T) {
	s := ship{}
	s.move("F10")
	s.move("N3")
	s.move("F7")
	s.move("R90")
	s.move("F11")
	if s.x != 17 {
		t.Errorf("s.x: got %d want 17", s.x)
	}
	if s.y != 8 {
		t.Errorf("s.x: got %d want 8", s.x)
	}
}
