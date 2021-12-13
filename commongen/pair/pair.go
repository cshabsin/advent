package pair

type Pair[T, U any] struct {
	a T
	b U
}

func (p Pair[T, U]) First() T {
	return p.a
}

func (p Pair[T, U]) Second() U {
	return p.b
}

func Make[T, U any](a T, b U) Pair[T, U] {
	return Pair[T, U]{a, b}
}
