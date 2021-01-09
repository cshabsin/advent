package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var entries []entry
	allIngredents := map[string]bool{}
	allAllergens := map[string]bool{}
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		ingredients, allergens := parseLine(line.Value())
		allergenSet := map[string]bool{}
		for _, allergen := range allergens {
			allergenSet[allergen] = true
		}
		ingredientSet := map[string]bool{}
		for _, ingredient := range ingredients {
			ingredientSet[ingredient] = true
		}
		entries = append(entries, entry{ingredientSet, allergenSet})
		for _, ingredient := range ingredients {
			allIngredents[ingredient] = true
		}
		for _, allergen := range allergens {
			allAllergens[allergen] = true
		}
	}

	// allergen -> set of ingredients
	possibleAllergens := map[string]map[string]bool{}
	for _, e := range entries {
		for a := range e.allergens {
			newPossibleAllergens := map[string]bool{}
			if aMap, ok := possibleAllergens[a]; ok {
				for ing := range aMap {
					if e.ingredientSet[ing] {
						newPossibleAllergens[ing] = true
					}
				}
			} else {
				newPossibleAllergens = e.ingredientSet
			}
			possibleAllergens[a] = newPossibleAllergens
		}
	}
	fmt.Println(possibleAllergens)

	allPossibleAllergens := map[string]bool{}
	for _, ingMap := range possibleAllergens {
		for ing := range ingMap {
			allPossibleAllergens[ing] = true
		}
	}
	fmt.Println(allPossibleAllergens)

	count := 0
	for _, ent := range entries {
		for ing := range ent.ingredientSet {
			if !allPossibleAllergens[ing] {
				count++
			}
		}
	}
	fmt.Println(count)
}

type entry struct {
	ingredientSet map[string]bool
	allergens     map[string]bool
}

var lineRe = regexp.MustCompile("^([a-z ]*) \\(contains ([a-z ,]*)\\)$")

func parseLine(line string) ([]string, []string) {
	fields := lineRe.FindStringSubmatch(line)
	return strings.Split(fields[1], " "), strings.Split(fields[2], ", ")
}
