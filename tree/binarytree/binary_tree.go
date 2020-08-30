package binarytree

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

func (n *Node) Insert(data int) {
	if n == nil {
		return
	}
	if data <= n.Data {
		if n.Left != nil {
			n.Left.Insert(data)
		} else {
			n.Left = &Node{Data: data}
		}
	} else {
		if n.Right != nil {
			n.Right.Insert(data)
		} else {
			n.Right = &Node{Data: data}
		}
	}
}

type BinaryTree struct {
	Root *Node
}

func (t *BinaryTree) Insert(data int) *BinaryTree {
	if t.Root == nil {
		t.Root = &Node{Data: data}
	} else {
		t.Root.Insert(data)
	}
	return t
}
func (n *Node) Print() {
	if n == nil {
		return
	}
	fmt.Println("Node", n.Data)
	if n.Left != nil {
		fmt.Println("Left Nodes of ", n.Data)
		n.Left.Print()
	}
	if n.Right != nil {
		fmt.Println("Right Nodes of ", n.Data)
		n.Right.Print()
	}
}
func (t *BinaryTree) Print() {
	if t.Root == nil {
		return
	}
	fmt.Println("Binary Tree")
	t.Root.Print()
}
