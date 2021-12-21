package main

import "fmt"

func main() {
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
