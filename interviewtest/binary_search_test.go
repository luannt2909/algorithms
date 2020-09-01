package interviewtest

import (
	"testing"
)

//find-the-maximum-element-in-an-array-which-is-first-increasing-and-then-decreasing
//Solution: Apply binary search => O(nlog(n))
//Find middle of array, if arr[mid] < arr[mid+1] => continue find half of right array and vice versa
func findMax(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		} else {
			return arr[1]
		}
	}
	mid := len(arr) / 2
	if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
		return arr[mid]
	}
	if arr[mid] < arr[mid+1] {
		return findMax(arr[mid:])
	}

	return findMax(arr[:mid])
}

func TestFindMax(t *testing.T) {
	testcases := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name: "max value on the left of array",
			input:  []int{1, 3, 30, 25, 24, 23, 20, 10},
			output: 30,
		},
		{
			name: "max value on the right of array",
			input:  []int{1, 3, 20, 25, 40, 54, 20, 10},
			output: 54,
		},
		{
			name: "max value is last of array",
			input:  []int{1, 3, 20, 25, 40, 54},
			output: 54,
		},
		{
			name: "max value is first of array",
			input:  []int{60, 54, 40, 32, 15, 10, 1},
			output: 60,
		},
		{
			name: "array is only 1 item",
			input:  []int{60},
			output: 60,
		},
		{
			name: "array is 2 items",
			input:  []int{20, 30},
			output: 30,
		},
		{
			name: "array is 3 items",
			input:  []int{20, 30, 20},
			output: 30,
		},
	}
	for _, tc := range testcases {
		max := findMax(tc.input)
		if max != tc.output {
			t.Fail()
		}
	}
}
