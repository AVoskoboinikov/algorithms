package main

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
	"strings"
	"github.com/magento-mcom/coursera/sort"
)

type GraphNode struct {
	explored bool
	label int
	edges []int
}

type Graph map[int]*GraphNode

type FinishTime map[int]int

// 434821, 968, 459, 313, 211
func main() {
	g, err := buildGraph()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to build graph: %s", err))
	}

	ft := firstPass(reverseGraph(g))
	fmt.Println(ft)

	ng := relabelGraph(g, ft)

	ldrs := secondPath(ng)
	szs := getSizes(ldrs)

	fmt.Println(sort.MergeSort(szs))
}

func getSizes(ldrs map[int]int) []int {
	a := make([]int, len(ldrs))

	i := 0
	for _, e := range ldrs {
		a[i] = e
		i++
	}

	return a
}

func secondPath(g Graph) map[int]int {
	leaders := make(map[int]int)

	for i := (len(g)); i != 0; i-- {
		if _, ok := g[i]; !ok {
			continue
		}

		if g[i].explored {
			continue
		}

		var s int
		g, s = dfs2(g, i)

		leaders[i] = s
	}

	return leaders
}

func relabelGraph(g Graph, ft FinishTime) Graph {
	ng := make(Graph)

	for _, n := range g {
		nl := ft[n.label]
		nEdges := make([]int, len(n.edges))

		for i, e := range n.edges {
			nEdges[i] = ft[e]
		}

		ng[nl] = &GraphNode{
			explored:false,
			label:nl,
			edges:nEdges,
		}
	}

	return ng
}

func firstPass(g Graph) FinishTime {
	t := 0
	ft := make(FinishTime)

	for i := (len(g)); i != 0; i-- {
		if _, ok := g[i]; !ok {
			continue
		}

		if g[i].explored {
			continue
		}

		g, ft, t = dfs1(g, i, ft, t)
	}

	return ft
}

func dfs2(g Graph, start int) (Graph, int) {
	g[start].explored = true
	size := 1

	for _, e := range g[start].edges {
		if _, ok := g[e]; !ok {
			continue
		}

		if g[e].explored {
			continue
		}

		var s int
		g, s = dfs2(g, e)
		size += s
	}

	return g, size
}


func dfs1(g Graph, start int, ft FinishTime, t int) (Graph, FinishTime, int) {
	g[start].explored = true

	for _, e := range g[start].edges {
		if _, ok := g[e]; !ok {
			continue
		}

		if g[e].explored {
			continue
		}

		g, ft, t = dfs1(g, e, ft, t)
	}

	t++
	ft[start] = t

	return g, ft, t
}

func reverseGraph(g Graph) Graph {
	rg := make(Graph)

	for _, n := range g {
		if _, ok := rg[n.label]; !ok {
			rg[n.label] = &GraphNode{
				explored:false,
				label:n.label,
				edges:[]int{},
			}
		}

		for _, e := range n.edges {
			if node, ok := rg[e]; ok {
				node.edges = append(node.edges, n.label)
			} else {
				rg[e] = &GraphNode{
					explored:false,
					label:e,
					edges:[]int{n.label},
				}
			}
		}
	}

	return rg
}

func buildGraph() (Graph, error) {
	//file, err := os.Open("/Users/andrii/go/src/github.com/magento-mcom/coursera/dfs_scc/dfs_test_3.txt")
	file, err := os.Open("/Users/andrii/go/src/github.com/magento-mcom/coursera/dfs_scc/dfs_scc.txt")
	if err != nil {
		return nil, err
	}

	graph := make(Graph)

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		lr := strings.NewReader(sc.Text())
		innerSc := bufio.NewScanner(lr)
		innerSc.Split(bufio.ScanWords)

		innerSc.Scan()
		lbl, err := strconv.Atoi(innerSc.Text())
		if err != nil {
			return nil, err
		}

		innerSc.Scan()
		edge, err := strconv.Atoi(innerSc.Text())
		if err != nil {
			return nil, err
		}

		if node, ok := graph[lbl]; ok {
			node.edges = append(node.edges, edge)
		} else {
			graph[lbl] = &GraphNode{
				explored:false,
				label:lbl,
				edges:[]int{edge},
			}
		}

		if _, ok := graph[edge]; !ok {
			graph[edge] = &GraphNode{
				explored:false,
				label:edge,
				edges:[]int{},
			}
		}
	}

	return graph, nil
}
