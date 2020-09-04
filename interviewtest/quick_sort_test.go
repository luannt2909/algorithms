package interviewtest

import (
	"fmt"
	"testing"
)

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	low, high := 0, len(arr)-1
	pivot := arr[high]
	for i := range arr {
		if arr[i] < pivot {
			arr[low], arr[i] = arr[i], arr[low]
			low++
		}
	}
	arr[low], arr[high] = arr[high], arr[low]
	quickSort(arr[:low])
	quickSort(arr[low+1:])
}

func TestQuickSort(t *testing.T) {
	arr := []int{1, 5, 3, 6, 4, 9, 2}
	//arr := []int{10, 80, 30, 90, 40, 50, 70}
	fmt.Println("array before sorting: ", arr)
	quickSort(arr)
	fmt.Println("array after sorting: ", arr)
}
