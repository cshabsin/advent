package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	Day18("input.txt", false)
	Day18("input.txt", true)
}

type foo [100]bool

func parseFoo(s string) (foo, error) {
	var f foo
	for i := 0; i < 100; i++ {
		if s[i] == '#' {
			f[i] = true
		}
	}
	return f, nil
}

// Day18 solves part 1 of day 18
func Day18(fn string, part bool) {
	ch, err := readinp.Read(fn, parseFoo)
	if err != nil {
		log.Fatal(err)
	}
	var board [100][100]bool
	var i int
	for linestr := range ch {
		line, err := linestr.Get()
		if err != nil {
			log.Fatal(err)
		}
		board[i] = line
		i++
	}

	for step := 0; step < 100; step++ {
		var newboard [100][100]bool
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				switch countN(board, x, y) {
				case 2:
					newboard[x][y] = board[x][y]
				case 3:
					newboard[x][y] = true
				}
			}
		}
		if part {
			newboard[0][0] = true
			newboard[0][99] = true
			newboard[99][0] = true
			newboard[99][99] = true
		}
		board = newboard
	}
	var c int
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			if board[x][y] {
				c++
			}
		}
	}
	fmt.Println(c)
}

func printBoard(board [100][100]bool) {
	for _, r := range board {
		for _, c := range r {
			ch := "."
			if c {
				ch = "#"
			}
			fmt.Print(ch)
		}
		fmt.Print("\n")
	}
	fmt.Println("---")
}

func countN(board [100][100]bool, x, y int) int {
	var cnt int
	for dx := -1; dx <= 1; dx++ {
		if x+dx < 0 || x+dx >= 100 {
			continue
		}
		for dy := -1; dy <= 1; dy++ {
			if y+dy < 0 || y+dy >= 100 {
				continue
			}
			if dy == 0 && dx == 0 {
				continue
			}
			if board[x+dx][y+dy] {
				cnt++
			}
		}
	}
	return cnt
}
