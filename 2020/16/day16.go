package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

var ruleRegexp = regexp.MustCompile(`^([a-z ]*): (\d*)-(\d*) or (\d*)-(\d*)$`)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var rules ruleSet
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		c := strings.TrimSpace(*line.Contents)
		if c == "" {
			break
		}
		rules = append(rules, makeRule(c))
	}
	<-ch // "your ticket:"
	line := <-ch
	if line.Error != nil {
		log.Fatal(line.Error)
	}
	myTicket := makeTicket(strings.TrimSpace(*line.Contents))
	rules.filterTicket(myTicket)
	<-ch // empty line
	<-ch // "nearby tickets:"
	total := 0
	for line = range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		total += rules.filterTicket(makeTicket(strings.TrimSpace(*line.Contents)))
	}
	fmt.Println(total)

	ans := 1
	for _, f := range []int{12, 6, 7, 18, 13, 14} {
		ans *= myTicket.fields[f]
	}
	fmt.Println("ans:", ans)

	sort.Slice(rules, func(i, j int) bool { return len(rules[i].disqualifiedFields) > len(rules[j].disqualifiedFields) })
	for {
		fmt.Println("---")
		for _, rule := range rules {
			fmt.Println(len(rule.disqualifiedFields), rule.name, rule.selectedField, rule.disqualifiedFields)
		}
		needToSearch := false
		for _, rule := range rules {
			if rule.selectedField != -1 {
				continue
			}
			needToSearch = true
			if len(rule.disqualifiedFields) == len(rules)-1 {
				qual := rule.onlyQualifiedField()
				if qual >= 0 {
					rule.selectedField = qual
					for _, rule := range rules {
						rule.disqualifiedFields[qual] = true
					}
				}
				break
			}
		}
		if !needToSearch {
			fmt.Println("foo")
			break
		}
	}
}

type rule struct {
	name                       string
	r1min, r1max, r2min, r2max int

	disqualifiedFields map[int]bool
	selectedField      int
}

func makeRule(c string) *rule {
	fields := ruleRegexp.FindStringSubmatch(c)
	r1min, err := strconv.Atoi(fields[2])
	if err != nil {
		log.Fatal(err)
	}
	r1max, err := strconv.Atoi(fields[3])
	if err != nil {
		log.Fatal(err)
	}
	r2min, err := strconv.Atoi(fields[4])
	if err != nil {
		log.Fatal(err)
	}
	r2max, err := strconv.Atoi(fields[5])
	if err != nil {
		log.Fatal(err)
	}
	return &rule{
		name:               fields[1],
		r1min:              r1min,
		r1max:              r1max,
		r2min:              r2min,
		r2max:              r2max,
		disqualifiedFields: map[int]bool{},
		selectedField:      -1,
	}
}

func (r rule) onlyQualifiedField() int {
	if len(r.disqualifiedFields) != 19 {
		return -1
	}
	for i := 0; i < 20; i++ {
		if !r.disqualifiedFields[i] {
			return i
		}
	}
	return -2
}

type ticket struct {
	fields []int
}

func makeTicket(c string) ticket {
	fieldStrings := strings.Split(c, ",")
	var fields []int
	for _, fieldStr := range fieldStrings {
		field, err := strconv.Atoi(fieldStr)
		if err != nil {
			log.Fatal(err)
		}
		fields = append(fields, field)
	}
	return ticket{fields}
}

type ruleSet []*rule

func (rs *ruleSet) filterTicket(t ticket) int {
	badVals := 0
	for _, val := range t.fields {
		anyGood := false
		for _, rule := range *rs {
			if rule.r1min <= val && rule.r1max >= val {
				anyGood = true
				continue
			}
			if rule.r2min <= val && rule.r2max >= val {
				anyGood = true
				continue
			}
			//			rule.disqualifiedFields[index] = true
		}
		if !anyGood {
			badVals += val
		}
	}
	if badVals == 0 {
		for index, val := range t.fields {
			for _, rule := range *rs {
				if rule.r1min <= val && rule.r1max >= val {
					continue
				}
				if rule.r2min <= val && rule.r2max >= val {
					continue
				}
				rule.disqualifiedFields[index] = true
			}
		}
	}
	return badVals
}
