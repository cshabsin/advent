package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	accessible := 0
	var board [][]bool
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		var row []bool
		for _, ch := range line.Value() {
			row = append(row, ch == '@')
		}
		board = append(board, row)
	}
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if !get(board, r, c) {
				continue
			}
			if isAccessible(board, r, c) {
				accessible++
			}
		}
	}
	fmt.Println(accessible)

	var totalRemoved int
	for {
		var toRemove [][]bool
		var removedThisIteration int
		for r := 0; r < len(board); r++ {
			var removeLine []bool
			for c := 0; c < len(board[r]); c++ {
				removeLine = append(removeLine, get(board, r, c) && isAccessible(board, r, c))
			}
			toRemove = append(toRemove, removeLine)
		}

		for r := 0; r < len(board); r++ {
			for c := 0; c < len(board[r]); c++ {
				if toRemove[r][c] {
					board[r][c] = false
					totalRemoved++
					removedThisIteration++
				}
			}
		}
		if removedThisIteration == 0 {
			break
		}
	}
	fmt.Println(totalRemoved)
}

func isAccessible(board [][]bool, r, c int) bool {
	var neighbors int
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			if get(board, r+dr, c+dc) {
				neighbors++
			}
		}
	}
	return neighbors < 4
}

func get(board [][]bool, r, c int) bool {
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[r]) {
		return false
	}
	return board[r][c]
}
