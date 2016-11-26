package graph

import (
	"errors"
	"math"
)

func (g *Graph) Dist(id1, id2 int) (int, error) {
	node1, err := g.Node(id1)
	if err != nil {
		return 0, err
	}
	node2, err := g.Node(id2)
	if err != nil {
		return 0, err
	}

	if id1 == id2 {
		return 0, nil
	}
	flags := make(map[int]bool)
	flags[id1] = true
	q := NewQueue()
	q.Enqueue(node1)
	dist := make(map[int]int)
	for _, node := range g.Nodes {
		dist[node.ID] = math.MaxInt64
	}
	dist[id1] = 0
	for !q.IsEmpty() {
		node, err := q.Dequeue()
		if err != nil {
			return 0, err
		}
		for _, neighbour := range node.Neighbours {
			if neighbour == node2 {
				return dist[node.ID] + 1, nil
			}
			q.Enqueue(neighbour)
			dist[neighbour.ID] = dist[node.ID] + 1
		}
	}
	return 0, errors.New("Not connected")
}

func (g *Graph) Components() []Component {
	explored := make(map[int]bool)
	var components []Component
	for _, node := range g.Nodes {
		if explored[node.ID] {
			continue
		}
		components = append(components, g.Component(node, explored))
	}
	return components
}

func (g *Graph) Component(node *Node, explored map[int]bool) Component {
	q := NewQueue()
	q.Enqueue(node)
	comp := Component{}
	for !q.IsEmpty() {
		node, _ := q.Dequeue()
		comp = append(comp, node.ID)
		explored[node.ID] = true
		for _, neighbour := range node.Neighbours {
			if explored[neighbour.ID] {
				continue
			}
			q.Enqueue(neighbour)
		}
	}

	return comp
}
