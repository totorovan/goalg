package graph

import (
	"reflect"
	"sort"
	"testing"
)

func TestGraph_Dist(t *testing.T) {
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node3 := &Node{ID: 3}
	node4 := &Node{ID: 4}
	node5 := &Node{ID: 5}
	node6 := &Node{ID: 6}
	node1.Neighbours = []*Node{node2, node3}
	node2.Neighbours = []*Node{node1, node4, node5}
	node3.Neighbours = []*Node{node1, node5}
	node4.Neighbours = []*Node{node2, node5, node6}
	node5.Neighbours = []*Node{node2, node3, node4, node6}
	node6.Neighbours = []*Node{node4, node5}
	g := &Graph{Nodes: map[int]*Node{
		node1.ID: node1,
		node2.ID: node2,
		node3.ID: node3,
		node4.ID: node4,
		node5.ID: node5,
		node6.ID: node6}}

	dist, err := g.Dist(1, 6)
	if err != nil {
		t.Error("Error happened")
	}
	if dist != 3 {
		t.Errorf("Incorrect dist: %d", dist)
	}

}

func TestGraph_Components(t *testing.T) {
	node1 := &Node{ID: 1}
	node2 := &Node{ID: 2}
	node3 := &Node{ID: 3}
	node4 := &Node{ID: 4}
	node5 := &Node{ID: 5}

	node1.Neighbours = []*Node{node2, node3}
	node2.Neighbours = []*Node{node1}
	node3.Neighbours = []*Node{node1}
	node4.Neighbours = []*Node{node5}
	node5.Neighbours = []*Node{node4}

	g := &Graph{Nodes: map[int]*Node{
		node1.ID: node1,
		node2.ID: node2,
		node3.ID: node3,
		node4.ID: node4,
		node5.ID: node5}}

	components := g.Components()

	if len(components) != 2 {
		t.Errorf("Incorrect number of components: %d", len(components))
	}

	sort.Ints(components[0])
	sort.Ints(components[1])

	if !reflect.DeepEqual(components[0], Component{1, 2, 3}) || !reflect.DeepEqual(components[1], Component{4, 5}) {
		t.Error()
	}
}
