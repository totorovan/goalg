package rcontr

import (
	"fmt"
	"math/rand"
	"time"
)

func (g *Graph) MinCut(numIter int, inParallel int) int {
	var minCut = len(g.nodes) * (len(g.nodes) - 1)
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
	for len(g.nodes) > 2 {
		i, j := g.randEdge()
		g.nodes = g.mergeNodes(i, j)
		g.removeLoops(i)
	}
	c <- len(g.nodes[0].neighboursIDs)
}

func (gr *Graph) randEdge() (int, int) {
	i := rand.Intn(len(gr.nodes))
	j := rand.Intn(len(gr.nodes[i].neighboursIDs))
	return gr.nodes[i].ID, gr.nodes[i].neighboursIDs[j]
}
