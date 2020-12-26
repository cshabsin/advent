package main

import (
	"fmt"
	"log"
	"regexp"
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
	for _, rule := range rules {
		fmt.Println(rule.name, len(rule.disqualifiedFields), rule.disqualifiedFields)
	}
}

type rule struct {
	name                       string
	r1min, r1max, r2min, r2max int

	disqualifiedFields map[int]bool
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
	}
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
	for index, val := range t.fields {
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
			rule.disqualifiedFields[index] = true
		}
		if !anyGood {
			badVals += val
		}
	}
	return badVals
}
