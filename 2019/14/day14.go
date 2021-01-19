package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	formulas := map[string]*formula{}
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		formula := parseFormula(line)
		if formulas[formula.output.name] != nil {
			log.Fatal("duplicate formula for output", formula.output)
		}
		formulas[formula.output.name] = formula
	}
	fmt.Println(*formulas["FUEL"])
}

type entry struct {
	name  string
	count int
}

type formula struct {
	output entry
	inputs []entry
}

func parseFormula(line string) *formula {
	fp := formParser{line: line}
	return fp.Parse()
}

type formParser struct {
	line string
	i    int
}

func (fp *formParser) readChar() rune {
	if fp.i >= len(fp.line) {
		return ' '
	}
	r := rune(fp.line[fp.i])
	fp.i++
	return r
}

func (fp *formParser) readInt() int {
	var numStr string
	for {
		r := fp.readChar()
		if !unicode.IsDigit(r) {
			break
		}
		numStr += string(r)
	}
	val, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err, fp.line, fp.i)
	}
	return val
}

func (fp *formParser) readName() (string, rune) {
	var r rune
	var name string
	for {
		r = fp.readChar()
		if !unicode.IsLetter(r) {
			break
		}
		name += string(r)
	}
	return name, r
}

func (fp *formParser) readEntry() (entry, rune) {
	count := fp.readInt()
	name, next := fp.readName()
	return entry{name: name, count: count}, next
}

func (fp *formParser) Parse() *formula {
	var inputs []entry
	for {
		entry, next := fp.readEntry()
		inputs = append(inputs, entry)
		if next != ',' {
			break
		}
		fp.readChar() // ' '
	}
	fp.readChar() // '='
	fp.readChar() // '>'
	fp.readChar() // ' '
	output, _ := fp.readEntry()
	return &formula{
		inputs: inputs,
		output: output,
	}
}
