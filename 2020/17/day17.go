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
	b := board{vals: map[int]map[int]map[int]bool{}}
	row := 0
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		col := 0
		for _, c := range line.Value() {
			if c == '#' {
				b.set(col, row, 0, true)
			}
			col++
		}
		row++
	}
	fmt.Println("initial", b)
	for i := 0; i < 6; i++ {
		b = b.advance()
		fmt.Println(i, b)
	}
	total := 0
	for _, col := range b.vals {
		for _, row := range col {
			for _, alive := range row {
				if alive {
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

type board struct {
	minX, minY, minZ int
	maxX, maxY, maxZ int
	vals             map[int]map[int]map[int]bool
}

func (b board) get(x, y, z int) bool {
	if b.vals[x] == nil {
		return false
	}
	if b.vals[x][y] == nil {
		return false
	}
	return b.vals[x][y][z]
}

func (b *board) set(x, y, z int, alive bool) {
	if x < b.minX {
		b.minX = x
	}
	if y < b.minY {
		b.minY = y
	}
	if z < b.minZ {
		b.minZ = z
	}
	if x > b.maxX {
		b.maxX = x
	}
	if y > b.maxY {
		b.maxY = y
	}
	if z > b.maxZ {
		b.maxZ = z
	}
	if b.vals[x] == nil {
		if !alive {
			return
		}
		b.vals[x] = map[int]map[int]bool{}
	}
	if b.vals[x][y] == nil {
		if !alive {
			return
		}
		b.vals[x][y] = map[int]bool{}
	}
	b.vals[x][y][z] = alive
}

func (b board) advance() board {
	newBoard := board{vals: map[int]map[int]map[int]bool{}}
	for x := b.minX - 1; x <= b.maxX+1; x++ {
		for y := b.minY - 1; y <= b.maxY+1; y++ {
			for z := b.minZ - 1; z <= b.maxZ+1; z++ {
				neighbors := b.countNeighbors(x, y, z)
				if neighbors == 3 {
					newBoard.set(x, y, z, true)
				} else if neighbors == 2 {
					newBoard.set(x, y, z, b.get(x, y, z))
				}
			}
		}
	}
	return newBoard
}

func (b board) countNeighbors(x, y, z int) int {
	neighbors := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				if b.get(x+dx, y+dy, z+dz) {
					neighbors++
				}
			}
		}
	}
	return neighbors
}
