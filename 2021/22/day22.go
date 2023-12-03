package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cshabsin/advent/commongen/cube"
	"github.com/cshabsin/advent/commongen/matrix"
	"github.com/cshabsin/advent/commongen/readinp"
)

// start over from scratch

func main() {
	part2("input.txt")
}

func part2(fn string) {
	ch := readinp.MustRead(fn, parse)
	var cubes []cube.Cube
	for in := range ch {
		cb := in.MustGet()
		var newCubes []cube.Cube
		for _, oldCube := range cubes {
			newCubes = append(newCubes, oldCube.Subtract(cb)...)
		}
		newCubes = append(newCubes, cb)
		cubes = newCubes
	}
	var vol int64
	for _, cb := range cubes {
		if cb.On {
			vol += cb.Volume()
		}
	}
	fmt.Println(vol)
}

var coordRE = regexp.MustCompile(`x=(-?\d*)\.\.(-?\d*),y=(-?\d*)\.\.(-?\d*),z=(-?\d*)\.\.(-?\d*)`)

func parse(line string) (cube.Cube, error) {
	words := strings.Split(line, " ")
	var on bool
	if words[0] == "on" {
		on = true
	}
	matches := coordRE.FindStringSubmatch(line)
	return cube.Cube{
		On:  on,
		Min: matrix.Point3{readinp.Atoi(matches[1]), readinp.Atoi(matches[3]), readinp.Atoi(matches[5])},
		Max: matrix.Point3{readinp.Atoi(matches[2]), readinp.Atoi(matches[4]), readinp.Atoi(matches[6])},
	}, nil
}
