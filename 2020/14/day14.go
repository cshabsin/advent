package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	c := computer{bitmaskAnd: -1, memory: make([]int, 100000)}
	for line := range ch {
		if line.Error != nil {
			log.Fatal(err)
		}
		if err := c.execute(strings.TrimSpace(*line.Contents)); err != nil {
			log.Fatal(err)
		}
	}
	total := 0
	for _, val := range c.memory {
		total += val
	}
	fmt.Println(total)
}

type computer struct {
	bitmaskOr, bitmaskAnd int
	memory                []int
}

var memsetRegexp = regexp.MustCompile(`^mem\[(\d*)\] = (\d*)$`)

func (c *computer) execute(line string) error {
	if strings.HasPrefix(line, "mask = ") {
		maskChars := strings.TrimPrefix(line, "mask = ")
		bitmaskZeroes := 0
		bitmaskOr := 0
		for _, c := range maskChars {
			bitmaskZeroes *= 2
			bitmaskOr *= 2
			switch c {
			case 'X':
				continue
			case '1':
				bitmaskOr++
			case '0':
				bitmaskZeroes++
			}
		}
		c.bitmaskOr = bitmaskOr
		c.bitmaskAnd = ^bitmaskZeroes
		return nil
	}
	fields := memsetRegexp.FindStringSubmatch(line)
	if fields == nil {
		return fmt.Errorf("line didn't match: %q", line)
	}
	addr, err := strconv.Atoi(fields[1])
	if err != nil {
		return err
	}
	val, err := strconv.Atoi(fields[2])
	if err != nil {
		return err
	}
	val |= c.bitmaskOr
	val &= c.bitmaskAnd
	c.memory[addr] = val
	return nil
}
