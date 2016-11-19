package main

import (
	"bufio"
	"fmt"
	"github.com/totorovan/alg/rcontr"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	inParallel := runtime.NumCPU()
	runtime.GOMAXPROCS(inParallel)

	graph := prepare()

	numIter := iterations(len(graph.Nodes))
	res := graph.MinCut(numIter, inParallel)
	fmt.Println("Mincut: ", res)
}

func prepare() *rcontr.Graph {
	file, err := os.Open("graph.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	graph := rcontr.Graph{}
	for i := 1; i <= 200; i++ {
		graph.AddNode(rcontr.Node{ID: i})
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
				graph.AddEdge(arrInt[0], arrInt[i])
			}
		}
		bytes, _, err = r.ReadLine()
	}

	return &graph
}

func iterations(n int) int {
	return n * n * int(math.Ceil(math.Log(float64(n))))
}
