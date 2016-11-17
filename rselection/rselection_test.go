package rselection

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	testCases := []struct {
		inputArr []int
		pos      int
		expected int
	}{
		{[]int{2, 1, 3, 5, 4}, 1, 2},
		{[]int{2, 1, 3, 5, 4}, 0, 1},
		{[]int{2, 1, 3, 5, 4}, 2, 3},
		{[]int{2, 1, 3, 5, 4}, 3, 4},
		{[]int{2, 1, 3, 5, 4}, 4, 5},
	}

	for _, tc := range testCases {
		act, _ := Select(tc.inputArr, tc.pos)
		if !reflect.DeepEqual(act, tc.expected) {
			t.Errorf("Expected %s but got %s", tc.expected, tc.inputArr)
		}
	}
}
