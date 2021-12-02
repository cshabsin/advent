package main

import (
	"fmt"
)

// Boss
// Hit Points: 100
// Damage: 8
// Armor: 2

type item struct {
	name             string
	cost, dmg, armor int
}

var (
	weapons = []item{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	}
	armors = []item{
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
	}
	rings = []item{
		{"Damage +1", 25, 1, 0},
		{"Damage +2", 50, 2, 0},
		{"Damage +3", 100, 3, 0},
		{"Defense +1", 20, 0, 1},
		{"Defense +2", 40, 0, 2},
		{"Defense +3", 80, 0, 3},
	}
)

type boss struct {
	hp, dmg, armor int
}

func main() {
	inv := inventory{weapons[1], armors[4]}
	fmt.Println(inv.dmg(), inv.armor())
	fmt.Println(inv.beats(8, boss{12, 7, 2}))

	Day21(boss{100, 8, 2})
	fmt.Println("---")
	Day21b(boss{100, 8, 2})
}

func Day21(b boss) {
	cost := 999999
	var winning inventory
	for _, weapon := range weapons {
		inv := inventory{weapon}
		if inv.beats(100, b) {
			if inv.cost() < cost {
				cost = inv.cost()
				winning = inv
			}
		}
		for _, armor := range armors {
			inv := inventory{weapon, armor}
			if inv.beats(100, b) {
				if inv.cost() < cost {
					cost = inv.cost()
					winning = inv
				}
			} else {
				for _, ring := range rings {
					inv := inventory{weapon, armor, ring}
					if inv.beats(100, b) {
						if inv.cost() < cost {
							cost = inv.cost()
							winning = inv
						}
					}
					for _, ring2 := range rings {
						if ring2 == ring {
							continue
						}
						inv := inventory{weapon, armor, ring, ring2}
						if inv.beats(100, b) {
							if inv.cost() < cost {
								cost = inv.cost()
								winning = inv
							}
						}
					}
				}
			}
		}
		for _, ring := range rings {
			inv := inventory{weapon, ring}
			if inv.beats(100, b) {
				if inv.cost() < cost {
					cost = inv.cost()
					winning = inv
				}
			}
			for _, ring2 := range rings {
				if ring2 == ring {
					continue
				}
				inv := inventory{weapon, ring, ring2}
				if inv.beats(100, b) {
					if inv.cost() < cost {
						cost = inv.cost()
						winning = inv
					}
				}
			}
		}
	}
	fmt.Println(cost, winning)
}

func Day21b(b boss) {
	cost := 0
	var losing inventory
	for _, weapon := range weapons {
		inv := inventory{weapon}
		if !inv.beats(100, b) {
			if inv.cost() > cost {
				cost = inv.cost()
				losing = inv
			}
		}
		for _, armor := range armors {
			inv := inventory{weapon, armor}
			if !inv.beats(100, b) {
				if inv.cost() > cost {
					cost = inv.cost()
					losing = inv
				}
			} else {
				for _, ring := range rings {
					inv := inventory{weapon, armor, ring}
					if !inv.beats(100, b) {
						if inv.cost() > cost {
							cost = inv.cost()
							losing = inv
						}
					}
					for _, ring2 := range rings {
						if ring2 == ring {
							continue
						}
						inv := inventory{weapon, armor, ring, ring2}
						if !inv.beats(100, b) {
							if inv.cost() > cost {
								cost = inv.cost()
								losing = inv
							}
						}
					}
				}
			}
		}
		for _, ring := range rings {
			inv := inventory{weapon, ring}
			if !inv.beats(100, b) {
				if inv.cost() > cost {
					cost = inv.cost()
					losing = inv
				}
			}
			for _, ring2 := range rings {
				if ring2 == ring {
					continue
				}
				inv := inventory{weapon, ring, ring2}
				if !inv.beats(100, b) {
					if inv.cost() > cost {
						cost = inv.cost()
						losing = inv
					}
				}
			}
		}
	}
	fmt.Println(cost, losing)
	fmt.Println(losing.beats(100, b))
}

type inventory []item

func (inv inventory) cost() int {
	var cost int
	for _, item := range inv {
		cost += item.cost
	}
	return cost
}

func (inv inventory) dmg() int {
	var dmg int
	for _, item := range inv {
		dmg += item.dmg
	}
	return dmg
}

func (inv inventory) armor() int {
	var armor int
	for _, item := range inv {
		armor += item.armor
	}
	return armor
}

func (inv inventory) beats(playerHP int, b boss) bool {
	bossDmg := b.dmg - inv.armor()
	if bossDmg <= 0 {
		return true
	}
	playerDuration := playerHP / bossDmg
	playerDmg := inv.dmg() - b.armor
	if playerDmg <= 0 {
		return false
	}
	bossDuration := b.hp / playerDmg
	return playerDuration >= bossDuration
}
