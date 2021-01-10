package main

import (
	"fmt"
	"hash/fnv"
	"log"
	"strconv"

	"github.com/cshabsin/advent/common/readinp"
)

func main() {
	ch, err := readinp.Read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	players := make([]deck, 2)
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
	fmt.Println("regular game score:", score(out))

	fmt.Println(playRecurse([]deck{
		deck{9, 2, 6, 3, 1},
		deck{5, 8, 4, 7, 10},
	}))
	fmt.Println(playRecurse([]deck{
		deck{43, 19},
		deck{2, 29, 14},
	}))
	fmt.Println("---")
	fmt.Println("recursive game score:", playRecurse(players))
}

type deck []int

func (d deck) hash() uint64 {
	h := fnv.New64a()
	for _, c := range d {
		h.Write([]byte(strconv.Itoa(c)))
	}
	return h.Sum64()
}

func (d deck) pop() (int, deck) {
	var r deck
	for i := 1; i < len(d); i++ {
		r = append(r, d[i])
	}
	return d[0], r
}

func score(out deck) int {
	score := 0
	for i := range out {
		score += out[i] * (len(out) - i)
	}
	return score
}

func play(players []deck) deck {
	for {
		if len(players[0]) == 0 {
			return players[1]
		} else if len(players[1]) == 0 {
			return players[0]
		}
		players = playRound(players)
	}
}

func playRound(players []deck) []deck {
	newPlayers := make([]deck, 2)
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

func playRecurse(players []deck) int {
	game := newRecursiveGame()
	_, score := game.play(players)
	return score
}

type recursiveGame struct {
	previousRounds map[uint64]map[uint64]bool
}

func newRecursiveGame() recursiveGame {
	return recursiveGame{map[uint64]map[uint64]bool{}}
}

func (r *recursiveGame) addRound(h1, h2 uint64) {
	if r.previousRounds[h1] == nil {
		r.previousRounds[h1] = map[uint64]bool{}
	}
	r.previousRounds[h1][h2] = true
	if r.previousRounds[h2] == nil {
		r.previousRounds[h2] = map[uint64]bool{}
	}
	r.previousRounds[h2][h1] = true
}

// play returns the winning player (0 or 1) and their score
func (r *recursiveGame) play(players []deck) (int, int) {
	for {
		if len(players[0]) == 0 {
			return 1, score(players[1])
		} else if len(players[1]) == 0 {
			return 0, score(players[0])
		}
		h1 := players[0].hash()
		h2 := players[1].hash()
		if r.previousRounds[h1][h2] || r.previousRounds[h2][h1] {
			fmt.Println("winner by loop:", players[0])
			return 0, score(players[0])
		}
		r.addRound(h1, h2)

		c0, d0 := players[0].pop()
		c1, d1 := players[1].pop()
		var w int
		if len(d0) >= c0 && len(d1) >= c1 {
			subgame := newRecursiveGame()
			w, _ = subgame.play([]deck{d0, d1})
		} else {
			if c0 > c1 {
				w = 0
			} else {
				w = 1
			}
		}
		if w == 0 {
			d0 = append(d0, c0, c1)
		} else {
			d1 = append(d1, c1, c0)
		}
		players[0] = d0
		players[1] = d1
	}
}
