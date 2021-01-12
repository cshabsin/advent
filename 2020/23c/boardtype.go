package main

import (
	"fmt"
	"log"
	"strings"
)

type boardType struct {
	entries []boardEntry
}

func (b boardType) Len() int {
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

func (b boardType) Get(loc int) int {
	index, offset := b.indexAndOffsetForLocation(loc)
	return b.entries[index].Get(offset)
}

func (b boardType) indexAndOffsetForLocation(tgt int) (int, int) {
	tgt = tgt % b.Len()
	var loc int
	for index, ent := range b.entries {
		locBegin := loc
		locEnd := loc
		if ent.card == 0 {
			locEnd += ent.rangeEnd - ent.rangeBegin
		}
		if tgt >= locBegin && tgt <= locEnd {
			return index, tgt - locBegin
		}
		loc = locEnd + 1
	}
	log.Fatalf("couldn't find offset for loc %d (board %v)", tgt, b)
	return 0, 0
}

func (b boardType) nextIndexAndOffset(index, offset int) (int, int) {
	if b.entries[index].card != 0 {
		index++
	} else if offset == b.entries[index].rangeEnd-b.entries[index].rangeBegin {
		index++
		offset = 0
	} else {
		offset++
	}
	if index == len(b.entries) {
		return 0, 0
	}
	return index, offset
}

// find returns the index of the given card, or -1 if not present
func (b boardType) Find(card int) int {
	var loc int
	for _, ent := range b.entries {
		if ent.card == card {
			return loc
		}
		if card >= ent.rangeBegin && card <= ent.rangeEnd {
			return loc + card - ent.rangeBegin
		}
		if ent.card == 0 {
			loc += ent.rangeEnd - ent.rangeBegin
		}
		loc++
	}
	return loc
}

// extract3 deletes 3 cards at given location, and returns the card values and
// the new location equivalent to the previous location (in the event that
// cards were deleted at the start of the board, shifting the location).
func (b *boardType) Extract3(loc int) ([]int, int) {
	var foundZero bool
	firstIndex, firstOffset := b.indexAndOffsetForLocation(loc)
	index, offset := firstIndex, firstOffset
	var vals []int
	var shift int // number of entries deleted from start (to subtract from loc)
	for i := 0; i < 3; i++ {
		vals = append(vals, b.entries[index].Get(offset))
		if index == 0 && offset == 0 {
			foundZero = true
		}
		if foundZero {
			shift++
		}
		index, offset = b.nextIndexAndOffset(index, offset)
	}
	var newEntries []boardEntry
	// if we found a 0 then we skip the entry(ies) at the beginning.
	if foundZero {
		if b.entries[index].card == 0 {
			newEntry := boardEntry{
				rangeBegin: b.entries[index].rangeBegin + offset,
				rangeEnd:   b.entries[index].rangeEnd,
			}
			// if newEntry.rangeEnd >= newEntry.rangeBegin {
			newEntries = append(newEntries, newEntry)
			// }
		}
		newEntries = append(newEntries, b.entries[index+1:firstIndex]...)
		if firstOffset != 0 {
			newEntries = append(newEntries, boardEntry{
				rangeBegin: b.entries[firstIndex].rangeBegin,
				rangeEnd:   b.entries[firstIndex].rangeBegin + firstOffset - 1,
			})
		}
	} else {
		newEntries = append(newEntries, b.entries[0:firstIndex]...)
		if firstOffset != 0 {
			newEntries = append(newEntries, boardEntry{
				rangeBegin: b.entries[firstIndex].rangeBegin,
				rangeEnd:   b.entries[firstIndex].rangeBegin + firstOffset - 1,
			})
		}
		if !(index == 0 && offset == 0) {
			// there is content to copy after the end of the extract
			newEntry := boardEntry{
				rangeBegin: b.entries[index].rangeBegin + offset,
				rangeEnd:   b.entries[index].rangeEnd,
			}
			// if newEntry.rangeEnd >= newEntry.rangeBegin {
			newEntries = append(newEntries, newEntry)
			// }
			if index+1 < len(b.entries) {
				newEntries = append(newEntries, b.entries[index+1:len(b.entries)]...)
			}
		}
	}
	b.entries = newEntries
	return vals, loc - shift // TODO what's the location?
}

func (b *boardType) Insert(loc int, vals []int) {

}

func (b boardType) String() string {
	return b.Render(-1)
}

func (b boardType) Render(highlight int) string {
	bld := strings.Builder{}
	bld.WriteString("{")
	for i, ent := range b.entries {
		if i != 0 {
			bld.WriteString(" ")
		}
		if i == highlight {
			bld.WriteString("(")
		}
		bld.WriteString(fmt.Sprintf("%v", ent))
		if i == highlight {
			bld.WriteString(")")
		}
	}
	bld.WriteString("}")
	return bld.String()

}
