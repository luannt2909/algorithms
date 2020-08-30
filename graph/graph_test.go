package graph

import "testing"

func TestGraph(t *testing.T) {
	g := &Graph{Edges: make([][]Edge,0, 5)}
	g.AddEdged(0,1,2).
		AddEdged(1,2,3).
		AddEdged(1,3,6).
		AddEdged(1,4,8).
		AddEdged(2,4,9).
		AddEdged(0,3,6).
		AddEdged(3,4,9)
	g.Print()
}