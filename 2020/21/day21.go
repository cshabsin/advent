package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
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

	fmt.Println(canonicalizeAllergens(possibleAllergens))
}

// allergen -> set of ingredients
func canonicalizeAllergens(possibleAllergens map[string]map[string]bool) string {
	var sortedAllergens []string
	allergenMap := map[string]string{}
	for {
		if len(possibleAllergens) == 0 {
			break
		}
		var allergen, ingredient string
		for aIter, ingredientMap := range possibleAllergens {
			if len(ingredientMap) == 1 {
				allergen = aIter
				ingredient = getOnlyEntry(ingredientMap)
				break
			}
		}
		allergenMap[allergen] = ingredient
		sortedAllergens = append(sortedAllergens, allergen)
		delete(possibleAllergens, allergen)
		for _, ingredientMap := range possibleAllergens {
			delete(ingredientMap, ingredient)
		}
	}
	fmt.Println(allergenMap)
	sort.Slice(sortedAllergens, func(i, j int) bool { return sortedAllergens[i] < sortedAllergens[j] })
	var ingredientList []string
	for _, allergen := range sortedAllergens {
		ingredientList = append(ingredientList, allergenMap[allergen])
	}
	return strings.Join(ingredientList, ",")
}

func getOnlyEntry(ingredientMap map[string]bool) string {
	for k := range ingredientMap {
		return k
	}
	return ""
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
