package main

import "fmt"

func main() {
	//	c(7), c(8), c(4), c(2), c(3), c(5), c(9), c(1), c(6),
	head := &card{7, nil}
	tail := head.append(8).append(4).append(2).append(3).append(5).append(9).append(1).append(6)
	for i := 10; i <= 1000000; i++ {
		tail = tail.append(i)
	}
	tail.next = head
	fmt.Println(head.len())
	current := head
	for i := 0; i < 10000000; i++ {
		next3 := current.next
		next3tail := next3.next.next
		current.next = next3tail.next

		destination := current.find(getDestination(current.val-1, []int{next3.val, next3.next.val, next3.next.next.val}))
		next3tail.next = destination.next
		destination.next = next3
		if i%1000 == 0 {
			fmt.Println(i)
		}
		current = current.next
	}
	one := current.find(1)
	fmt.Println(one.next.val * one.next.next.val)
}

type card struct {
	val  int
	next *card
}

func (c *card) append(val int) *card {
	next := &card{val, nil}
	c.next = next
	return next
}

func (c *card) len() int {
	i := 0
	cur := c
	for {
		cur = cur.next
		i++
		if cur == c {
			return i
		}
	}
}

func (c *card) find(val int) *card {
	for {
		if c.val == val {
			return c
		}
		c = c.next
	}
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
