package quicksort

import (
	"math/rand"
	"time"
)

func Sort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	p := getPivot(arr)
	i := 1
	for j := 1; j < len(arr); j++ {
		if arr[j] < p {
			if i != j {
				Swap(i, j, arr)
			}
			i++
		}
	}
	Swap(0, i - 1, arr)
	Sort(arr[: i - 1])
	Sort(arr[i:])
}

func getPivot(arr []int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	ind := r1.Intn(len(arr))
	Swap(0, ind, arr)
	return arr[0]
}

func Swap(i int, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}
