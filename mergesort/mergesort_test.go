package mergesort

import (
	"math/rand"
	"reflect"
	"runtime"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		left           []int
		right          []int
		expectedResult []int
		actualResult   []int
	}{
		{[]int{1, 3}, []int{2, 4}, []int{1, 2, 3, 4}, make([]int, 4)},
		{[]int{3}, []int{2, 4}, []int{2, 3, 4}, make([]int, 3)},
		{[]int{1, 3}, []int{2}, []int{1, 2, 3}, make([]int, 3)},
	}

	for _, tc := range testCases {
		Merge(tc.left, tc.right, tc.actualResult)

		if !reflect.DeepEqual(tc.actualResult, tc.expectedResult) {
			t.Errorf("Expected %s but got %s", tc.expectedResult, tc.actualResult)
		}
	}
}

func TestSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 1, 3, 5, 4}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 1, 3, 5}, []int{1, 2, 3, 5}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{2, 1, 1, 5, 6, 4, 7, 3, 5}, []int{1, 1, 2, 3, 4, 5, 5, 6, 7}},
	}

	for _, tc := range testCases {
		Sort(tc.input)
		if !reflect.DeepEqual(tc.input, tc.expected) {
			t.Errorf("Expected %s but got %s", tc.expected, tc.input)
		}
	}
}

func TestMultiThreadedSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 1, 3, 5, 4}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 1, 3, 5}, []int{1, 2, 3, 5}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{2, 1, 1, 5, 6, 4, 7, 3, 5}, []int{1, 1, 2, 3, 4, 5, 5, 6, 7}},
	}

	for _, tc := range testCases {
		MultiThreadedSort(tc.input, nil, true)
		if !reflect.DeepEqual(tc.input, tc.expected) {
			t.Errorf("Expected %s but got %s", tc.expected, tc.input)
		}
	}
}

const numberOfElems = 10000000

func BenchmarkSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sort(rand.Perm(numberOfElems))
	}
}

func BenchmarkMultiThreadedSort(b *testing.B) {
	runtime.GOMAXPROCS(4)
	for n := 0; n < b.N; n++ {
		MultiThreadedSort(rand.Perm(numberOfElems), nil, true)
	}
}
