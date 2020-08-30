package linklist

import "testing"

func TestDoubleLinkList(t *testing.T) {
	ll := &DoubleLinkList{}
	ll.InsertTail(1).
		InsertTail(2).
		InsertTail(3).
		InsertTail(-1).
		InsertTail(10)
	ll.Print()
}