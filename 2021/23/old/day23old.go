package main

func main() {
	makeBoard()
}

func makeBoard() *board {
	var hall []*space
	for i := 0; i < 11; i++ {
		noStop := i == 2 || i == 4 || i == 6 || i == 8
		hall = append(hall, &space{noStop: noStop})
	}
	for i := 0; i < 10; i++ {
		connect(hall[i], hall[i+1])
	}
	a1 := &space{contents: apod, target: apod}
	connect(hall[2], a1)
	a2 := &space{contents: dpod, target: apod}
	connect(a1, a2)
	a3 := &space{contents: dpod, target: apod}
	connect(a2, a3)
	a4 := &space{contents: cpod, target: apod}
	connect(a3, a4)

	b1 := &space{contents: dpod, target: bpod}
	connect(hall[4], b1)
	b2 := &space{contents: cpod, target: bpod}
	connect(b1, b2)
	b3 := &space{contents: bpod, target: bpod}
	connect(b2, b3)
	b4 := &space{contents: dpod, target: bpod}
	connect(b3, b4)

	c1 := &space{contents: apod, target: cpod}
	connect(hall[6], c1)
	c2 := &space{contents: bpod, target: cpod}
	connect(c1, c2)
	c3 := &space{contents: apod, target: cpod}
	connect(c2, c3)
	c4 := &space{contents: bpod, target: cpod}
	connect(c3, c4)

	d1 := &space{contents: cpod, target: dpod}
	connect(hall[8], d1)
	d2 := &space{contents: apod, target: dpod}
	connect(d1, d2)
	d3 := &space{contents: cpod, target: dpod}
	connect(d2, d3)
	d4 := &space{contents: bpod, target: dpod}
	connect(d3, d4)

	return &board{
		hall:    hall,
		aTarget: []*space{a1, a2, a3, a4},
		bTarget: []*space{b1, b2, b3, b4},
		cTarget: []*space{c1, c2, c3, c4},
		dTarget: []*space{d1, d2, d3, d4},
	}
}

func connect(a, b *space) {
	a.neighbors = append(a.neighbors, b)
	b.neighbors = append(b.neighbors, a)
}

type content int

const (
	blank content = iota
	apod
	bpod
	cpod
	dpod
)

type space struct {
	contents  content
	neighbors []*space
	noStop    bool
	target    content
}

type board struct {
	hall    []*space
	aTarget []*space
	bTarget []*space
	cTarget []*space
	dTarget []*space
	aPod    []*pod
	bPod    []*pod
	cPod    []*pod
	dPod    []*pod
}

type pod struct {
	target   content
	location *space
}
