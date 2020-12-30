package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	tiles := map[int]*tile{}
	edgeMap := map[int][]int{}
	for {
		decl, err := read(ch)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		tile, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(decl, "Tile "), ":"))
		if err != nil {
			log.Fatal(err)
		}
		tiles[tile] = readTile(ch)
		for i := 0; i < 4; i++ {
			edge := tiles[tile].readEdge(i)
			edgeMap[edge] = append(edgeMap[edge], tile)
			edgeMap[edgeDual(edge)] = append(edgeMap[edgeDual(edge)], tile)
		}
		_, err = read(ch) // skip blank line
		if err != nil {
			log.Fatal(err)
		}
	}
	numMatches := map[int][]int{} // number of matches for each tile's four edges
	for tileNum, tile := range tiles {
		for i := 0; i < 4; i++ {
			numMatches[tileNum] = append(numMatches[tileNum], len(edgeMap[tile.readEdge(i)])-1)
		}
	}
	total := 1
	for tile, matchList := range numMatches {
		matches := 0
		for _, num := range matchList {
			matches += num
		}
		if matches == 2 {
			total *= tile
		}
	}
	fmt.Println(total)
}

type tile struct {
	pixels [][]bool
	edges  []int
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
	return t.edges[e%4]
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
