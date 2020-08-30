package linklist

import "fmt"

type DoubleNode struct {
	Data     int
	Next     *DoubleNode
	Previous *DoubleNode
}

func (n *DoubleNode) Insert(data int) {
	if n == nil {
		return
	}
	if n.Next == nil {
		n.Next = &DoubleNode{Data: data, Previous: n}
		return
	}
	n.Next.Insert(data)
}

func (n *DoubleNode) Print() {
	if n == nil {
		return
	}
	fmt.Println("Node: ", n.Data)
	if n.Next != nil {
		fmt.Println("Next Node of ", n.Data)
		n.Next.Print()
	}
}

type DoubleLinkList struct {
	Head *DoubleNode
}

func (ll *DoubleLinkList) InsertTail(data int) *DoubleLinkList {
	if ll.Head == nil {
		ll.Head = &DoubleNode{Data: data}
		return ll
	}
	ll.Head.Insert(data)
	return ll
}

func (ll *DoubleLinkList) InsertHead(data int) *DoubleLinkList {
	if ll.Head == nil {
		ll.Head = &DoubleNode{Data: data}
		return ll
	}
	temp := ll.Head
	ll.Head = &DoubleNode{Data: data, Next: temp}
	temp.Previous = ll.Head
	return ll
}

func (ll *DoubleLinkList) Print() {
	fmt.Println("Double link list: ")
	ll.Head.Print()
}
