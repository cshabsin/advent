package main

import (
	"fmt"

	"github.com/cshabsin/advent/commongen/board"
	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	part1("sample.txt")
}

func part1(fn string) {
	ch := readinp.MustReadConsumer[image](fn, &parser{})
	img := (<-ch).MustGet()
	fmt.Print(img.board)
}

type image struct {
	algo  [512]bool
	board board.Board[boolS]
}

type parser struct {
	current   *image
	wantBlank bool
}

func (p *parser) Parse(line string) (image, bool, error) {
	if p.current == nil {
		p.current = &image{}
		p.wantBlank = true
		for i, c := range line {
			if c == '#' {
				p.current.algo[i] = true
			}
		}
		return image{}, false, nil
	}
	if line == "" {
		return image{}, false, nil
	}
	if p.wantBlank {
		p.wantBlank = false
		for i := 0; i < 2; i++ {
			var blankLine []boolS
			for j := 0; j < len(line)+4; j++ {
				blankLine = append(blankLine, false)
			}
			p.current.board = append(p.current.board, blankLine)
		}
	}
	boardLine := []boolS{false, false}
	for _, c := range line {
		boardLine = append(boardLine, c == '#')
	}
	boardLine = append(boardLine, false, false)
	p.current.board = append(p.current.board, boardLine)
	return image{}, false, nil
}

func (p *parser) Done() (*image, bool, error) {
	for i := 0; i < 2; i++ {
		var blankLine []boolS
		for j := 0; j < len(p.current.board[0]); j++ {
			blankLine = append(blankLine, false)
		}
		p.current.board = append(p.current.board, blankLine)
	}
	return p.current, true, nil
}

type boolS bool

func (b boolS) String() string {
	if b {
		return "#"
	}
	return "."
}

func (b boolS) Delimiter() string {
	return ""
}
