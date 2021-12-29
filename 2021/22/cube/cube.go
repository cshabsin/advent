package cube

import (
	"github.com/cshabsin/advent/commongen/matrix"
)

type Cube struct {
	On  bool
	Min matrix.Point3
	Max matrix.Point3
}

func (c Cube) Volume() int64 {
	return int64(c.Max.X()-c.Min.X()+1) * int64(c.Max.Y()-c.Min.Y()+1) * int64(c.Max.Z()-c.Min.Z()+1)
}

func (c Cube) Overlaps(d Cube) bool {
	for i := 0; i < 3; i++ {
		if c.Max[i] < d.Min[i] {
			return false
		}
		if c.Min[i] > d.Max[i] {
			return false
		}
	}
	return true
}

func (c Cube) Clone() Cube {
	return Cube{
		On:  c.On,
		Min: c.Min.Clone(),
		Max: c.Max.Clone(),
	}
}

func (c Cube) Subtract(d Cube) []Cube {
	if !c.Overlaps(d) {
		return []Cube{c.Clone()}
	}
	var cleaves []Cube
	overlap := c.Clone()
	for i := 0; i < 3; i++ {
		if d.Min[i] > overlap.Min[i] && d.Min[i] <= overlap.Max[i] {
			cleave := overlap.Clone()
			cleave.Max[i] = d.Min[i] - 1
			cleaves = append(cleaves, cleave)
			overlap.Min[i] = d.Min[i]
		}
		if d.Max[i] >= overlap.Min[i] && d.Max[i] < overlap.Max[i] {
			cleave := overlap.Clone()
			cleave.Min[i] = d.Max[i] + 1
			cleaves = append(cleaves, cleave)
			overlap.Max[i] = d.Max[i]
		}
	}
	return cleaves
}
