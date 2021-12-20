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
				sf.isFound.Add(i)
			} else {
				fmt.Println("find", i, "false")
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
	allScanners [][][][]matrix.Vector3

	foundScanners  []int              // each entry is an index into allScanners
	isFound        set.Set[int]       // each allScanners index that has been added to foundScanners
	foundOffsets   []matrix.Vector3   // each entry is the location of the given scanner relative to the "origin"
	foundRotations []matrix.Matrix3x3 // each entry is the rotation vector of the given scanner
}

func newScannerFinder(rawScanners []scanner) *scannerFinder {
	rc := &scannerFinder{
		foundScanners:  []int{0},
		isFound:        set.Set[int]{0: true},
		foundOffsets:   []matrix.Vector3{matrix.Vector3{0, 0, 0}},
		foundRotations: []matrix.Matrix3x3{matrix.Ident()},
	}
	var allScanners [][][][]matrix.Vector3
	for _, scanner := range rawScanners {
		var byScanner [][][]matrix.Vector3
		for _, rot := range matrix.AllRotations() {
			var rotBeacons []matrix.Point3
			for _, beacon := range scanner.beacons {
				rotBeacons = append(rotBeacons, beacon.Mul(rot))
			}
			var byOrigin [][]matrix.Vector3
			for _, origin := range rotBeacons {
				var beacons []matrix.Vector3
				for _, target := range rotBeacons {
					beacons = append(beacons, target.Sub(origin))
				}
				byOrigin = append(byOrigin, beacons)
			}
			byScanner = append(byScanner, byOrigin)
		}
		allScanners = append(allScanners, byScanner)
	}
	rc.allScanners = allScanners
	return rc
}

func (s scannerFinder) find(tgtScanner int) bool {
	if s.isFound[tgtScanner] {
		return true
	}

	for _, cmpScanner := range s.foundScanners {
		cmpScanners := s.allScanners[cmpScanner][0][0]
		for rot, byOrigin := range s.allScanners[tgtScanner] {
			for origin, tgtBeacons := range byOrigin {
				var matches int
				for _, tgtBeacon := range tgtBeacons {
					for _, cmpBeacon := range cmpScanners {
						if tgtBeacon.Eq(cmpBeacon) {
							matches++
						}
					}
				}
				if matches > 12 {
					fmt.Println("matched", cmpScanner, "to", tgtScanner, "with rot", rot, "and origin", origin)
					return true
				}
			}
		}
	}
	return false
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
	beacons []matrix.Point3
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
	p.current.beacons = append(p.current.beacons, matrix.Point3{
		readinp.Atoi(strs[0]),
		readinp.Atoi(strs[1]),
		readinp.Atoi(strs[2]),
	})
	return scanner{}, false, nil
}

func (p *parser) Done() (*scanner, bool, error) {
	return &p.current, true, nil
}
