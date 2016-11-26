package graph

import (
	"reflect"
	"sort"
	"testing"
)

func TestGraph_TopologicalSort(t *testing.T) {
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node3 := &Node{ID: 3}
	node4 := &Node{ID: 4}

	node1.Neighbours = []*Node{node2, node3}
	node2.Neighbours = []*Node{node4}
	node3.Neighbours = []*Node{node4}

	g := &Graph{Nodes: map[int]*Node{
		node1.ID: node1,
		node2.ID: node2,
		node3.ID: node3,
		node4.ID: node4}, IsDirected: true}

	expected := make(map[int]int)
	expected[1] = 1
	expected[2] = 3
	expected[3] = 2
	expected[4] = 4

	f := g.TopologicalSort()

	if !reflect.DeepEqual(f, expected) {
		t.Error()
	}
}

func TestGraph_DFSLoop2(t *testing.T) {
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node3 := &Node{ID: 3}
	node4 := &Node{ID: 4}
	node5 := &Node{ID: 5}

	node1.Neighbours = []*Node{node3}
	node2.Neighbours = []*Node{node1}
	node3.Neighbours = []*Node{node2}
	node4.Neighbours = []*Node{node2, node5}
	node5.Neighbours = []*Node{node4}

	g := &Graph{Nodes: map[int]*Node{
		node1.ID: node1,
		node2.ID: node2,
		node3.ID: node3,
		node4.ID: node4,
		node5.ID: node5}, IsDirected: true}
	time := []int{3, 5, 4, 2, 1}

	expected := map[int][]int{
		1: []int{1, 2, 3},
		4: []int{4, 5},
	}

	result := g.DFSLoop2(time)

	sort.Ints(result[1])
	sort.Ints(result[4])

	if !reflect.DeepEqual(result, expected) {
		t.Error()
	}
}

func TestGraph_SCC(t *testing.T) {
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node3 := &Node{ID: 3}
	node4 := &Node{ID: 4}
	node5 := &Node{ID: 5}

	node1.Neighbours = []*Node{node3}
	node2.Neighbours = []*Node{node1}
	node3.Neighbours = []*Node{node2}
	node4.Neighbours = []*Node{node2, node5}
	node5.Neighbours = []*Node{node4}

	g := &Graph{Nodes: map[int]*Node{
		node1.ID: node1,
		node2.ID: node2,
		node3.ID: node3,
		node4.ID: node4,
		node5.ID: node5}, IsDirected: true}

	expected := []Component{{1, 2, 3}, {4, 5}}

	result := g.SCC()
	sort.Ints(result[0])
	sort.Ints(result[1])

	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected: ", expected, " got: ", result)
	}
}
