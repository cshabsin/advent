package main

import (
	"strings"
)

type data struct {
	line     string
	patterns []string
	output   []string
}

func parse(line string) (data, error) {
	var d data
	d.line = line
	parts := strings.Split(line, "|")
	for _, out := range strings.Split(parts[0], " ") {
		d.patterns = append(d.patterns, strings.TrimSpace(out))
	}
	for _, out := range strings.Split(parts[1], " ") {
		d.output = append(d.output, strings.TrimSpace(out))
	}
	return d, nil
}
