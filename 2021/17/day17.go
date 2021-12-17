package main

import "fmt"

//target area: x=269..292, y=-68..-44

// x: sum(0..n) = n*(n+1)/2
// 275 -> 550 = n^2+n
// n^2+n-550 = 0

func main() {
	var cnt int
	for dx := 5; dx < 2000; dx++ {
		for dy := -500; dy < 2000; dy++ {
			if step(dx, dy, 269, 292, -68, -44) {
				// if step(dx, dy, 20, 30, -10, -5) {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

func step(dx, dy, tgtXmin, tgtXmax, tgtYmin, tgtYmax int) bool {
	initdx, initdy := dx, dy
	var x, y, maxY int
	for {
		if x >= tgtXmin && x <= tgtXmax && y >= tgtYmin && y <= tgtYmax {
			fmt.Println("matched", initdx, initdy, "(peak", maxY, ")")
			return true
		}
		if x > tgtXmax || y < tgtYmin {
			return false
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
