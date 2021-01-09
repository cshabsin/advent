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
	tiles := tileMap{}
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
		tile.id = tile
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
	// cornerTiles := day20a(tiles, edgeMap)

	numMatches := map[int][]int{} // number of matches for each tile's four edges
	for tileNum, tile := range tiles {
		for i := 0; i < 4; i++ {
			numMatches[tileNum] = append(numMatches[tileNum], len(edgeMap[tile.readEdge(i)])-1)
		}
	}
	tileGrid := [12][12]*tile{}
	for tile, matchList := range numMatches {
		matchCount := 0
		for _, matches := range matchList {
			matchCount += matches
		}
		if matchCount == 2 {
			tileGrid[0][0] = tiles[tile]
			if matchList[0] != 0 && matchList[3] != 0 {
				tiles[tile].rotation = 1
			} else if matchList[3] != 0 && matchList[2] != 0 {
				tiles[tile].rotation = 2
			} else if matchList[2] != 0 && matchList[1] != 0 {
				tiles[tile].rotation = 3
			}
			break
		}
	}
	col := 1
	row := 0
	for {
		if col == 0 {
			// match up
			nmBottom := numMatches[]
		} else {

		}
	}
}

type tileMap map[int]*tile

func day20a(tiles tileMap, edgeMap map[int][]int) []int {
	// day 20a
	numMatches := map[int][]int{} // number of matches for each tile's four edges
	for tileNum, tile := range tiles {
		for i := 0; i < 4; i++ {
			numMatches[tileNum] = append(numMatches[tileNum], len(edgeMap[tile.readEdge(i)])-1)
		}
	}
	total := 1
	var corners []int
	for tile, matchList := range numMatches {
		matches := 0
		for _, num := range matchList {
			matches += num
		}
		if matches == 2 {
			total *= tile
			corners = append(corners, tile)
		}
	}
	fmt.Println(total)
	return corners
}
