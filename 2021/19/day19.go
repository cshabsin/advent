package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cshabsin/advent/commongen/matrix"
	"github.com/cshabsin/advent/commongen/readinp"
	"github.com/cshabsin/advent/commongen/set"
)

func main() {
	fmt.Println(len(allRotations))
	fmt.Println(allRotations)
	//	part1("input.txt")
}

func part1(fn string) {
	fmt.Println(readScanners(fn))
}

func readScanners(fn string) []scanner {
	ch := readinp.MustReadConsumer[scanner](fn, &parser{})
	var scanners []scanner
	for l := range ch {
		scanners = append(scanners, l.MustGet())
	}
	return scanners
}

type scanner struct {
	num     int
	beacons [][3]int
}
type parser struct {
	current scanner
}

var headerRE = regexp.MustCompile(`--- scanner (\d*) ---`)

func (p *parser) Parse(line string) (scanner, bool, error) {
	if line == "" {
		rc := p.current
		p.current = scanner{}
		return rc, true, nil
	}
	if snum := headerRE.FindStringSubmatch(line); snum != nil {
		p.current.num = readinp.Atoi(snum[1])
		return scanner{}, false, nil
	}
	strs := strings.Split(line, ",")
	p.current.beacons = append(p.current.beacons, [3]int{
		readinp.Atoi(strs[0]),
		readinp.Atoi(strs[1]),
		readinp.Atoi(strs[2]),
	})
	return scanner{}, false, nil
}

func (p *parser) Done() (*scanner, bool, error) {
	return &p.current, true, nil
}

var (
	allRotations = calcAllRotations()
)

func calcAllRotations() []matrix.Matrix {
	zRot := matrix.Matrix{ // rotate around z axis
		{0, -1, 0},
		{1, 0, 0},
		{0, 0, 1},
	}
	yRot := matrix.Matrix{ //rotate around y axis
		{0, 0, -1},
		{0, 1, 0},
		{1, 0, 0},
	}
	xRot := matrix.Matrix{ // rotate around x axis
		{1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	}
	rotSet := set.Set[matrix.Matrix]{}
	rotSet.Add(matrix.Ident())
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			for z := 0; z < 4; z++ {
				val := xRot.Pow(x).Mul(yRot.Pow(y)).Mul(zRot.Pow(z))
				rotSet.Add(val)
			}
		}
	}
	return rotSet.AsSlice()
}
