package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap[int]()

	_, err := heap.ExtractMin()
	assert.True(t, heap.IsEmpty())
	assert.Equal(t, err, ErrMinHeapIsEmpty)
	_, err = heap.Peek()
	assert.Equal(t, err, ErrMinHeapIsEmpty)
	assert.Equal(t, heap.String(),  "MinHeap vazio")

	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(7)
	heap.Insert(1)

	assert.False(t, heap.IsEmpty())

	v, err := heap.ExtractMin()
	assert.NoError(t, err)
	assert.Equal(t, v, 1)
	assert.False(t, heap.IsEmpty())

	v, err = heap.Peek()
	assert.NoError(t, err)
	assert.Equal(t, v, 7)

	size := heap.Size()
	assert.Equal(t, size, 3)

	assert.Equal(t, heap.Parent(3), 1)
	assert.Equal(t, heap.String(),  "└── 7\n    ├── 20\n    └── 10\n")
}
