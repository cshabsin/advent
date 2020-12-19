package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Reindeer struct {
	Name string
	Speed int
	BurnDuration int
	RestDuration int
}

func MakeReindeer(line string) Reindeer {
	tokens := strings.Split(strings.Trim(line, "\n"), " ")
	name := tokens[0]
	speed, err := strconv.ParseInt(tokens[3], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	burnDuration, err := strconv.ParseInt(tokens[6], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	restDuration, err := strconv.ParseInt(tokens[13], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return Reindeer{name, int(speed), int(burnDuration), int(restDuration)}
}

func DeerDistance(r Reindeer, seconds int) int {
	period := r.BurnDuration + r.RestDuration
	numPeriods := seconds/period
	totalKm := numPeriods * r.Speed * r.BurnDuration
	remainder := int(math.Mod(float64(seconds), float64(period)))
	totalKm += int(math.Min(float64(remainder), float64(r.BurnDuration))) * r.Speed
	return totalKm
}

func main() {
	var infile = flag.String("infile", "input14.txt", "Input file")
	var seconds = flag.Int("seconds", 2503, "Number of seconds")
	flag.Parse()

	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	reindeer := make([]Reindeer, 0, 9)
	for {
		line, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		reindeer = append(reindeer, MakeReindeer(line))
	}
	scores := make([]int, len(reindeer))
	for secs:=1; secs<=*seconds; secs++ {
		leader_index := 0
		leader_distance := 0
		for i, r := range reindeer {
			d := DeerDistance(r, secs)
			if d > leader_distance {
				leader_index = i
				leader_distance = d
			}
		}
		scores[leader_index]++
	}
	for i, r := range reindeer {
		fmt.Print(r.Name, ": ", scores[i], "\n")
	}
}
