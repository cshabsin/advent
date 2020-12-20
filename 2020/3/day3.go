package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	b := board{}
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if err := b.Read(line); err != nil {
			log.Fatal(err)
		}
	}
	tot := 1
	for _, slope := range [][2]int{
		[2]int{1, 1},
		[2]int{3, 1},
		[2]int{5, 1},
		[2]int{7, 1},
		[2]int{1, 2},
	} {
		ck := b.checkSlope(slope)
		fmt.Println(slope, ":", ck)
		tot = tot * ck
	}
	fmt.Println("tot:", tot)
}

func (b board) checkSlope(slope [2]int) int {
	trees := 0
	var x, y int
	for y < b.Len() {
		if b.IsTree(x, y) {
			trees++
		}
		x += slope[0]
		y += slope[1]
	}
	return trees
}

type board struct {
	width int
	lines [][]bool
}

func (b *board) Read(line string) error {
	line = strings.TrimSpace(line)
	if b.width == 0 {
		b.width = len(line)
	}
	var newLines []bool
	for _, c := range line {
		if c != '#' && c != '.' {
			return fmt.Errorf("unexpected char %c in line %q", c, line)
		}
		newLines = append(newLines, c == '#')
	}
	b.lines = append(b.lines, newLines)
	return nil
}

func (b board) IsTree(x, y int) bool {
	return b.lines[y][x%b.width]
}

func (b board) Len() int {
	return len(b.lines)
}
