package main

import (
	"fmt"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	part1("sample.txt")
	part1("input.txt")
}

func part1(fn string) {
	ch := readinp.MustRead(fn, parse)
	var p program
	for l := range ch {
		step := l.MustGet()
		p = append(p, step)
	}
	fmt.Println(p)
	a, b := p.run()
	fmt.Println(fn, ":", a, b)
}

type program []*step

func (p program) run() (int, int) {
	var a, b int
	var i int
	for {
		if i >= len(p) {
			break
		}
		step := p[i]
		fmt.Println(i, "executing", step)
		switch step.i {
		case hlf:
			if step.reg == "a" {
				a = a / 2
			} else {
				b = b / 2
			}
		case tpl:
			if step.reg == "a" {
				a *= 3
			} else {
				b *= 3
			}
		case inc:
			if step.reg == "a" {
				a++
			} else {
				b++
			}
		case jmp:
			i += step.offset
			continue
		case jie:
			var val int
			if step.reg == "a" {
				val = a
			} else {
				val = b
			}
			if val%2 == 0 {
				i += step.offset
				continue
			}
		case jio:
			var val int
			if step.reg == "a" {
				val = a
			} else {
				val = b
			}
			if val%2 == 1 {
				i += step.offset
				continue
			}
		}
		i++
	}
	return a, b
}

type instruction int

const (
	hlf instruction = iota
	tpl
	inc
	jmp
	jie
	jio
)

var instrMap = map[string]instruction{
	"hlf": hlf,
	"tpl": tpl,
	"inc": inc,
	"jmp": jmp,
	"jie": jie,
	"jio": jio,
}

func (i instruction) String() string {
	for sm, im := range instrMap {
		if im == i {
			return sm
		}
	}
	return fmt.Sprintf("unknown instruction %d", i)
}

type register int

type step struct {
	i      instruction
	offset int
	reg    string
}

func (r step) String() string {
	switch r.i {
	case hlf, tpl, inc:
		return fmt.Sprintf("%v %s", r.i, r.reg)
	case jie, jio:
		return fmt.Sprintf("%v %s, %+d", r.i, r.reg, r.offset)
	case jmp:
		return fmt.Sprintf("%v %+d", r.i, r.offset)
	}
	return fmt.Sprintf("unknown instruction %v", r.i)
}

func parse(line string) (*step, error) {
	i, found := instrMap[string(line[0:3])]
	if !found {
		return nil, fmt.Errorf("unrecognized instruction in %v", line)
	}
	rec := &step{i: i}
	switch i {
	case hlf, tpl, inc:
		// register only
		rec.reg = string(line[4])
	case jie, jio:
		rec.reg = string(line[4])
		rec.offset = readinp.Atoi(line[7:])
	case jmp:
		rec.offset = readinp.Atoi(line[4:])
	}
	return rec, nil
}
