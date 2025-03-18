package dsa

import "cmp"

type Ordered[T cmp.Ordered] interface {
	Priority() T
}

type PriorityQueue[T cmp.Ordered] []Ordered[T]

func NewPriorityQueue[T cmp.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

func (pq *PriorityQueue[T]) Len() int { return len(*pq) }

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].Priority() < (*pq)[j].Priority()
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	*pq = append(*pq, x.(Ordered[T]))
}

func (pq *PriorityQueue[T]) Pop() any {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}
