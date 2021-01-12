package main

import (
	"fmt"
	"log"
)

// boardEntry represents either a single card, or a range of cards starting with rangeBegin and ending with rangeEnd (inclusive).
type boardEntry struct {
	card                 int
	rangeBegin, rangeEnd int
}

func (be boardEntry) Get(index int) int {
	if be.card != 0 {
		if index != 0 {
			log.Fatalf("be.Get(%d) on card %d", index, be.card)
		}
		return be.card
	}
	if index > be.rangeEnd-be.rangeBegin {
		log.Fatalf("be.Get(%d) on range %d-%d", index, be.rangeBegin, be.rangeEnd)
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
