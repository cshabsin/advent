package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	// part1("sample1.txt")
	// part1("sample2.txt")
	// part1("sample3.txt")
	// part1("input.txt")
	// part2("sample1.txt")
	// part2("sample2.txt")
	// part2("sample3.txt")
	part2("input.txt")
}

func part2(fn string) {
	ch := readinp.MustRead(fn, parse)
	var cuboids []*cuboid
	for l := range ch {
		cuboids = append(cuboids, l.MustGet())
		fmt.Println(cuboids[len(cuboids)-1].volume())
	}
	for i, cu := range cuboids {
		for j := i + 1; i < len(cuboids); i++ {
			cu.subtract(cuboids[j])
		}
	}
}

func part1(fn string) {
	ch := readinp.MustRead(fn, parse)
	var cuboids []*cuboid
	for l := range ch {
		cuboids = append(cuboids, l.MustGet())
	}
	var cube [101][101][101]bool
	for _, oid := range cuboids {
		for x := oid.xmin + 50; x <= oid.xmax+50; x++ {
			if x < 0 || x > 100 {
				continue
			}
			for y := oid.ymin + 50; y <= oid.ymax+50; y++ {
				if y < 0 || y > 100 {
					continue
				}
				for z := oid.zmin + 50; z <= oid.zmax+50; z++ {
					if z < 0 || z > 100 {
						continue
					}
					cube[x][y][z] = oid.on
				}
			}
		}
	}
	var count int
	for x := range cube {
		for y := range cube[x] {
			for z := range cube[x][y] {
				if cube[x][y][z] {
					count++
				}
			}
		}
	}
	fmt.Println(fn, count)
}

type cuboid struct {
	order                              int
	on                                 bool
	xmin, xmax, ymin, ymax, zmin, zmax int
	exemptions                         []*cuboid
}

func (c cuboid) volume() int {
	return (c.xmax - c.xmin + 1) * (c.ymax - c.ymin + 1) * (c.zmax - c.zmin + 1)
}

func (c *cuboid) Overlaps(d *cuboid) bool {
	if c.xmax < d.xmin || c.xmin > d.xmax || c.ymax < d.ymin || c.xmin > d.ymax || c.zmax < d.zmin || c.zmin > d.zmax {
		return false
	}
	return true
}

func (c *cuboid) subtract(d *cuboid) {
	if !c.Overlaps(d) {
		return
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c cuboid) String() string {
	s := "off"
	if c.on {
		s = "on"
	}
	return fmt.Sprintf("%d: %s x=%d..%d,y=%d..%d,z=%d..%d", c.order, s, c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax)
}

var coordRE = regexp.MustCompile(`x=(-?\d*)\.\.(-?\d*),y=(-?\d*)\.\.(-?\d*),z=(-?\d*)\.\.(-?\d*)`)

func parse(line string) (*cuboid, error) {
	words := strings.Split(line, " ")
	var on bool
	if words[0] == "on" {
		on = true
	}
	matches := coordRE.FindStringSubmatch(line)
	return &cuboid{
		on:   on,
		xmin: readinp.Atoi(matches[1]),
		xmax: readinp.Atoi(matches[2]),
		ymin: readinp.Atoi(matches[3]),
		ymax: readinp.Atoi(matches[4]),
		zmin: readinp.Atoi(matches[5]),
		zmax: readinp.Atoi(matches[6]),
	}, nil
}
