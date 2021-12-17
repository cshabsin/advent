package main

import "fmt"

//target area: x=269..292, y=-68..-44

// x: sum(0..n) = n*(n+1)/2
// 275 -> 550 = n^2+n
// n^2+n-550 = 0

func main() {
	dx := 23
	for dy := 0; dy < 2000; dy++ {
		step(dx, dy, 269, 292, -68, -44)
	}
}

func step(dx, dy, tgtXmin, tgtXmax, tgtYmin, tgtYmax int) {
	initdx, initdy := dx, dy
	var x, y, maxY int
	for {
		if x >= tgtXmin && x <= tgtXmax && y >= tgtYmin && y <= tgtYmax {
			fmt.Println("matched", initdx, initdy, "(peak", maxY, ")")
			return
		}
		if x > tgtXmax || y < tgtYmin {
			return
		}
		x += dx
		if dx > 0 {
			dx--
		} else if dx < 0 {
			dx++
		}
		y += dy
		dy--
		if y > maxY {
			maxY = y
		}
	}
}
