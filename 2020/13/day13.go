package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	inp, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	line := <-inp
	if line.Error != nil {
		log.Fatal(err)
	}
	_, err = strconv.Atoi(strings.TrimSpace(*line.Contents))
	if err != nil {
		log.Fatal(err)
	}
	line = <-inp
	if line.Error != nil {
		log.Fatal(err)
	}
	buses, err := getBuses(strings.TrimSpace(*line.Contents))
	if err != nil {
		log.Fatal(err)
	}
	// day13a(minTime, buses)
	fmt.Println(buses)
	v := sieveBuses(buses, 991*449*41*37*29, 6820035189)
	fmt.Println(v)
}

func sieveBuses(buses map[int]int, overrideMult, overrideFirst int) int {
	var busesSorted []int
	max := 1
	for mod := range buses {
		busesSorted = append(busesSorted, mod)
		max *= mod
	}
	sort.Sort(sort.IntSlice(busesSorted))
	var ch chan int
	for i := len(busesSorted) - 1; i >= 0; i-- {
		mod := busesSorted[i]
		a := buses[mod]
		if ch == nil {
			fmt.Println("kicking off multiples", mod, a)
			if overrideFirst == 0 {
				overrideFirst = mod - a
			}
			if overrideMult == 0 {
				overrideMult = mod
			}
			ch = multiples(overrideMult, overrideFirst, max)
			continue
		}
		fmt.Println("kicking off sieve", mod, a)
		ch = sieve(ch, mod, a)
	}
	val := <-ch
	return val
}

func multiples(mod, a, max int) chan int {
	ch := make(chan int)
	go func() {
		val := a
		for {
			ch <- val
			val += mod
			if val > max {
				break
			}
		}
		close(ch)
	}()
	return ch
}

func sieve(in chan int, mod, a int) chan int {
	out := make(chan int)
	go func() {
		for val := range in {
			if (val+a)%mod == 0 {
				out <- val
			}
		}
		close(out)
	}()
	return out
}

func day13a(minTime int, buses map[int]int) {
	t := minTime
	id := 0
outer:
	for {
		for bus := range buses {
			if t%bus == 0 {
				id = bus
				break outer
			}
		}
		t++
	}
	fmt.Println("minTime", minTime)
	fmt.Println("t", t)
	fmt.Println("id", id)
	fmt.Println("wait*id", (t-minTime)*id)
}

func getBuses(line string) (map[int]int, error) {
	buses := map[int]int{}
	i := 0
	for _, s := range strings.Split(line, ",") {
		if s == "x" {
			i++
			continue
		}
		bus, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		buses[bus] = i
		i++
	}
	return buses, nil
}
