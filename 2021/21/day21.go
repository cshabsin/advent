package main

import (
	"fmt"
)

func main() {
	// part1()
	part2()
}

func part2() {
	dg := diracGame{
		positions: map[diracState]int{
			{p1Position: 1, p2Position: 2}: 1,
		},
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
	p2turn              bool
}

func (d diracState) after(roll int) diracState {
	p1Position, p1Score, p2Position, p2Score := d.p1Position, d.p1Score, d.p2Position, d.p2Score
	if d.p2turn {
		p2Position = d.p2Position + roll
		for p2Position > 10 {
			p2Position -= 10
		}
		p2Score = d.p2Score + p2Position

	} else {
		p1Position = d.p1Position + roll
		for p1Position > 10 {
			p1Position -= 10
		}
		p1Score = d.p1Score + p1Position
	}
	return diracState{p1Position, p1Score, p2Position, p2Score, !d.p2turn}
}

type diracGame struct {
	positions    map[diracState]int // map of (position,score):# of universes in that state
	p1win, p2win int
}

func (d diracGame) Summary() string {
	p1locations := map[int]int{}
	p2locations := map[int]int{}
	p1scores := map[int]int{}
	p2scores := map[int]int{}
	for ds, count := range d.positions {
		p1locations[ds.p1Position] += count
		p2locations[ds.p2Position] += count
		p1scores[ds.p1Score] += count
		p2scores[ds.p2Score] += count
	}
	return fmt.Sprintf("p1: {%d} %v\np2: {%d} %v", d.p1win, p1scores, d.p2win, p2scores)
}

func (d *diracGame) doTurn() {
	newPositions := map[diracState]int{}
	for state, count := range d.positions {
		for _, roll := range []int{3, 4, 4, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 8, 9} {
			var newState diracState
			newState = state.after(roll)
			if newState.p1Score >= 21 {
				d.p1win += count
			} else if newState.p2Score >= 21 {
				d.p2win += count
			} else {
				newPositions[newState] += count
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
