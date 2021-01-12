package main

import "fmt"

func main() {
	// from input.txt
	v1 := 14082811
	v2 := 5249543

	num := 1
	// subject := 7
	subject := v2
	for loops := 0; loops < 17188728; loops++ {
		num *= subject
		num = num % 20201227
		if num == v1 {
			fmt.Println(v1, ":", loops)
		}
		if num == v2 {
			fmt.Println(v2, ":", loops)
		}
		if num == 5764801 {
			fmt.Println(5764801, ":", loops)
		}
	}
	fmt.Println(num)
}

// 7271248 too high
