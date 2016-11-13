package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/totorovan/alg/quicksort/quicksort"
)

func main() {
	file, err := os.Open("list.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(file)
	buf, _, err := r.ReadLine()
	ul := make([]int, 0, 10000)
	for err == nil {
		n, _ := strconv.Atoi(string(buf))
		ul = append(ul, n)
		buf, _, err = r.ReadLine()
	}
	ul1 := make([]int, 10000)
	ul2 := make([]int, 10000)
	ul3 := make([]int, 10000)
	copy(ul1, ul)
	copy(ul2, ul)
	copy(ul3, ul)
	countFirst := quicksort.CountSwap(ul1, getFirst)
	fmt.Println(countFirst)
	countLast := quicksort.CountSwap(ul2, getLast)
	fmt.Println(countLast)
	countMed := quicksort.CountSwap(ul3, getMedian)
	fmt.Println(countMed)
}

func getFirst(arr []int) int {
	return arr[0]
}

func getLast(arr []int) int {
	quicksort.Swap(0, len(arr)-1, arr)
	return arr[0]
}

func getMedian(arr []int) int {
	first := arr[0]
	last := arr[len(arr)-1]
	midInd := getMInd(arr)
	mid := arr[midInd]
	var ind int
	switch {
	case (first < mid && mid < last) || (last < mid && mid < first):
		ind = midInd
	case (first < last && last < mid) || (mid < last && last < first):
		ind = len(arr) - 1
	default:
		ind = 0
	}
	quicksort.Swap(0, ind, arr)
	return arr[0]
}

func getMInd(arr []int) int {
	l := len(arr)
	if l%2 == 0 {
		return l/2 - 1
	} else {
		return l / 2
	}
}
