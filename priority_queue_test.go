package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

type testItem struct {
	priority float64
}

func (ti testItem) Priority() float64 {
	return ti.priority
}
func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[float64]()
	var item testItem
	pq.Push(item)
	assert.Equal(t, pq.Len(), 1)
	item = testItem{priority: 5.31983313874312}
	pq.Push(item)
	assert.Equal(t, pq.Len(), 2)
	assert.True(t, pq.Less(0, 1))
	assert.Equal(t, pq.Pop().(testItem).Priority(), item.priority)
	assert.Equal(t, pq.Len(), 1)
	assert.Equal(t, pq.Pop().(testItem).Priority(), float64(0))
	assert.Equal(t, pq.Len(), 0)
}
