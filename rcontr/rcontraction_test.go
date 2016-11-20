package rcontr

import (
	"reflect"
	"sort"
	"testing"
)

func TestAddNode(t *testing.T) {
	gr := Graph{}
	node := Node{ID: 1}

	gr.AddNode(node)

	if len(gr.nodes) != 1 || !reflect.DeepEqual(gr.nodes[0], node) {
		t.Error("Node with ID 1 doesn't exist")
	}
	gr.AddNode(Node{ID: 1})
	if len(gr.nodes) != 1 {
		t.Error("Added dublicated node")
	}
}

func TestRemoveNode(t *testing.T) {
	gr := Graph{}
	gr.nodes = []Node{{1, []int{2}}, {2, []int{1}}}

	gr.nodes = gr.RemoveNode(1)

	if len(gr.nodes) != 1 || len(gr.nodes[0].neighboursIDs) != 0 {
		t.Error("Node with ID 1 hasn't been deleted")
	}
}

func TestNode(t *testing.T) {
	gr := Graph{}
	gr.nodes = []Node{{ID: 1}}

	if node, i := gr.Node(1); node == nil || node.ID != 1 || i != 0 {
		t.Error("Node hasn't found")
	}
	if node, _ := gr.Node(2); node != nil {
		t.Error("Node with ID 2 shouldn't exist")
	}
}

func TestAddEdge(t *testing.T) {
	gr := Graph{}
	gr.nodes = []Node{{1, []int{2}}, {2, []int{1, 3}}, {3, []int{2}}}

	gr.AddEdge(1, 3)

	node1 := gr.nodes[0]
	node3 := gr.nodes[2]
	if len(node1.neighboursIDs) != 2 || node1.neighboursIDs[1] != 3 || len(node3.neighboursIDs) != 2 || node3.neighboursIDs[1] != 1 {
		t.Error("Edge from 1 to 3 hasn't been added")
	}
}

func TestRemoveEdge(t *testing.T) {
	gr := Graph{}
	gr.nodes = []Node{{1, []int{1, 2, 3}}, {2, []int{1, 3}}, {3, []int{1, 2}}}

	gr.removeEdges(1, 2)

	node1 := gr.nodes[0]
	node2 := gr.nodes[1]
	if len(node1.neighboursIDs) != 2 || len(node2.neighboursIDs) != 1 || node2.neighboursIDs[0] != 3 {
		t.Error("Edge from 1 to 2 hasn't been removed")
	}
}

func TestRemoveLoops(t *testing.T) {
	gr := Graph{}
	gr.nodes = []Node{{1, []int{1, 1}}}

	gr.removeLoops(1)

	node := gr.nodes[0]
	if len(node.neighboursIDs) != 0 {
		t.Error("Loop hasn't been removed")
	}
}

func TestMergeNodes_Undirected(t *testing.T) {
	gr := Graph{}
	gr.nodes = []Node{{1, []int{1, 2}}, {2, []int{1, 3}}, {3, []int{2}}}

	gr.nodes = gr.mergeNodes(1, 2)

	if len(gr.nodes) != 2 {
		t.Error("Node 2 hasn't been removed")
	}

	node1 := gr.nodes[0]
	node3 := gr.nodes[1]

	if !reflect.DeepEqual(node1.neighboursIDs, []int{1, 1, 3}) || !reflect.DeepEqual(node3.neighboursIDs, []int{1}) {
		t.Error("Incorrect edges for IDs 1 and 3")
	}
}

func TestMergeNodes_Directed(t *testing.T) {
	gr := Graph{isDirected: true}
	gr.nodes = []Node{{1, []int{2, 3}}, {2, []int{3}}, {3, []int{}}}

	gr.nodes = gr.mergeNodes(1, 2)

	if len(gr.nodes) != 2 {
		t.Error("Node 2 hasn't been removed")
	}

	node1 := gr.nodes[0]
	node3 := gr.nodes[1]

	sort.Ints(node1.neighboursIDs)
	if !reflect.DeepEqual(node1.neighboursIDs, []int{1, 3, 3}) || !reflect.DeepEqual(node3.neighboursIDs, []int{}) {
		t.Error("Incorrect edges for IDs 1 and 3")
	}
}

func TestRemove(t *testing.T) {
	ids := []int{1, 1, 2, 1, 1}
	ids = remove(1, ids)
	if len(ids) != 1 || ids[0] != 2 {
		t.Error("Id 1 hasn't been removed")
	}
}

func TestGraph_Len(t *testing.T) {
	g := Graph{nodes: []Node{{ID: 1}}}

	if g.Len() != 1 {
		t.Error("Incorrect length of graph")
	}
}

func TestGraph_New(t *testing.T) {
	g1 := NewGraph(true)

	if len(g1.nodes) != 0 || !g1.isDirected {
		t.Error("Incorrect construction of graph")
	}

	g2 := NewGraph(false)

	if len(g2.nodes) != 0 || g2.isDirected {
		t.Error("Incorrect construction of graph")
	}
}

func TestGraph_Copy(t *testing.T) {
	g := &Graph{nodes: []Node{{1, []int{2}}, {2, []int{}}}}

	gCopy := g.Copy()

	if g == gCopy {
		t.Error("Copy must not return the same reference")
	}

	if !reflect.DeepEqual(g.nodes, gCopy.nodes) {
		t.Error("Nodes are not the same")
	}
}

func TestGraph_MinCut(t *testing.T) {
	expected := 2

	gr := Graph{
		nodes: []Node{{1, []int{2, 3, 4}}, {2, []int{1, 3}}, {3, []int{1, 2, 4}}, {4, []int{1, 3}}},
	}

	res := gr.MinCut(32, 4)

	if expected != res {
		t.Errorf("Expected: %d, actual: %d", expected, res)
	}
}
