package rcontr

import (
	"reflect"
	"testing"
)

func TestAddNode(t *testing.T) {
	gr := Graph{}
	node := Node{ID: 1}

	gr.AddNode(node)

	if len(gr.Nodes) != 1 || !reflect.DeepEqual(gr.Nodes[0], node) {
		t.Error("Node with ID 1 doesn't exist")
	}
	gr.AddNode(Node{ID: 1})
	if len(gr.Nodes) != 1 {
		t.Error("Added dublicated node")
	}
}

func TestRemoveNode(t *testing.T) {
	gr := Graph{}
	gr.Nodes = []Node{{1, []int{2}}, {2, []int{1}}}

	gr.Nodes = gr.RemoveNode(1)

	if len(gr.Nodes) != 1 || len(gr.Nodes[0].NeighboursIDs) != 0 {
		t.Error("Node with ID 1 hasn't been deleted")
	}
}

func TestNode(t *testing.T) {
	gr := Graph{}
	gr.Nodes = []Node{{ID: 1}}

	if node, i := gr.Node(1); node == nil || node.ID != 1 || i != 0 {
		t.Error("Node hasn't found")
	}
	if node, _ := gr.Node(2); node != nil {
		t.Error("Node with ID 2 shouldn't exist")
	}
}

func TestAddEdge(t *testing.T) {
	gr := Graph{}
	gr.Nodes = []Node{{1, []int{2}}, {2, []int{1, 3}}, {3, []int{2}}}

	gr.AddEdge(1, 3)

	node1 := gr.Nodes[0]
	node3 := gr.Nodes[2]
	if len(node1.NeighboursIDs) != 2 || node1.NeighboursIDs[1] != 3 || len(node3.NeighboursIDs) != 2 || node3.NeighboursIDs[1] != 1 {
		t.Error("Edge from 1 to 3 hasn't been added")
	}
}

func TestRemoveEdge(t *testing.T) {
	gr := Graph{}
	gr.Nodes = []Node{{1, []int{1, 2, 3}}, {2, []int{1, 3}}, {3, []int{1, 2}}}

	gr.removeEdges(1, 2)

	node1 := gr.Nodes[0]
	node2 := gr.Nodes[1]
	if len(node1.NeighboursIDs) != 2 || len(node2.NeighboursIDs) != 1 || node2.NeighboursIDs[0] != 3 {
		t.Error("Edge from 1 to 2 hasn't been removed")
	}
}

func TestRemoveLoops(t *testing.T) {
	gr := Graph{}
	gr.Nodes = []Node{{1, []int{1, 1}}}

	gr.removeLoops(1)

	node := gr.Nodes[0]
	if len(node.NeighboursIDs) != 0 {
		t.Error("Loop hasn't been removed")
	}
}

func TestMergeNodes(t *testing.T) {
	gr := Graph{}
	gr.Nodes = []Node{{1, []int{1, 2}}, {2, []int{1, 3}}, {3, []int{2}}}

	gr.Nodes = gr.mergeNodes(1, 2)

	if len(gr.Nodes) != 2 {
		t.Error("Node 2 hasn't been removed")
	}

	node1 := gr.Nodes[0]
	node3 := gr.Nodes[1]

	if !reflect.DeepEqual(node1.NeighboursIDs, []int{1, 1, 3}) || !reflect.DeepEqual(node3.NeighboursIDs, []int{1}) {
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
