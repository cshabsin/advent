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

func count(line string) (code_len int, value_len int) {
	line = strings.Trim(line, "\n")
	// value, err := strconv.Unquote(line)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return len(strconv.Quote(line)), len(line)
}

func main() {
	var infile = flag.String("infile", "input8.txt", "Input filename")
	flag.Parse()
	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	code_len := 0
	value_len := 0
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		current_code_len, current_value_len := count(line)
		code_len += current_code_len
		value_len += current_value_len
	}
	fmt.Printf("Code: %d\nValue: %d\nDiff: %d\n", code_len, value_len,
		code_len - value_len)
}
