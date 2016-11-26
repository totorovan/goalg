package graph

import (
	"errors"
	"fmt"
)

type Graph struct {
	Nodes      []*Node
	IsDirected bool
}

func NewGraph(isDirected bool) *Graph {
	return &Graph{IsDirected: isDirected}
}

func (g *Graph) Reverse() *Graph {
	if !g.IsDirected {
		panic("Undirected graph!")
	}
	rG := &Graph{IsDirected: true}
	nodesMap := make(map[int]*Node)
	for _, node := range g.Nodes {
		n := &Node{ID: node.ID}
		nodesMap[node.ID] = n
		rG.AddNode(n)
	}
	for _, node := range g.Nodes {
		for _, neighbor := range node.Neighbours {
			rG.AddEdge(nodesMap[neighbor.ID], nodesMap[node.ID])
		}
	}
	return rG
}

func (g *Graph) Copy() *Graph {
	nodes := make([]*Node, len(g.Nodes))
	nodesMap := make(map[int]*Node)
	for i, node := range g.Nodes {
		nodes[i] = &Node{ID: node.ID}
		nodesMap[node.ID] = nodes[i]
	}
	for i, node := range g.Nodes {
		nodes[i].Neighbours = make([]*Node, len(node.Neighbours))
		for j, neighbor := range node.Neighbours {
			nodes[i].Neighbours[j] = nodesMap[neighbor.ID]
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
	for _, node := range g.Nodes {
		if node.ID == id {
			return node, nil
		}
	}
	return nil, errors.New("Not found")
}

func (g *Graph) getIndex(node *Node) (int, error) {
	for i := range g.Nodes {
		if g.Nodes[i] == node {
			return i, nil
		}
	}
	return 0, errors.New("Not found")
}

func (g *Graph) AddNode(node *Node) {
	_, err := g.getIndex(node)
	if err == nil {
		panic(fmt.Sprintf("Node with id %d already exists", node.ID))
	}
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) RemoveNode(node *Node) []*Node {
	i, err := g.getIndex(node)
	if err != nil {
		panic(fmt.Sprintf("Node with id %d doesn't exist", node.ID))
	}
	for _, neighbor := range node.Neighbours {
		neighbor.Neighbours = remove(node, neighbor.Neighbours)
	}

	return append(g.Nodes[:i], g.Nodes[i+1:]...)
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

func (g *Graph) MergeNodes(node1, node2 *Node) []*Node {
	if node1 == node2 {
		return g.Nodes
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

	return g.RemoveNode(node2)
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
