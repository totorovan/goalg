package mergesort

import (
	"errors"
	"sync"
)

func Merge(left []int, right []int, to []int) error {
	if len(left)+len(right) != len(to) {
		return errors.New("Size mismath")
	}

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
				copy(to[k+1:], rightCopy[j:])
				break
			}
		} else {
			to[k] = rightCopy[j]
			j++
			if j == len(rightCopy) {
				copy(to[k+1:], leftCopy[i:])
				break
			}
		}
	}

	return nil
}

func Sort(arr []int) {
	if len(arr) == 0 || len(arr) == 1 {
		return
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			swap(0, 1, arr)
		}
		return
	}
	Sort(arr[:len(arr)/2])
	Sort(arr[len(arr)/2:])
	Merge(arr[:len(arr)/2], arr[len(arr)/2:], arr)
}

func swap(i int, j int, arr []int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func MultiThreadedSort(arr []int, wg *sync.WaitGroup, multiThreaded bool) {
	if wg != nil {
		defer wg.Done()
	}

	if len(arr) == 0 || len(arr) == 1 {
		return
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			swap(0, 1, arr)
		}
		return
	}

	if multiThreaded {
		var localWG sync.WaitGroup
		localWG.Add(2)

		go MultiThreadedSort(arr[:len(arr)/2], &localWG, false)
		go MultiThreadedSort(arr[len(arr)/2:], &localWG, false)

		localWG.Wait()
	} else {
		MultiThreadedSort(arr[:len(arr)/2], nil, false)
		MultiThreadedSort(arr[len(arr)/2:], nil, false)
	}

	Merge(arr[:len(arr)/2], arr[len(arr)/2:], arr)
}
