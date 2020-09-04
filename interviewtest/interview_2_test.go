package interviewtest

import (
	"fmt"
	"testing"
)

/**
Split the array elements into strictly increasing and decreasing sequence
Given an array of N elements. The task is to split the elements into two arrays say a1[] and a2[]
such that one contains strictly increasing elements and the other contains strictly decreasing elements and a1.size() + a2.size() = a.size().
If it is not possible to do so, print -1 or else print both the arrays.
Note: There can be multiple answers and the order of elements needs not to be the same in the child arrays.
Example:
Input: a[] = {7, 2, 7, 3, 3, 1, 4}
Output: a1[] = {1, 2, 3, 4, 7} , a2[] = {7, 3}

Input: a[] = {1, 2, 2, 1, 1}
Output: -1
It is not possible
*/

func splitElements(arr []int) ([]int, []int, int) {
	var arr1 []int
	var arr2 []int
	for _, a := range arr {
		if exist(arr1, a) && exist(arr2, a) {
			return nil, nil, -1
		}
		if exist(arr1, a) {
			arr2 = append(arr2, a)
			continue
		}
		arr1 = append(arr1, a)
	}
	quickSort(arr1)
	quickSort(arr2)
	reverseArr(arr2)
	return arr1, arr2, 1
}

func reverseArr(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func exist(arr []int, value int) bool {
	for _, a := range arr {
		if a == value {
			return true
		}
	}
	return false
}

func TestSplitElements(t *testing.T) {
	arr := []int{7, 2, 3, 7, 3, 1, 4}
	arr1, arr2, valid := splitElements(arr)
	fmt.Println(arr1, arr2, valid)
}
