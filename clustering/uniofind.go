package main

import "fmt"

type UnionFind struct {
	clustersCount 	int
	nodes 			[]int
}

func (uf *UnionFind) GetParent(n int) int {
	fmt.Printf(".")

	if uf.nodes[n] != n {
		uf.nodes[n] = uf.GetParent(uf.nodes[n])
		return uf.nodes[n]
	}

	return n
}

func (uf *UnionFind) Merge(n1 int, n2 int) {
	p1 := uf.GetParent(n1)
	p2 := uf.GetParent(n2)

	if p1 != p2 {
		uf.nodes[p1] = p2
		uf.clustersCount--
	}
}