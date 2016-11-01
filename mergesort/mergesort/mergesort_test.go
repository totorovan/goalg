package mergesort

import (
	"errors"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		left           []int
		right          []int
		expectedResult []int
		actualResult   []int
		expectedErr    error
	}{
		{[]int{1, 3}, []int{2, 4}, []int{1, 2, 3, 4}, make([]int, 4), nil},
		{[]int{3}, []int{2, 4}, []int{2, 3, 4}, make([]int, 3), nil},
		{[]int{1, 3}, []int{2}, []int{1, 2, 3}, make([]int, 3), nil},
		{[]int{1}, []int{2}, make([]int, 1), make([]int, 1), errors.New("Size mismath")},
	}

	for _, tc := range testCases {
		actualErr := Merge(tc.left, tc.right, tc.actualResult)

		if !reflect.DeepEqual(tc.expectedErr, actualErr) {
			t.Errorf("Expected %s but got %s", tc.expectedErr, actualErr)
		}
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
	}

	for _, tc := range testCases {
		MultiThreadedSort(tc.input, nil)
		if !reflect.DeepEqual(tc.input, tc.expected) {
			t.Errorf("Expected %s but got %s", tc.expected, tc.input)
		}
	}
}
