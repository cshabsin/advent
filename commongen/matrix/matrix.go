package matrix

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
