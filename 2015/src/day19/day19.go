package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func Search(line string, transformations_sorted []string,
	transformations map[string]string) int {
//	fmt.Println(line)
	if line == "e" {
		fmt.Println("got one.")
		return 0
	}
	min_steps := -1
	for _, transform := range transformations_sorted {
		for line_index, _ := range(line) {
			if line_index + len(transform) > len(line) {
				continue
			}
			if line[line_index:line_index+len(transform)] == transform {
				source := transformations[transform]
//				fmt.Println(source, "->", transform)
				steps := Search(
					line[:line_index] +
					source +
					line[line_index + len(transform):],
					transformations_sorted,
					transformations)
				if min_steps == -1 || min_steps > steps {
					min_steps = steps
				}
			}
		}
	}
	return min_steps
}

func main() {
	var infile = flag.String("infile", "input19.txt", "Input file")
	flag.Parse()

	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)

	// transformations is a reverse map, target: source
	transformations := make(map[string]string)
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if line == "\n" {
			break  // Next read will be the input.
		}
		tokens := strings.Split(strings.Trim(line, "\n"), " => ")
		from := tokens[0]
		if v, ok := transformations[tokens[1]]; ok {
			log.Fatal("assumption false: ", tokens[1], " can result from ", from, " or ", v)
		}
		transformations[tokens[1]] = from
	}
	line, err := rdr.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	line = strings.Trim(line, "\n")

	transformations_sorted := make([]string, len(transformations))
	i := 0
	for k, _ := range transformations {
		transformations_sorted[i] = k
		i++
	}
	sort.Sort(ByLength(transformations_sorted))
	fmt.Println(Search(line, transformations_sorted, transformations))

// 	outputs := make(map[string]bool)
// 	for i, _ := range line {
// 		keys := make([]string, 0, 2)
// 		if i < len(line)-1 {
// 			keys = append(keys, line[i:i+1])
// 		}
// 		if i < len(line)-2 {
// 			keys = append(keys, line[i:i+2])
// 		}
// 		for _, key := range keys {
// 			if values, ok := expansions[key]; ok {
// 				for _, r := range values {
// 					output := line[:i] + r + line[i+len(key):]
// //					fmt.Print("line: \"", line, "\"\ni: ", i, "\nr: \"", r, "\"\noutput: \"", output, "\"\n")
// 					outputs[output] = true
// 				}
// 			}
// 		}
// 	}
// 	fmt.Print(len(outputs))
}
