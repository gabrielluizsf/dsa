package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestMergeSort(t *testing.T) {
	head := &ListNode{Val: 4}
	head.Add(10)
	head.Add(2)
	head.Add(1)
	head.Add(3)
	head.Add(5)
	head.Add(8)
	head.Add(6)
	head.Add(7)
	head.Add(9)
	expected := "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> nil"
	result := MergeSort(head)
	assert.Equal(t, result.String(), expected)
}
