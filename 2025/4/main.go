// Package main solves a grid-based puzzle.
// The puzzle involves a board of cells, which can be either occupied ('@') or empty.
// The program first calculates how many occupied cells are "accessible". A cell
// is accessible if it has fewer than 4 occupied neighbors (out of 8 possible).
// Then, it simulates a process where all accessible cells are removed in discrete
// steps, until no more cells can be removed. It reports the total number of
// cells removed during this simulation.
package main

import (
	"fmt"
	"log"

	"github.com/cshabsin/advent/common/readinp"
)

const (
	// occupiedCell represents the character for an occupied cell in the input.
	occupiedCell = '@'
	// accessibilityThreshold is the number of neighbors at which a cell is no longer accessible.
	accessibilityThreshold = 4
)

func main() {
	board, err := parseBoard("input.txt")
	if err != nil {
		log.Fatalf("failed to parse board: %v", err)
	}

	// Part 1: Count initially accessible cells.
	initiallyAccessible := countAccessible(board)
	fmt.Println(initiallyAccessible)

	// Part 2: Simulate removal of accessible cells and count them.
	// This simulation will modify the board.
	totalRemoved := runRemovalSimulation(board)
	fmt.Println(totalRemoved)
}

// parseBoard reads a file and converts it into a boolean grid.
func parseBoard(filename string) ([][]bool, error) {
	lineCh, err := readinp.Read(filename)
	if err != nil {
		return nil, err
	}

	var board [][]bool
	for line := range lineCh {
		if line.Error != nil {
			return nil, line.Error
		}
		var row []bool
		for _, char := range line.Value() {
			row = append(row, char == occupiedCell)
		}
		board = append(board, row)
	}
	return board, nil
}

// countAccessible iterates over the board and counts how many occupied cells are accessible.
func countAccessible(board [][]bool) int {
	accessibleCount := 0
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if getCell(board, r, c) && isAccessible(board, r, c) {
				accessibleCount++
			}
		}
	}
	return accessibleCount
}

// runRemovalSimulation simulates the process of removing accessible cells until a stable state is reached.
// It returns the total number of cells removed.
func runRemovalSimulation(board [][]bool) int {
	var totalRemoved int
	for {
		type coord struct{ r, c int }
		var toRemove []coord

		// First pass: identify all cells to be removed in this step.
		// We can't modify the board while iterating, as it would affect neighbor counts for subsequent cells in the same step.
		for r := 0; r < len(board); r++ {
			for c := 0; c < len(board[r]); c++ {
				if getCell(board, r, c) && isAccessible(board, r, c) {
					toRemove = append(toRemove, coord{r, c})
				}
			}
		}

		if len(toRemove) == 0 {
			// The board is stable, no more cells can be removed.
			break
		}

		// Second pass: apply the removals.
		for _, cell := range toRemove {
			board[cell.r][cell.c] = false
		}
		totalRemoved += len(toRemove)
	}
	return totalRemoved
}

// isAccessible checks if a cell at a given coordinate is accessible.
// A cell is accessible if it has fewer than accessibilityThreshold occupied neighbors.
// The 8 neighbors (horizontally, vertically, and diagonally) are checked.
func isAccessible(board [][]bool, r, c int) bool {
	var neighborCount int
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				// Don't count the cell itself.
				continue
			}
			if getCell(board, r+dr, c+dc) {
				neighborCount++
			}
		}
	}
	return neighborCount < accessibilityThreshold
}

// getCell safely retrieves the state of a cell on the board.
// It returns false for coordinates that are out of bounds.
func getCell(board [][]bool, r, c int) bool {
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[r]) {
		return false
	}
	return board[r][c]
}
