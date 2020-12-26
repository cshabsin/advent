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
	c := computer{version: 2, bitmaskAnd: -1, memory: map[int]int{}}
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
	version               int
	bitmaskOr, bitmaskAnd int
	addressOr             int
	floatBits             []int
	memory                map[int]int
}

var memsetRegexp = regexp.MustCompile(`^mem\[(\d*)\] = (\d*)$`)

func (c *computer) execute(line string) error {
	if strings.HasPrefix(line, "mask = ") {
		maskChars := strings.TrimPrefix(line, "mask = ")
		bitmaskZeroes := 0
		bitmaskOr := 0
		var floatBits []int
		for bit, c := range maskChars {
			bitmaskZeroes *= 2
			bitmaskOr *= 2
			switch c {
			case '1':
				bitmaskOr++
			case '0':
				bitmaskZeroes++
			case 'X':
				floatBits = append(floatBits, 35-bit)
			}
		}
		c.bitmaskOr = bitmaskOr
		c.bitmaskAnd = ^bitmaskZeroes
		c.floatBits = floatBits
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
	var floatBits []int
	switch c.version {
	case 1:
		val |= c.bitmaskOr
		val &= c.bitmaskAnd
	case 2:
		addr |= c.bitmaskOr
		floatBits = c.floatBits
	default:
		return fmt.Errorf("bad version %d", c.version)
	}
	for addr := range floatVals(addr, floatBits) {
		c.memory[addr] = val
	}
	return nil
}

func floatVals(addr int, floatBits []int) chan int {
	out := make(chan int)
	go func() {
		if len(floatBits) == 0 {
			out <- addr
		} else {
			bit := 1 << floatBits[0]
			for floatedAddrs := range floatVals(addr, floatBits[1:len(floatBits)]) {
				out <- floatedAddrs | bit
				out <- floatedAddrs & ^bit
			}
		}
		close(out)
	}()
	return out
}
