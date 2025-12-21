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
		// v := depthFirstJoltages(machine)
		// v := breadthFirstJoltages(machine)
		// v := matrixVoltages(machine)
		v := euclideanJoltages(machine)
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

func euclideanJoltages(machine Machine) int {
	rows := len(machine.Joltages)
	cols := len(machine.Buttons)

	// Build augmented matrix for the linear system Ax = b.
	// Rows correspond to joltage constraints, columns to buttons.
	// The last column is the target joltage values.
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols+1)
		for j, btn := range machine.Buttons {
			if btn.Includes(i) {
				matrix[i][j] = 1
			}
		}
		matrix[i][cols] = machine.Joltages[i]
	}

	// Gaussian elimination with Euclidean steps to reduce to Row Echelon Form over integers.
	// Instead of division (which requires fields like Reals), we use the Euclidean algorithm
	// (repeated subtraction) to compute the GCD of the column entries, ensuring integer integrity.
	pivotRow := 0
	pivots := make([]int, rows)
	for i := range pivots {
		pivots[i] = -1
	}
	colToPivotRow := make(map[int]int)

	for j := 0; j < cols && pivotRow < rows; j++ {
		// Find row with non-zero entry
		sel := -1
		for i := pivotRow; i < rows; i++ {
			if matrix[i][j] != 0 {
				sel = i
				break
			}
		}
		if sel == -1 {
			continue
		}

		// Swap to pivotRow
		matrix[pivotRow], matrix[sel] = matrix[sel], matrix[pivotRow]

		// Eliminate entries below the pivot using Euclidean algorithm steps.
		// This effectively computes the GCD of the column values at the pivot position.
		for i := pivotRow + 1; i < rows; i++ {
			// Repeat until the entry below the pivot is 0
			for matrix[i][j] != 0 {
				// 1. Safety Swap: Ensure we aren't dividing by zero.
				if matrix[pivotRow][j] == 0 {
					matrix[pivotRow], matrix[i] = matrix[i], matrix[pivotRow]
					continue
				}

				// 2. Calculate the multiple (quotient)
				factor := matrix[i][j] / matrix[pivotRow][j]

				// 3. Swap if magnitude is too small
				// If factor is 0, it means abs(matrix[i][j]) < abs(matrix[pivotRow][j]).
				// We swap rows so we are always dividing the larger magnitude by the smaller,
				// ensuring the value shrinks.
				if factor == 0 {
					matrix[pivotRow], matrix[i] = matrix[i], matrix[pivotRow]
					continue
				}

				// 4. Row Operation: R_i = R_i - factor * R_pivot
				// This effectively performs: matrix[i][j] = matrix[i][j] % matrix[pivotRow][j]
				// but applies the operation to the entire row to keep the equation valid.
				for k := j; k <= cols; k++ {
					matrix[i][k] -= factor * matrix[pivotRow][k]
				}
			}
		}

		// Ensure positive pivot for easier back-substitution logic later.
		if matrix[pivotRow][j] < 0 {
			for k := j; k <= cols; k++ {
				matrix[pivotRow][k] = -matrix[pivotRow][k]
			}
		}

		pivots[pivotRow] = j
		colToPivotRow[j] = pivotRow
		pivotRow++
	}

	// Check for consistency. If a row is all zeros but the augmented part is non-zero,
	// the system has no solution (0 = k where k != 0).
	for r := pivotRow; r < rows; r++ {
		if matrix[r][cols] != 0 {
			return 0 // Impossible
		}
	}

	// Identify free variables (columns that do not contain a pivot).
	// These variables can take on multiple values, creating a solution space we must search.
	var freeVars []int
	for j := 0; j < cols; j++ {
		if _, ok := colToPivotRow[j]; !ok {
			freeVars = append(freeVars, j)
		}
	}

	// Calculate upper bounds for free variables to limit the search space.
	// Since all variables must be non-negative, a variable cannot exceed the smallest target joltage
	// it contributes to (assuming coefficients are non-negative, which they are here: 0 or 1).
	bounds := make([]int, cols)
	for j := 0; j < cols; j++ {
		minLimit := -1
		for i := 0; i < rows; i++ {
			if machine.Buttons[j].Includes(i) {
				limit := machine.Joltages[i]
				if minLimit == -1 || limit < minLimit {
					minLimit = limit
				}
			}
		}
		if minLimit == -1 {
			bounds[j] = 0
		} else {
			bounds[j] = minLimit
		}
	}

	minTotal := -1

	// Recursive solver to iterate through valid assignments of free variables.
	// For each assignment, we use back-substitution to find the values of pivot variables.
	var solve func(freeIdx int, currentAssignment []int)
	solve = func(freeIdx int, currentAssignment []int) {
		if freeIdx == len(freeVars) {
			fullAssignment := make([]int, cols)
			for k, v := range currentAssignment {
				fullAssignment[freeVars[k]] = v
			}

			currentSum := 0
			for _, v := range currentAssignment {
				currentSum += v
			}

			possible := true
			// Back substitution: solve for pivot variables from bottom up.
			// equation: coeff * x_pivot + sum(other_terms) = rhs
			for r := pivotRow - 1; r >= 0; r-- {
				pCol := pivots[r]
				rhs := matrix[r][cols]
				sumOther := 0
				for c := pCol + 1; c < cols; c++ {
					sumOther += matrix[r][c] * fullAssignment[c]
				}
				val := rhs - sumOther
				coeff := matrix[r][pCol]

				// Check if integer solution exists for this pivot variable
				if val%coeff != 0 {
					possible = false
					break
				}
				x_p := val / coeff
				// Check non-negativity constraint
				if x_p < 0 {
					possible = false
					break
				}
				fullAssignment[pCol] = x_p
				currentSum += x_p
			}

			if possible {
				if minTotal == -1 || currentSum < minTotal {
					minTotal = currentSum
				}
			}
			return
		}

		fVar := freeVars[freeIdx]
		limit := bounds[fVar]
		for val := 0; val <= limit; val++ {
			solve(freeIdx+1, append(currentAssignment, val))
		}
	}

	solve(0, []int{})

	if minTotal == -1 {
		return 0
	}
	return minTotal
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
