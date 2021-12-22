package set

type Set[T comparable] map[T]bool

func Make[T comparable](vs ...T) Set[T] {
	m := Set[T]{}
	for _, v := range vs {
		m[v] = true
	}
	return m
}

func (s Set[T]) Contains(value T) bool {
	return s[value]
}

func (s Set[T]) Add(value T) {
	s[value] = true
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) Clone() Set[T] {
	clone := Set[T]{}
	for k := range s {
		clone.Add(k)
	}
	return clone
}

func (s Set[T]) AnyValue() *T {
	for k := range s {
		return &k
	}
	return nil
}

func (s Set[T]) AsSlice() []T {
	var slice []T
	for v := range s {
		slice = append(slice, v)
	}
	return slice
}

func Intersect[T comparable](m Set[T], only Set[T]) Set[T] {
	r := Set[T]{}
	for v := range m {
		if only.Contains(v) {
			r.Add(v)
		}
	}
	return r
}
