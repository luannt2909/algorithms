package graph

import "fmt"

type Edge struct {
	From   int
	To     int
	Weight int
}

type Graph struct {
	Edges [][]Edge
}

func (g *Graph) AddEdged(u, v, w int) *Graph {
	g.Edges[u] = append(g.Edges[u], Edge{u, v, w})
	g.Edges[v] = append(g.Edges[v], Edge{v, u, w})
	return g
}

func (g *Graph) Print() {
	for u, edges := range g.Edges {
		fmt.Println("Edged of ", u)
		for _, edge := range edges {
			fmt.Printf("Edge: %v -> %v = %v",edge.From, edge.To, edge.Weight)
		}
	}
}
