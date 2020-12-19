package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func CountNotRed(m interface{}) int64 {
	var rc int64
	switch val := m.(type) {
	case float64:
		return int64(val)
	case []interface{}:
		for _, u := range(val) {
			rc += CountNotRed(u)
		}
	case map[string]interface{}:
		for _, v := range(val) {
			switch inner_val := v.(type) {
			case string:
				if inner_val == "red" {
					return 0
				}
			default:
				rc += CountNotRed(inner_val)
			}
		}
	}
	return rc
}

func main() {
	var infile = flag.String("infile", "input12.txt", "Input file")
	flag.Parse()

	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	var total int64
	for {
		bytes, err := rdr.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var m interface{}
		err = json.Unmarshal(bytes, &m)
		total += CountNotRed(m)

	}
	fmt.Print("total: ", total, "\n")
}
