package dsa

import (
	"cmp"
	"errors"
	"fmt"
)

type MinHeap[T cmp.Ordered] struct {
	items []T
}

func NewMinHeap[T cmp.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{items: make([]T, 0)}
}

func (h *MinHeap[T]) Len() int {
	return len(h.items)
}

func (h *MinHeap[T]) Less(i, j int) bool {
	return h.items[i] < h.items[j]
}

func (h *MinHeap[T]) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *MinHeap[T]) Parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap[T]) Left(i int) int {
	return 2*i + 1
}

func (h *MinHeap[T]) Right(i int) int {
	return 2*i + 2
}

func (h *MinHeap[T]) HeapfyUp(i int) {
	for h.items[i] < h.items[h.Parent(i)] {
		h.Swap(i, h.Parent(i))
		i = h.Parent(i)
	}
}

func (h *MinHeap[T]) HeapfyDown(i int) {
	smallest := i
	left := h.Left(i)
	right := h.Right(i)
	if left < h.Size() && h.items[left] < h.items[smallest] {
		smallest = left
	}
	if right < h.Size() && h.items[right] < h.items[smallest] {
		smallest = right
	}
	if smallest != i {
		h.Swap(i, smallest)
		h.HeapfyDown(smallest)
	}
}

func (h *MinHeap[T]) Insert(value T) {
	h.items = append(h.items, value)
	h.HeapfyUp(h.Size() - 1)
}

var (
	ErrMinHeapIsEmpty = errors.New("min heap is empty")
)

func (h *MinHeap[T]) ExtractMin() (T, error) {
	if h.Size() == 0 {
		var zero T
		return zero, ErrMinHeapIsEmpty
	}
	min := h.items[0]
	h.items[0] = h.items[h.Size()-1]
	h.items = h.items[:h.Size()-1]
	h.HeapfyDown(0)
	return min, nil
}

func (h *MinHeap[T]) Peek() (T, error) {
	if h.Size() == 0 {
		var zero T
		return zero, ErrMinHeapIsEmpty
	}
	return h.items[0], nil
}

func (h *MinHeap[T]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *MinHeap[T]) Size() int {
	return h.Len()
}

func (h *MinHeap[T]) String() string {
	if h.IsEmpty() {
		return "MinHeap vazio"
	}
	return h.treeString(0, "", true)
}

func (h *MinHeap[T]) treeString(index int, prefix string, isTail bool) string {
	if index >= h.Size() {
		return ""
	}

	connector := "├── "
	if isTail {
		connector = "└── "
	}

	result := prefix + connector + fmt.Sprintf("%v\n", h.items[index])
	newPrefix := prefix
	if isTail {
		newPrefix += "    "
	} else {
		newPrefix += "│   "
	}

	left := h.Left(index)
	right := h.Right(index)

	children := []int{}
	if left < h.Size() {
		children = append(children, left)
	}
	if right < h.Size() {
		children = append(children, right)
	}

	for i, childIdx := range children {
		isLast := i == len(children)-1
		result += h.treeString(childIdx, newPrefix, isLast)
	}

	return result
}

