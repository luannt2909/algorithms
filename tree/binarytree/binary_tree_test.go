package binarytree

import "testing"

func TestBinaryTree(t *testing.T) {
	bt := &BinaryTree{}
	bt.Insert(1).
		Insert(2).
		Insert(3).
		Insert(-1).
		Insert(-2).
		Insert(30).
		Insert(35).
		Insert(27).
		Insert(-10).
		Insert(-15).
		Insert(-5)
	bt.Print()
}
