package graph

import (
	"errors"
	"fmt"
)

type Graph struct {
	Nodes      map[int]*Node
	IsDirected bool
}

func NewGraph(isDirected bool) *Graph {
	g := &Graph{IsDirected: isDirected}
	g.Nodes = make(map[int]*Node)
	return g
}

func (g *Graph) Reverse() *Graph {
	if !g.IsDirected {
		panic("Undirected graph!")
	}
	rG := NewGraph(true)
	for id := range g.Nodes {
		n := &Node{ID: id}
		rG.AddNode(n)
	}
	for _, node := range g.Nodes {
		for _, neighbor := range node.Neighbours {
			rG.AddEdge(rG.Nodes[neighbor.ID], rG.Nodes[node.ID])
		}
	}
	return rG
}

func (g *Graph) Copy() *Graph {
	nodes := make(map[int]*Node)
	for id := range g.Nodes {
		nodes[id] = &Node{ID: id}
	}
	for id, node := range g.Nodes {
		nodes[id].Neighbours = make([]*Node, len(node.Neighbours))
		for j, neighbor := range node.Neighbours {
			nodes[id].Neighbours[j] = nodes[neighbor.ID]
		}
	}
	return &Graph{Nodes: nodes, IsDirected: g.IsDirected}
}

func (g *Graph) Len() int {
	return len(g.Nodes)
}

type Node struct {
	ID         int
	Neighbours []*Node
}

type Component []int

func (g *Graph) Node(id int) (*Node, error) {
	node, ok := g.Nodes[id]
	if ok {
		return node, nil
	}
	return nil, errors.New("Not found")
}

func (g *Graph) AddNode(node *Node) {
	_, ok := g.Nodes[node.ID]
	if ok {
		return
	}
	g.Nodes[node.ID] = node
}

func (g *Graph) RemoveNode(id int) {
	node, ok := g.Nodes[id]
	if !ok {
		panic(fmt.Sprintf("Node with id %d doesn't exist", node.ID))
	}
	for _, neighbor := range node.Neighbours {
		neighbor.Neighbours = remove(node, neighbor.Neighbours)
	}

	delete(g.Nodes, id)
}

func (g *Graph) AddEdge(node1, node2 *Node) {
	if node1 == node2 {
		g.addLoop(node1)
		return
	}

	node1.Neighbours = append(node1.Neighbours, node2)
	if !g.IsDirected {
		node2.Neighbours = append(node2.Neighbours, node1)
	}
}

func (g *Graph) AddEdgeByIDs(id1, id2 int) {
	node1, ok := g.Nodes[id1]
	if !ok {
		node1 = &Node{ID: id1}
		g.AddNode(node1)
	}
	node2, ok := g.Nodes[id2]
	if !ok {
		node2 = &Node{ID: id2}
		g.AddNode(node2)
	}
	g.AddEdge(node1, node2)
}

func (g *Graph) AddNodeByID(id int) {
	_, ok := g.Nodes[id]
	if ok {
		return
	}
	node := &Node{ID: id}
	g.AddNode(node)
}

func (gr *Graph) addLoop(node *Node) {
	node.Neighbours = append(node.Neighbours, node)
}

func (g *Graph) removeEdges(node1, node2 *Node) {
	node1.Neighbours = remove(node2, node1.Neighbours)
	if !g.IsDirected {
		node2.Neighbours = remove(node1, node2.Neighbours)
	}
}

func (gr *Graph) removeLoops(node *Node) {
	node.Neighbours = remove(node, node.Neighbours)
}

func (g *Graph) MergeNodes(node1, node2 *Node) {
	if node1 == node2 {
		return
	}

	for _, neighbor := range node2.Neighbours {
		g.AddEdge(node1, neighbor)
	}

	if g.IsDirected {
		for _, node := range g.Nodes {
			for i, neighbor := range node.Neighbours {
				if neighbor == node2 {
					node.Neighbours[i] = node1
				}
			}
		}
	}

	g.RemoveNode(node2.ID)
}

func remove(node *Node, nodes []*Node) []*Node {
	i := 0
	for i < len(nodes) {
		if nodes[i] == node {
			nodes = append(nodes[:i], nodes[i+1:]...)
			continue
		}
		i++
	}

	return nodes
}
