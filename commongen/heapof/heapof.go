package heapof

import "container/heap"

type coster interface {
	Cost() int
}

func Make[T coster](initElems []T) *Heap[T] {
	h := &Heap[T]{elems: initElems}
	heap.Init(h)
	return h
}

type Heap[T coster] struct {
	elems []T
}

func (h *Heap[T]) Len() int {
	return len(h.elems)
}

func (h *Heap[T]) Less(i, j int) bool {
	return h.elems[i].Cost() < h.elems[j].Cost()
}

func (h *Heap[T]) Swap(i, j int) {
	tmp := h.elems[i]
	h.elems[i] = h.elems[j]
	h.elems[j] = tmp
}

func (h *Heap[T]) Push(x interface{}) {
	h.elems = append(h.elems, x.(T))
}

func (h *Heap[T]) PushHeap(t T) {
	heap.Push(h, t)
}

func (h *Heap[T]) PopHeap() T {
	return heap.Pop(h).(T)
}

func (h *Heap[T]) Pop() interface{} {
	tmp := h.elems[h.Len()-1]
	h.elems = h.elems[:h.Len()-1]
	return tmp
}
