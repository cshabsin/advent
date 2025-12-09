package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

// CustomSplitFunc splits the input by a specific delimiter, ignoring empty tokens.
func CustomSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Find the index of the delimiter
	if i := strings.IndexByte(string(data), ','); i >= 0 {
		return i + 1, data[:i], nil // Return the token and advance past the delimiter
	}

	// If at EOF and there's data left, return it as a token
	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}

	// No delimiter found yet, and not at EOF, so ask for more data
	return 0, nil, nil
}

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var total int
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		val := line.Value()
		reader := strings.NewReader(val)
		scanner := bufio.NewScanner(reader)

		scanner.Split(CustomSplitFunc)
		for scanner.Scan() {
			token := scanner.Text()
			vals := strings.Split(token, "-")
			first, err := strconv.Atoi(vals[0])
			if err != nil {
				log.Fatal(err)
			}
			second, err := strconv.Atoi(vals[1])
			if err != nil {
				log.Fatal(err)
			}
			for i := first; i <= second; i++ {
				if isInvalid(i) {
					fmt.Println(i)
					total += i
				}
			}
		}
	}
	fmt.Println("total:", total)
}

// isInvalid checks if a number is "invalid". A number is considered invalid
// if its string representation can be formed by repeating a smaller string
// multiple times. For example, 1212 is invalid because it's "12" repeated twice.
func isInvalid(i int) bool {
	s := fmt.Sprintf("%d", i)
	length := len(s)

	// We check for all possible numbers of repetitions, from the length of the string down to 2.
	for numRepeats := length; numRepeats > 1; numRepeats-- {
		// If the length isn't divisible by numRepeats, then it can't be formed by repetition.
		if length%numRepeats != 0 {
			continue
		}
		partLen := length / numRepeats
		substring := s[:partLen]
		if strings.Repeat(substring, numRepeats) == s {
			return true
		}
	}
	return false
}
