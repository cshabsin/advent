package main

import (
	"fmt"
	"math"
)

func main() {
	Day20b(33100000)
}

func Day20a(min int) {
	i := 15000
	for {
		if total(i) > min {
			fmt.Println(i, total(i))
			return
		}
		i++
		if i%1000 == 0 {
			fmt.Println("trying", i)
		}
	}
}

func Day20b(min int) {
	i := 15000
	for {
		if total2(i) > min {
			fmt.Println(i, total2(i))
			return
		}
		i++
		if i%1000 == 0 {
			fmt.Println("trying", i)
		}
	}
}

func total(v int) int {
	var tot int
	for i := 1; i <= int(math.Sqrt(float64(v))); i++ {
		if v%i == 0 {
			tot += i
			if v/i != i {
				tot += v / i
			}
		}
	}
	return tot * 10
}

func total2(v int) int {
	var tot int
	for i := 1; i <= int(math.Sqrt(float64(v))); i++ {
		if v%i == 0 {
			if v/i <= 50 {
				tot += i
			}
			if v/i != i && i <= 50 {
				tot += v / i
			}
		}
	}
	return tot * 11
}
