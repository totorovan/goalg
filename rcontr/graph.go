package rcontr

import "fmt"

type Graph struct {
	nodes      []Node
	isDirected bool
}

func NewGraph(isDirected bool) *Graph {
	return &Graph{isDirected: isDirected}
}

func (g *Graph) Copy() *Graph {
	nodes := make([]Node, len(g.nodes))
	for i, node := range g.nodes {
		nodes[i] = *node.copyNode()
	}
	return &Graph{nodes: nodes, isDirected: g.isDirected}
}

func (g *Graph) Len() int {
	return len(g.nodes)
}

type Node struct {
	ID            int
	neighboursIDs []int
}

func (n *Node) copyNode() *Node {
	ids := make([]int, len(n.neighboursIDs))
	copy(ids, n.neighboursIDs)
	return &Node{n.ID, ids}
}

func (g *Graph) Node(id int) (*Node, int) {
	for i := range g.nodes {
		if g.nodes[i].ID == id {
			return &g.nodes[i], i
		}
	}
	return nil, 0
}

func (g *Graph) AddNode(node Node) {
	if n, _ := g.Node(node.ID); n != nil {
		return
	}
	g.nodes = append(g.nodes, node)
}

func (gr *Graph) RemoveNode(id int) []Node {
	node, i := gr.Node(id)
	if node == nil {
		return gr.nodes
	}
	for _, neighborID := range node.neighboursIDs {
		n, _ := gr.Node(neighborID)
		n.neighboursIDs = remove(id, n.neighboursIDs)
	}

	return append(gr.nodes[:i], gr.nodes[i+1:]...)
}

func (g *Graph) AddEdge(id1, id2 int) {
	if id1 == id2 {
		g.addLoop(id1)
		return
	}

	node1, node2 := g.edge(id1, id2)
	node1.neighboursIDs = append(node1.neighboursIDs, id2)
	if !g.isDirected {
		node2.neighboursIDs = append(node2.neighboursIDs, id1)
	}
}

func (gr *Graph) addLoop(id int) {
	node, _ := gr.Node(id)
	if node == nil {
		panic(fmt.Sprintf("Nodes with ID %s dosn't exist", id))
	}
	node.neighboursIDs = append(node.neighboursIDs, id)
}

func (g *Graph) removeEdges(id1, id2 int) {
	node1, node2 := g.edge(id1, id2)
	node1.neighboursIDs = remove(id2, node1.neighboursIDs)
	if !g.isDirected {
		node2.neighboursIDs = remove(id1, node2.neighboursIDs)
	}
}

func (gr *Graph) removeLoops(id int) {
	node, _ := gr.Node(id)
	if node == nil {
		panic(fmt.Sprintf("No node with id %s", id))
	}
	node.neighboursIDs = remove(id, node.neighboursIDs)
}

func (g *Graph) mergeNodes(id1, id2 int) []Node {
	if id1 == id2 {
		return g.nodes
	}

	_, node2 := g.edge(id1, id2)
	for _, id := range node2.neighboursIDs {
		g.AddEdge(id1, id)
	}

	if g.isDirected {
		for _, node := range g.nodes {
			for i, id := range node.neighboursIDs {
				if id == id2 {
					node.neighboursIDs[i] = id1
				}
			}
		}
	}

	return g.RemoveNode(id2)
}

func (g *Graph) edge(id1, id2 int) (*Node, *Node) {
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
