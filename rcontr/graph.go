package rcontr

import "fmt"

type Graph struct {
	Nodes []Node
}

func (gr *Graph) Copy() *Graph {
	nodes := make([]Node, len(gr.Nodes))
	for i, node := range gr.Nodes {
		nodes[i] = *node.Copy()
	}
	return &Graph{nodes}
}

type Node struct {
	ID            int
	NeighboursIDs []int
}

func (n *Node) Copy() *Node {
	ids := make([]int, len(n.NeighboursIDs))
	copy(ids, n.NeighboursIDs)
	return &Node{n.ID, ids}
}

func (g *Graph) Node(id int) (*Node, int) {
	for i := range g.Nodes {
		if g.Nodes[i].ID == id {
			return &g.Nodes[i], i
		}
	}
	return nil, 0
}

func (g *Graph) AddNode(node Node) {
	if n, _ := g.Node(node.ID); n != nil {
		return
	}
	g.Nodes = append(g.Nodes, node)
}

func (gr *Graph) RemoveNode(id int) []Node {
	node, i := gr.Node(id)
	if node == nil {
		return gr.Nodes
	}
	for _, neighborID := range node.NeighboursIDs {
		n, _ := gr.Node(neighborID)
		n.NeighboursIDs = remove(id, n.NeighboursIDs)
	}

	return append(gr.Nodes[:i], gr.Nodes[i+1:]...)
}

func (g *Graph) AddEdge(id1, id2 int) {
	if id1 == id2 {
		g.addLoop(id1)
		return
	}

	node1, node2 := g.Edge(id1, id2)
	node1.NeighboursIDs = append(node1.NeighboursIDs, id2)
	node2.NeighboursIDs = append(node2.NeighboursIDs, id1)
}

func (gr *Graph) addLoop(id int) {
	node, _ := gr.Node(id)
	if node == nil {
		panic(fmt.Sprintf("Nodes with ID %s dosn't exist", id))
	}
	node.NeighboursIDs = append(node.NeighboursIDs, id)
}

func (g *Graph) removeEdges(id1, id2 int) {
	node1, node2 := g.Edge(id1, id2)
	node1.NeighboursIDs = remove(id2, node1.NeighboursIDs)
	node2.NeighboursIDs = remove(id1, node2.NeighboursIDs)
}

func (gr *Graph) removeLoops(id int) {
	node, _ := gr.Node(id)
	if node == nil {
		panic(fmt.Sprintf("No node with id %s", id))
	}
	node.NeighboursIDs = remove(id, node.NeighboursIDs)
}

func (g *Graph) mergeNodes(id1, id2 int) []Node {
	if id1 == id2 {
		return g.Nodes
	}

	_, node2 := g.Edge(id1, id2)
	for _, id := range node2.NeighboursIDs {
		g.AddEdge(id1, id)
	}

	return g.RemoveNode(id2)
}

func (g *Graph) Edge(id1, id2 int) (*Node, *Node) {
	node1, _ := g.Node(id1)
	node2, _ := g.Node(id2)
	if node1 == nil || node2 == nil {
		panic(fmt.Sprintf("Nodes with IDs %s or %s don't exist", id1, id2))
	}

	return node1, node2
}

func remove(id int, nodeIDs []int) []int {
	i := 0
	for i < len(nodeIDs) {
		if nodeIDs[i] == id {
			nodeIDs = append(nodeIDs[:i], nodeIDs[i+1:]...)
			continue
		}
		i++
	}

	return nodeIDs
}
