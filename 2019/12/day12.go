package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	var moons []*Moon
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		line = strings.TrimSuffix(line, "\n")
		moons = append(moons, parseMoon(line))
	}
	printMoons(moons)
	fmt.Println()
	for i := 0; i < 1000; i++ {
		for first := 0; first < len(moons); first++ {
			for second := first + 1; second < len(moons); second++ {
				if moons[first].X < moons[second].X {
					moons[first].VX++
					moons[second].VX--
				} else if moons[first].X > moons[second].X {
					moons[first].VX--
					moons[second].VX++
				}
				if moons[first].Y < moons[second].Y {
					moons[first].VY++
					moons[second].VY--
				} else if moons[first].Y > moons[second].Y {
					moons[first].VY--
					moons[second].VY++
				}
				if moons[first].Z < moons[second].Z {
					moons[first].VZ++
					moons[second].VZ--
				} else if moons[first].Z > moons[second].Z {
					moons[first].VZ--
					moons[second].VZ++
				}
			}
		}
		for _, moon := range moons {
			moon.Step()
		}
		fmt.Println("After", i+1, "steps:")
		printMoons(moons)
		fmt.Println()
	}
	printMoons(moons)
	fmt.Println("energy:", energy(moons))
}

func printMoons(moons []*Moon) {
	for _, moon := range moons {
		fmt.Println(moon)
	}
}

func energy(moons []*Moon) int {
	var e int
	for _, moon := range moons {
		e += moon.Potential() * moon.Kinetic()
	}
	return e
}

// Moon is data about a moon's position and velocity.
type Moon struct {
	X, Y, Z    int
	VX, VY, VZ int
}

func (m Moon) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>", m.X, m.Y, m.Z, m.VX, m.VY, m.VZ)
}

func (m *Moon) Step() {
	m.X += m.VX
	m.Y += m.VY
	m.Z += m.VZ
}

func (m Moon) Potential() int {
	var e int
	if m.X > 0 {
		e += m.X
	} else {
		e -= m.X
	}
	if m.Y > 0 {
		e += m.Y
	} else {
		e -= m.Y
	}
	if m.Z > 0 {
		e += m.Z
	} else {
		e -= m.Z
	}
	return e
}

func (m Moon) Kinetic() int {
	var e int
	if m.VX > 0 {
		e += m.VX
	} else {
		e -= m.VX
	}
	if m.VY > 0 {
		e += m.VY
	} else {
		e -= m.VY
	}
	if m.VZ > 0 {
		e += m.VZ
	} else {
		e -= m.VZ
	}
	return e
}

var lineRE = regexp.MustCompile("<x=(-?\\d*), y=(-?\\d*), z=(-?\\d*)>")

func parseMoon(line string) *Moon {
	submatches := lineRE.FindStringSubmatch(line)
	x, err := strconv.Atoi(submatches[1])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(submatches[2])
	if err != nil {
		log.Fatal(err)
	}
	z, err := strconv.Atoi(submatches[3])
	if err != nil {
		log.Fatal(err)
	}

	return &Moon{X: x, Y: y, Z: z}
}
