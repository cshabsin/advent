package main

import (
	"fmt"
	"sort"

	"github.com/cshabsin/advent/commongen/readinp"
	"github.com/cshabsin/advent/commongen/set"
)

func part2stillwrong(fn string) {
	ch := readinp.MustRead(fn, parse)
	var cuboids []*cuboid
	for l := range ch {
		cu := l.MustGet()
		cuboids = append(cuboids, cu)
	}
	cl := clumpsFrom(cuboids)
	const clumpSize = 500
	cube := make([][][]bool, clumpSize, clumpSize)
	for x := range cube {
		cube[x] = make([][]bool, clumpSize, clumpSize)
		for y := range cube[x] {
			cube[x][y] = make([]bool, clumpSize, clumpSize)
		}
	}
	var total int
	for xmin := cl.xmin; xmin < cl.xmax+clumpSize*2; xmin += clumpSize {
		fmt.Println(xmin)
		for ymin := cl.ymin; ymin < cl.ymax+clumpSize*2; ymin += clumpSize {
			for zmin := cl.zmin; zmin < cl.zmax+clumpSize*2; zmin += clumpSize {
				total += run(cube, cuboids, xmin, xmin+clumpSize-1, ymin, ymin+clumpSize-1, zmin, zmin+clumpSize-1)
			}
		}
	}
	fmt.Println(total)
}

func part2alsobad(fn string) {
	ch := readinp.MustRead(fn, parse)
	var cuboids [][]*cuboid
	var i int
	for l := range ch {
		cu := l.MustGet()
		cu.order = i
		i++
		toAdd := set.Set[int]{}
		for i, cuArr := range cuboids {
			for _, oid := range cuArr {
				if oid.Overlaps(cu) {
					toAdd.Add(i)
					break
				}
			}
		}
		switch len(toAdd) {
		case 0:
			cuboids = append(cuboids, []*cuboid{cu})
		case 1:
			cuboids[*toAdd.AnyValue()] = append(cuboids[*toAdd.AnyValue()], cu)
		default:
			// merge the ones to add
			var newEntry []*cuboid
			for a := range toAdd {
				newEntry = append(newEntry, cuboids[a]...)
			}
			newCuboids := [][]*cuboid{newEntry}
			for i, cu := range cuboids {
				if toAdd.Contains(i) {
					continue
				}
				newCuboids = append(newCuboids, cu)
			}
			cuboids = newCuboids
		}
	}
	var clumps []*clump
	for _, cuList := range cuboids {
		clumps = append(clumps, clumpsFrom(cuList))
		fmt.Println(cuList, clumps[len(clumps)-1])
	}
	var count int
	fmt.Println(len(clumps))
	for _, clump := range clumps {
		fmt.Println(clump)
		count += clump.run()
		fmt.Println(count)
	}
	fmt.Println(count)
}

type clump struct {
	cuboids                            []*cuboid
	xmin, xmax, ymin, ymax, zmin, zmax int
}

func run(cube [][][]bool, cuboids []*cuboid, xmin, xmax, ymin, ymax, zmin, zmax int) int {
	var found bool
	for _, cu := range cuboids {
		if cu.xmax >= xmin && cu.xmin <= xmax && cu.ymax >= ymin && cu.xmin <= ymax && cu.zmax >= zmin && cu.zmin <= zmax {
			found = true
			break
		}
	}
	if !found {
		// fmt.Print(".")
		return 0
	}
	// fmt.Println(xmin, xmax, ymin, ymax, zmin, zmax)
	width := xmax - xmin + 1
	height := ymax - ymin + 1
	depth := zmax - zmin + 1
	for x := range cube {
		for y := range cube[x] {
			for z := range cube[x][y] {
				cube[x][y][z] = false
			}
		}
	}
	for _, oid := range cuboids {
		if oid.xmax < xmin || oid.xmin > xmax || oid.ymax < ymin || oid.xmin > ymax || oid.zmax < zmin || oid.zmin > zmax {
			continue
		}
		for x := oid.xmin - xmin; x <= oid.xmax-xmin; x++ {
			if x < 0 || x >= width {
				continue
			}
			for y := oid.ymin - ymin; y <= oid.ymax-ymin; y++ {
				if y < 0 || y >= height {
					continue
				}
				for z := oid.zmin - zmin; z <= oid.zmax-zmin; z++ {
					if z < 0 || z >= depth {
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

	return count

}

func (c clump) run() int {
	return run(nil, c.cuboids, c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax)
}

func clumpsFrom(cuboids []*cuboid) *clump {
	xmin := cuboids[0].xmin
	xmax := cuboids[0].xmax
	ymin := cuboids[0].ymin
	ymax := cuboids[0].ymax
	zmin := cuboids[0].zmin
	zmax := cuboids[0].zmax
	for _, cu := range cuboids {
		if cu.xmin < xmin {
			xmin = cu.xmin
		}
		if cu.xmax > xmax {
			xmax = cu.xmax
		}
		if cu.ymin < ymin {
			ymin = cu.ymin
		}
		if cu.ymax > ymax {
			ymax = cu.ymax
		}
		if cu.zmin < zmin {
			zmin = cu.zmin
		}
		if cu.ymax > zmax {
			zmax = cu.zmax
		}
	}
	return &clump{
		cuboids: cuboids,
		xmin:    xmin,
		xmax:    xmax,
		ymin:    ymin,
		ymax:    ymax,
		zmin:    zmin,
		zmax:    zmax,
	}
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

func part2obsolete(fn string) {
	ch := readinp.MustRead(fn, parse)
	var cuboidsRaw []*cuboid
	bld := newBoundariesBuilder()
	for l := range ch {
		c := l.MustGet()
		if c.xmin >= -50 && c.xmin <= 50 {
			continue // ignore the core, we can get its numbers from part1
		}
		cuboidsRaw = append(cuboidsRaw, c)
		bld.add(c)
	}
	bounds := bld.build()
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
