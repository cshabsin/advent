package main

import "fmt"

func main() {
	//	c(7), c(8), c(4), c(2), c(3), c(5), c(9), c(1), c(6),
	n := make(nexts, 1000000) // n[0] = the card that comes after card 1
	n.SetCardNext(7, 8)
	n.SetCardNext(8, 4)
	n.SetCardNext(4, 2)
	n.SetCardNext(2, 3)
	n.SetCardNext(3, 5)
	n.SetCardNext(5, 9)
	n.SetCardNext(9, 1)
	n.SetCardNext(1, 6)
	n.SetCardNext(6, 10)
	n.SetCardNext(1000000, 7)
	for i := 10; i < 1000000; i++ {
		n.SetCardNext(i, i+1)
	}
	current := 7
	for i := 0; i < 10000000; i++ {
		var extract []int
		nextCard := n.GetNextCard(current)
		extract = append(extract, nextCard)
		extract = append(extract, n.GetNextCard(nextCard))
		extract = append(extract, n.GetNextCard(n.GetNextCard(nextCard)))
		destination := getDestination(current-1, extract)

		n.SetCardNext(current, n.GetNextCard(extract[2]))
		n.SetCardNext(extract[2], n.GetNextCard(destination))
		n.SetCardNext(destination, extract[0])
		current = n.GetNextCard(current)
	}
	fmt.Println(n.GetNextCard(1), n.GetNextCard(n.GetNextCard(1)))
}

type nexts []int

func (n *nexts) SetCardNext(card int, next int) {
	(*n)[card-1] = next - 1
}

func (n nexts) GetNextCard(card int) int {
	return n[card-1] + 1
}

func getDestination(val int, extract []int) int {
	if val < 1 {
		val = 1000000
	}
	for _, e := range extract {
		if e == val {
			return getDestination(val-1, extract)
		}
	}
	return val
}
