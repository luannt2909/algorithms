package interviewtest

import (
	"testing"
)

/**
Remove minimum elements from array such that no three consecutive element are either increasing or decreasing
Example:
Input : arr[] = {5, 2, 3, 6, 1}
Output : 1
Given arr[] is not in required form
(2 < 3 < 6). So, after removal of
6 or 3, array will be in required manner.

Input : arr[] = { 4, 2, 6, 3, 10, 1}
Output : 0
*/

func findMinElements(arr []int) int {
	if len(arr) <= 2 {
		return 0
	}
	var count int
	for i := 1; i < len(arr)-1; i++ {
		//Three consecutive increasing
		consecutiveInc := arr[i-1] < arr[i] && arr[i] < arr[i+1]
		consecutiveDesc := arr[i-1] > arr[i] && arr[i] > arr[i+1]
		if consecutiveDesc || consecutiveInc {
			count++
		}
	}
	return count
}

func TestFindMinElements(t *testing.T) {
	testcases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 3},
			output: 0,
		},
		{
			input:  []int{4, 2, 6, 3, 10, 1},
			output: 0,
		},
		{
			input:  []int{1, 3, 20, 25, 40, 54, 20, 10},
			output: 5,
		},
	}
	for _, tc := range testcases {
		minElements := findMinElements(tc.input)
		if minElements != tc.output {
			t.Fail()
			t.Log("Failed test case with input", tc.input, "output", tc.output, "expected", minElements)
		}
	}
}
