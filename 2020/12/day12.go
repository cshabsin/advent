package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := relship{}
	s.move("E10")
	s.move("N1")
	for line := range ch {
		if line.Error != nil {
			log.Fatal(err)
		}
		if err := s.move(strings.TrimSpace(*line.Contents)); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(s)
	fmt.Println(abs(s.x) + abs(s.y))
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

type dir int

const (
	east dir = iota
	south
	west
	north
)

func (d *dir) left(deg int) {
	for deg > 0 {
		if *d == east {
			*d = north
		} else {
			*d--
		}
		deg -= 90
	}
}

func (d *dir) right(deg int) {
	for deg > 0 {
		if *d == north {
			*d = east
		} else {
			*d++
		}
		deg -= 90
	}
}

func (d dir) moves() (int, int, error) {
	switch d {
	case east:
		return 1, 0, nil
	case west:
		return -1, 0, nil
	case north:
		return 0, -1, nil
	case south:
		return 0, 1, nil
	default:
		return 0, 0, fmt.Errorf("invalid dir %d", d)
	}
}

type ship struct {
	dir
	x, y int
}

func (s *ship) move(mv string) error {
	val, err := strconv.Atoi(mv[1:len(mv)])
	if err != nil {
		return err
	}
	switch mv[0] {
	case 'N':
		s.y -= val
	case 'S':
		s.y += val
	case 'E':
		s.x += val
	case 'W':
		s.x -= val
	case 'L':
		s.left(val)
	case 'R':
		s.right(val)
	case 'F':
		if err := s.forward(val); err != nil {
			return err
		}
	}
	return nil
}

func (s *ship) forward(val int) error {
	dx, dy, err := s.moves()
	if err != nil {
		return err
	}
	s.x += dx * val
	s.y += dy * val
	return nil
}

type relship struct {
	x, y   int
	wx, wy int
}

func (rs *relship) move(mv string) error {
	val, err := strconv.Atoi(mv[1:len(mv)])
	if err != nil {
		return err
	}
	switch mv[0] {
	case 'N':
		rs.wy -= val
	case 'S':
		rs.wy += val
	case 'E':
		rs.wx += val
	case 'W':
		rs.wx -= val
	case 'L':
		for val > 0 {
			rs.wx, rs.wy = rs.wy, -rs.wx
			val -= 90
		}
	case 'R':
		for val > 0 {
			rs.wx, rs.wy = -rs.wy, rs.wx
			val -= 90
		}
	case 'F':
		rs.x += rs.wx * val
		rs.y += rs.wy * val
	}
	return nil
}
