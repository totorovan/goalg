package quicksort

import (
	"math/rand"
	"time"
)

func Sort(a []int) {
	if len(a) <= 1 {
		return
	}
	if len(a) == 2 {
		if a[0] > a[1] {
			Swap(0, 1, a)
		}
		return
	}
	p := getPivot(a)
	i, j := 1, 1
	k := len(a)
	for j < k {
		switch {
		case a[j] < p:
			{
				if i != j {
					Swap(i, j, a)
				}
				i++
				j++
			}
		case a[j] == p:
			{
				Swap(j, k-1, a)
				k--
			}
		default:
			j++
		}
	}
	Swap(0, i-1, a)
	Sort(a[:i-1])
	l := i + len(a) - k
	j = len(a) - 1
	for i < k && j >= k {
		Swap(i, j, a)
		i++
		j--
	}
	Sort(a[l:])
}

func getPivot(arr []int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	ind := r.Intn(len(arr))
	Swap(0, ind, arr)
	return arr[0]
}

func Swap(i int, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}
