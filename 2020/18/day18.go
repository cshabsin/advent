package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	total := 0
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		val := calc2(line.Value())
		if err != nil {
			fmt.Printf("error parsing line %q: %v\n", line.Value(), err)
		}
		total += val
	}
	fmt.Println(total)
}

func calc(line string) (int, int, error) {
	fmt.Printf("calc(%q)\n", line)
	var off, val int
	operator := '+'
	for {
		var subval, l int
		if line[off] == '(' {
			var err error
			subval, l, err = calc(line[off+1 : len(line)])
			fmt.Printf("line %q: subval %d, len %d\n", line[off+1:len(line)], subval, l)
			if err != nil {
				return 0, 0, err
			}
			l += 2 // add one for the '(' and skip the ')'
		} else {
			var err error
			subval, err = strconv.Atoi(line[off : off+1])
			if err != nil {
				return 0, 0, fmt.Errorf("parsing digit from %q: %v", line, err)
			}
			l = 1
		}
		switch operator {
		case '+':
			fmt.Println("applying + to", val, subval)
			val += subval
		case '-':
			fmt.Println("applying - to", val, subval)
			val -= subval
		case '*':
			fmt.Println("applying * to", val, subval)
			val *= subval
		}
		off += l
		if off == len(line) {
			return val, off, nil
		}
		if line[off] == ')' {
			return val, off, nil
		}
		off++ // skip the space
		operator = rune(line[off])
		off += 2
	}
}

type term struct {
	parenthetical expr

	// used if parenthetical is nil
	literal int
}
type operator rune
type entry struct {
	operator
	term
}

// expression with an implicit "0" as starting value (first term should have operator '+')
type expr []entry

// parse returns an expression and a length (to either the end of the string or the close paren)
func parse(line string) (expr, int) {
	fmt.Printf("parse %q\n", line)
	off := 0
	var rc expr
	op := operator('+')
	for {
		if line[off] == '(' {
			subexp, l := parse(line[off+1 : len(line)])
			rc = append(rc, entry{op, term{parenthetical: subexp}})
			off += l + 2
		} else {
			subval, err := strconv.Atoi(line[off : off+1])
			if err != nil {
				log.Fatalf("parsing digit from %q: %v", line, err)
			}
			rc = append(rc, entry{op, term{literal: subval}})
			off++
		}
		if off == len(line) || line[off] == ')' {
			return rc, off
		}
		off++
		op = operator(line[off])
		off += 2
	}
}

func calc2(line string) int {
	expr, _ := parse(line)
	fmt.Printf("calc2(%q): %v\n", line, expr)
	return calcExpr(expr)
}

func calcExpr(e expr) int {
	// collapse subexpressions
	newExpr := collapseSubExp(e)
	fmt.Println("collapseSubexp", newExpr)
	newExpr = collapseAdds(newExpr)
	fmt.Println("collapseAdds", newExpr)
	newExpr = collapseMult(newExpr)
	fmt.Println("collapseMult", newExpr)
	return newExpr[0].term.literal
}

func collapseSubExp(e expr) expr {
	var newExpr expr
	for _, ent := range e {
		if ent.term.parenthetical != nil {
			newExpr = append(newExpr, entry{
				operator: ent.operator,
				term:     term{literal: calcExpr(ent.term.parenthetical)},
			})
		} else {
			newExpr = append(newExpr, ent)
		}
	}
	return newExpr
}

func collapseAdds(e expr) expr {
	var newExpr expr
	i := 0
	for i < len(e) {
		val := e[i].term.literal
		j := i + 1
		for j < len(e) {
			if e[j].operator == '*' {
				break
			}
			if e[j].operator == '+' {
				val += e[j].term.literal
			}
			if e[j].operator == '-' {
				val -= e[j].term.literal
			}
			j++
		}
		newExpr = append(newExpr, entry{
			operator: e[i].operator,
			term:     term{literal: val},
		})
		i = j
	}
	return newExpr
}

func collapseMult(e expr) expr {
	var newExpr expr
	i := 0
	for i < len(e) {
		val := e[i].term.literal
		j := i + 1
		for j < len(e) {
			if e[j].operator == '*' {
				val *= e[j].term.literal
			}
			if e[j].operator == '+' {
				log.Fatalf("unexpected operator + in %v", e)
			}
			if e[j].operator == '-' {
				log.Fatalf("unexpected operator - in %v", e)
			}
			j++
		}
		newExpr = append(newExpr, entry{
			operator: e[i].operator,
			term:     term{literal: val},
		})
		i = j
	}
	return newExpr
}
