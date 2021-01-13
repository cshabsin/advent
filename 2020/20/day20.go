package main

import (
	"fmt"
	"io"
	"log"

	"github.com/cshabsin/advent/2020/tile"
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
		nextTile, err := tile.Read(ch)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		tid := nextTile.Id()
		tiles[tid] = nextTile
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < 4; i++ {
			edge := nextTile.ReadEdge(i)
			edgeMap[edge] = append(edgeMap[edge], tid)
			edgeMap[tile.EdgeDual(edge)] = append(edgeMap[tile.EdgeDual(edge)], tid)
		}
		_, err = tile.ReadLine(ch) // skip blank line
		if err != nil {
			log.Fatal(err)
		}
	}
	// cornerTiles := day20a(tiles, edgeMap)

	numMatches := map[int][]int{} // number of matches for each tile's four edges
	for tileNum, tile := range tiles {
		for i := 0; i < 4; i++ {
			numMatches[tileNum] = append(numMatches[tileNum], len(edgeMap[tile.ReadEdge(i)])-1)
		}
	}
	tileGrid := [12][12]*tile.Tile{}
	for tile, matchList := range numMatches {
		matchCount := 0
		for _, matches := range matchList {
			matchCount += matches
		}
		if matchCount == 2 {
			tileGrid[0][0] = tiles[tile]
			if matchList[0] != 0 && matchList[3] != 0 {
				tiles[tile].Rotate(1)
			} else if matchList[3] != 0 && matchList[2] != 0 {
				tiles[tile].Rotate(2)
			} else if matchList[2] != 0 && matchList[1] != 0 {
				tiles[tile].Rotate(3)
			}
			break
		}
	}

	// top left is set, now fill in the rest.
	fmt.Println("topleft\n", tileGrid[0][0].String())
	col := 1
	usedTiles := map[int]bool{tileGrid[0][0].Id(): true}
	usedEdges := map[int]bool{}
	for col < 12 {
		rightEdge := tileGrid[col-1][0].ReadEdge(1)
		usedEdges[rightEdge] = true
		matches := edgeMap[rightEdge]
		var match int
		for _, match = range matches {
			if usedTiles[match] {
				continue
			}
			break
		}

		edge := 0
		for {
			if tiles[match].ReadEdge(edge) == rightEdge {
				break
			}
			edge++
		}
		tiles[match].Rotate(1 - edge)
		fmt.Println(tiles[match])
		tileGrid[col][0] = tiles[match]
		fmt.Println(tileGrid[col][0])
		col++
	}
}

type tileMap map[int]*tile.Tile

func day20a(tiles tileMap, edgeMap map[int][]int) []int {
	// day 20a
	numMatches := map[int][]int{} // number of matches for each tile's four edges
	for tileNum, tile := range tiles {
		for i := 0; i < 4; i++ {
			numMatches[tileNum] = append(numMatches[tileNum], len(edgeMap[tile.ReadEdge(i)])-1)
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
