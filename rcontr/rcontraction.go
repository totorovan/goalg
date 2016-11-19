package rcontr

import (
	"fmt"
	"math/rand"
	"time"
)

func (g *Graph) Cut(c chan int) {
	rand.Seed(time.Now().UTC().UnixNano())
	for len(g.Nodes) > 2 {
		i, j := g.randEdge()
		g.Nodes = g.mergeNodes(i, j)
		g.removeLoops(i)
	}
	c <- len(g.Nodes[0].NeighboursIDs)
}

func (g *Graph) MinCut(numIter int, inParallel int) int {
	var minCut = len(g.Nodes) * (len(g.Nodes) - 1)
	logStep := 1000 / inParallel
	for i := 0; i < numIter/inParallel; i++ {
		c := make(chan int)
		for j := 0; j < inParallel; j++ {
			go func() {
				grCopy := g.Copy()
				grCopy.Cut(c)
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

func (gr *Graph) randEdge() (int, int) {
	i := rand.Intn(len(gr.Nodes))
	j := rand.Intn(len(gr.Nodes[i].NeighboursIDs))
	return gr.Nodes[i].ID, gr.Nodes[i].NeighboursIDs[j]
}
