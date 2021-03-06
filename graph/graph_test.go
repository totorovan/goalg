package graph

import (
	"reflect"
	"testing"
)

func TestAddNode(t *testing.T) {
	g := NewGraph(false)
	node := &Node{ID: 1}

	g.AddNode(node)

	if len(g.Nodes) != 1 || !reflect.DeepEqual(g.Nodes[1], node) {
		t.Error("Node with ID 1 doesn't exist")
	}
}

func TestRemoveNode(t *testing.T) {
	gr := Graph{}
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node1.Neighbours = []*Node{node2}
	node2.Neighbours = []*Node{node1}
	gr.Nodes = map[int]*Node{
		node1.ID: node1,
		node2.ID: node2}

	gr.RemoveNode(node1.ID)

	if len(gr.Nodes) != 1 || len(gr.Nodes[node2.ID].Neighbours) != 0 {
		t.Error("Node with ID 1 hasn't been deleted")
	}
}

func TestAddEdge(t *testing.T) {
	gr := Graph{}
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node3 := &Node{ID: 3}
	node1.Neighbours = []*Node{node2}
	node2.Neighbours = []*Node{node1, node3}
	node3.Neighbours = []*Node{node2}

	gr.AddEdge(node1, node3)

	if len(node1.Neighbours) != 2 || node1.Neighbours[1] != node3 || len(node3.Neighbours) != 2 || node3.Neighbours[1] != node1 {
		t.Error("Edge from 1 to 3 hasn't been added")
	}
}

func TestRemoveEdge(t *testing.T) {
	gr := Graph{}
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node3 := &Node{ID: 3}
	node1.Neighbours = []*Node{node1, node2, node3}
	node2.Neighbours = []*Node{node1, node3}
	node3.Neighbours = []*Node{node1, node2}

	gr.removeEdges(node1, node2)

	if len(node1.Neighbours) != 2 || len(node2.Neighbours) != 1 || node2.Neighbours[0] != node3 {
		t.Error("Edge from 1 to 2 hasn't been removed")
	}
}

func TestRemoveLoops(t *testing.T) {
	gr := Graph{}
	node1 := &Node{ID: 1}
	node1.Neighbours = []*Node{node1}

	gr.removeLoops(node1)

	if len(node1.Neighbours) != 0 {
		t.Error("Loop hasn't been removed")
	}
}

func TestMergeNodes_Undirected(t *testing.T) {
	gr := Graph{}
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node3 := &Node{ID: 3}
	node1.Neighbours = []*Node{node1, node2}
	node2.Neighbours = []*Node{node1, node3}
	node3.Neighbours = []*Node{node2}
	gr.Nodes = map[int]*Node{
		node1.ID: node1,
		node2.ID: node2,
		node3.ID: node3}

	gr.MergeNodes(node1, node2)

	if len(gr.Nodes) != 2 {
		t.Error("Node 2 hasn't been removed")
	}

	if !reflect.DeepEqual(node1.Neighbours, []*Node{node1, node1, node3}) || !reflect.DeepEqual(node3.Neighbours, []*Node{node1}) {
		t.Error("Incorrect edges for IDs 1 and 3")
	}
}

func TestMergeNodes_Directed(t *testing.T) {
	gr := Graph{IsDirected: true}
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node3 := &Node{ID: 3}
	node1.Neighbours = []*Node{node2, node3}
	node2.Neighbours = []*Node{node3}
	node3.Neighbours = []*Node{}
	gr.Nodes = map[int]*Node{
		node1.ID: node1,
		node2.ID: node2,
		node3.ID: node3}

	gr.MergeNodes(node1, node2)

	if len(gr.Nodes) != 2 {
		t.Error("Node 2 hasn't been removed")
	}
}

func TestRemove(t *testing.T) {
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	nodes := []*Node{node1, node1, node2, node1, node1}
	nodes = remove(node1, nodes)
	if len(nodes) != 1 || nodes[0] != node2 {
		t.Error("Id 1 hasn't been removed")
	}
}

func TestGraph_Len(t *testing.T) {
	g := Graph{Nodes: map[int]*Node{
		1: {ID: 1}}}

	if g.Len() != 1 {
		t.Error("Incorrect length of graph")
	}
}

func TestGraph_New(t *testing.T) {
	g1 := NewGraph(true)

	if len(g1.Nodes) != 0 || !g1.IsDirected {
		t.Error("Incorrect construction of graph")
	}

	g2 := NewGraph(false)

	if len(g2.Nodes) != 0 || g2.IsDirected {
		t.Error("Incorrect construction of graph")
	}
}

func TestGraph_Copy(t *testing.T) {
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	g := &Graph{Nodes: map[int]*Node{
		1: {1, []*Node{node2}},
		2: {2, []*Node{node1}}}}

	gCopy := g.Copy()

	if g == gCopy {
		t.Error("Copy must not return the same reference")
	}

	if g.Nodes[node1.ID] == gCopy.Nodes[node1.ID] || g.Nodes[node2.ID] == gCopy.Nodes[node2.ID] {
		t.Error("The same node pointers")
	}
}

func TestGraph_Reverse(t *testing.T) {
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node1.Neighbours = []*Node{node2}
	g := &Graph{Nodes: map[int]*Node{
		node1.ID: node1,
		node2.ID: node2}, IsDirected: true}

	gR := g.Reverse()

	if !gR.IsDirected {
		t.Error("Graph should be directed")
	}
	if len(gR.Nodes) != 2 {
		t.Error("Incorrect number of nodes")
	}
	if gR.Nodes[node1.ID].ID != 1 {
		t.Error("Incorrect ID of first node")
	}
	if len(gR.Nodes[node1.ID].Neighbours) != 0 {
		t.Error("First node shouldn't have outgoing edges")
	}
	if gR.Nodes[node2.ID].ID != 2 {
		t.Error("Incorrect ID of the second node")
	}
	if len(gR.Nodes[node2.ID].Neighbours) != 1 {
		t.Error("Incorrect number of edges outgoing from 2nd node")
	}
	if gR.Nodes[node2.ID].Neighbours[0] != gR.Nodes[node1.ID] {
		t.Error("2nd node points to incorrect node")
	}
}
