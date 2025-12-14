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

type Point struct {
	X, Y, Z int
}

type Edge struct {
	U, V int
	Dist int
}

func main() {
	points, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	edges := generateEdges(points)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Dist < edges[j].Dist
	})

	// Part 1: Connect the 1000 closest pairs.
	uf := NewUnionFind(len(points))
	limit := 1000
	if len(edges) < limit {
		limit = len(edges)
	}

	for i := 0; i < limit; i++ {
		uf.Union(edges[i].U, edges[i].V)
	}

	sizes := uf.ComponentSizes()
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	if len(sizes) >= 3 {
		fmt.Printf("Part 1: %d\n", sizes[0]*sizes[1]*sizes[2])
	}

	// Part 2: Continue connecting until one circuit remains.
	var lastEdge Edge
	for i := limit; i < len(edges); i++ {
		if uf.Count() == 1 {
			break
		}
		e := edges[i]
		if uf.Union(e.U, e.V) {
			lastEdge = e
		}
	}

	p1 := points[lastEdge.U]
	p2 := points[lastEdge.V]
	fmt.Printf("Part 2: %d\n", p1.X*p2.X)
}

func readInput(filename string) ([]Point, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var points []Point
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) < 3 {
			continue
		}
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		z, _ := strconv.Atoi(fields[2])
		points = append(points, Point{x, y, z})
	}
	return points, scanner.Err()
}

func generateEdges(points []Point) []Edge {
	var edges []Edge
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			dist := (p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y) + (p1.Z-p2.Z)*(p1.Z-p2.Z)
			edges = append(edges, Edge{U: i, V: j, Dist: dist})
		}
	}
	return edges
}

type UnionFind struct {
	parent []int
	size   []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size, count: n}
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	return uf.parent[i]
}

func (uf *UnionFind) Union(i, j int) bool {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)
	if rootI == rootJ {
		return false
	}
	if uf.size[rootI] < uf.size[rootJ] {
		rootI, rootJ = rootJ, rootI
	}
	uf.parent[rootJ] = rootI
	uf.size[rootI] += uf.size[rootJ]
	uf.count--
	return true
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func (uf *UnionFind) ComponentSizes() []int {
	var sizes []int
	for i, p := range uf.parent {
		if i == p {
			sizes = append(sizes, uf.size[i])
		}
	}
	return sizes
}
