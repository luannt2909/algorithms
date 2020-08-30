package linklist

import "testing"

func TestSingleLinkList(t *testing.T) {
	ll := &SingleLinkList{}
	ll.InsertTail(1).
		InsertTail(2).
		InsertTail(3).
		InsertTail(-1).
		InsertTail(10).
		InsertHead(100)
	ll.Print()
}