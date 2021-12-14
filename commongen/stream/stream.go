package stream

import "constraints"

type Stream[T any] chan T

func (s Stream[T]) Filter(filt func(T) bool) Stream[T] {
	out := make(Stream[T])
	go func() {
		for v := range s {
			if filt(v) {
				out <- v
			}
		}
		close(out)
	}()
	return out
}

func Map[T, U any](s Stream[T], m func(T) U) Stream[U] {
	out := make(Stream[U])
	go func() {
		for v := range s {
			out <- m(v)
		}
		close(out)
	}()
	return out

}

// Methods may not take type parameters.
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#No-parameterized-methods
//
// func (s Stream[T]) Map[U any](m func(T) U) Stream[U] {
// 	return Map(s, m)
// }

func Sum[T constraints.Ordered](s Stream[T]) T {
	var n T
	for v := range s {
		n += v
	}
	return n
}
