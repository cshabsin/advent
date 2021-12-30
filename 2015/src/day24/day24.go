package main

import (
	"fmt"

	"github.com/cshabsin/advent/commongen/set"
)

func main() {
	var tot int
	for _, i := range input {
		tot += i
	}
	fmt.Println("part 1:", compute(set.Make(input...), 5, tot/3))
	fmt.Println("part 2:", compute(set.Make(input...), 3, tot/4))
}

func compute(in set.Set[int], levels, remaining int) int64 {
	s := search(in, levels, remaining)
	var minQE int64
	for _, weights := range s {
		var qe int64 = 1
		for _, wt := range weights {
			qe *= int64(wt)
		}
		if minQE == 0 || qe < minQE {
			minQE = qe
		}
	}
	return minQE
}

func search(in set.Set[int], levels, remaining int) [][]int {
	if levels == 0 {
		if in.Contains(remaining) {
			return [][]int{{remaining}}
		}
		return nil
	}
	var results [][]int
	for i := range in {
		if i > remaining {
			continue
		}
		res := search(in.Without(i), levels-1, remaining-i)
		if res == nil {
			continue
		}
		for _, subres := range res {
			results = append(results, append(subres, i))
		}
	}
	return results
}

var input = []int{
	1, 2, 3, 7, 11, 13, 17, 19, 23, 31, 37, 41, 43, 47, 53,
	59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}
