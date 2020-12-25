package main

import (
	"fmt"
	"testing"
)

func TestDay12(t *testing.T) {
	s := relship{wx: 10, wy: -1}
	s.move("F10")
	fmt.Println(s)
	s.move("N3")
	fmt.Println(s)
	s.move("F7")
	fmt.Println(s)
	s.move("R90")
	fmt.Println(s)
	s.move("L90")
	fmt.Println(s)
	s.move("R90")
	fmt.Println(s)
	s.move("F11")
	fmt.Println(s)
	if s.x != 214 {
		t.Errorf("s.x: got %d want 17", s.x)
	}
	if s.y != 72 {
		t.Errorf("s.y: got %d want 8", s.y)
	}
}

func TestRot(t *testing.T) {
	testcases := []struct {
		mv    string
		wantX int
		wantY int
	}{
		{
			mv:    "L90",
			wantX: 1,
			wantY: -2,
		},
		{
			mv:    "R90",
			wantX: -1,
			wantY: 2,
		},
		{
			mv:    "L180",
			wantX: -2,
			wantY: -1,
		},
		{
			mv:    "R180",
			wantX: -2,
			wantY: -1,
		},
		{
			mv:    "R270",
			wantX: 1,
			wantY: -2,
		},
		{
			mv:    "L270",
			wantX: -1,
			wantY: 2,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.mv, func(t *testing.T) {
			s := relship{wx: 2, wy: 1}
			s.move(tc.mv)
			if s.wx != tc.wantX {
				t.Errorf("s.wx: got %d, want %d", s.wx, tc.wantX)
			}
			if s.wy != tc.wantY {
				t.Errorf("s.wy: got %d, want %d", s.wy, tc.wantY)
			}
		})

	}
}
