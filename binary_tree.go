package dsa

import (
	"cmp"
	"fmt"
	"strings"
)

type TreeNode[T cmp.Ordered] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewTreeNode[T cmp.Ordered](val T) *TreeNode[T] {
	return &TreeNode[T]{Val: val}
}

type TraversalVisualization string

const (
	InOrder   TraversalVisualization = "inOrder"
	PostOrder TraversalVisualization = "postOrder"
	PreOrder  TraversalVisualization = "preOrder"
)

type BinaryTree[T cmp.Ordered] struct {
	Root *TreeNode[T]
	TraversalVisualization
}

func (bt *BinaryTree[T]) Insert(val T) {
	newNode := NewTreeNode(val)
	if bt.Root == nil {
		bt.Root = newNode
		return
	}

	node := bt.Root
	for {
		if val < node.Val {
			if node.Left == nil {
				node.Left = newNode
				return
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = newNode
				return
			}
			node = node.Right
		}
	}
}

func (bt *BinaryTree[T]) Search(val T) bool {
	node := bt.Root
	for node != nil {
		if node.Val == val {
			return true
		}
		if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return false
}

func (bt *BinaryTree[T]) String() string {
	switch bt.TraversalVisualization {
	case PreOrder:
		return fmt.Sprintf("PreOrder Traversal: %v", bt.PreOrderTraversal())
	case PostOrder:
		return fmt.Sprintf("PostOrder Traversal: %v", bt.PostOrderTraversal())
	case InOrder:
		return fmt.Sprintf("InOrder Traversal: %v", bt.InOrderTraversal())
	default:
		return bt.treeVisualization()
	}
}

func (bt *BinaryTree[T]) treeVisualization() string {
	if bt.Root == nil {
		return "Empty Tree"
	}

	var sb strings.Builder
	type stackItem struct {
		node   *TreeNode[T]
		prefix string
		isLeft bool
	}

	stack := []stackItem{{bt.Root, "", false}}

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if item.node == nil {
			continue
		}

		sb.WriteString(item.prefix)
		if item.isLeft {
			sb.WriteString("├── ")
		} else {
			sb.WriteString("└── ")
		}
		sb.WriteString(fmt.Sprintf("%v\n", item.node.Val))

		newPrefix := item.prefix + "│   "
		if item.node.Right == nil {
			newPrefix = item.prefix + "    "
		}

		stack = append(stack, stackItem{item.node.Right, newPrefix, false})
		stack = append(stack, stackItem{item.node.Left, newPrefix, true})
	}

	return "\n" + sb.String()
}

func (bt *BinaryTree[T]) PreOrderTraversal() []T {
	var result []T
	var traverse func(*TreeNode[T])
	traverse = func(node *TreeNode[T]) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(bt.Root)
	return result
}

func (bt *BinaryTree[T]) InOrderTraversal() []T {
	var result []T
	var traverse func(*TreeNode[T])
	traverse = func(node *TreeNode[T]) {
		if node == nil {
			return
		}
		traverse(node.Left)
		result = append(result, node.Val)
		traverse(node.Right)
	}
	traverse(bt.Root)
	return result
}

func (bt *BinaryTree[T]) PostOrderTraversal() []T {
	var result []T
	var traverse func(*TreeNode[T])
	traverse = func(node *TreeNode[T]) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		result = append(result, node.Val)
	}
	traverse(bt.Root)
	return result
}

func (bt *BinaryTree[T]) DFS(val T) bool {
	return bt.dfsRecursive(bt.Root, val)
}

func (bt *BinaryTree[T]) dfsRecursive(node *TreeNode[T], val T) bool {
	if node == nil {
		return false
	}
	if node.Val == val {
		return true
	}
	if bt.dfsRecursive(node.Left, val) {
		return true
	}

	if bt.dfsRecursive(node.Right, val) {
		return true
	}
	return false
}

func (bt *BinaryTree[T]) BFS(val T) bool {
	if bt.Root == nil {
		return false
	}
	queue := NewDeque()
	queue.Append(bt.Root)
	for queue.Len() > 0 {
		v, _ := queue.Popleft()
		node, ok := v.(*TreeNode[T])
		if !ok {
			return false
		}
		if node.Val == val {
			return true
		}
		if node.Left != nil {
			queue.Append(node.Left)
		}
		if node.Right != nil {
			queue.Append(node.Right)
		}
	}
	return false
}

//https://leetcode.com/problems/binary-tree-level-order-traversal 
func levelOrder(root *TreeNode[int]) [][]int {
	matrix := make([][]int, 0)
	queue := NewDeque()
	queue.Append(root)

	for queue.Len() > 0 {
		vector := make([]int, 0)
		for range queue.Len() {
			v, ok := queue.Popleft()
			node, _ := v.(*TreeNode[int])
			if ok {
				vector = append(vector, node.Val)
				if node.Left != nil {
					queue.Append(node.Left)
				}
				if node.Right != nil {
					queue.Append(node.Right)
				}
			}
		}
		if len(vector) > 0 {
			matrix = append(matrix, vector)
		}
	}
	return matrix
}

// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
func buildTree(inorder, postorder []int) *TreeNode[int] {
	indexMap := make(map[int]int)
	for i, v := range inorder {
		indexMap[v] = i
	}

	var helper func(int, int, int, int) *TreeNode[int]
	helper = func(inLeft, inRight, postLeft, postRight int) *TreeNode[int] {
		if inLeft > inRight || postLeft > postRight {
			return nil
		}

		rootVal := postorder[postRight]
		root := &TreeNode[int]{Val: rootVal}

		inorderIndex := indexMap[rootVal]
		leftSize := inorderIndex - inLeft

		root.Left = helper(inLeft, inorderIndex-1, postLeft, postLeft+leftSize-1)
		root.Right = helper(inorderIndex+1, inRight, postLeft+leftSize, postRight-1)

		return root
	}

	return helper(0, len(inorder)-1, 0, len(postorder)-1)
}

/**
 * Definition for a binary tree node.
 * https://leetcode.com/problems/binary-tree-inorder-traversal/
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode[int]) []int {
	return inorder(root)
}
func inorder(root *TreeNode[int]) []int {
	if root != nil {
		return append(append(inorder(root.Left), root.Val), inorder(root.Right)...)
	}
	return []int{}
}

// https://leetcode.com/problems/path-sum/
func hasPathSum(root *TreeNode[int], targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
