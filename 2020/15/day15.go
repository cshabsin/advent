package main

import "fmt"

var input = []int{6, 13, 1, 15, 2, 0}

func main() {
	whenLast := map[int]int{}
	var prev int
	for i, v := range input {
		if i == len(input)-1 {
			prev = v
		} else {
			whenLast[v] = i
		}
	}
	fmt.Println(whenLast, prev)
	var next int
	for i := len(input); i < 30000000; /* 2020 */ i++ {
		if when, found := whenLast[prev]; found {
			next = i - when - 1
		} else {
			next = 0
		}
		whenLast[prev] = i - 1
		//fmt.Println(next, prev, whenLast[prev])
		prev = next
	}
	fmt.Println(next)
}
