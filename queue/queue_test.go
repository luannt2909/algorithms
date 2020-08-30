package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := &Queue{}
	queue.Enqueue(1).
		Enqueue(2).
		Enqueue(3)
	for !queue.Empty() {
		fmt.Print(queue.Dequeue()," - ")
	}
	fmt.Println()
}
