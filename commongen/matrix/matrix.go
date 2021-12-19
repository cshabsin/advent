package matrix

import "github.com/cshabsin/advent/commongen/set"

var (
	ident = Matrix{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
)

type Matrix [3][3]int

func Ident() Matrix {
	return ident.Clone()
}

func (m Matrix) Clone() Matrix {
	rc := Matrix{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rc[i][j] = m[i][j]
		}
	}
	return rc
}

func (m Matrix) Mul(n Matrix) Matrix {
	rc := Matrix{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rc[i][j] = m.dot(n, i, j)
		}
	}
	return rc
}

func (m Matrix) Pow(n int) Matrix {
	rc := ident
	for i := 0; i < n; i++ {
		rc = rc.Mul(m)
	}
	return rc
}

func (m Matrix) dot(n Matrix, i, j int) int {
	var rc int
	for k := 0; k < 3; k++ {
		rc += m[i][k] * n[k][j]
	}
	return rc
}

func AllRotations() []Matrix {
	if allRotations == nil {
		allRotations = calcAllRotations()
	}
	return allRotations
}

var (
	allRotations []Matrix
)

func calcAllRotations() []Matrix {
	zRot := Matrix{ // rotate around z axis
		{0, -1, 0},
		{1, 0, 0},
		{0, 0, 1},
	}
	yRot := Matrix{ //rotate around y axis
		{0, 0, -1},
		{0, 1, 0},
		{1, 0, 0},
	}
	xRot := Matrix{ // rotate around x axis
		{1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	}
	rotSet := set.Set[Matrix]{}
	rotSet.Add(Ident())
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			for z := 0; z < 4; z++ {
				val := xRot.Pow(x).Mul(yRot.Pow(y)).Mul(zRot.Pow(z))
				rotSet.Add(val)
			}
		}
	}
	return rotSet.AsSlice()
}
