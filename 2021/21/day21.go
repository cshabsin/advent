package main

import (
	"fmt"
	"math/big"
)

func main() {
	// part1()
	part2()
}

func part2() {
	dg := diracGame{
		positions: map[diracState]*big.Int{
			{p1Position: 1, p2Position: 2}: big.NewInt(1),
		},
		p1win: big.NewInt(0),
		p2win: big.NewInt(0),
	}
	fmt.Println(dg.Summary())
	for len(dg.positions) != 0 {
		dg.doTurn()
		fmt.Println("---")
		fmt.Println(dg.Summary())
	}
	fmt.Println("---")
	fmt.Println(dg.Summary())
}

type diracState struct {
	p1Position, p1Score int
	p2Position, p2Score int
}

func (d diracState) after(p1Roll, p2Roll int) diracState {
	p1Position := d.p1Position + p1Roll
	for p1Position > 10 {
		p1Position -= 10
	}
	p1Score := d.p1Score + p1Position

	p2Position := d.p2Position
	p2Score := d.p2Score
	if p1Score < 21 { // p2 doesn't go if p1 won
		p2Position = d.p2Position + p2Roll
		for p2Position > 10 {
			p2Position -= 10
		}
		p2Score = d.p2Score + p2Position
	}
	return diracState{p1Position, p1Score, p2Position, p2Score}
}

type diracGame struct {
	positions    map[diracState]*big.Int // map of (position,score):# of universes in that state
	p1win, p2win *big.Int
}

func (d diracGame) Summary() string {
	p1locations := map[int]*big.Int{}
	p2locations := map[int]*big.Int{}
	p1scores := map[int]*big.Int{}
	p2scores := map[int]*big.Int{}
	for ds, count := range d.positions {
		if p1locations[ds.p1Position] != nil {
			p1locations[ds.p1Position].Add(p1locations[ds.p1Position], count)
		} else {
			p1locations[ds.p1Position] = new(big.Int).Set(count)
		}
		if p2locations[ds.p2Position] != nil {
			p2locations[ds.p2Position].Add(p2locations[ds.p2Position], count)
		} else {
			p2locations[ds.p2Position] = new(big.Int).Set(count)
		}
		if p1scores[ds.p1Score] != nil {
			p1scores[ds.p1Score].Add(p1scores[ds.p1Score], count)
		} else {
			p1scores[ds.p1Score] = new(big.Int).Set(count)
		}
		if p2scores[ds.p2Score] != nil {
			p2scores[ds.p2Score].Add(p2scores[ds.p2Score], count)
		} else {
			p2scores[ds.p2Score] = new(big.Int).Set(count)
		}
	}
	return fmt.Sprintf("p1: {%d} %v\np2: {%d} %v", d.p1win, p1scores, d.p2win, p2scores)
}

func (d *diracGame) doTurn() {
	newPositions := map[diracState]*big.Int{}
	for state, count := range d.positions {
		for p1Roll1 := 1; p1Roll1 <= 3; p1Roll1++ {
			for p1Roll2 := 1; p1Roll2 <= 3; p1Roll2++ {
				for p1Roll3 := 1; p1Roll3 <= 3; p1Roll3++ {
					for p2Roll1 := 1; p2Roll1 <= 3; p2Roll1++ {
						for p2Roll2 := 1; p2Roll2 <= 3; p2Roll2++ {
							for p2Roll3 := 1; p2Roll3 <= 3; p2Roll3++ {
								newState := state.after(p1Roll1+p1Roll2+p1Roll3, p2Roll1+p2Roll2+p2Roll3)
								if newState.p1Score >= 21 {
									d.p1win.Add(d.p1win, count)
								} else if newState.p2Score >= 21 {
									d.p2win.Add(d.p2win, count)
								} else {
									if newPositions[newState] != nil {
										newPositions[newState].Add(newPositions[newState], count)
									} else {
										newPositions[newState] = new(big.Int).Set(count)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	d.positions = newPositions
}

func part1() {
	var d die
	p1 := &player{position: 1}
	p2 := &player{position: 2}
	var loser *player
	for {
		if p1.doTurn(&d) {
			loser = p2
			break
		}
		if p2.doTurn(&d) {
			loser = p1
			break
		}
	}
	fmt.Println(loser.score, d.rolls, loser.score*d.rolls)
}

type player struct {
	position int
	score    int
}

func (p *player) doTurn(d *die) bool {
	r := d.roll() + d.roll() + d.roll()
	p.position += r
	for p.position > 10 {
		p.position -= 10
	}
	p.score += p.position
	return p.score >= 1000
}

type die struct {
	rolls int
	val   int
}

func (d *die) roll() int {
	d.rolls++
	d.val++
	if d.val == 11 {
		d.val = 1
	}
	return d.val
}
