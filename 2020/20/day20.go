package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cshabsin/advent/2020/tile"
)

func main() {
	tiles, err := tile.ReadFile("testinput.txt")
	if err != nil {
		log.Fatal(err)
	}

	tileNum := 1951 // 3413
	for rot := 0; rot < 4; rot++ {
		gf := gridFiller{
			tiles:     tiles,
			usedTiles: map[int]bool{},
		}
		tiles.Rotate(tileNum, 3)
		tg := gf.Fill(tileNum)
		fmt.Println(tg)
		fmt.Println(tg.tiles)

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
		tiles.Rotate(tileNum, 1)
	}
}

func day20a(tiles map[int]*tile.Tile, edgeMap map[int][]int) []int {
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
	tiles                          map[int]map[int]int
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
	if tg.tiles == nil {
		tg.tiles = map[int]map[int]int{}
	}
	if tg.tiles[r] == nil {
		tg.tiles[r] = map[int]int{}
	}
	tg.tiles[r][c] = t.ID()
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
	tiles     *tile.Map
}

func (g gridFiller) Fill(tileNum int) *tileGrid {
	tg := &tileGrid{grid: map[int]map[int]bool{}}
	g.fillTile(tg, 0, 0, tileNum)
	return tg
}

func (g gridFiller) fillTile(tg *tileGrid, r, c, tileNum int) {
	g.usedTiles[tileNum] = true
	tile := g.tiles.GetTile(tileNum)
	fmt.Printf("placing tile at %d, %d:\n%v\n", r, c, tile)
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
	fmt.Printf("doing neighbor %d from tile %d [%d, %d]\n", e, tile.ID(), r, c)
	edgeMatch := tile.ReadEdge(e)
	oppositeEdge := (e + 2) % 4
	fmt.Printf("ensuring edge %d matches %d\n", oppositeEdge, edgeMatch)
	for !g.tiles.GetTile(n).EdgeMatches(oppositeEdge, edgeMatch) {
		g.tiles.Rotate(n, 1)
	}

	g.fillTile(tg, r, c, n)
}
