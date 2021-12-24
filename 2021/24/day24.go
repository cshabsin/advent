package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/2021/24/pgm"
	"github.com/cshabsin/advent/commongen/readinp"
)

func oldMmain() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var statements []statement
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		var needB bool
		var st statement
		switch fields[0] {
		case "inp":
			st = statement{in: inp}
		case "add":
			needB = true
			st = statement{in: add}
		case "mul":
			needB = true
			st = statement{in: mul}
		case "div":
			needB = true
			st = statement{in: div}
		case "mod":
			needB = true
			st = statement{in: mod}
		case "eql":
			needB = true
			st = statement{in: eql}
		}
		st.a = strings.TrimSpace(fields[1])
		if needB {
			i, err := strconv.Atoi(fields[2])
			if err != nil {
				st.bStr = strings.TrimSpace(fields[2])
			} else {
				st.bIsInt = true
				st.bVal = i
			}
		}
		statements = append(statements, st)
	}
	i := 91457289443995
	for {
		if run(statements, strconv.Itoa(i)) {
			fmt.Println(i)
			break
		}
		i-- // "13579246899999")
		if i%100000 == 0 {
			fmt.Println("...", i)
		}
	}
}

type registers struct {
	w, x, y, z int
}

func (r registers) get(reg string) int {
	if reg == "w" {
		return r.w
	} else if reg == "x" {
		return r.x
	} else if reg == "y" {
		return r.y
	} else if reg == "z" {
		return r.z
	}
	fmt.Printf("unknown register %q\n", reg)
	return 0
}

func (r *registers) set(reg string, val int) {
	if reg == "w" {
		r.w = val
	} else if reg == "x" {
		r.x = val
	} else if reg == "y" {
		r.y = val
	} else if reg == "z" {
		r.z = val
	} else {
		fmt.Printf("unknown register %q\n", reg)
	}
}

func getB(st statement, reg registers) int {
	if st.bIsInt {
		return st.bVal
	}
	return reg.get(st.bStr)
}

func run(program []statement, input string) bool {
	var reg registers
	for _, st := range program {
		switch st.in {
		case inp:
			reg.set(st.a, int(input[0]-'0'))
			input = input[1:]
		case add:
			reg.set(st.a, reg.get(st.a)+getB(st, reg))
		case mul:
			reg.set(st.a, reg.get(st.a)*getB(st, reg))
		case div:
			reg.set(st.a, reg.get(st.a)/getB(st, reg))
		case mod:
			reg.set(st.a, reg.get(st.a)%getB(st, reg))
		case eql:
			if reg.get(st.a) == getB(st, reg) {
				reg.set(st.a, 1)
			} else {
				reg.set(st.a, 0)
			}
		}
	}
	// fmt.Println(reg)
	return reg.z == 0
}

type statement struct {
	in     instruction
	a      string
	bIsInt bool
	bStr   string
	bVal   int
}

type instruction int

const (
	inp instruction = iota
	add
	mul
	div
	mod
	eql
)

// codegen
func pgm1(i int) bool {
	in := strconv.Itoa(i)
	inIdx := 0
	var w, x, y, z int
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	y = w + 4 // 5-13
	z = y     // 5-13
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x = 1
	y = 26
	z *= y     //(5-13)*26
	y = w + 16 //17-25
	z += y     //(5-13)*26 + (17-25)
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x = y + 11 // (28-36)
	x = 1
	y = 25*x + 1
	z *= y
	y *= 0
	y += w
	y += 14
	y *= x
	z += y
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x *= 0
	x += z
	x = x % 26
	z = z / 26
	x += -13
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 3
	y *= x
	z += y
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x *= 0
	x += z
	x = x % 26
	z = z / 1
	x += 14
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 11
	y *= x
	z += y
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x *= 0
	x += z
	x = x % 26
	z = z / 1
	x += 15
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 13
	y *= x
	z += y
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x *= 0
	x += z
	x = x % 26
	z = z / 26
	x += -7
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 11
	y *= x
	z += y
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x *= 0
	x += z
	x = x % 26
	z = z / 1
	x += 10
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 7
	y *= x
	z += y
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x *= 0
	x += z
	x = x % 26
	z = z / 26
	x += -12
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 12
	y *= x
	z += y
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x *= 0
	x += z
	x = x % 26
	z = z / 1
	x += 15
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 15
	y *= x
	z += y
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x *= 0
	x += z
	x = x % 26
	z = z / 26
	x += -16
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 13
	y *= x
	z += y
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x *= 0
	x += z
	x = x % 26
	z = z / 26
	x += -9
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 1
	y *= x
	z += y
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x *= 0
	x += z
	x = x % 26
	z = z / 26
	x += -8
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 15
	y *= x
	z += y
	w = readinp.Atoi(string(in[inIdx]))
	inIdx++
	x *= 0
	x += z
	x = x % 26
	z = z / 26
	x += -8
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 4
	y *= x
	z += y
	return z == 0
}

type step struct {
	xadd, yadd int
}

func calculate(s step, z int) map[int]int {
	out := map[int]int{}
	for w := 1; w <= 9; w++ {
		c := &pgm.Compy{
			W: w,
			Z: z,
		}
		c.XBlock(s.xadd)
		c.YBlock(s.yadd)
		out[c.Z] = w
	}
	return out
}

func main() {
	steps := []step{
		{15, 4}, {14, 16}, {11, 14}, {-13, 3},
		{14, 11}, {15, 13}, {-7, 11}, {10, 7},
		{-12, 12}, {15, 15}, {-16, 13}, {-9, 1},
		{-8, 15}, {-8, 4},
	}
	vals := map[int]int{0: 0}
	for _, st := range steps {
		afterVals := map[int]int{}
		var valKeys []int
		for z := range vals {
			valKeys = append(valKeys, z)
		}
		sort.Sort(sort.IntSlice(valKeys))
		for _, z := range valKeys {
			max := vals[z]
			next := calculate(st, z)
			for newZ, digit := range next {
				afterVals[newZ] = max*10 + digit
			}
		}
		vals = afterVals
		fmt.Println(len(vals))
	}
	// 61896178354892
	// 71673396132593
	// 91457289443995
	fmt.Println(vals[0])
	// for i := 99990898000000; i > 10000000000000; i-- {
	// 	if pgm1(i) {
	// 		fmt.Println(i)
	// 	}
	// 	if i%100000 == 0 {
	// 		fmt.Println("...", i)
	// 	}
	// }

}

// maybe not actually 5
func try5() {
	for i := 0; i < 10000; i++ {
		c := pgm.Compy{In: fmt.Sprintf("%04d", i)}
		c.XBlock(15)
		c.YBlock(4)
		c.XBlock(14)
		c.XBlock(16)
		fmt.Println(i, ":", c.Z)
	}
}
