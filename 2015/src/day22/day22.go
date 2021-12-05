package main

import (
	"fmt"
	"log"
)

func main() {
	s := state{
		playerHP:    15,
		playerMana:  500,
		playerArmor: 0,
		bossHP:      500,
		effects: map[string]effect{
			"Shield": {
				armor:    5,
				duration: 4,
			},
			"Poison": {
				dmg:      5,
				duration: 6,
			},
		},
	}
	fmt.Println(s)
	for len(s.effects) != 0 {
		s = s.applyEffects()
		fmt.Println(s)
	}
	Day22a()
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

type effect struct {
	dmg, heal, armor, mana, duration int
}

type spell struct {
	name                                   string
	cost, dmg, heal, armor, mana, duration int
}

func (s spell) effect() effect {
	if s.duration == 0 {
		log.Fatal("effect called on", s.name)
	}
	return effect{s.dmg, s.heal, s.armor, s.mana, s.duration}
}

var (
	spells = []spell{
		{
			name: "Magic Missile",
			cost: 53,
			dmg:  4,
		},
		{
			name: "Drain",
			cost: 73,
			dmg:  2,
			heal: 2,
		},
		{
			name:     "Shield",
			cost:     113,
			armor:    7,
			duration: 6,
		},
		{
			name:     "Poison",
			cost:     173,
			dmg:      3,
			duration: 6,
		},
		{
			name:     "Recharge",
			cost:     229,
			mana:     101,
			duration: 5,
		},
	}
)

// player: 50 hp, 500 mana
// boss: 55 hp, 8 dmg

type state struct {
	spells                            []string
	playerHP, playerMana, playerArmor int
	bossHP, bossDmg                   int
	effects                           map[string]effect
}

func (s state) applyEffects() state {
	newState := state{
		spells:     s.spells,
		playerHP:   s.playerHP,
		playerMana: s.playerMana,
		bossHP:     s.bossHP,
		bossDmg:    s.bossDmg,
		effects:    map[string]effect{},
	}
	for name, eff := range s.effects {
		newState.playerHP += eff.heal
		newState.playerMana += eff.mana
		newState.bossHP -= eff.dmg
		if eff.armor != 0 {
			newState.playerArmor = eff.armor
		}
		duration := eff.duration - 1
		if duration > 0 {
			newState.effects[name] = effect{
				dmg:      eff.dmg,
				heal:     eff.heal,
				armor:    eff.armor,
				mana:     eff.mana,
				duration: duration,
			}
		}
	}
	return newState
}

func (s state) win() bool {
	if s.bossHP <= 0 && s.playerHP > 0 {
		return true
	}
	return false
}

func (s state) lose() bool {
	return s.playerHP <= 0
}

// trySpell attempts to cast the spell in the game state. Then it
// - runs the player's turn.
// - checks for a win. If so, it returns true and the total cost (input cost plus cost of the spell that won).
// - runs the boss's turn.
// - checks for a loss. If so, it returns false and 0.
// - tries each spell recursively.
func (s state) trySpell(cost, prune int, sp spell) (bool, int) {
	if sp.cost > prune {
		return false, 0
	}
	if sp.cost > s.playerMana {
		print("too expensive:", sp.name)
		return false, 0
	}
	print("casting", sp.name)
	s.spells = append(s.spells, sp.name)
	s = s.applyEffects()
	if _, ok := s.effects[sp.name]; ok {
		return false, 0 // can't cast a spell that's in effect.
	}
	if s.win() {
		print("win with cost", cost)
		return true, cost
	}
	if s.lose() {
		print("lose")
		return false, 0
	}
	if sp.duration != 0 {
		s.effects[sp.name] = sp.effect()
	} else {
		s.bossHP -= sp.dmg
		s.playerHP += sp.heal
	}
	cost += sp.cost
	if s.win() {
		print("win with cost", cost)
		return true, cost
	}
	if s.lose() {
		print("lose")
		return false, 0
	}
	s = s.applyEffects()
	if s.win() {
		print("win with cost", cost)
		return true, cost
	}
	if s.lose() {
		print("lose")
		return false, 0
	}
	s.playerHP -= s.bossDmg - s.playerArmor
	if s.win() {
		print("win with cost", cost)
		return true, cost
	}
	if s.lose() {
		print("lose")
		return false, 0
	}
	return true, s.try(cost, prune)
}

func (s state) try(parentCost, prune int) int {
	print(fmt.Sprintf("trying state %+v", s))
	minCost := prune
	for _, spell := range spells {
		win, cost := s.trySpell(0, minCost, spell)
		if win && cost < minCost {
			minCost = cost
		}
	}
	return parentCost + minCost
}

func Day22a() {
	s := state{
		playerHP:   50,
		playerMana: 500,
		bossHP:     55,
		bossDmg:    8,
		effects:    map[string]effect{},
	}
	fmt.Println(s.try(0, 99999))
}
