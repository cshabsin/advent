package matrix

import "github.com/cshabsin/advent/commongen/set"

var (
	ident = Matrix3x3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
)

// Point3 is a point in 3-space.
type Point3 [3]int

func (p Point3) Sub(q Point3) Vector3 {
	return Vector3{q[0] - p[0], q[1] - p[1], q[2] - p[2]}
}

func (p Point3) Mul(rot Matrix3x3) Point3 {
	return Point3(Vector3(p).Mul(rot))
}

// Vector3 is a vector in 3-space, or a difference between two points.
type Vector3 [3]int

func (v Vector3) Mul(rot Matrix3x3) Vector3 {
	return Vector3{
		v[0]*rot[0][0] + v[1]*rot[1][0] + v[2]*rot[2][0],
		v[0]*rot[0][1] + v[1]*rot[1][1] + v[2]*rot[2][1],
		v[0]*rot[0][2] + v[1]*rot[1][2] + v[2]*rot[2][2],
	}
}

func (v Vector3) Eq(w Vector3) bool {
	return v[0] == w[0] && v[1] == w[1] && v[2] == w[2]
}

type Matrix3x3 [3][3]int

func Ident() Matrix3x3 {
	return ident.Clone()
}

func (m Matrix3x3) Transpose() Matrix3x3 {
	rc := Matrix3x3{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rc[i][j] = m[j][i]
		}
	}
	return rc
}

func (m Matrix3x3) Clone() Matrix3x3 {
	rc := Matrix3x3{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rc[i][j] = m[i][j]
		}
	}
	return rc
}

func (m Matrix3x3) Mul(n Matrix3x3) Matrix3x3 {
	rc := Matrix3x3{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rc[i][j] = m.dot(n, i, j)
		}
	}
	return rc
}

func (m Matrix3x3) Pow(n int) Matrix3x3 {
	rc := ident
	for i := 0; i < n; i++ {
		rc = rc.Mul(m)
	}
	return rc
}

func (m Matrix3x3) dot(n Matrix3x3, i, j int) int {
	var rc int
	for k := 0; k < 3; k++ {
		rc += m[i][k] * n[k][j]
	}
	return rc
}

func AllRotations() []Matrix3x3 {
	if allRotations == nil {
		allRotations = calcAllRotations()
	}
	return allRotations
}

func Rotation(i int) Matrix3x3 {
	return AllRotations()[i]
}

var (
	allRotations []Matrix3x3
)

func calcAllRotations() []Matrix3x3 {
	zRot := Matrix3x3{ // rotate around z axis
		{0, -1, 0},
		{1, 0, 0},
		{0, 0, 1},
	}
	yRot := Matrix3x3{ //rotate around y axis
		{0, 0, -1},
		{0, 1, 0},
		{1, 0, 0},
	}
	xRot := Matrix3x3{ // rotate around x axis
		{1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	}
	rotSet := set.Set[Matrix3x3]{}
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			for z := 0; z < 4; z++ {
				val := xRot.Pow(x).Mul(yRot.Pow(y)).Mul(zRot.Pow(z))
				rotSet.Add(val)
			}
		}
	}
	return append([]Matrix3x3{Ident()}, rotSet.AsSlice()...)
}
