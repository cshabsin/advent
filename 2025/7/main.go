package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	row, col int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input.txt: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	splitters := map[point]bool{}
	var row, maxRow int
	var start point
	for scanner.Scan() {
		line := scanner.Text()
		for col, ch := range line {
			if ch == 'S' {
				start = point{row, col}
			}
			if ch == '^' {
				splitters[point{row, col}] = true
				maxRow = row
			}
		}
		row++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	fmt.Println(countSplits(start, splitters, maxRow))
	fmt.Println(countSplits2(start, splitters, maxRow))
}

func countSplits(start point, splitters map[point]bool, maxRow int) int {
	var count int
	row := 0
	current := map[int]bool{start.col: true}
	for {
		if row > maxRow {
			break
		}
		next := map[int]bool{}
		row++
		for col := range current {
			if splitters[point{row, col}] {
				count++
				next[col-1] = true
				next[col+1] = true
			} else {
				next[col] = true
			}
		}
		current = next
	}
	return count
}

func countSplits2(start point, splitters map[point]bool, maxRow int) int {
	row := 0
	current := map[int]int{start.col: 1}
	for {
		if row > maxRow {
			break
		}
		next := map[int]int{}
		row++
		for col, num := range current {
			fmt.Println("checking ", point{row, col})
			if splitters[point{row, col}] {
				next[col-1] = next[col-1] + num
				next[col+1] = next[col+1] + num
			} else {
				next[col] = next[col] + num
			}
		}
		fmt.Println(next)
		current = next
	}
	var count int
	for _, n := range current {
		count += n
	}
	return count
}
