package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

var bagRuleRegex = regexp.MustCompile(`^([a-z ]*) bags contain (.*)$`)
var subbagRegex = regexp.MustCompile(`(\d) ([a-z ]*) bag`)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	bagMap := map[string][]string{}
	rbm := reverseBagMap{}
	for line := range ch {
		if line.Error != nil {
			log.Fatal(err)
		}
		cont := strings.TrimSpace(*line.Contents)
		fields := bagRuleRegex.FindStringSubmatch(cont)
		color := fields[1]
		subBagEntries := strings.Split(fields[2], ",")
		var subBags []subBag
		for _, subBagEntry := range subBagEntries {
			if subBagEntry == "no other bags." {
				continue
			}
			subBagFields := subbagRegex.FindStringSubmatch(subBagEntry)
			if len(subBagFields) < 3 {
				log.Fatalf("%q len(subBagFields): %d", subBagEntry, len(subBagFields))
			}
			num, err := strconv.Atoi(subBagFields[1])
			if err != nil {
				log.Fatal(err)
			}
			sb := subBag{num, subBagFields[2]}
			subBags = append(subBags, sb)
		}
		for _, subBag := range subBags {
			bagMap[color] = append(bagMap[color], subBag.color)
			rbm[subBag.color] = append(rbm[subBag.color], color)
		}
	}
	canContainShinyGold := map[string]bool{}
	rbm.findAllChildren(canContainShinyGold, "shiny gold")
	fmt.Println(len(canContainShinyGold))
}

type subBag struct {
	num   int
	color string
}

type reverseBagMap map[string][]string

func (rbm reverseBagMap) findAllChildren(m map[string]bool, color string) {
	for _, subColor := range rbm[color] {
		m[subColor] = true
		rbm.findAllChildren(m, subColor)
	}
}
