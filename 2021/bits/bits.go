package bits

import (
	"fmt"
	"log"
	"strconv"
)

func Calculate(pkt *Packet) int {
	switch pkt.Type {
	case PTSum:
		var sum int
		for _, subp := range pkt.Subpackets {
			sum += Calculate(subp)
		}
		return sum
	case PTProduct:
		prod := 1
		for _, subp := range pkt.Subpackets {
			prod *= Calculate(subp)
		}
		return prod
	case PTMin:
		min := Calculate(pkt.Subpackets[0])
		for _, subp := range pkt.Subpackets {
			cur := Calculate(subp)
			if cur < min {
				min = cur
			}
		}
		return min
	case PTMax:
		max := Calculate(pkt.Subpackets[0])
		for _, subp := range pkt.Subpackets {
			cur := Calculate(subp)
			if cur > max {
				max = cur
			}
		}
		return max
	case PTLiteral:
		return pkt.Literal
	case PTGT:
		if Calculate(pkt.Subpackets[0]) > Calculate(pkt.Subpackets[1]) {
			return 1
		} else {
			return 0
		}
	case PTLT:
		if Calculate(pkt.Subpackets[0]) < Calculate(pkt.Subpackets[1]) {
			return 1
		} else {
			return 0
		}
	case PTEq:
		if Calculate(pkt.Subpackets[0]) == Calculate(pkt.Subpackets[1]) {
			return 1
		} else {
			return 0
		}
	}
	log.Fatal("bad type", pkt.Type)
	return 0
}

const (
	PTSum     = 0
	PTProduct = 1
	PTMin     = 2
	PTMax     = 3
	PTLiteral = 4
	PTGT      = 5
	PTLT      = 6
	PTEq      = 7
)

type Packet struct {
	Version    int
	Type       int
	Literal    int // only if pType == ptLiteral
	Subpackets []*Packet
}

func (p *Packet) String() string {
	if p.Type == PTLiteral {
		return fmt.Sprintf("[v%d l%d]", p.Version, p.Literal)
	}
	return fmt.Sprintf("[v%d t%d %v]", p.Version, p.Type, p.Subpackets)
}

// ReadPacket consumes and returns a packet, as well as the total number of bits read or any error.
// It does *not* skip to the next digit boundary (as would be desirable for a top-level read), so it
// can be used recursively with operators. Do this manually with br.discardRestOfDigit()
func ReadPacket(br *BitReader) (*Packet, int, error) {
	version := br.Consume(3)
	pType := br.Consume(3)
	bits := 6
	if pType == PTLiteral {
		var val int
		for {
			last := br.Consume(1) == 0
			val = val*16 + br.Consume(4)
			bits += 5
			if last {
				break
			}
		}
		return &Packet{
			Version: version,
			Type:    pType,
			Literal: val,
		}, bits, nil
	}
	lengthType := br.Consume(1)
	bits++
	var totalBits, numPackets int
	if lengthType == 0 {
		totalBits = br.Consume(15)
		bits += 15
	} else {
		numPackets = br.Consume(11)
		bits += 11
	}
	var subpackets []*Packet
	var subBits int
	for {
		pkt, pktLen, _ := ReadPacket(br)
		subpackets = append(subpackets, pkt)
		subBits += pktLen
		bits += pktLen
		if lengthType == 0 {
			if subBits == totalBits {
				break
			}
			if subBits > totalBits {
				log.Fatal("too many bits?", bits, totalBits, *pkt)
			}
		} else {
			if len(subpackets) == numPackets {
				break
			}
		}
	}

	return &Packet{
		Version:    version,
		Type:       pType,
		Subpackets: subpackets,
	}, bits, nil
}

type BitReader struct {
	S           string
	digitCursor int // index into s
	bitCursor   int // 0-2, offset into current digit
}

func (r *BitReader) ConsumeBit() int {
	digit, err := strconv.ParseInt(string(r.S[r.digitCursor]), 16, 8)
	if err != nil {
		log.Fatal(err)
	}
	digit = (digit >> (3 - r.bitCursor)) & 1
	r.bitCursor++
	if r.bitCursor == 4 {
		r.bitCursor = 0
		r.digitCursor++
	}
	return int(digit)
}

func (r *BitReader) Consume(n int) int {
	var v int
	for i := 0; i < n; i++ {
		v = v*2 + r.ConsumeBit()
	}
	return v
}

func (r *BitReader) DiscardRestOfDigit() {
	if r.bitCursor != 0 {
		r.bitCursor = 0
		r.digitCursor++
	}
}
