package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	fmt.Println("-----")
	part2()
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input.txt: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var problems [][]int
	var line string
	// Part 1: Parse and merge ranges.
	// The input format expects ranges (e.g., "5-10") until an empty line.
	for scanner.Scan() {
		line = scanner.Text()
		var entries []int
		for _, s := range strings.Fields(line) {
			i, err := strconv.Atoi(s)
			if err != nil {
				break // it's an operator
			}
			entries = append(entries, i)
		}
		for i, v := range entries {
			if len(problems) <= i {
				problems = append(problems, nil)
			}
			problems[i] = append(problems[i], v)
		}
	}
	var total int
	for i, op := range strings.Fields(line) {
		problem := problems[i]
		switch op {
		case "+":
			answer := 0
			for _, v := range problem {
				answer += v
			}
			fmt.Println("answer ", i, ": ", answer)
			total += answer
		case "*":
			answer := 1
			for _, v := range problem {
				answer *= v
			}
			fmt.Println("answer ", i, ": ", answer)
			total += answer
		default:
			log.Fatal("bad operator")
		}
	}
	fmt.Println(total)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input.txt: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var valuesByColumn [][]byte
	var line string
	// Part 1: Parse and merge ranges.
	// The input format expects ranges (e.g., "5-10") until an empty line.
	for scanner.Scan() {
		line = scanner.Text()
		if line[0] == '+' || line[0] == '*' {
			break
		}
		for i, c := range line {
			if len(valuesByColumn) <= i {
				valuesByColumn = append(valuesByColumn, nil)
			}
			if c == ' ' {
				continue
			}
			valuesByColumn[i] = append(valuesByColumn[i], byte(c))
		}
	}

	var total, current int
	var op rune
	for i, c := range line {
		if c == '*' || c == '+' {
			total += current
			op = c
			if op == '*' {
				current = 1
			} else {
				current = 0
			}
		}
		col := string(valuesByColumn[i])
		if col != "" {
			val, err := strconv.Atoi(col)
			if err != nil {
				log.Fatal(err)
			}
			if op == '*' {
				current *= val
			} else {
				current += val
			}
		}
	}
	total += current
	fmt.Println(total)
}
