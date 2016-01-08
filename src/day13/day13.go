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
)

func AddEdge(edges map[string]map[string]int64, line string) {
	tokens := strings.Split(strings.Trim(line, "\n"), " ")
	source := tokens[0]
	target := strings.Trim(tokens[10], ".")
	sign := tokens[2] != "gain"  // true = negative
	value, err := strconv.ParseInt(tokens[3], 10, 32)
	if sign {
		value = -value
	}
	if err != nil {
		log.Fatal(err)
	}
	if edges[source] == nil {
		edges[source] = make(map[string]int64)
	}
	edges[source][target] = value
}

func RemoveElement(list []string, elem string) []string {
	out := make([]string, 0, len(list))
	list_index := 0
	for ; list_index < len(list); {
		if list[list_index] != elem {
			out = append(out, list[list_index])
		}
		list_index++
	}
	return out
}

type SearchState struct {
	Edges map[string]map[string]int64
	InitialNode string
	CurrentNode string
	RemainingNodes []string
	Depth int
}

func initializeSearch(edges map[string]map[string]int64, nodes []string) SearchState {
	return SearchState{edges, nodes[0], nodes[0], nodes[1:], 0}
}

func findHappiest(state SearchState) int64 {
	if len(state.RemainingNodes) == 0 {
		return (state.Edges[state.CurrentNode][state.InitialNode] +
			state.Edges[state.InitialNode][state.CurrentNode])
	}
	var maxTotal int64
	var maxNode string
	for _, nextNode := range state.RemainingNodes {
		var curTotal int64
		curTotal = state.Edges[state.CurrentNode][nextNode] +
			state.Edges[nextNode][state.CurrentNode]
		nextState := SearchState{state.Edges, state.InitialNode,
			nextNode, RemoveElement(state.RemainingNodes, nextNode),
			state.Depth + 1}
		curTotal += findHappiest(nextState)
		if maxTotal == 0 || curTotal > maxTotal {
			maxTotal = curTotal
			maxNode = nextNode
		}
	}
	fmt.Print("maxNode: ", maxNode, " (", state.Depth, ")\n")
	return maxTotal
}

func FindHappiest(edges map[string]map[string]int64, nodes []string) int64 {
	return findHappiest(initializeSearch(edges, nodes))
}

func main() {
	var infile = flag.String("infile", "input13.txt", "Input file")
	flag.Parse()

	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	edges := make(map[string]map[string]int64)
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		AddEdge(edges, line)
	}
	names := make([]string, len(edges))
	i := 0
	for name, _ := range(edges) {
		names[i] = name
		i++
		if edges["me"] == nil {
			edges["me"] = make(map[string]int64)
		}
		edges["me"][name] = 0
		edges[name]["me"] = 0
	}
	names = append(names, "me")
	fmt.Print("names: ", names, "\n")
	fmt.Print(FindHappiest(edges, names))
}
