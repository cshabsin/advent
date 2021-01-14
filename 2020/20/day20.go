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

	for rot := 0; rot < 4; rot++ {
		gf := gridFiller{
			tiles:     tiles,
			usedTiles: map[int]bool{},
		}
		tg := gf.Fill()
		//fmt.Println(tg)

		monster := []struct{ r, c int }{
			{0, 18},
			{1, 0},
			{1, 5},
			{1, 6},
			{1, 11},
			{1, 12},
			{1, 17},
			{1, 18},
			{1, 19},
			{2, 1},
			{2, 4},
			{2, 7},
			{2, 10},
			{2, 13},
			{2, 16},
		}
		// monsterSpot[r*10000+c] = true if it's a monster spot
		monsterSpots := map[int]bool{}
		for y := tg.minRow; y < tg.maxRow; y++ {
			for x := tg.minCol; x < tg.maxCol; x++ {
				found := true
				for _, rc := range monster {
					if !tg.get(y+rc.r, x+rc.c) {
						found = false
						break
					}
				}
				if found {
					for _, rc := range monster {
						monsterSpots[(y+rc.r)*10000+x+rc.c] = true
					}
				}
			}
		}
		fmt.Println(tg.allRoughness(), len(monsterSpots))
		tiles[3413].Rotate(1)
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

// map[row]map[col]pixel
type tileGrid struct {
	minRow, maxRow, minCol, maxCol int
	grid                           map[int]map[int]bool
}

func (tg tileGrid) get(r, c int) bool {
	if tg.grid[r] == nil {
		return false
	}
	return tg.grid[r][c]
}

func (tg tileGrid) allRoughness() int {
	roughness := 0
	for _, row := range tg.grid {
		roughness += len(row)
	}
	return roughness
}

func (tg *tileGrid) setTile(t *tile.Tile, r, c int) {
	if tg.minRow > r*8 {
		tg.minRow = r * 8
	}
	if tg.maxRow < (r+1)*8 {
		tg.maxRow = (r + 1) * 8
	}
	if tg.minCol > c*8 {
		tg.minCol = c * 8
	}
	if tg.maxCol < (c+1)*8 {
		tg.maxCol = (c + 1) * 8
	}
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if t.Get(x, y) {
				if row := tg.grid[r*8+y]; row == nil {
					tg.grid[r*8+y] = map[int]bool{c*8 + x: true}
				} else {
					row[c*8+x] = true
				}
			}
		}
	}
}

func (tg tileGrid) String() string {
	var b strings.Builder
	for row := tg.minRow; row < tg.maxRow; row++ {
		for col := tg.minCol; col < tg.maxCol; col++ {
			if tg.grid[row] != nil && tg.grid[row][col] {
				b.WriteString("X ")
			} else {
				b.WriteString(". ")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

type gridFiller struct {
	usedTiles map[int]bool
	tiles     tileMap
}

func (g gridFiller) Fill() *tileGrid {
	tg := &tileGrid{grid: map[int]map[int]bool{}}
	tileNum := 3413
	g.usedTiles[tileNum] = true
	g.fillTile(tg, 0, 0, tileNum)
	return tg
}

func (g gridFiller) fillTile(tg *tileGrid, r, c, tileNum int) {
	tile := g.tiles[tileNum]
	tg.setTile(tile, r, c)
	g.doNeighbor(tg, r-1, c, tile, 0)
	g.doNeighbor(tg, r, c-1, tile, 1)
	g.doNeighbor(tg, r+1, c, tile, 2)
	g.doNeighbor(tg, r, c+1, tile, 3)
}

func (g gridFiller) doNeighbor(tg *tileGrid, r, c int, tile *tile.Tile, e int) {
	if !tile.HasNeighbor(e) {
		return
	}
	n := tile.GetNeighbor(e)
	if g.usedTiles[n] {
		return
	}
	edgeMatch := tile.ReadEdge(e)
	oppositeEdge := (e + 2) % 4
	for !g.tiles[n].EdgeMatches(oppositeEdge, edgeMatch) {
		g.tiles[n].Rotate(1)
	}

	g.usedTiles[n] = true
	g.fillTile(tg, r, c, n)
}
