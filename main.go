package main

import (
	"bufio"
	"fmt"
	"github.com/totorovan/alg/graph"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	g := prepare()
	components := g.SCC()
	cL := make([]int, len(components))
	for i, comp := range components {
		cL[i] = len(comp)
	}
	sort.Ints(cL)
	l := len(cL)
	fmt.Println(cL[l-5:])
}

func prepare() *graph.Graph {
	file, err := os.Open("SCC.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	g := graph.NewGraph(true)

	r := bufio.NewReader(file)
	bytes, _, err := r.ReadLine()
	l := 1
	for l < 5105044 {
		line := string(bytes)
		arrStr := strings.Split(line, " ")
		id1, err := strconv.Atoi(arrStr[0])
		if err != nil {
			panic(fmt.Sprint("Error on line %d", l))
		}
		id2, err := strconv.Atoi(arrStr[1])
		if err != nil {
			panic(fmt.Sprint("Error on line %d", l))
		}
		g.AddEdgeByIDs(id1, id2)
		bytes, _, err = r.ReadLine()
		l++
	}

	return g
}

func iterations(n int) int {
	return n * n * int(math.Ceil(math.Log(float64(n))))
}
