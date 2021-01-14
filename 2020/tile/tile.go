package tile

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

// Tile is a tile from day 20 of advent of code 2020.
type Tile struct {
	id           int
	pixels       [][]bool
	edges        []int
	matches      []bool
	rotation     int
	neighbors    []int // same indeces as edges, 0 for none
	numNeighbors int
}

// Read reads a tile from the given readinp channel and returns it.
func Read(ch chan readinp.Line) (*Tile, error) {
	var lines []string
	for i := 0; i < 11; i++ {
		line, err := ReadLine(ch)
		if err != nil {
			return nil, err
		}
		lines = append(lines, line)
	}
	return ParseLines(lines)
}

func ParseLines(lines []string) (*Tile, error) {
	tid, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(lines[0], "Tile "), ":"))
	if err != nil {
		return nil, err
	}
	allVals := make([][]bool, 10)
	for i := 0; i < 10; i++ {
		if err != nil {
			log.Fatal(err)
		}
		allValLine := make([]bool, 10)
		for j, c := range lines[i+1] {
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
		pixels = append(pixels, allVals[i][1:9])
	}

	return &Tile{
		id:     tid,
		pixels: pixels,
		edges:  edges,
	}, nil
}

// ID returns the ID of the tile.
func (t Tile) ID() int {
	return t.id
}

// ReadEdge reads the binary value of the bits on the given edge.
// edges are: 0(top), 1(left), 2(bottom), 3(right)
func (t Tile) ReadEdge(e int) int {
	if len(t.edges) != 4 {
		log.Fatalf("Tile %d has too few edges", t.id)
	}
	// fmt.Println("reading edge", e, "with rotation", t.rotation, "as", (e-t.rotation+4)%4, ":", t.edges[(e-t.rotation+4)%4])
	return t.edges[(e-t.rotation+4)%4]
}

func (t Tile) Get(x, y int) bool {
	switch t.rotation % 4 {
	case 1:
		x, y = 7-y, x
	case 2:
		x, y = 7-x, 7-y
	case 3:
		x, y = y, 7-x
	}
	return t.pixels[y][x]
}

// Rotate rotates the tile counterclockwise n times.
func (t *Tile) Rotate(n int) {
	// fmt.Println("rotating", n, "to", (t.rotation+n)%4)
	t.rotation = (t.rotation + n) % 4
}

func (t Tile) EdgeMatches(e, val int) bool {
	if t.ReadEdge(e) == val {
		return true
	}
	if t.ReadEdge(e) == EdgeDual(val) {
		return true
	}
	return false
}

// SetNeighborFromEdgeMap sets the neighbors from a collected map of edge value to matching tiles.
func (t *Tile) SetNeighborFromEdgeMap(edgeMap map[int][]int) {
	neighborCount := 0
	t.neighbors = make([]int, 4)
	for i := 0; i < 4; i++ {
		edge := t.ReadEdge(i)
		edgeMatches := edgeMap[edge]
		if len(edgeMatches) == 1 {
			continue // no neighbor in that direction
		}
		for _, neighbor := range edgeMatches {
			if neighbor != t.id {
				neighborCount++
				t.neighbors[i] = neighbor
				break
			}
		}
	}
	t.numNeighbors = neighborCount
}

func (t Tile) NumNeighbors() int {
	return t.numNeighbors
}

func (t Tile) GetNeighbor(e int) int {
	if t.neighbors == nil {
		return -1
	}
	return t.neighbors[(e-t.rotation+4)%4]
}

func (t Tile) HasNeighbor(e int) bool {
	return t.GetNeighbor(e) != 0
}

func (t Tile) String() string {
	leftEdge := strconv.Itoa(t.ReadEdge(1))
	leftDual := fmt.Sprintf("(%d)", EdgeDual(t.ReadEdge(1)))
	spacer := " " + strconv.Itoa(t.rotation)
	for len(leftEdge) < len(leftDual) {
		leftEdge += " "
	}
	for range leftDual {
		spacer += " "
	}
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%d%s%d (%d) ^ %d\n", t.id, spacer, t.ReadEdge(0), EdgeDual(t.ReadEdge(0)), t.GetNeighbor(0)))
	for y := 0; y < 8; y++ {
		if y == 3 {
			b.WriteString(leftEdge + " ")
		} else if y == 4 {
			b.WriteString(leftDual + " ")
		} else {
			b.WriteString(spacer)
		}
		for x := 0; x < 8; x++ {
			if t.Get(x, y) {
				b.WriteString("X ")
			} else {
				b.WriteString(". ")
			}
		}
		if y == 3 {
			b.WriteString(" " + strconv.Itoa(t.ReadEdge(3)))
		} else if y == 4 {
			b.WriteString(fmt.Sprintf(" (%d)", EdgeDual(t.ReadEdge(3))))
		}
		b.WriteString("\n")
	}
	b.WriteString(fmt.Sprintf("<- %d  %d (%d) v %d -> %d\n", t.GetNeighbor(1), t.ReadEdge(2), EdgeDual(t.ReadEdge(2)), t.GetNeighbor(2), t.GetNeighbor(3)))
	return b.String()
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

func EdgeDual(a int) int {
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

func ReadLine(ch chan readinp.Line) (string, error) {
	line, ok := <-ch
	if !ok {
		return "", io.EOF
	}
	if line.Error != nil {
		return "", line.Error
	}
	return strings.TrimSpace(*line.Contents), nil
}
