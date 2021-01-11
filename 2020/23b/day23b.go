package main

import (
	"fmt"
	"log"
)

func main() {
	board := boardType{
		entries: []boardEntry{
			c(3), c(8), c(9), c(1), c(2), c(5), c(4), c(6), c(7),
			boardEntry{rangeBegin: 10, rangeEnd: 1000000},
		},
	}
	fmt.Println(board.len())
	var current int
	for i := 0; i < 10000000; i++ {
		board.move(current)
		current++
	}
}

type boardType struct {
	entries []boardEntry
}

func (b boardType) len() int {
	var len int
	for _, ent := range b.entries {
		if ent.card != 0 {
			len++
		} else {
			len += ent.rangeEnd - ent.rangeBegin + 1
		}
	}
	return len
}

func (b boardType) get(index int) int {
	index = index % b.len()
	var i int
	for _, ent := range b.entries {
		iBegin := i
		iEnd := i
		if ent.card == 0 {
			iEnd += ent.rangeEnd - ent.rangeBegin
		}
		if index >= iBegin && index <= iEnd {
			return ent.get(index - iBegin)
		}
		i = iEnd + 1
	}
	return 0
}

type cursor struct {
	b         *boardType
	entIndex  int
	i         int
	iEnd      int
	entOffset int
}

func newCursor(b *boardType) *cursor {
	return &cursor{
		b:    b,
		iEnd: b.entries[0].rangeEnd - b.entries[0].rangeBegin,
	}
}

func (c cursor) entry() *boardEntry {
	return &c.b.entries[c.entIndex]
}

func (c cursor) contains(index int) bool {
	return index >= c.i && index <= c.iEnd
}

func (c *cursor) next() {
	c.entIndex++
	c.i = c.iEnd + 1
	if c.entIndex >= len(c.b.entries) {
		c.entIndex = 0
		c.i = 0
	}
	c.iEnd = c.i + c.entry().rangeEnd - c.entry().rangeBegin
}

// NOTE: not an actual seek, but it's currently only ever used from the beginning of an entry.
// if we need actual seek, write code more like advance()
func (c *cursor) seek(index int) {
	if !c.contains(index) {
		log.Fatalf("seek(%d) but entry did not contain index", index)
	}
	c.entOffset = index - c.i
}

func (c *cursor) seekTo(val int) {
	for {
		if c.entry().card == val {
			return
		}
		if c.entry().card == 0 && c.entry().rangeBegin <= val && c.entry().rangeEnd >= val {
			c.entOffset = val - c.entry().rangeBegin
			return
		}
		c.next()
	}
}

func (c *cursor) get() int {
	return c.entry().get(c.entOffset)
}

func (c *cursor) advance() {
	c.entOffset++
	if c.entOffset > c.iEnd-c.i {
		c.next()
		c.entOffset = 0
	}
}

func (c *cursor) write(val int) {
	if c.entry().card != 0 {
		c.entry().card = val
		return
	}
	// we need to split the range.
	var before, after *boardEntry
	if c.entOffset == 1 {
		before = &boardEntry{card: c.entry().rangeBegin}
	} else if c.entOffset != 0 {
		begin := c.entry().rangeBegin
		before = &boardEntry{
			rangeBegin: begin,
			rangeEnd:   begin + c.entOffset - 1,
		}
	}

	var newEntries []boardEntry
	for i := 0; i < c.entIndex; i++ {
		newEntries = append(newEntries, c.b.entries[i])
	}
	if before != nil {
		newEntries = append(newEntries, *before)
		c.entIndex++ // we've inserted one before the current entry now
	}
	newEntries = append(newEntries, boardEntry{card: val})
	c.iEnd = c.i
	c.entOffset = 0
	if after != nil {
		newEntries = append(newEntries, *after)
	}
	for i := c.entIndex + 1; i < len(c.b.entries); i++ {
		newEntries = append(newEntries, c.b.entries[i])
	}
	c.b.entries = newEntries
}

// extract3 removes next 3 cards, returns them, leaves cursor pointing at the same spot (the next card after the third removed card).
func (c *cursor) extract3() []int {
	var newEntries []boardEntry
	for i := 0; i < c.entIndex; i++ {
		newEntries = append(newEntries, c.b.entries[i])
	}

	var extract []int
	newEntIndex := c.entIndex
	for len(extract) < 3 {
		if c.entry().card != 0 {
			extract = append(extract, c.entry().card)
			c.next()
			continue
		}
		if c.entOffset != 0 {
			newEntries = append(newEntries, boardEntry{
				rangeBegin: c.entry().rangeBegin,
				rangeEnd:   c.entry().rangeBegin + c.entOffset - 1,
			})
			newEntIndex++ // inserting before the new location
		}
		for len(extract) < 3 && c.entry().rangeBegin+c.entOffset <= c.entry().rangeEnd {
			extract = append(extract, c.entry().rangeBegin+c.entOffset)
			c.entOffset++
		}
		if c.entry().rangeBegin+c.entOffset <= c.entry().rangeEnd {
			newEntries = append(newEntries, boardEntry{
				rangeBegin: c.entry().rangeBegin + c.entOffset,
				rangeEnd:   c.entry().rangeEnd,
			})
		}
		c.next()
	}

	newEntries = append(newEntries, c.b.entries[c.entIndex:len(c.b.entries)]...)
	c.b.entries = newEntries
	c.entIndex = newEntIndex
	return extract
}

func (c *cursor) insert(vals []int) {
	var newEntries []boardEntry
	for i := 0; i < c.entIndex; i++ {
		newEntries = append(newEntries, c.b.entries[i])
	}
	newEntIndex := c.entIndex
	if c.entOffset != 0 {
		newEntries = append(newEntries, boardEntry{
			rangeBegin: c.entry().rangeBegin,
			rangeEnd:   c.entry().rangeBegin + c.entOffset - 1,
		})
		newEntIndex++ // inserting before the new location
	}
	for _, v := range vals {
		newEntries = append(newEntries, boardEntry{card: v})
	}
	if c.entry().card != 0 {
		newEntries = append(newEntries, c.b.entries[c.entIndex:len(c.b.entries)]...)
	} else {
		newEntries = append(newEntries, boardEntry{
			rangeBegin: c.entry().rangeBegin + c.entOffset,
			rangeEnd:   c.entry().rangeEnd,
		})
		newEntries = append(newEntries, c.b.entries[c.entIndex+1:len(c.b.entries)]...)
	}
	c.b.entries = newEntries
}

func (b *boardType) move(current int) {
	fmt.Println("move", current, *b)
	c := newCursor(b)
	for {
		if c.contains(current) {
			c.seek(current) // set entOffset
			break
		}
		c.next()
	}
	// cursor now points at current
	destinationVal := c.get() - 1 // tentative
	c.advance()
	extract := c.extract3()
	destinationVal = getDestination(destinationVal, extract)
	fmt.Println("looking for destination", destinationVal, "(extract: ", extract, ")")
	c.seekTo(destinationVal)
	c.advance()
	c.insert(extract)
	fmt.Printf("%v\n", c.b)
}

// boardEntry represents either a single card, or a range of cards starting with rangeBegin and ending with rangeEnd (inclusive).
type boardEntry struct {
	card                 int
	rangeBegin, rangeEnd int
}

func (be boardEntry) get(index int) int {
	if be.card != 0 {
		if index != 0 {
			log.Fatalf("be.get(%d) on card %d", index, be.card)
		}
		return be.card
	}
	if index > be.rangeEnd-be.rangeBegin {
		log.Fatalf("be.get(%d) on range %d-%d", index, be.rangeBegin, be.rangeEnd)
	}
	return be.rangeBegin + index
}

func (be boardEntry) String() string {
	if be.card != 0 {
		return fmt.Sprintf("%d", be.card)
	}
	return fmt.Sprintf("%d-%d", be.rangeBegin, be.rangeEnd)
}

func c(card int) boardEntry {
	return boardEntry{card: card}
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
