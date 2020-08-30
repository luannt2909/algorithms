package linklist

import "fmt"

type SingleNode struct {
	Data int
	Next *SingleNode
}

func (n *SingleNode) Insert(data int){
	if n == nil{
		return
	}
	if n.Next == nil{
		n.Next = &SingleNode{Data: data}
		return
	}
	n.Next.Insert(data)
}

func (n *SingleNode) Print() {
	if n == nil{
		return
	}
	fmt.Println("SingleNode: ", n.Data)
	if n.Next != nil{
		fmt.Println("Next SingleNode of ", n.Data)
		n.Next.Print()
	}
}

type SingleLinkList struct {
	Head *SingleNode
}

func (ll *SingleLinkList) InsertTail(data int) *SingleLinkList {
	if ll.Head == nil{
		ll.Head = &SingleNode{Data: data}
		return ll
	}
	ll.Head.Insert(data)
	return ll
}

func (ll *SingleLinkList) InsertHead(data int) *SingleLinkList {
	if ll.Head == nil{
		ll.Head = &SingleNode{Data: data}
		return ll
	}
	temp := ll.Head
	ll.Head = &SingleNode{Data: data, Next:  temp}
	return ll
}

func (ll *SingleLinkList) Print()  {
	fmt.Println("Single link list: ")
	ll.Head.Print()
}
