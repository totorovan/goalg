package quicksort

import (
	"testing"
	"reflect"
	"math/rand"
)

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

const numberOfElems = 100000000

func BenchmarkSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sort(rand.Perm(numberOfElems))
	}
}
