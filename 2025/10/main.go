package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input.txt: %v", err)
	}
	defer f.Close()

	var machines []Machine
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		machines = append(machines, ParseLine(scanner.Text()))
	}
	var total int
	var vals []int
	for _, machine := range machines {
		bf := breadthFirst(machine)
		vals = append(vals, bf)
		total += bf
	}
	fmt.Println(vals)
	fmt.Println(total)
}

func depthFirst(machine Machine, lights []bool, depth int) int {
	if slices.Equal(lights, machine.Goal) {
		return depth
	}
	var best int
	for _, button := range machine.Buttons {
		newLights := slices.Clone(lights)
		for _, idx := range button {
			newLights[idx] = !newLights[idx]
		}
		bf := depthFirst(machine, newLights, depth+1)
		if best == 0 || bf < best {
			best = bf
		}
	}
	return best
}

func breadthFirst(machine Machine) int {
	depth := 0
	currentLights := []Lights{machine.NewLights()}
	for {
		for _, lights := range currentLights {
			if slices.Equal(lights, machine.Goal) {
				return depth
			}
		}
		var nextLights []Lights
		for _, lights := range currentLights {
			for _, button := range machine.Buttons {
				newLights := slices.Clone(lights)
				for _, idx := range button {
					newLights[idx] = !newLights[idx]
				}
				nextLights = append(nextLights, newLights)
			}
		}
		currentLights = nextLights
		depth++
	}
}

type Lights []bool

type Machine struct {
	Goal     Lights
	Buttons  []Button
	Joltages []int
}

func (m Machine) NewLights() Lights {
	return make([]bool, len(m.Goal))
}

type Button []int

func ParseLine(line string) Machine {
	chunks := strings.Split(line, " ")
	goal := ParseGoal(chunks[0])
	var buttons []Button
	for i := 1; i < len(chunks)-1; i++ {
		button := ParseButton(chunks[i])
		buttons = append(buttons, button)
	}
	joltages := ParseButton(chunks[len(chunks)-1])
	return Machine{Goal: goal, Buttons: buttons, Joltages: joltages}
}

func ParseGoal(s string) Lights {
	var goal Lights
	for _, c := range s {
		if c == '[' || c == ']' {
			continue
		}
		goal = append(goal, c == '#')
	}
	return goal
}

func ParseButton(s string) Button {
	vals := strings.Split(s[1:len(s)-1], ",")
	var button Button
	for _, v := range vals {
		vi, _ := strconv.Atoi(v)
		button = append(button, vi)
	}
	return button
}
