package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	board := board{}
	for line := range ch {
		if line.Error != nil {
			log.Fatal(err)
		}
		if err := board.addLine(strings.TrimSpace(*line.Contents)); err != nil {
			log.Fatal(err)
		}
	}
	board.advance()
	fmt.Println(board)
	for {
		changed := board.advance()
		if !changed {
			break
		}
	}
	fmt.Println(board.countOccupied())
}

type cell int

const (
	floor cell = iota
	chair
	occupied
)

type board [][]cell

func (b *board) addLine(line string) error {
	var newBoardLine []cell
	for _, c := range line {
		switch c {
		case '.':
			newBoardLine = append(newBoardLine, floor)
		case 'L':
			newBoardLine = append(newBoardLine, chair)
		case '#':
			newBoardLine = append(newBoardLine, occupied)
		default:
			return fmt.Errorf("invalid character %v", c)
		}
	}
	if len(*b) != 0 && len((*b)[0]) != len(newBoardLine) {
		return fmt.Errorf("invalid line length for %q", line)
	}
	*b = append(*b, newBoardLine)
	return nil
}

func (b board) copy() (out [][]cell) {
	for _, boardLine := range b {
		var newBoardLine []cell
		for _, cell := range boardLine {
			newBoardLine = append(newBoardLine, cell)
		}
		out = append(out, newBoardLine)
	}
	return
}

func (b board) String() string {
	var buf bytes.Buffer
	for _, boardLine := range b {
		for _, cell := range boardLine {
			switch cell {
			case floor:
				buf.WriteRune('.')
			case chair:
				buf.WriteRune('L')
			case occupied:
				buf.WriteRune('#')
			default:
				buf.WriteRune('?')
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

func (b board) numRows() int {
	return len(b)
}

func (b board) numCols() int {
	return len(b[0])
}

func (b board) at(r, c int) cell {
	if r < 0 || c < 0 || r >= b.numRows() || c >= b.numCols() {
		return floor
	}
	return b[r][c]
}

func (b board) isOccupied(r, c int) bool {
	return b.at(r, c) == occupied
}

func (b board) countNeighbors(r, c int) (count int) {
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			if b.isOccupied(r+dr, c+dc) {
				count++
			}
		}
	}
	return
}

func (b *board) advance() (changed bool) {
	newBoard := b.copy()
	for r := 0; r < b.numRows(); r++ {
		for c := 0; c < b.numCols(); c++ {
			if b.at(r, c) == floor {
				continue
			}
			n := b.countNeighbors(r, c)
			if b.isOccupied(r, c) && n >= 4 {
				changed = true
				newBoard[r][c] = chair
			} else if !b.isOccupied(r, c) && n == 0 {
				changed = true
				newBoard[r][c] = occupied
			}
		}
	}
	*b = newBoard
	return
}

func (b board) countOccupied() (occupied int) {
	for r := 0; r < b.numRows(); r++ {
		for c := 0; c < b.numCols(); c++ {
			if b.isOccupied(r, c) {
				occupied++
			}
		}
	}
	return
}
