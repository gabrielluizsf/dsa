package dsa

import (
	"cmp"
)

type GraphNode[T cmp.Ordered] struct {
	Val       T
	Neighbors []*GraphNode[T]
}

//https://leetcode.com/problems/clone-graph/
func cloneGraph(node *GraphNode[int]) *GraphNode[int] {
	if node == nil {
        return nil
    }
    
    queue := NewDeque()
    clones := make(map[int]*GraphNode[int])
    clones[node.Val] = newGraphNode(node.Val)

    queue.Append(node)
    
    for queue.Len() > 0 {
        current, _ := queue.Popleft()
        
        currentNode, ok := current.(*GraphNode[int])
        if !ok {
            return nil 
        }

        currentClone := clones[currentNode.Val]
        
        for _, n := range currentNode.Neighbors {
            if _, ok := clones[n.Val]; !ok {
                clones[n.Val] = newGraphNode(n.Val)
                queue.Append(n)
            }
            
            currentClone.Neighbors = append(currentClone.Neighbors, clones[n.Val])
        }
    }

    return clones[node.Val]
}

func newGraphNode[T cmp.Ordered](val T) *GraphNode[T] {
    return &GraphNode[T]{Val: val, Neighbors: make([]*GraphNode[T], 0)}
}