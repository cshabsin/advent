package pgm

import (
	"strconv"

	"github.com/cshabsin/advent/commongen/readinp"
)

type Compy struct {
	W, X, Y, Z int
	In         string
	inIdx      int
}

func (c *Compy) inp() int {
	inp := readinp.Atoi(string(c.In[c.inIdx]))
	c.inIdx++
	return inp
}

func (c *Compy) XBlock(addToX int) {
	c.X = c.Z % 26
	if addToX < 0 {
		c.Z = c.Z / 26
	}
	c.X += addToX
	if c.X == c.W {
		c.X = 1
	} else {
		c.X = 0
	}
	if c.X == 0 {
		c.X = 1
	} else {
		c.X = 0
	}
}

func (c *Compy) YBlock(addToY int) {
	c.Y = 25*c.X + 1
	c.Z *= c.Y
	c.Y = (c.W + addToY) * c.X
	c.Z += c.Y
}

func Run(i int) bool {
	c := Compy{In: strconv.Itoa(i)}

	c.W = c.inp()
	c.XBlock(15)
	c.YBlock(4)

	c.W = c.inp()
	c.XBlock(14)
	c.YBlock(16)

	c.W = c.inp()
	c.XBlock(11)
	c.YBlock(14)

	c.W = c.inp()
	c.XBlock(-13)
	c.YBlock(3)

	c.W = c.inp()
	c.XBlock(14)
	c.YBlock(11)

	c.W = c.inp()
	c.XBlock(15)
	c.YBlock(13)

	c.W = c.inp()
	c.XBlock(-7)
	c.YBlock(11)

	c.W = c.inp()
	c.XBlock(10)
	c.YBlock(7)

	c.W = c.inp()
	c.XBlock(-12)
	c.YBlock(12)

	c.W = c.inp()
	c.XBlock(15)
	c.YBlock(15)

	c.W = c.inp()
	c.XBlock(-16)
	c.YBlock(13)

	c.W = c.inp()
	c.XBlock(-9)
	c.YBlock(1)

	c.W = c.inp()
	c.XBlock(-8)
	c.YBlock(15)

	c.W = c.inp()
	c.XBlock(-8)
	c.YBlock(4)

	return c.Z == 0
}
