package main

import (
	"fmt"
	"log"
	"strconv"
)

const input = "A052E04CFD9DC0249694F0A11EA2044E200E9266766AB004A525F86FFCDF4B25DFC401A20043A11C61838600FC678D51B8C0198910EA1200010B3EEA40246C974EF003331006619C26844200D414859049402D9CDA64BDEF3C4E623331FBCCA3E4DFBBFC79E4004DE96FC3B1EC6DE4298D5A1C8F98E45266745B382040191D0034539682F4E5A0B527FEB018029277C88E0039937D8ACCC6256092004165D36586CC013A008625A2D7394A5B1DE16C0E3004A8035200043220C5B838200EC4B8E315A6CEE6F3C3B9FFB8100994200CC59837108401989D056280803F1EA3C41130047003530004323DC3C860200EC4182F1CA7E452C01744A0A4FF6BBAE6A533BFCD1967A26E20124BE1920A4A6A613315511007A4A32BE9AE6B5CAD19E56BA1430053803341007E24C168A6200D46384318A6AAC8401907003EF2F7D70265EFAE04CCAB3801727C9EC94802AF92F493A8012D9EABB48BA3805D1B65756559231917B93A4B4B46009C91F600481254AF67A845BA56610400414E3090055525E849BE8010397439746400BC255EE5362136F72B4A4A7B721004A510A7370CCB37C2BA0010D3038600BE802937A429BD20C90CCC564EC40144E80213E2B3E2F3D9D6DB0803F2B005A731DC6C524A16B5F1C1D98EE006339009AB401AB0803108A12C2A00043A134228AB2DBDA00801EC061B080180057A88016404DA201206A00638014E0049801EC0309800AC20025B20080C600710058A60070003080006A4F566244012C4B204A83CB234C2244120080E6562446669025CD4802DA9A45F004658527FFEC720906008C996700397319DD7710596674004BE6A161283B09C802B0D00463AC9563C2B969F0E080182972E982F9718200D2E637DB16600341292D6D8A7F496800FD490BCDC68B33976A872E008C5F9DFD566490A14"

func main() {
	parse("D2FE28")
	parse("38006F45291200")
	parse(input)
}

func parse(input string) {
	br := &bitReader{s: input}
	pkt, bits, _ := readPacket(br)
	fmt.Println(input, "(", bits, "bits):")
	fmt.Println(pkt)
	fmt.Println("versions:", sumVersions(pkt))
	fmt.Println("---")
}

func sumVersions(pkt *packet) int {
	vers := pkt.version
	for _, subp := range pkt.subpackets {
		vers += sumVersions(subp)
	}
	return vers
}

const (
	ptLiteral = 4
)

type packet struct {
	version    int
	pType      int
	literal    int // only if pType == ptLiteral
	subpackets []*packet
}

func (p *packet) String() string {
	if p.pType == ptLiteral {
		return fmt.Sprintf("[v%d l%d]", p.version, p.literal)
	}
	return fmt.Sprintf("[v%d t%d %v]", p.version, p.pType, p.subpackets)
}

// readPacket consumes and returns a packet, as well as the total number of bits read or any error.
// It does *not* skip to the next digit boundary (as would be desirable for a top-level read), so it
// can be used recursively with operators. Do this manually with br.discardRestOfDigit()
func readPacket(br *bitReader) (*packet, int, error) {
	version := br.consume(3)
	pType := br.consume(3)
	bits := 6
	if pType == ptLiteral {
		var val int
		for {
			last := br.consume(1) == 0
			val = val*16 + br.consume(4)
			bits += 5
			if last {
				break
			}
		}
		return &packet{
			version: version,
			pType:   pType,
			literal: val,
		}, bits, nil
	}
	lengthType := br.consume(1)
	bits++
	var totalBits, numPackets int
	if lengthType == 0 {
		totalBits = br.consume(15)
		bits += 15
	} else {
		numPackets = br.consume(11)
		bits += 11
	}
	var subpackets []*packet
	var subBits int
	for {
		pkt, pktLen, _ := readPacket(br)
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

	return &packet{
		version:    version,
		pType:      pType,
		subpackets: subpackets,
	}, bits, nil
}

type bitReader struct {
	s           string
	digitCursor int // index into s
	bitCursor   int // 0-2, offset into current digit
}

func (r *bitReader) consumeBit() int {
	digit, err := strconv.ParseInt(string(r.s[r.digitCursor]), 16, 8)
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

func (r *bitReader) consume(n int) int {
	var v int
	for i := 0; i < n; i++ {
		v = v*2 + r.consumeBit()
	}
	return v
}

func (r *bitReader) discardRestOfDigit() {
	if r.bitCursor != 0 {
		r.bitCursor = 0
		r.digitCursor++
	}
}
