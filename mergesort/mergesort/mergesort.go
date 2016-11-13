package mergesort

import (
	"sync"
)

func Merge(left []int, right []int, to []int) {
	leftCopy := make([]int, len(left))
	rightCopy := make([]int, len(right))
	copy(leftCopy, left)
	copy(rightCopy, right)

	var i, j int
	for k := range to {
		if leftCopy[i] <= rightCopy[j] {
			to[k] = leftCopy[i]
			i++
			if i == len(leftCopy) {
				copy(to[k + 1:], rightCopy[j:])
				break
			}
		} else {
			to[k] = rightCopy[j]
			j++
			if j == len(rightCopy) {
				copy(to[k + 1:], leftCopy[i:])
				break
			}
		}
	}
}

func Sort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			swap(0, 1, arr)
		}
		return
	}
	mid := len(arr) / 2
	Sort(arr[:mid])
	Sort(arr[mid:])
	Merge(arr[:mid], arr[mid:], arr)
}

func swap(i int, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func MultiThreadedSort(arr []int, wg *sync.WaitGroup, multiThreaded bool) {
	if wg != nil {
		defer wg.Done()
	}

	if len(arr) <= 1 {
		return
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			swap(0, 1, arr)
		}
		return
	}

	mid := len(arr) / 2

	if multiThreaded {
		var localWG sync.WaitGroup
		localWG.Add(2)

		go MultiThreadedSort(arr[:mid], &localWG, false)
		go MultiThreadedSort(arr[mid:], &localWG, false)

		localWG.Wait()
	} else {
		MultiThreadedSort(arr[:mid], nil, false)
		MultiThreadedSort(arr[mid:], nil, false)
	}

	Merge(arr[:mid], arr[mid:], arr)
}
