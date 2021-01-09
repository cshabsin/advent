package main

import (
	"io"
	"log"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

type tile struct {
	id       int
	pixels   [][]bool
	edges    []int
	matches  []bool
	rotation int
}

func readTile(ch chan readinp.Line) *tile {
	allVals := make([][]bool, 10)
	for i := 0; i < 10; i++ {
		line, err := read(ch)
		if err != nil {
			log.Fatal(err)
		}
		allValLine := make([]bool, 10)
		for j, c := range line {
			allValLine[j] = c == '#'
		}
		allVals[i] = allValLine
	}
	var edges []int
	for i := 0; i < 4; i++ {
		edges = append(edges, readEdge(allVals, i))
	}

	var pixels [][]bool
	for i := 1; i < 9; i++ {
		pixels = append(pixels, allVals[i][1:8])
	}

	return &tile{
		pixels: pixels,
		edges:  edges,
	}
}

func (t tile) readEdge(e int) int {
	return t.edges[(e+t.rotation)%4]
}

func readEdge(allVals [][]bool, e int) int {
	switch e {
	case 0: // top
		total := 0
		for j := 0; j < 10; j++ {
			total *= 2
			if allVals[0][j] {
				total++
			}
		}
		return total
	case 1: // left
		total := 0
		for j := 9; j >= 0; j-- {
			total *= 2
			if allVals[j][0] {
				total++
			}
		}
		return total
	case 2: // bottom
		total := 0
		for j := 9; j >= 0; j-- {
			total *= 2
			if allVals[9][j] {
				total++
			}
		}
		return total
	case 3: // right
		total := 0
		for j := 0; j < 10; j++ {
			total *= 2
			if allVals[j][9] {
				total++
			}
		}
		return total
	}
	log.Fatal("bad e", e)
	return -1
}

func edgeDual(a int) int {
	var b int
	for i := 0; i < 10; i++ {
		b = b*2 + (a >> i & 1)
	}
	return b
}

func edgesMatch(a, b int) bool {
	for i := 0; i < 10; i++ {
		if a>>i&1 != b>>(9-i)&1 {
			return false
		}
	}
	return true
}

func read(ch chan readinp.Line) (string, error) {
	line, ok := <-ch
	if !ok {
		return "", io.EOF
	}
	if line.Error != nil {
		return "", line.Error
	}
	return strings.TrimSpace(*line.Contents), nil
}
