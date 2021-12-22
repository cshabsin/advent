package main

import (
	"fmt"
	"regexp"
	"sort"
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
	var cuboidsRaw []*cuboid
	bld := newBoundariesBuilder()
	for l := range ch {
		c := l.MustGet()
		cuboidsRaw = append(cuboidsRaw, c)
		bld.add(c)
	}
	bounds := bld.build()
	fmt.Println(bounds)
	var cuboids []*cuboid
	for _, oid := range cuboidsRaw {
		cuboids = append(cuboids, oid.split(bounds)...)
	}
	cubMap := map[cuboid]bool{}
	for _, oid := range cuboids {
		on := oid.on
		oid.on = false // make them all match to overlap in the map
		cubMap[*oid] = on
	}
	var volume int
	for oid, on := range cubMap {
		if !on {
			continue
		}
		volume += (oid.xmax - oid.xmin + 1) * (oid.ymax - oid.ymin + 1) * (oid.zmax - oid.zmin + 1)
	}
}

// boundaries defines the various lines on which any cuboid must be split.
// sorted from lowest to highest
type boundaries struct {
	x []int
	y []int
	z []int
}

// boudariesBuilder builds boundaries
type boundariesBuilder struct {
	x map[int]bool
	y map[int]bool
	z map[int]bool
}

func newBoundariesBuilder() *boundariesBuilder {
	return &boundariesBuilder{
		x: make(map[int]bool),
		y: make(map[int]bool),
		z: make(map[int]bool),
	}
}

func (bld *boundariesBuilder) add(c *cuboid) {
	bld.x[c.xmax] = true
	bld.x[c.xmin] = true
	bld.y[c.ymax] = true
	bld.y[c.ymin] = true
	bld.z[c.zmax] = true
	bld.z[c.zmin] = true
}

func (bld *boundariesBuilder) build() *boundaries {
	var x, y, z []int
	for xval := range bld.x {
		x = append(x, xval)
	}
	for yval := range bld.y {
		y = append(y, yval)
	}
	for zval := range bld.z {
		z = append(z, zval)
	}
	sort.Sort(sort.IntSlice(x))
	sort.Sort(sort.IntSlice(y))
	sort.Sort(sort.IntSlice(z))
	return &boundaries{x, y, z}
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
	on                                 bool
	xmin, xmax, ymin, ymax, zmin, zmax int
}

func (c *cuboid) Overlaps(d *cuboid) bool {
	if c.xmax < d.xmin || c.xmin > d.xmax || c.ymax < d.ymin || c.xmin > d.ymax || c.zmax < d.zmin || c.zmin > d.zmax {
		return false
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c *cuboid) split(b *boundaries) []*cuboid {
	var rc []*cuboid
	xmin := c.xmin
	for _, xbound := range b.x {
		if xbound < xmin {
			continue
		}
		xmax := min(c.xmax, xbound)
		ymin := c.ymin
		for _, ybound := range b.y {
			if ybound < ymin {
				continue
			}
			ymax := min(c.ymax, ybound)
			zmin := c.zmin
			for _, zbound := range b.z {
				if zbound < zmin {
					continue
				}
				zmax := min(c.zmax, zbound)
				newcub := &cuboid{
					on:   c.on,
					xmin: xmin,
					xmax: xmax,
					ymin: ymin,
					ymax: ymax,
					zmin: zmin,
					zmax: zmax,
				}
				fmt.Println(newcub)
				rc = append(rc, newcub)
				zmin = zmax
				if zmin == c.zmax {
					break
				}
			}
			ymin = ymax
			if ymin == c.ymax {
				break
			}
		}
		xmin = xmax
		if xmin == c.xmax {
			break
		}
	}
	return rc
}

func (c cuboid) String() string {
	s := "off"
	if c.on {
		s = "on"
	}
	return fmt.Sprintf("%s x=%d..%d,y=%d..%d,z=%d..%d", s, c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax)
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
