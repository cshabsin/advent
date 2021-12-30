package main

import "fmt"

func main() {
	// row 2981, col 3075
	// (row-1)*row/2 = 4441690
	// 4441690+3075-1 = 4444764
	// (20151125*(252533)^4444764)%33554393
	// 20151125*15053307 % 33554393
	// 303341071020375 % 33554393

	// fuck it we're doing it live
	row, col := 1, 1
	var val int64 = 20151125
	for {
		if row == 2981 && col == 3075 {
			fmt.Println(val)
			return
		}
		val = (val * 252533) % 33554393
		if row == 1 {
			row = col + 1
			col = 1
			fmt.Println(row)
			continue
		}
		row--
		col++
	}
}
