package graph

func (g *Graph) TopologicalSort() map[int]int {
	if !g.IsDirected {
		panic("Not directed grapgh")
	}
	explored := make(map[int]bool)
	f := make(map[int]int)
	label := g.Len()
	for _, node := range g.Nodes {
		if explored[node.ID] {
			continue
		}
		g.topSubRoutine(node, explored, f, &label)
	}
	return f
}

func (g *Graph) topSubRoutine(node *Node, explored map[int]bool, f map[int]int, label *int) {
	explored[node.ID] = true
	for _, node2 := range node.Neighbours {
		if explored[node2.ID] {
			continue
		}
		g.topSubRoutine(node2, explored, f, label)
	}
	f[node.ID] = *label
	*label--
}

func (g *Graph) DFSLoop1() []int {
	a := 0
	t := &a
	time := make([]int, g.Len())
	explored := make(map[int]bool)
	for _, node := range g.Nodes {
		if explored[node.ID] {
			continue
		}
		g.loop1(node, explored, time, t)
	}
	return time
}

func (g *Graph) loop1(node *Node, explored map[int]bool, time []int, t *int) {
	explored[node.ID] = true
	for _, neighbor := range node.Neighbours {
		if explored[neighbor.ID] {
			continue
		}
		g.loop1(neighbor, explored, time, t)
	}
	time[*t] = node.ID
	*t++
}

func (g *Graph) DFSLoop2(time []int) map[int][]int {
	leadersToChildren := make(map[int][]int)
	explored := make(map[int]bool)
	var leader *Node
	for i := len(time) - 1; i >= 0; i-- {
		node := g.Nodes[time[i]]
		if explored[node.ID] {
			continue
		}
		leader = node
		g.loop2(node, leader, explored, leadersToChildren)
	}
	return leadersToChildren
}

func (g *Graph) loop2(node *Node, leader *Node, explored map[int]bool, leadersToChildren map[int][]int) {
	explored[node.ID] = true
	leadersToChildren[leader.ID] = append(leadersToChildren[leader.ID], node.ID)
	for _, neighbor := range node.Neighbours {
		if explored[neighbor.ID] {
			continue
		}
		g.loop2(neighbor, leader, explored, leadersToChildren)
	}
}

func (g *Graph) SCC() []Component {
	if !g.IsDirected {
		panic("Not directed graph!")
	}
	gR := g.Reverse()
	time := gR.DFSLoop1()
	leadersToChildren := g.DFSLoop2(time)

	components := make([]Component, len(leadersToChildren))
	i := 0
	for leader := range leadersToChildren {
		components[i] = leadersToChildren[leader]
		i++
	}

	return components
}
