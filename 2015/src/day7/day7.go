package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Formula struct {
	Target string
	Type string  // NOT, OR, AND, RSHIFT, LSHIFT, ->
	Operand1 string
	Operand2 string
	HasValue bool
	Value uint
}

func (f Formula) SetValue(v uint) Formula {
	f.HasValue = true
	f.Value = v
	return f
}

type Forms struct {
	F map[string]Formula
	stack []string
}

func NewForms() Forms {
	var forms Forms
	forms.F = make(map[string]Formula)
	return forms
}

func (forms Forms) AddFormula(line string) {
	tokens := strings.Split(strings.Trim(line, "\n"), " ")
	target := tokens[len(tokens)-1]
	var f Formula
	if tokens[1] == "->" {
		// straight assignment
		f = Formula{target, tokens[1], tokens[0], "", false, 0}
	} else if tokens[0] == "NOT" {
		f = Formula{target, "NOT", tokens[1], "", false, 0}
	} else {
		f = Formula{target, tokens[1], tokens[0], tokens[2], false, 0}
	}
	forms.F[target] = f
}

func (forms Forms) Evaluate(sym string) uint {
	lit, err := strconv.ParseInt(sym, 10, 64)
	if err == nil {
		return uint(lit) & 0xffff
	}

	f := forms.F[sym]
	if f.HasValue {
		return f.Value
	}

	forms.stack = append(forms.stack, sym)
	r := uint(forms.evalinner(f)) & 0xffff
	forms.F[sym] = f.SetValue(r)
	if (forms.stack[len(forms.stack)-1] != sym) {
		log.Fatal("Returned from stack for sym ", sym,
			" but at the top found ",
			forms.stack[len(forms.stack)-1])
	}
	fmt.Printf("%s\n", forms.stack)
	forms.stack = forms.stack[:len(forms.stack)-1]
	return r
}

func (forms Forms) evalinner(f Formula) uint {
	if f.Type == "->" {
		return forms.Evaluate(f.Operand1)
	}
	if f.Type == "NOT" {
		return ^forms.Evaluate(f.Operand1)
	}
	if f.Type == "AND" {
		return forms.Evaluate(f.Operand1) &
			forms.Evaluate(f.Operand2)
	}
	if f.Type == "OR" {
		return forms.Evaluate(f.Operand1) |
			forms.Evaluate(f.Operand2)
	}
	if f.Type == "LSHIFT" {
		return forms.Evaluate(f.Operand1) <<
			forms.Evaluate(f.Operand2)
	}
	if f.Type == "RSHIFT" {
		return forms.Evaluate(f.Operand1) >>
			forms.Evaluate(f.Operand2)
	}
	log.Fatal("Unknown type " + f.Type)
	return 0
}

func main() {
	var infile = flag.String("infile", "input7.txt", "Input filename")
	var field = flag.String("field", "a", "Field to evaluate")
	flag.Parse()
	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	forms := NewForms()
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		forms.AddFormula(line)
	}
	fmt.Printf("%s: ", *field)
	fmt.Printf("%v\n", forms.Evaluate(*field))
}
