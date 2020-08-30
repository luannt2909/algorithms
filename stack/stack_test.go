package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := &Stack{}
	stack.Push(1).
		Push(2).
		Push(2)
	for !stack.Empty() {
		fmt.Print(stack.Pop()," - ")
	}
	fmt.Println()
}
