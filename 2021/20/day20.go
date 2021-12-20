package main

import (
	"fmt"

	"github.com/cshabsin/advent/commongen/board"
	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	// part1("sample.txt")
	part1("input.txt")
}

func part1(fn string) {
	ch := readinp.MustReadConsumer[image](fn, &parser{})
	img := (<-ch).MustGet()

	for i := 0; i < 50; i++ {
		img = img.expand()
		fmt.Println(i, img.count())
	}
}

type image struct {
	algo  [512]boolS
	board board.Board[boolS]
}

func (i image) expand() image {
	img := image{algo: i.algo}
	for range i.board {
		var line []boolS
		for range i.board[0] {
			line = append(line, false)
		}
		img.board = append(img.board, line)
	}

	for boardR := 1; boardR < len(i.board)-1; boardR++ {
		for boardC := 1; boardC < len(i.board[0])-1; boardC++ {
			var b int
			for r := boardR - 1; r <= boardR+1; r++ {
				for c := boardC - 1; c <= boardC+1; c++ {
					var bit int
					if i.board.Get(r, c) {
						bit = 1
					}
					b = b*2 + bit
				}
			}
			img.board.Set(boardR, boardC, i.algo[b])
		}
	}
	return img
}

func (i image) count() int {
	var count int
	for _, co := range i.board.AllCoords() {
		if co.R() == 0 || co.R() == len(i.board)-1 || co.C() == 0 || co.C() == len(i.board[0])-1 {
			continue
		}
		if i.board.GetCoord(co) {
			count++
		}
	}
	return count
}

type parser struct {
	current   *image
	wantBlank bool
}

const padding = 5000

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
		for i := 0; i < padding; i++ {
			var blankLine []boolS
			for j := 0; j < len(line)+(2*padding); j++ {
				blankLine = append(blankLine, false)
			}
			p.current.board = append(p.current.board, blankLine)
		}
	}
	var boardLine []boolS
	for i := 0; i < padding; i++ {
		boardLine = append(boardLine, false)
	}
	for _, c := range line {
		boardLine = append(boardLine, c == '#')
	}
	for i := 0; i < padding; i++ {
		boardLine = append(boardLine, false)
	}
	p.current.board = append(p.current.board, boardLine)
	return image{}, false, nil
}

func (p *parser) Done() (*image, bool, error) {
	for i := 0; i < padding; i++ {
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
