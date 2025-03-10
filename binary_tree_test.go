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

	t.Run("https://leetcode.com/problems/binary-tree-inorder-traversal/", func (t *testing.T)  {
		tree := &TreeNode[int]{
			Val: 1,
			Right: &TreeNode[int]{
				Val: 2,
				Left: &TreeNode[int]{
					Val: 3,
				},
			},
		}
		result := inorder(tree)
		assert.Equal(t, result, []int{1, 3, 2})
	
		tree = &TreeNode[int]{
			Val: 1,
			Left: &TreeNode[int]{
				Val: 2,
				Left: &TreeNode[int]{
					Val: 4,
				},
				Right: &TreeNode[int]{
					Val: 5,
					Left: &TreeNode[int]{Val: 6},
					Right: &TreeNode[int]{Val: 7},
				},
			},
			Right: &TreeNode[int]{
				Val: 3,
		        Right: &TreeNode[int]{
					Val: 8,
					Left: &TreeNode[int]{Val: 9},
				},	
			},
		}
		result = inorder(tree)
		assert.Equal(t, result, []int{4,2,6,5,7,1,3,9,8})
	})

	t.Run("https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal", func (t *testing.T)  {
		out := buildTree([]int{9,3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
		assert.Equal(t, out.Val , 3)
		tree := BinaryTree[int]{Root: out}
		tree.TraversalVisualization = PreOrder
		assert.Equal(t, tree.String(), "PreOrder Traversal: [3 9 20 15 7]")
	})
}
