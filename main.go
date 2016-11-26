package main

import (
	"bufio"
	"fmt"
	"github.com/totorovan/alg/graph"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*inParallel := runtime.NumCPU()
	runtime.GOMAXPROCS(inParallel)

	g := prepare()

	iter := iterations(g.Len())
	res := g.MinCut(iter, 4)
	fmt.Println("Mincut: ", res)*/
	node1 := &graph.Node{ID: 1}
	node2 := &graph.Node{ID: 2}
	node3 := &graph.Node{ID: 3}
	node4 := &graph.Node{ID: 7}
	node5 := &graph.Node{ID: 5}
	node6 := &graph.Node{ID: 6}

	node1.Neighbours = []*graph.Node{node3}
	node2.Neighbours = []*graph.Node{node1, node6}
	node3.Neighbours = []*graph.Node{node2}
	node4.Neighbours = []*graph.Node{node2, node5}
	node5.Neighbours = []*graph.Node{node4}
	node6.Neighbours = []*graph.Node{node1}

	g := &graph.Graph{Nodes: []*graph.Node{node1, node2, node3, node4, node5, node6}, IsDirected: true}
	result := g.SCC()
	fmt.Println(result)
}

func prepare() *graph.Graph {
	file, err := os.Open("graph.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	nodes := make(map[int]*graph.Node)

	g := graph.NewGraph(false)
	for i := 1; i <= 200; i++ {
		node := &graph.Node{ID: i}
		nodes[i] = node
		g.AddNode(node)
	}

	r := bufio.NewReader(file)
	bytes, _, err := r.ReadLine()
	for err == nil {
		line := string(bytes)
		arrStr := strings.Split(line, "\t")
		arrInt := make([]int, len(arrStr))
		for i, str := range arrStr {
			arrInt[i], err = strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
		}
		for i := 1; i < len(arrInt); i++ {
			if arrInt[i] > arrInt[0] {
				g.AddEdge(nodes[arrInt[0]], nodes[arrInt[i]])
			}
		}
		bytes, _, err = r.ReadLine()
	}

	return g
}

func iterations(n int) int {
	return n * n * int(math.Ceil(math.Log(float64(n))))
}
