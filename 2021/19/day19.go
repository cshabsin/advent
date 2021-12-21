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
	// // look at the rotations and multiplications:
	// v := matrix.Vector3{1, 2, 3}
	// vs := set.Set[matrix.Vector3]{}
	// for _, rot := range matrix.AllRotations() {
	// 	vm := v.Mul(rot)
	// 	fmt.Println(rot, vm)
	// 	vs.Add(vm)
	// }
	// fmt.Println(len(vs))

	// sc := readScanners("simple.txt")
	// for rot := range matrix.AllRotations() {
	// 	fmt.Println(sc[0].beaconPoints(rot))
	// 	fmt.Println("+++")
	// }

	// sf := newScannerFinder(readScanners("sample.txt"))
	// fmt.Println(sf.allScanners[0].beaconVecs(0, 0))
	// fmt.Println("---")
	// fmt.Println(sf.allScanners[0].beaconVecs(1, 0))
	// fmt.Println("---")
	// fmt.Println(sf.allScanners[0].beaconVecs(0, 1))
	// fmt.Println("---")
	// fmt.Println(sf.allScanners[0].beaconVecs(1, 1))
	// fmt.Println("---")
	// fmt.Println(overlap(sf.allScanners[0], sf.allScanners[1]))

	part1("sample.txt")
	part1("input.txt")
}

func part1(fn string) {
	sf := newScannerFinder(readScanners(fn))
	fmt.Println("scannerfinder initialized")
	for {
		var numMatches int
		for i := range sf.allScanners {
			if sf.find(i) {
				numMatches++
				if !sf.isFound[i] {
					sf.isFound.Add(i)
					fmt.Println("found", i)
				}
				// } else {
				// fmt.Println("find", i, "false")
			}
		}
		if numMatches == len(sf.allScanners) {
			break
		}
	}
}

type scannerFinder struct {
	// all beacon vectors precalculated for each scanner.
	// arranged by scanner, then rotation (0-23), then by origin beacon.
	// Note: allBeacons[*][*][x][x] == 0,0,0
	allScanners []*scanner

	foundScanners  []int              // each entry is an index into allScanners
	isFound        set.Set[int]       // each allScanners index that has been added to foundScanners
	foundOffsets   []matrix.Vector3   // each entry is the location of the given scanner relative to the "origin"
	foundRotations []matrix.Matrix3x3 // each entry is the rotation vector of the given scanner
}

func newScannerFinder(rawScanners []*scanner) *scannerFinder {
	rc := &scannerFinder{
		allScanners:    rawScanners,
		foundScanners:  []int{0},
		isFound:        set.Set[int]{0: true},
		foundOffsets:   []matrix.Vector3{{0, 0, 0}},
		foundRotations: []matrix.Matrix3x3{matrix.Ident()},
	}
	return rc
}

func overlap(cmp, tgt *scanner) bool {
	for cmpOrigin := range cmp.beacons {
		for cmpRot := range matrix.AllRotations() {
			cmpBeaconVecs := cmp.beaconVecs(cmpOrigin, cmpRot)
			for rot := range matrix.AllRotations() {
				for origin := range tgt.beacons {
					var matches int
					for _, tgtBeacon := range tgt.beaconVecs(origin, rot) {
						for _, cmpBeacon := range cmpBeaconVecs {
							if tgtBeacon.Eq(cmpBeacon) {
								matches++
							}
						}
					}
					if matches >= 12 {
						fmt.Println("matched", cmp.num, "to", tgt.num, "with rot", rot, "and origin", origin)
						return true
					}
				}
			}
		}
	}
	return false
}

func (s scannerFinder) find(tgtI int) bool {
	if s.isFound[tgtI] {
		return true
	}

	tgtScanner := s.allScanners[tgtI]
	for _, cmpI := range s.foundScanners {
		cmpScanner := s.allScanners[cmpI]
		if overlap(cmpScanner, tgtScanner) {
			return true
		}
	}
	return false
}

func readScanners(fn string) []*scanner {
	ch := readinp.MustReadConsumer[*scanner](fn, &parser{})
	var scanners []*scanner
	for l := range ch {
		scanners = append(scanners, l.MustGet())
	}
	return scanners
}

type scanner struct {
	num            int
	beacons        []matrix.Point3
	rotatedBeacons [][]matrix.Point3
	beaconVecCache [][][]matrix.Vector3
}

func (s *scanner) beaconPoints(rotation int) []matrix.Point3 {
	if s.rotatedBeacons == nil {
		s.rotatedBeacons = make([][]matrix.Point3, len(matrix.AllRotations()), len(matrix.AllRotations()))
	}
	if s.rotatedBeacons[rotation] == nil {
		rotMatrix := matrix.Rotation(rotation)
		var rotated []matrix.Point3
		for _, p := range s.beacons {
			rotated = append(rotated, p.Mul(rotMatrix))
		}
		s.rotatedBeacons[rotation] = rotated
	}
	return s.rotatedBeacons[rotation]
}

func (s *scanner) beaconVecs(origin int, rotation int) []matrix.Vector3 {
	if s.beaconVecCache == nil {
		s.beaconVecCache = make([][][]matrix.Vector3, len(s.beacons), len(s.beacons))
	}
	if s.beaconVecCache[origin] == nil {
		s.beaconVecCache[origin] = make([][]matrix.Vector3, len(matrix.AllRotations()), len(matrix.AllRotations()))
	}
	if s.beaconVecCache[origin][rotation] != nil {
		return s.beaconVecCache[origin][rotation]
	}
	rotatedPts := s.beaconPoints(rotation)
	var vecs []matrix.Vector3
	for _, target := range rotatedPts {
		vecs = append(vecs, target.Sub(rotatedPts[origin]))
	}
	s.beaconVecCache[origin][rotation] = vecs
	return vecs
}

type parser struct {
	current scanner
}

var headerRE = regexp.MustCompile(`--- scanner (\d*) ---`)

func (p *parser) Parse(line string) (*scanner, bool, error) {
	if line == "" {
		rc := p.current
		p.current = scanner{}
		return &rc, true, nil
	}
	if snum := headerRE.FindStringSubmatch(line); snum != nil {
		p.current.num = readinp.Atoi(snum[1])
		return nil, false, nil
	}
	strs := strings.Split(line, ",")
	p.current.beacons = append(p.current.beacons, matrix.Point3{
		readinp.Atoi(strs[0]),
		readinp.Atoi(strs[1]),
		readinp.Atoi(strs[2]),
	})
	return nil, false, nil
}

func (p *parser) Done() (**scanner, bool, error) {
	s := &p.current
	return &s, true, nil
}
