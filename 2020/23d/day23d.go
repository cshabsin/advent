package main

import (
	"fmt"
	"sync"
)

var findStarts map[int]*card

func main() {
	//	c(7), c(8), c(4), c(2), c(3), c(5), c(9), c(1), c(6),
	head := &card{7, nil}
	tail := head.append(8).append(4).append(2).append(3).append(5).append(9).append(1).append(6)
	findStarts = map[int]*card{7: head}
	for i := 10; i <= 1000000; i++ {
		tail = tail.append(i)
		if i%100000 == 0 {
			findStarts[i] = tail
		}
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
	var rc *card
	m := &sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(len(findStarts))
	for first, firstCard := range findStarts {
		go func(first int, firstCard *card) {
			defer wg.Done()
			if first == val {
				m.Lock()
				rc = firstCard
				m.Unlock()
				return
			}
			cur := firstCard.next
			i := 0
			for {
				i++
				if i%50 == 0 {
					m.Lock()
					if rc != nil {
						m.Unlock()
						return
					}
					m.Unlock()
				}
				if cur.val == val {
					m.Lock()
					rc = cur
					m.Unlock()
					return
				}
				if findStarts[cur.val] != nil {
					return
				}
				cur = cur.next
			}
		}(first, firstCard)
	}
	wg.Wait()
	return rc
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
