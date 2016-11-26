package graph

import (
	"reflect"
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

	g := &Graph{Nodes: []*Node{node1, node2, node3, node4}, IsDirected: true}

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

func TestGraph_DFSLoop1(t *testing.T) {
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node3 := &Node{ID: 3}
	node4 := &Node{ID: 4}
	node5 := &Node{ID: 5}

	node1.Neighbours = []*Node{node2}
	node2.Neighbours = []*Node{node3, node4}
	node3.Neighbours = []*Node{node1}
	node4.Neighbours = []*Node{node5}
	node5.Neighbours = []*Node{node4}

	g := &Graph{Nodes: []*Node{node1, node2, node3, node4, node5}, IsDirected: true}

	expected := []int{3, 5, 4, 2, 1}

	f := g.DFSLoop1()

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

	g := &Graph{Nodes: []*Node{node1, node2, node3, node4, node5}, IsDirected: true}
	time := []int{3, 5, 4, 2, 1}

	expected := make(map[int][]int)
	expected[1] = []int{1, 3, 2}
	expected[4] = []int{4, 5}

	result := g.DFSLoop2(time)

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

	g := &Graph{Nodes: []*Node{node1, node2, node3, node4, node5}, IsDirected: true}

	expected := []Component{{1, 3, 2}, {4, 5}}

	result := g.SCC()

	if !reflect.DeepEqual(result, expected) {
		t.Error()
	}
}
