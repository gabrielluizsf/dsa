package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestBinaryTree(t *testing.T) {
	tree := BinaryTree[int]{}
	numbers := []int{10, 5, 15, 3, 7, 12, 18}
	for _, val := range numbers {
		tree.Insert(val)
	}
	assert.True(t, tree.Search(10))
	assert.True(t, tree.Search(5))
	assert.True(t, tree.Search(15))
	assert.False(t, tree.Search(20))
	assert.True(t, tree.Search(18))
	assert.True(t, tree.Search(12))
	assert.True(t, tree.Search(7))
	assert.True(t, tree.Search(3))
	t.Log(tree.String())
}
