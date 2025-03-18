package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestCloneGraph(t *testing.T) {
	t.Run("single node", func(t *testing.T) {
		node := &GraphNode[int]{Val: 1}
		assert.NotNil(t, node)
		nodeClone := cloneGraph(node)
		assert.NotNil(t, nodeClone)
		assert.NotEqual(t, node, nodeClone)
		assert.Equal(t, node.Val, nodeClone.Val)
	})

	t.Run("multiple nodes", func(t *testing.T) {
		node1 := &GraphNode[int]{Val: 1}
		node2 := &GraphNode[int]{Val: 2}
		node3 := &GraphNode[int]{Val: 3}

		node1.Neighbors = []*GraphNode[int]{node2}
		node2.Neighbors = []*GraphNode[int]{node1, node3}
		node3.Neighbors = []*GraphNode[int]{node2}

		clonedGraph := cloneGraph(node1)

		assert.NotNil(t, clonedGraph)
		assert.Equal(t, clonedGraph.Val, node1.Val)
		assert.Equal(t, len(clonedGraph.Neighbors), len(node1.Neighbors))

		for i, neighbor := range clonedGraph.Neighbors {
			assert.Equal(t, neighbor.Val, node1.Neighbors[i].Val)
		}
	})
}

func TestDijkstra(t *testing.T) {
	graph := map[string][]Edge[int]{
		"A": {{"B", 1}, {"C", 4}},
		"B": {{"A", 1}, {"C", 2}, {"D", 5}},
		"C": {{"A", 4}, {"B", 2}, {"D", 1}},
		"D": {{"B", 5}, {"C", 1}},
	}

	distances := Dijkstra(graph, "A")
	assert.Equal(t, 0, distances["A"])
	assert.Equal(t, 1, distances["B"])
	assert.Equal(t, 3, distances["C"])
	assert.Equal(t, 4, distances["D"])
}

func TestNumberOfIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	assert.Equal(t, numIslands(grid), 1)

	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	assert.Equal(t, numIslands(grid), 3)
}
