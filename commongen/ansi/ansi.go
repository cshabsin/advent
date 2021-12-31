package ansi

import "fmt"

func Clear() {
	fmt.Print("\x1b[2J\x1b[H")
}

func Loc(r, c int) {
	fmt.Printf("\x1b[%d;%dH", r+1, c+1)
}

func Color(i int, s string) {
	fmt.Printf("\x1b[%d;1m%s\x1b[0m", i, s)
}
