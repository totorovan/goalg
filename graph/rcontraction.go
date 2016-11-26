package graph

import (
	"fmt"
	"math/rand"
	"time"
)

func (g *Graph) MinCut(numIter int, inParallel int) int {
	var minCut = len(g.Nodes) * (len(g.Nodes) - 1)
	logStep := 1000 / inParallel
	for i := 0; i < numIter/inParallel; i++ {
		c := make(chan int)
		for j := 0; j < inParallel; j++ {
			go func() {
				grCopy := g.Copy()
				grCopy.cut(c)
			}()
		}
		for j := 0; j < inParallel; j++ {
			cut := <-c
			if cut < minCut {
				minCut = cut
			}
		}
		if i%logStep == 0 {
			fmt.Printf("Step: %d\\%d. MinCut: %d\n", i*inParallel, numIter, minCut)
		}
	}
	return minCut
}

func (g *Graph) cut(c chan int) {
	rand.Seed(time.Now().UTC().UnixNano())
	for len(g.Nodes) > 2 {
		node1, node2 := g.randEdge()
		g.MergeNodes(node1, node2)
		g.removeLoops(node1)
	}
	c <- len(g.Nodes[0].Neighbours)
}

func (gr *Graph) randEdge() (*Node, *Node) {
	i := rand.Intn(len(gr.Nodes))
	j := rand.Intn(len(gr.Nodes[i].Neighbours))
	return gr.Nodes[i], gr.Nodes[i].Neighbours[j]
}
