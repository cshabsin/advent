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
	ingEntries := map[string]*ingredientEntry{}
	for ingredient := range allIngredents {
		ingEntries[ingredient] = createIngredientEntry(ingredient, allAllergens)
	}
	for _, entry := range entries {
		for ingredient := range entry.ingredientSet {
			ingEntries[ingredient].disqualifyFromAllergens(entry.allergens)
		}
	}
	allergyFree := map[string]bool{}
	for _, ent := range ingEntries {
		anyAllergen := false
		for _, allergic := range ent.possibleAllergens {
			if allergic {
				anyAllergen = true
				break
			}
		}
		if !anyAllergen {
			allergyFree[ent.name] = true
		}
	}
	count := 0
	for _, ent := range entries {
		foundAllergyFree := false
		for ing := range ent.ingredientSet {
			if allergyFree[ing] {
				foundAllergyFree = true
				break
			}
		}
		if foundAllergyFree {
			count++
		}
	}
	fmt.Println(count)
}

type entry struct {
	ingredientSet map[string]bool
	allergens     map[string]bool
}

type ingredientEntry struct {
	name              string
	possibleAllergens map[string]bool
}

func (ie *ingredientEntry) disqualifyFromAllergens(allergens map[string]bool) {
	for a := range ie.possibleAllergens {
		if !allergens[a] {
			ie.possibleAllergens[a] = false
		}
	}
}

func createIngredientEntry(ingredient string, allAllergens map[string]bool) *ingredientEntry {
	possibleAllergens := map[string]bool{}
	for allergen := range allAllergens {
		possibleAllergens[allergen] = true
	}
	return &ingredientEntry{ingredient, possibleAllergens}
}

var lineRe = regexp.MustCompile("^([a-z ]*) \\(contains ([a-z ,]*)\\)$")

func parseLine(line string) ([]string, []string) {
	fields := lineRe.FindStringSubmatch(line)
	return strings.Split(fields[1], " "), strings.Split(fields[2], ", ")
}
