package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
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
	var totalLights, totalJoltages int
	var lightVals, joltageVals []int
	for _, machine := range machines {
		l := breadthFirstLights(machine)
		lightVals = append(lightVals, l)
		totalLights += l
		v := depthFirstJoltages(machine)
		// v := breadthFirstJoltages(machine)
		// v := matrixVoltages(machine)
		fmt.Println(machine, "-> ", v)
		joltageVals = append(joltageVals, v)
		totalJoltages += v
	}
	fmt.Println(lightVals)
	fmt.Println(totalLights)
	fmt.Println(joltageVals)
	fmt.Println(totalJoltages)
	fmt.Println(totalLights + totalJoltages)
}

type JValues [12]int

func (j JValues) Equal(other JValues) bool {
	return j == other
}

func (j JValues) Press(b Button) JValues {
	newJ := j // copy the array
	for _, idx := range b {
		newJ[idx]++
	}
	return newJ
}

func (j JValues) EqualsTarget(target []int) bool {
	for i, v := range target {
		if j[i] != v {
			return false
		}
	}
	return true
}

// TooBig returns true if the JValues have grown larger than the target, and the search can be pruned.
func (j JValues) TooBig(target []int) bool {
	for i, v := range target {
		if j[i] > v {
			return true
		}
	}
	return false
}

func depthFirstJoltages(machine Machine) int {
	return depthFirstJoltagesHelper(machine, JValues{}, map[JValues]int{})
}

func depthFirstJoltagesHelper(machine Machine, joltages JValues, cache map[JValues]int) int {
	if joltages.EqualsTarget(machine.Joltages) { // success!
		return 0
	}
	if v, ok := cache[joltages]; ok {
		return v
	}
	best := -1
	for _, button := range machine.Buttons {
		newJoltages := joltages.Press(button)
		if newJoltages.TooBig(machine.Joltages) {
			continue
		}
		res := depthFirstJoltagesHelper(machine, newJoltages, cache)
		if res != -1 {
			steps := 1 + res
			if best == -1 || steps < best {
				best = steps
			}
		}
	}
	cache[joltages] = best
	return best
}

func breadthFirstLights(machine Machine) int {
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

func breadthFirstJoltages(machine Machine) int {
	depth := 0
	currentJoltages := [][]int{make([]int, len(machine.Joltages))}
	seen := map[string]bool{}
	for {
		fmt.Println(depth, len(currentJoltages))
		for _, joltages := range currentJoltages {
			if slices.Equal(joltages, machine.Joltages) {
				return depth
			}
		}
		var nextJoltages [][]int
		for _, joltages := range currentJoltages {
			for _, button := range machine.Buttons {
				newJoltages := slices.Clone(joltages)
				for _, idx := range button {
					newJoltages[idx]++
				}
				if seen[sliceToString(newJoltages)] {
					continue
				}
				seen[sliceToString(newJoltages)] = true
				var tooBig bool
				for i, v := range newJoltages {
					if v > machine.Joltages[i] {
						tooBig = true
						break
					}
				}
				if tooBig {
					continue
				}
				nextJoltages = addUnique(nextJoltages, newJoltages)
			}
		}
		currentJoltages = nextJoltages
		depth++
	}
}

func matrixVoltages(machine Machine) int {
	// build coefficient matrix
	// b0, b1, b2... are number of presses of button 0, 1, 2...
	//
	fmt.Println("buttons:", machine.Buttons)
	var coeffMatrix []float64
	for i := range machine.Joltages {
		var coeffRow []float64
		for _, button := range machine.Buttons {
			if button.Includes(i) {
				coeffRow = append(coeffRow, 1)
			} else {
				coeffRow = append(coeffRow, 0)
			}
		}
		coeffMatrix = append(coeffMatrix, coeffRow...)
	}
	coeff := mat.NewDense(len(machine.Joltages), len(machine.Buttons), coeffMatrix)
	fmt.Printf("Coefficient matrix:\n%v\n", mat.Formatted(coeff, mat.Prefix("")))

	fmt.Println(machine.Joltages)
	var joltageList []float64
	for _, j := range machine.Joltages {
		joltageList = append(joltageList, float64(j))
	}
	joltages := mat.NewVecDense(len(machine.Joltages), joltageList)

	fmt.Printf("Joltage vector:\n%v\n", mat.Formatted(joltages, mat.Prefix("")))

	var output mat.VecDense
	if err := output.SolveVec(coeff, joltages); err != nil {
		log.Fatalf("failed to solve matrix: %v", err)
	}
	fmt.Printf("Solution vector:\n%v\n", mat.Formatted(&output, mat.Prefix("")))
	return 0
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

func (b Button) Includes(i int) bool {
	return slices.Contains(b, i)
}

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

func addUnique(a [][]int, b []int) [][]int {
	for _, aa := range a {
		if slices.Equal(aa, b) {
			return a
		}
	}
	return append(a, b)
}

func sliceToString(a []int) string {
	var s string
	for i, v := range a {
		if i > 0 {
			s += ", "
		}
		s += strconv.Itoa(v)
	}
	return s
}
