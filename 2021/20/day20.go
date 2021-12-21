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
	fmt.Println(img)
	fmt.Println("---")

	for i := 0; i < 50; i++ {
		img = img.expand()
		fmt.Println(img)
		fmt.Println(i, img.count())
		fmt.Println("---")
	}
}

type image struct {
	algo                           [512]boolS
	minRow, maxRow, minCol, maxCol int
	board                          map[board.Coord]boolS
	worldIsLive                    bool // gross hack for part 2
}

func (i image) String() string {
	out := fmt.Sprintf("[%d-%d][%d-%d]\n", i.minRow, i.maxRow, i.minCol, i.maxCol)
	for row := i.minRow; row <= i.maxRow; row++ {
		var line string
		for col := i.minCol; col <= i.maxCol; col++ {
			if i.get(row, col) {
				line += "#"
			} else {
				line += "."
			}
		}
		out += line + "\n"
	}
	return out
}

func (i *image) set(r, c int, val boolS) {
	if r < i.minRow {
		i.minRow = r
	}
	if r > i.maxRow {
		i.maxRow = r
	}
	if c < i.minCol {
		i.minCol = c
	}
	if c > i.maxCol {
		i.maxCol = c
	}
	i.board[board.MakeCoord(r, c)] = val
}

func (i image) get(r, c int) boolS {
	if r < i.minRow || r > i.maxRow || c < i.minCol || c > i.maxCol {
		return boolS(i.worldIsLive)
	}
	return i.board[board.MakeCoord(r, c)]
}

func (i image) expand() image {
	img := image{algo: i.algo, board: make(map[board.Coord]boolS)}

	for boardR := i.minRow - 3; boardR <= i.maxRow+3; boardR++ {
		for boardC := i.minCol - 3; boardC <= i.maxCol+3; boardC++ {
			var b int
			for r := boardR - 1; r <= boardR+1; r++ {
				for c := boardC - 1; c <= boardC+1; c++ {
					var bit int
					if i.get(r, c) {
						bit = 1
					}
					b = b*2 + bit
				}
			}
			if i.algo[b] {
				img.set(boardR, boardC, true)
			}
		}
	}
	img.worldIsLive = !i.worldIsLive
	return img
}

func (i image) count() int {
	var count int
	for r := i.minRow; r <= i.maxRow; r++ {
		for c := i.minCol; c <= i.maxCol; c++ {
			if i.get(r, c) {
				count++
			}
		}
	}
	return count
}

type parser struct {
	row     int
	current *image
}

func (p *parser) Parse(line string) (image, bool, error) {
	if p.current == nil {
		p.current = &image{
			board: make(map[board.Coord]boolS),
		}
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

	for col, ch := range line {
		p.current.set(p.row, col, ch == '#')
	}
	p.row++
	return image{}, false, nil
}

func (p *parser) Done() (*image, bool, error) {
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
