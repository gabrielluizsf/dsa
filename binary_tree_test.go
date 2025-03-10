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
	tree.TraversalVisualization = PreOrder
	assert.Equal(t, tree.String(), "PreOrder Traversal: [10 5 3 7 15 12 18]")
	tree.TraversalVisualization = InOrder
	assert.Equal(t, tree.String(), "InOrder Traversal: [3 5 7 10 12 15 18]")
	tree.TraversalVisualization = PostOrder
	assert.Equal(t, tree.String(), "PostOrder Traversal: [3 7 5 12 18 15 10]")

	treeEmpty := BinaryTree[int]{TraversalVisualization: PostOrder}
	assert.False(t, treeEmpty.Search(10))
	assert.Equal(t, treeEmpty.String(), "PostOrder Traversal: []")

	treeSingle := BinaryTree[int]{TraversalVisualization: InOrder}
	treeSingle.Insert(10)
	assert.Equal(t, treeSingle.String(), "InOrder Traversal: [10]")

}
