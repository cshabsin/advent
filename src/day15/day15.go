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
	"unicode"
)

func isFieldSeparator(r rune) bool {
	if r == ':' {
		return true
	}
	if r == ',' {
		return true
	}
	if unicode.IsSpace(r) {
		return true
	}
	return false
}

type Ingredient struct {
	Name string
	Capacity int
	Durability int
	Flavor int
	Texture int
	Calories int
}

func ParseIngredient(line string) Ingredient {
	tokens := strings.FieldsFunc(line, isFieldSeparator)
	cap, err := strconv.Atoi(tokens[2])
	if err != nil {
		log.Fatal(err)
	}
	dur, err := strconv.Atoi(tokens[4])
	if err != nil {
		log.Fatal(err)
	}
	fla, err := strconv.Atoi(tokens[6])
	if err != nil {
		log.Fatal(err)
	}
	tex, err := strconv.Atoi(tokens[8])
	if err != nil {
		log.Fatal(err)
	}
	cal, err := strconv.Atoi(tokens[10])
	if err != nil {
		log.Fatal(err)
	}
	return Ingredient{tokens[0], cap, dur, fla, tex, cal}
}

func sum(a []int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

func IncrementCounts(counts []int) error {
	for index := len(counts)-1; index >= 0; index-- {
		counts[index] = counts[index] + 1
		if sum(counts) <= 100 {
			return nil
		}
		counts[index] = 0
	}
	return io.EOF
}

func ScoreIngredients(ingredients []Ingredient, counts []int) int {
	var cap, dur, fla, tex, cal int
	used := 0
	i := 0
	for ; i < len(counts); i++ {
		cap += counts[i] * ingredients[i].Capacity
		dur += counts[i] * ingredients[i].Durability
		fla += counts[i] * ingredients[i].Flavor
		tex += counts[i] * ingredients[i].Texture
		cal += counts[i] * ingredients[i].Calories
		used += counts[i]
	}
	cap += (100-used) * ingredients[i].Capacity
	dur += (100-used) * ingredients[i].Durability
	fla += (100-used) * ingredients[i].Flavor
	tex += (100-used) * ingredients[i].Texture
	cal += (100-used) * ingredients[i].Calories
	if cal != 500 {
		return 0
	}
	if cap <= 0 || dur <= 0 || fla <= 0 || tex <= 0 {
		return 0
	}
	return cap*dur*fla*tex
}

func main() {
	var infile = flag.String("infile", "input15.txt", "Input file")
	flag.Parse()

	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	ingredients := make([]Ingredient, 0, 4)
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		ingredient := ParseIngredient(line)
		ingredients = append(ingredients, ingredient)
	}
	max_value := 0
	counts := make([]int, len(ingredients)-1)
	for {
		val := ScoreIngredients(ingredients, counts)
		if val > max_value {
			max_value = val
		}
		err := IncrementCounts(counts)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Print(max_value, "\n")
}
