package mergesort

import (
	"errors"
	"sync"
)

func Merge(left []int, right []int, to []int) error {
	if len(left)+len(right) != len(to) {
		return errors.New("Size mismath")
	}

	var i, j int
	for k := range to {
		if left[i] <= right[j] {
			to[k] = left[i]
			i++
			if i == len(left) {
				copy(to[k+1:], right[j:])
				break
			}
		} else {
			to[k] = right[j]
			j++
			if j == len(right) {
				copy(to[k+1:], left[i:])
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

func MultiThreadedSort(arr []int, wg *sync.WaitGroup) {
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

	var localWG sync.WaitGroup
	localWG.Add(2)

	go MultiThreadedSort(arr[:len(arr)/2], &localWG)
	go MultiThreadedSort(arr[len(arr)/2:], &localWG)

	localWG.Wait()

	Merge(arr[:len(arr)/2], arr[len(arr)/2:], arr)
}
