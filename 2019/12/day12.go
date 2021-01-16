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
	var moons, firstMoons []*Moon
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
		firstMoons = append(firstMoons, parseMoon(line))
	}
	printMoons(moons)
	fmt.Println()
	var i, xEquiv, yEquiv, zEquiv int64
	for {
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
		i++
		if equivalentX(moons, firstMoons) {
			if xEquiv == 0 {
				fmt.Println("xEquiv", i)
				xEquiv = i
			}
		}
		if equivalentY(moons, firstMoons) {
			if yEquiv == 0 {
				fmt.Println("yEquiv", i)
				yEquiv = i
			}
		}
		if equivalentZ(moons, firstMoons) {
			if zEquiv == 0 {
				fmt.Println("zEquiv", i)
				zEquiv = i
			}
		}
		if xEquiv != 0 && yEquiv != 0 && zEquiv != 0 {
			break
		}
	}
	var j int64
	for j = -2; j < 4; j++ {
		fmt.Println(j, "gcd", gcd(gcd(xEquiv+j, yEquiv+j), zEquiv+j))
		fmt.Println(j, "lcm", (xEquiv+j)*(yEquiv+j)*(zEquiv+j)/gcd(gcd(xEquiv+j, yEquiv+j), zEquiv+j))
	}
	// cshabsin@DESKTOP-LC7C1G1:~/go/src/github.com/cshabsin/advent/2019/12$ factor 108344
	// 108344: 2 2 2 29 467
	// cshabsin@DESKTOP-LC7C1G1:~/go/src/github.com/cshabsin/advent/2019/12$ factor 113028
	// 113028: 2 2 3 9419
	// cshabsin@DESKTOP-LC7C1G1:~/go/src/github.com/cshabsin/advent/2019/12$ factor 231614
	// 231614: 2 115807
	// cshabsin@DESKTOP-LC7C1G1:~/go/src/github.com/cshabsin/advent/2019/12$ dc
	// 2 2 2 3 29 467 9419 115807 *********p
	// dc: stack empty
	// dc: stack empty
	// 354540398381256
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

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func equivalentX(a []*Moon, b []*Moon) bool {
	for i := 0; i < len(a); i++ {
		if a[i].X != b[i].X {
			return false
		}
		if a[i].VX != b[i].VX {
			return false
		}
	}
	return true
}

func equivalentY(a []*Moon, b []*Moon) bool {
	for i := 0; i < len(a); i++ {
		if a[i].Y != b[i].Y {
			return false
		}
		if a[i].VY != b[i].VY {
			return false
		}
	}
	return true
}

func equivalentZ(a []*Moon, b []*Moon) bool {
	for i := 0; i < len(a); i++ {
		if a[i].Z != b[i].Z {
			return false
		}
		if a[i].VZ != b[i].VZ {
			return false
		}
	}
	return true
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
