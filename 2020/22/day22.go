package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	players := make([][]int, 2)
	player := 0
	for line := range ch {
		if line.Error != nil {
			log.Fatal(line.Error)
		}
		if line.Value() == "" {
			continue
		}
		if line.Value() == "Player 1:" {
			player = 0
			continue
		}
		if line.Value() == "Player 2:" {
			player = 1
			continue
		}
		value, err := strconv.Atoi(line.Value())
		if err != nil {
			log.Fatal(err)
		}
		players[player] = append(players[player], value)
	}
	// out := play([][]int{
	// 	[]int{9, 2, 6, 3, 1},
	// 	[]int{5, 8, 4, 7, 10},
	// })
	// fmt.Println("---")
	out := play(players)
	score := 0
	for i := range out {
		score += out[i] * (len(out) - i)
	}
	fmt.Println("regular game score:", score)
}

func play(players [][]int) []int {
	for {
		if len(players[0]) == 0 {
			return players[1]
		} else if len(players[1]) == 0 {
			return players[0]
		}
		players = playRound(players)
	}
}

func playRound(players [][]int) [][]int {
	newPlayers := make([][]int, 2)
	for i := 1; i < len(players[0]); i++ {
		newPlayers[0] = append(newPlayers[0], players[0][i])
	}
	for i := 1; i < len(players[1]); i++ {
		newPlayers[1] = append(newPlayers[1], players[1][i])
	}
	p0 := players[0][0]
	p1 := players[1][0]
	if p0 > p1 {
		newPlayers[0] = append(newPlayers[0], p0, p1)
	} else {
		newPlayers[1] = append(newPlayers[1], p1, p0)
	}
	return newPlayers
}
