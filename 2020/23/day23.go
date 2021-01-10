package main

import "fmt"

func main() {
	//board := boardType{7, 8, 4, 2, 3, 5, 9, 1, 6}
	board := boardType{3, 8, 9, 1, 2, 5, 4, 6, 7}
	for i := 10; i <= 1000000; i++ {
		board = append(board, i)
	}
	var current int
	for i := 0; i < 10000000; i++ {
		board = move(current, board)
		current++
		if current%100 == 0 {
			fmt.Println(current)
		}
	}
	fmt.Println(board)
}

type boardType []int

func (b boardType) get(index int) int {
	return b[index%len(b)]
}

func (b *boardType) set(index, val int) {
	(*b)[index%len(*b)] = val
}

func move(current int, board boardType) boardType {
	var newBoard boardType
	for _, v := range board {
		newBoard = append(newBoard, v)
	}
	var extract []int
	for i := current + 1; i < current+4; i++ {
		extract = append(extract, board.get(i))
	}
	destinationVal := getDestination(board.get(current)-1, extract)
	i := current + 1
	for {
		v := newBoard.get(i + 3)
		newBoard.set(i, v)
		if v == destinationVal {
			break
		}
		i++
	}
	newBoard.set(i+1, extract[0])
	newBoard.set(i+2, extract[1])
	newBoard.set(i+3, extract[2])
	return newBoard
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
