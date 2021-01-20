package main

import (
	"fmt"
	"log"
	"strconv"
)

const input = "59793513516782374825915243993822865203688298721919339628274587775705006728427921751430533510981343323758576985437451867752936052153192753660463974146842169169504066730474876587016668826124639010922391218906707376662919204980583671961374243713362170277231101686574078221791965458164785925384486127508173239563372833776841606271237694768938831709136453354321708319835083666223956618272981294631469954624760620412170069396383335680428214399523030064601263676270903213996956414287336234682903859823675958155009987384202594409175930384736760416642456784909043049471828143167853096088824339425988907292558707480725410676823614387254696304038713756368483311"

func main() {
	signal := parseSignal(input)
	// signal := parseSignal("80871224585914546619083218645595")
	repeat := 100
	// signal := parseSignal("12345678")
	// repeat := 4
	fmt.Println(signal)
	for i := 0; i < repeat; i++ {
		signal = fft(signal)
		fmt.Println(signal)
	}
	fmt.Printf("%d%d%d%d%d%d%d%d\n", signal[0], signal[1], signal[2], signal[3], signal[4], signal[5], signal[6], signal[7])
}

func parseSignal(input string) []int {
	var signal []int
	for _, c := range input {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal(err)
		}
		signal = append(signal, v)
	}
	return signal
}

func fft(signal []int) []int {
	var out []int

	for i := 0; i < len(signal); i++ {
		out = append(out, getDigit(signal, i))
	}

	return out
}

func getDigit(signal []int, index int) int {
	coeffs := []int{0, 1, 0, -1}
	coeff := 0
	patternOffset := 1
	out := 0
	for i := 0; i < len(signal); i++ {
		if patternOffset > index {
			patternOffset = 0
			coeff = (coeff + 1) % 4
		}
		out += signal[i] * coeffs[coeff]
		patternOffset++
	}
	if out < 0 {
		out = -out
	}
	return out % 10
}
