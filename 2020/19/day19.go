package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rules := ruleMap{}
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		rule, done := makeRule(line.Value())
		if done {
			break
		}
		rules[rule.index] = rule
		if line.Value() != rule.String() {
			fmt.Println(line.Value(), "-->", rule)
		}
	}

	var inputs []string
	count := 0
	for line := range ch {
		v := line.Value()
		inputs = append(inputs, v)
		if rules.matchTerminal(0, v) {
			count++
		}
	}
	fmt.Println(count)

	fmt.Println(rules.matchTerminal(120, "b"))
	for m := range rules.match(120, "ba") {
		fmt.Println("ba", m)
	}
}

type rule struct {
	index   int
	literal string
	a1, a2  int
	b1, b2  int // -1 if n/a
}

func (r rule) String() string {
	if r.literal != "" {
		return fmt.Sprintf("%d: %q", r.index, r.literal)
	}
	s := fmt.Sprintf("%d: %d", r.index, r.a1)
	if r.a2 != -1 {
		s += fmt.Sprintf(" %d", r.a2)
	}
	if r.b1 != -1 {
		s += fmt.Sprintf(" | %d", r.b1)
		if r.b2 != -1 {
			s += fmt.Sprintf(" %d", r.b2)
		}
	}
	return s
}

func makeRule(line string) (*rule, bool) {
	if line == "" {
		return nil, true
	}
	index, rest := getIndex(line)
	if rest[0] == '"' {
		return &rule{
			index:   index,
			literal: rest[1:2],
		}, false
	}
	subRules := strings.Split(rest, "|")
	a1, a2 := getPair(subRules[0])
	b1 := -1
	b2 := -1
	if len(subRules) > 1 {
		b1, b2 = getPair(subRules[1])
	}
	if len(subRules) > 2 {
		log.Fatalf("more than 2 subrules: %q", line)
	}
	return &rule{index, "", a1, a2, b1, b2}, false
}

func getIndex(line string) (int, string) {
	fields := strings.Split(line, ":")
	index, err := strconv.Atoi(fields[0])
	if err != nil {
		log.Fatal(err)
	}
	return index, strings.TrimSpace(fields[1])
}

func getPair(subRules string) (int, int) {
	fields := strings.Split(strings.TrimSpace(subRules), " ")
	a1, err := strconv.Atoi(strings.TrimSpace(fields[0]))
	if err != nil {
		log.Fatal(err)
	}
	a2 := -1
	if len(fields) > 1 {
		a2, err = strconv.Atoi(strings.TrimSpace(fields[1]))
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(fields) > 2 {
		log.Fatalf("more than two fields: %q", subRules)
	}
	return a1, a2
}

type ruleMap map[int]*rule

func (r ruleMap) matchTerminal(rule int, toMatch string) bool {
	remainders := r.match(rule, toMatch)
	for remainder := range remainders {
		if remainder == "" {
			return true
		}
	}
	return false
}

func (r ruleMap) match(rule int, toMatch string) chan string {
	out := make(chan string)
	go func() {
		if r[rule].literal != "" {
			if toMatch[0] == r[rule].literal[0] { // assumes literal is only one character, which is true in input
				out <- toMatch[1:len(toMatch)]
			}
			close(out)
			return
		}
		// Do the a's
		for a1Remainder := range r.match(r[rule].a1, toMatch) {
			if r[rule].a2 != -1 {
				for a2Remainder := range r.match(r[rule].a2, a1Remainder) {
					out <- a2Remainder
				}
			}
		}
		if r[rule].b1 != -1 {
			for b1Remainder := range r.match(r[rule].b1, toMatch) {
				if r[rule].b2 != -1 {
					for b2Remainder := range r.match(r[rule].b2, b1Remainder) {
						out <- b2Remainder
					}
				}
			}
		}
		close(out)
	}()
	return out
}
