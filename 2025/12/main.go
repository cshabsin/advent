package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Point represents a coordinate in the grid (row, col).
type Point struct {
	r, c int
}

// Shape is a collection of points representing a present.
type Shape []Point

func (s Shape) Width() int {
	maxC := 0
	for _, p := range s {
		if p.c > maxC {
			maxC = p.c
		}
	}
	return maxC + 1
}

func (s Shape) Height() int {
	maxR := 0
	for _, p := range s {
		if p.r > maxR {
			maxR = p.r
		}
	}
	return maxR + 1
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input.txt: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}

	shapes, regionsStart := parseShapes(lines)

	// Precompute orientations for all shapes.
	orientations := make([][]Shape, len(shapes))
	for i, s := range shapes {
		orientations[i] = generateOrientations(s)
	}

	count := 0
	for i := regionsStart; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		if solveRegion(line, orientations) {
			count++
		}
	}
	fmt.Println(count)
}

// parseShapes parses the shape definitions from the input lines.
// It returns the list of shapes and the line index where region definitions start.
func parseShapes(lines []string) ([]Shape, int) {
	var shapes []Shape
	i := 0
	for i < len(lines) {
		line := lines[i]
		// Check if we've reached the region definitions.
		// Region lines look like "12x5: ...", shape headers look like "0:".
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			if strings.Contains(parts[0], "x") {
				return shapes, i
			}
		}

		if strings.HasSuffix(line, ":") {
			// Start of a shape definition.
			i++
			var points []Point
			r := 0
			for i < len(lines) {
				l := lines[i]
				if l == "" {
					i++
					break
				}
				for c, char := range l {
					if char == '#' {
						points = append(points, Point{r, c})
					}
				}
				r++
				i++
			}
			shapes = append(shapes, normalize(points))
		} else {
			i++
		}
	}
	return shapes, i
}

// normalize shifts the shape so that its top-left bounding box corner is at (0,0).
func normalize(points []Point) Shape {
	if len(points) == 0 {
		return nil
	}
	minR, minC := points[0].r, points[0].c
	for _, p := range points {
		if p.r < minR {
			minR = p.r
		}
		if p.c < minC {
			minC = p.c
		}
	}
	var res Shape
	for _, p := range points {
		res = append(res, Point{p.r - minR, p.c - minC})
	}
	// Sort points to ensure canonical representation for deduplication.
	sort.Slice(res, func(i, j int) bool {
		if res[i].r == res[j].r {
			return res[i].c < res[j].c
		}
		return res[i].r < res[j].r
	})
	return res
}

// generateOrientations generates all unique rotations and flips of a shape.
func generateOrientations(s Shape) []Shape {
	unique := make(map[string]Shape)

	current := s
	for i := 0; i < 4; i++ {
		// Add current rotation
		norm := normalize(current)
		unique[shapeKey(norm)] = norm

		// Add flipped version of current rotation
		flipped := make(Shape, len(current))
		for j, p := range current {
			flipped[j] = Point{p.r, -p.c}
		}
		normFlipped := normalize(flipped)
		unique[shapeKey(normFlipped)] = normFlipped

		// Rotate 90 deg clockwise for next iteration: (r, c) -> (c, -r)
		next := make(Shape, len(current))
		for j, p := range current {
			next[j] = Point{p.c, -p.r}
		}
		current = next
	}

	var res []Shape
	for _, v := range unique {
		res = append(res, v)
	}
	// Sort orientations deterministically.
	sort.Slice(res, func(i, j int) bool {
		return shapeKey(res[i]) < shapeKey(res[j])
	})
	return res
}

func shapeKey(s Shape) string {
	var sb strings.Builder
	for _, p := range s {
		fmt.Fprintf(&sb, "%d,%d|", p.r, p.c)
	}
	return sb.String()
}

func solveRegion(line string, allOrientations [][]Shape) bool {
	parts := strings.Split(line, ":")
	dims := strings.Split(parts[0], "x")
	W, _ := strconv.Atoi(dims[0])
	H, _ := strconv.Atoi(dims[1])

	countsStr := strings.Fields(parts[1])
	var pieces []int // indices of shapes
	totalArea := 0
	for id, s := range countsStr {
		count, _ := strconv.Atoi(s)
		shapeArea := len(allOrientations[id][0])
		for k := 0; k < count; k++ {
			pieces = append(pieces, id)
			totalArea += shapeArea
		}
	}

	if totalArea > W*H {
		return false
	}

	// Sort pieces by size (area) descending to fail fast.
	sort.Slice(pieces, func(i, j int) bool {
		return len(allOrientations[pieces[i]][0]) > len(allOrientations[pieces[j]][0])
	})

	grid := make([][]bool, H)
	for r := 0; r < H; r++ {
		grid[r] = make([]bool, W)
	}

	return backtrack(grid, pieces, 0, W, H, allOrientations)
}

func backtrack(grid [][]bool, pieces []int, idx int, W, H int, allOrientations [][]Shape) bool {
	if idx == len(pieces) {
		return true
	}

	shapeIdx := pieces[idx]
	possibleShapes := allOrientations[shapeIdx]

	// Try to place the piece
	for _, s := range possibleShapes {
		sH := s.Height()
		sW := s.Width()

		// Optimization: Don't iterate if the shape is larger than the grid.
		if sH > H || sW > W {
			continue
		}

		for r := 0; r <= H-sH; r++ {
			for c := 0; c <= W-sW; c++ {
				if canPlace(grid, s, r, c) {
					place(grid, s, r, c, true)
					if backtrack(grid, pieces, idx+1, W, H, allOrientations) {
						return true
					}
					place(grid, s, r, c, false)
				}
			}
		}
	}

	return false
}

func canPlace(grid [][]bool, s Shape, r, c int) bool {
	for _, p := range s {
		if grid[r+p.r][c+p.c] {
			return false
		}
	}
	return true
}

func place(grid [][]bool, s Shape, r, c int, val bool) {
	for _, p := range s {
		grid[r+p.r][c+p.c] = val
	}
}
