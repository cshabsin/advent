package main

import (
	"fmt"
	"io"
	"log"
	"strings"

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
		tid := nextTile.ID()
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

	for _, tile := range tiles {
		for i := 0; i < 4; i++ {
			tile.SetNeighborFromEdgeMap(edgeMap)
		}
	}
	fmt.Println(tiles[2803])
	tiles[2803].Rotate(1)
	fmt.Println(tiles[2803])
	tiles[2803].Rotate(1)
	fmt.Println(tiles[2803])
	tiles[2803].Rotate(1)
	fmt.Println(tiles[2803])
	tiles[2803].Rotate(1)
	fmt.Println(tiles[2803])

	tileGrid := [12][12]*tile.Tile{}
	for _, tile := range tiles {
		if tile.NumNeighbors() != 2 {
			continue
		}
		tileGrid[0][0] = tile
		if tile.HasNeighbor(0) && tile.HasNeighbor(1) {
			tile.Rotate(2)
		} else if tile.HasNeighbor(0) && tile.HasNeighbor(3) {
			tile.Rotate(3)
		} else if tile.HasNeighbor(2) && tile.HasNeighbor(1) {
			tile.Rotate(1)
		}
		break
	}

	// top left is set, now fill in the rest.
	fmt.Println("topleft:\n", tileGrid[0][0].String())
	col := 1
	usedTiles := map[int]bool{tileGrid[0][0].ID(): true}
	for col < 12 {
		rightEdge := tileGrid[col-1][0].ReadEdge(3)
		matches := edgeMap[rightEdge]
		// fmt.Println("matches for", rightEdge, tile.EdgeDual(rightEdge), ":", matches)
		var match int
		for _, match = range matches {
			if usedTiles[match] {
				continue
			}
			break
		}
		if usedTiles[match] {
			log.Fatalf("no unused match tile for rightEdge %d", rightEdge)
		}
		usedTiles[match] = true

		edge := 0
		for edge < 4 {
			if tiles[match].ReadEdge(edge) == rightEdge {
				break
			}
			if tiles[match].ReadEdge(edge) == tile.EdgeDual(rightEdge) {
				break
			}
			edge++
		}
		if edge == 4 {
			log.Fatalf("no match for right edge %d (dual %d) in tile %v", rightEdge, tile.EdgeDual(rightEdge), tiles[match])
		}
		// fmt.Println("edge", edge, "match before rotate:\n", tiles[match])
		tiles[match].Rotate(5 - edge)
		// fmt.Println("match after rotate", 5-edge, ":\n", tiles[match].String())
		fmt.Println("[ 0 ,", col, "]:\n", tiles[match])
		tileGrid[col][0] = tiles[match]
		col++
	}
	row := 1
	for row < 12 {
		for col := 0; col < 12; col++ {
			topEdge := tileGrid[col][row-1].ReadEdge(2)
			matches := edgeMap[topEdge]
			fmt.Println("matches for", topEdge, ":", matches, "/", edgeMap[tile.EdgeDual(topEdge)])
			var match int
			for _, match = range matches {
				if usedTiles[match] {
					continue
				}
				break
			}
			if usedTiles[match] {
				log.Fatalf("no unused match tile for topEdge %d (%d, %d)\n%v\n%v", topEdge, row, col, tileGrid[col][row-1], tileGrid[0][0])
			}
			usedTiles[match] = true

			edge := 0
			for edge < 4 {
				if tiles[match].ReadEdge(edge) == topEdge {
					break
				}
				if tiles[match].ReadEdge(edge) == tile.EdgeDual(topEdge) {
					break
				}
				edge++
			}
			if edge == 4 {
				log.Fatalf("no match for top edge %d (dual %d) in tile %v (%d, %d)", topEdge, tile.EdgeDual(topEdge), tiles[match], row, col)
			}
			tiles[match].Rotate(4 - edge)
			tileGrid[col][row] = tiles[match]
		}
		row++
	}
	var grid [96][96]bool
	for row := 0; row < 12; row++ {
		for col := 0; col < 12; col++ {
			for x := 0; x < 8; x++ {
				for y := 0; y < 8; y++ {
					grid[row*8+y][col*8+x] = tileGrid[col][row].Get(x, y)
				}
			}
		}
	}
	for r := 0; r < 96; r++ {
		var b strings.Builder
		for c := 0; c < 96; c++ {
			if grid[r][c] {
				b.WriteString("X")
			} else {
				b.WriteString(".")
			}
		}
		fmt.Println(b.String)
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
