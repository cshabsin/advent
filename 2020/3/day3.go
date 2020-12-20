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
	trees := 0
	for y := 0; y < b.Len(); y++ {
		if b.IsTree(y*3, y) {
			trees++
		}
	}
	fmt.Println(trees)
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
