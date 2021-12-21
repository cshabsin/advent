package main

import "fmt"

func main() {
	part1()
	part2()
}

func part2() {
	dg := diracGame{
		positions: map[diracState]int{
			{p1Position: 4, p2Position: 8}: 1,
		},
	}
	fmt.Println(dg)
	for len(dg.positions) != 0 {
		dg.doTurn()
		fmt.Println(dg)
	}
	fmt.Println(dg.p1win)
	fmt.Println(dg.p2win)
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
	positions    map[diracState]int // map of (position,score):# of universes in that state
	p1win, p2win int
}

func (d *diracGame) doTurn() {
	newPositions := map[diracState]int{}
	for state, count := range d.positions {
		for p1Roll1 := 1; p1Roll1 <= 3; p1Roll1++ {
			for p1Roll2 := 1; p1Roll2 <= 3; p1Roll2++ {
				for p1Roll3 := 1; p1Roll3 <= 3; p1Roll3++ {
					for p2Roll1 := 1; p2Roll1 <= 3; p2Roll1++ {
						for p2Roll2 := 1; p2Roll2 <= 3; p2Roll2++ {
							for p2Roll3 := 1; p2Roll3 <= 3; p2Roll3++ {
								newState := state.after(p1Roll1+p1Roll2+p1Roll3, p2Roll1+p2Roll2+p2Roll3)
								if newState.p1Score >= 21 {
									d.p1win += count
								} else if newState.p2Score >= 21 {
									d.p2win += count
								} else {
									newPositions[newState] += count
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
