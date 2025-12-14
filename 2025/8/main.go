package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

type distanceValue struct {
	distance    int
	sourceIndex int
	targetIndex int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input.txt: %v", err)
	}
	defer f.Close()

	var points []point
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		z, _ := strconv.Atoi(fields[2])
		points = append(points, point{x, y, z})
	}
	topCount := 1000 // counting ten shortest connections in example, 1000 in real data
	sort.Slice(points, func(i, j int) bool {
		if points[i].x < points[j].x {
			return true
		}
		if points[i].x > points[j].x {
			return false
		}
		if points[i].y < points[j].y {
			return true
		}
		if points[i].y > points[j].y {
			return false
		}
		return points[i].z < points[j].z
	})
	var lengths []int       // top n lengths
	var dvs []distanceValue // length for each source to its nearest neighbor
	for i, p := range points {
		// do a triangle.
		for j := i + 1; j < len(points); j++ {
			dest := points[j]
			if i == j {
				continue
			}
			dist := (p.x - dest.x) * (p.x - dest.x)
			dist += (p.y - dest.y) * (p.y - dest.y)
			dist += (p.z - dest.z) * (p.z - dest.z)
			if len(lengths) < topCount || dist < lengths[topCount-1] {
				lengths = append(lengths, dist)
				if len(lengths) > topCount {
					sort.Ints(lengths)
					lengths = lengths[:topCount]
				}
				dvs = append(dvs, distanceValue{distance: dist, sourceIndex: i, targetIndex: j})
			}
		}
	}
	sort.Slice(dvs, func(i, j int) bool {
		return dvs[i].distance < dvs[j].distance
	})
	cs := newCircuitSet()
	for i := 0; i < topCount; i++ { // process top ten pairs
		cs.addDistanceValue(dvs[i])
	}
	fmt.Println(len(cs.circuits))
	sort.Slice(cs.circuits, func(i, j int) bool {
		return len(cs.circuits[i]) > len(cs.circuits[j])
	})
	fmt.Println(len(cs.circuits[0]) * len(cs.circuits[1]) * len(cs.circuits[2]))
}

type circuitSet struct {
	circuits  [][]int // list of indeces in each circuit
	inCircuit map[int]int
}

func newCircuitSet() *circuitSet {
	return &circuitSet{
		inCircuit: map[int]int{},
	}
}

func (c *circuitSet) addDistanceValue(pair distanceValue) {
	sourceCircuit, sourceInCircuit := c.inCircuit[pair.sourceIndex]
	targetCircuit, targetInCircuit := c.inCircuit[pair.targetIndex]
	if sourceInCircuit && targetInCircuit {
		if sourceCircuit != targetCircuit {
			fmt.Println("merging ", sourceCircuit, " and ", targetCircuit)
			c.circuits[sourceCircuit] = append(c.circuits[sourceCircuit], c.circuits[targetCircuit]...)
			for _, p := range c.circuits[targetCircuit] {
				c.inCircuit[p] = sourceCircuit
			}
			c.circuits = slices.Delete(c.circuits, targetCircuit, targetCircuit+1)
			for point, circuit := range c.inCircuit {
				if circuit > targetCircuit {
					c.inCircuit[point] = circuit - 1
				}
			}
		}
		return
	}
	if sourceInCircuit {
		c.circuits[sourceCircuit] = append(c.circuits[sourceCircuit], pair.targetIndex)
		c.inCircuit[pair.targetIndex] = sourceCircuit
	} else if targetInCircuit {
		c.circuits[targetCircuit] = append(c.circuits[targetCircuit], pair.sourceIndex)
		c.inCircuit[pair.sourceIndex] = targetCircuit
	} else {
		c.circuits = append(c.circuits, []int{pair.sourceIndex, pair.targetIndex})
		c.inCircuit[pair.sourceIndex] = len(c.circuits) - 1
		c.inCircuit[pair.targetIndex] = len(c.circuits) - 1
	}

}
