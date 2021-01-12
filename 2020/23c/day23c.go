package main

import "fmt"

func main() {
	// board := &boardType{[]boardEntry{c(3), c(8), c(9), c(1), c(2), c(5), c(4), c(6), c(7)}}
	board := &boardType{
		entries: []boardEntry{
			c(7), c(8), c(4), c(2), c(3), c(5), c(9), c(1), c(6),
			boardEntry{rangeBegin: 10, rangeEnd: 1000000},
		},
	}
	currentCard := board.Get(0)
	for i := 0; i < 10000000; i++ {
		current := board.Find(currentCard)
		fmt.Println(i, len(board.entries), board.Render(current))
		move(board, current)
		// if i%1000 == 0 {
		// fmt.Println(i, len(board.entries))
		// }
		currentCard = board.Get(board.Find(currentCard) + 1)
	}
	fmt.Println(board)
}

func move(b *boardType, current int) {
	destinationVal := b.Get(current) - 1
	extract := b.Extract3(current + 1)
	destinationVal = getDestination(destinationVal, extract)
	// fmt.Println("seeking", destinationVal)
	b.Insert(b.Find(destinationVal)+1, extract)
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