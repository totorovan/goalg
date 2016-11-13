package quicksort

func CountSwap(arr []int, getPivot func([]int) int) int {
	if len(arr) <= 1 {
		return 0
	}
	count := len(arr) - 1
	p := getPivot(arr)
	i := 1
	for j := 1; j < len(arr); j++ {
		if arr[j] < p {
			Swap(i, j, arr)
			i++
		}
	}
	Swap(0, i - 1, arr)
	return CountSwap(arr[:i - 1], getPivot) + CountSwap(arr[i:], getPivot) + count
}
