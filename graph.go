package dsa

import (
	"cmp"
	"container/heap"
	"math"
)

type GraphNode[T cmp.Ordered] struct {
	Val       T
	Neighbors []*GraphNode[T]
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type Edge[T Number] struct {
	Node   string
	Weight T
}

type Item struct {
	Node     string
	priority int
	Index    int
}

func (item Item) Priority() int {
	return item.priority
}


func Dijkstra(graph map[string][]Edge[int], start string) (distances map[string]int) {
	distances = make(map[string]int)
	for node := range graph {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0

	pq := NewPriorityQueue[int]()
	heap.Init(pq)
	heap.Push(pq, Item{Node: start, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Item)
		currentNode := current.Node
		currentDistance := current.priority

		if currentDistance > distances[currentNode] {
			continue
		}

		for _, edge := range graph[currentNode] {
			distance := currentDistance + edge.Weight
			if distance < distances[edge.Node] {
				distances[edge.Node] = distance
				heap.Push(pq, Item{Node: edge.Node, priority: distance})
			}
		}
	}

	return 
}

//https://leetcode.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }

    m, n := len(grid), len(grid[0])
    total := 0

    // Define the DFS function inline
    var dfs func(int, int)
    dfs = func(i, j int) {
        // Check boundaries and if the current cell is part of an island
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
            return
        }

        // Mark the cell as visited
        grid[i][j] = '0'

        // Explore all four directions
        dfs(i-1, j) // Up
        dfs(i+1, j) // Down
        dfs(i, j-1) // Left
        dfs(i, j+1) // Right
    }

    // Iterate through the grid to find islands
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                total++
                dfs(i, j)
            }
        }
    }

    return total
}

// https://leetcode.com/problems/clone-graph/
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
