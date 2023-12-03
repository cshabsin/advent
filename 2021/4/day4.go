package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/cshabsin/advent/commongen/readinp"
)

func main() {
	fmt.Println("sample a: ")
	Day4a("sample.txt")
	fmt.Println("---")
	fmt.Print("sample b: ")
	Day4b("sample.txt")
	fmt.Println("---")
	fmt.Println("real a:   ")
	Day4a("input.txt")
	fmt.Println("---")
	fmt.Print("real b:   ")
	Day4b("input.txt")
}

func parseFoo(s string) (string, error) {
	return s, nil
}

type board struct {
	nums  [][]int
	found [][]bool
}

func (b *board) callNum(n int) {
	for i := range b.nums {
		for j := range b.nums[i] {
			if b.nums[i][j] == n {
				b.found[i][j] = true
				return
			}
		}
	}
}

func (b *board) winner() bool {
	for i := range b.nums {
		foundRow := true
		for j := range b.nums[i] {
			if !b.found[i][j] {
				foundRow = false
				break
			}
		}
		if foundRow {
			return true
		}
	}
	for j := range b.nums[0] {
		foundCol := true
		for i := range b.nums {
			if !b.found[i][j] {
				foundCol = false
			}
		}
		if foundCol {
			return true
		}
	}
	// foundDiag := true
	// for i := range b.nums {
	// 	if !b.found[i][i] {
	// 		foundDiag = false
	// 		break
	// 	}
	// }
	// if foundDiag {
	// 	return true
	// }
	// foundDiag = true
	// for i := range b.nums {
	// 	if !b.found[i][4-i] {
	// 		foundDiag = false
	// 		break
	// 	}
	// }
	// if foundDiag {
	// 	return true
	// }
	return false
}

func (b *board) score() int {
	var sum int
	for i := range b.nums {
		for j := range b.nums[0] {
			if !b.found[i][j] {
				sum += b.nums[i][j]
			}
		}
	}
	return sum
}

func (b *board) String() string {
	var s string
	for i := range b.nums {
		for j := range b.nums[0] {
			if b.found[i][j] {
				s += "*"
			} else {
				s += " "
			}
			s += fmt.Sprintf("%2d ", b.nums[i][j])
		}
		s += "\n"
	}
	return s
}

func readBoard(ch chan readinp.Line[string]) (*board, bool, error) {
	var b [][]int
	for i := 0; i < 5; i++ {
		l := <-ch
		line, err := l.Get()
		if err == io.EOF {
			return nil, false, err
		} else if err != nil {
			log.Fatal(err)
		}
		var row []int
		for _, nstr := range strings.Split(line, " ") {
			if nstr == "" {
				continue
			}
			n, err := strconv.Atoi(nstr)
			if err != nil {
				log.Fatal(line, err)
			}
			row = append(row, n)
		}
		b = append(b, row)
	}
	_, more := <-ch
	var found [][]bool
	for range b {
		var l []bool
		for range b[0] {
			l = append(l, false)
		}
		found = append(found, l)
	}
	return &board{nums: b, found: found}, more, nil
}

// Day4a solves part 1 of day 4
func Day4a(fn string) {
	ch, err := readinp.Read(fn, parseFoo)
	if err != nil {
		log.Fatal(err)
	}
	line := <-ch
	first, err := line.Get()
	var balls []int
	for _, bstr := range strings.Split(first, ",") {
		b, err := strconv.Atoi(bstr)
		if err != nil {
			log.Fatal(err)
		}
		balls = append(balls, b)
	}
	<-ch
	var boards []*board
	for {
		b, more, err := readBoard(ch)
		if err == io.EOF {
			break
		}
		boards = append(boards, b)
		if !more {
			break
		}
	}

	for _, ball := range balls {
		for _, b := range boards {
			b.callNum(ball)
			if b.winner() {
				fmt.Println(b.String(), b.score(), ball, b.score()*ball)
				return
			}
		}
	}
	for _, b := range boards {
		fmt.Println(b.String())
	}
	fmt.Println("no winners")
}

// Day4b solves part 2 of day 4
func Day4b(fn string) {
	ch, err := readinp.Read(fn, parseFoo)
	if err != nil {
		log.Fatal(err)
	}
	line := <-ch
	first, err := line.Get()
	var balls []int
	for _, bstr := range strings.Split(first, ",") {
		b, err := strconv.Atoi(bstr)
		if err != nil {
			log.Fatal(err)
		}
		balls = append(balls, b)
	}
	<-ch
	var boards []*board
	for {
		b, more, err := readBoard(ch)
		if err == io.EOF {
			break
		}
		boards = append(boards, b)
		if !more {
			break
		}
	}

	won := map[int]int{}
	var lastBoard int
	for _, ball := range balls {
		for bindex, b := range boards {
			if b.winner() {
				continue
			}
			b.callNum(ball)
			if b.winner() {
				won[bindex] = ball
			}
			if len(won) == len(boards) {
				lastBoard = bindex
				break
			}
		}
	}
	fmt.Println(boards[lastBoard].String(), won[lastBoard], boards[lastBoard].score(), won[lastBoard]*boards[lastBoard].score())
}
