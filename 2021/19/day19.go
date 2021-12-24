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
	part1("sample.txt")
	part1("input.txt")
}

func part1(fn string) {
	sf := newScannerFinder(readScanners(fn))
	fmt.Println("scannerfinder initialized")
	for len(sf.foundScanners) != len(sf.allScanners) {
		for i := range sf.allScanners {
			sf.find(i)
		}
	}
	beacons := map[matrix.Point3]bool{}
	for i := range sf.foundScanners {
		var newPts int
		for _, pt := range sf.allScanners[sf.foundScanners[i]].beaconPoints(sf.foundRotations[i]) {
			pt = pt.Offset(sf.foundOffsets[i])
			if !beacons[pt] {
				newPts++
				beacons[pt] = true
			}
		}
		fmt.Println(i, len(sf.allScanners[sf.foundScanners[i]].beacons), newPts)
	}
}

type scannerFinder struct {
	// all beacon vectors precalculated for each scanner.
	// arranged by scanner, then rotation (0-23), then by origin beacon.
	// Note: allBeacons[*][*][x][x] == 0,0,0
	allScanners []*scanner

	foundScanners  []int            // each entry is an index into allScanners
	isFound        set.Set[int]     // each allScanners index that has been added to foundScanners (for efficient checking)
	foundOffsets   []matrix.Vector3 // each entry is the location of the given scanner relative to the "origin"
	foundRotations []int            // each entry is the rotation index of the given scanner
}

func newScannerFinder(rawScanners []*scanner) *scannerFinder {
	rc := &scannerFinder{
		allScanners:    rawScanners,
		foundScanners:  []int{0},
		isFound:        set.Set[int]{0: true},
		foundOffsets:   []matrix.Vector3{{0, 0, 0}},
		foundRotations: []int{0},
	}
	return rc
}

func overlap(cmp, tgt *scanner, cmpRot int) (bool, matrix.Vector3, int) {
	for cmpOrigin := range cmp.beacons {
		cmpBeaconVecs := cmp.beaconVecs(cmpOrigin, cmpRot)
		for rot := range matrix.AllRotations() {
			for origin := range tgt.beacons {
				var matches int
				var offset matrix.Vector3
				for i, tgtBeacon := range tgt.beaconVecs(origin, rot) {
					for j, cmpBeacon := range cmpBeaconVecs {
						if tgtBeacon.Eq(cmpBeacon) {
							matches++
							offset = tgt.beaconPoints(rot)[i].Sub(cmp.beaconPoints(0)[j])
						}
					}
				}
				if matches >= 12 {
					fmt.Println("matched", tgt.num, "to", cmp.num, "with rot", rot, "and origin", origin, "offset", offset)
					return true, offset, rot
				}
			}
		}
	}
	return false, matrix.Vector3{}, 0
}

func (s *scannerFinder) find(tgtI int) bool {
	if s.isFound[tgtI] {
		return true
	}

	tgtScanner := s.allScanners[tgtI]
	for i, cmpI := range s.foundScanners {
		cmpScanner := s.allScanners[cmpI]
		found, offset, rot := overlap(cmpScanner, tgtScanner, s.foundRotations[i])
		if found {
			s.isFound.Add(tgtI)
			s.foundScanners = append(s.foundScanners, tgtI)
			s.foundOffsets = append(s.foundOffsets, offset)
			s.foundRotations = append(s.foundRotations, rot)
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
