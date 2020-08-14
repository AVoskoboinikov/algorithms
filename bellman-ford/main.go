package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	from   int
	to     int
	weight int
}

func newGraph(edges []Edge) *Graph {
	g := Graph{
		edges:       edges,
		edgesToFrom: make(map[int][]int),
		cache:       make(map[int]map[int]int),
		vertices: 	 make([]int, 0),
	}

	indx := map[int]interface{}{}

	for i := 0; i < len(g.edges); i++ {
		if _, ok := g.edgesToFrom[g.edges[i].to]; !ok {
			g.edgesToFrom[g.edges[i].to] = make([]int, 0)
		}

		g.edgesToFrom[g.edges[i].to] = append(g.edgesToFrom[g.edges[i].to], i)

		indx[g.edges[i].to] = nil
		indx[g.edges[i].from] = nil
	}

	for k := range indx {
		g.vertices = append(g.vertices, k)
	}

	return &g
}

type Graph struct {
	vertices	[]int
	edges       []Edge
	edgesToFrom map[int][]int
	cache       map[int]map[int]int
}

func (g *Graph) findShortestPathFromSource(s int) {
	g.cache = make(map[int]map[int]int)

	for i := range g.edges {
		g.cache[i] = map[int]int{s: 0}
	}

	for _, v := range g.vertices {
		if s != v {
			g.cache[0][v] = 999999
		}
	}

	for i := 1; i <= (len(g.edges) - 1); i++ {
		if i%1000 == 0 {
			fmt.Println(fmt.Sprintf("%v out of %v", i, len(g.edges)))
		}

		//t := time.Now()
		for to := range g.edgesToFrom {
			g.cache[i][to] = g.findShortestPathToDestinationByBudget(to, i)
		}
		//fmt.Println(fmt.Sprintf("took: %v", time.Since(t)))
	}
}

func (g *Graph) findShortestPathToDestinationByBudget(d int, b int) int {
	if v1, ok := g.cache[b]; ok {
		if v2, ok := v1[d]; ok {
			return v2
		}
	}

	path := g.findShortestPathToDestinationByBudget(d, b-1)
	for _, e := range g.edgesToFrom[d] {
		newPath := g.findShortestPathToDestinationByBudget(g.edges[e].from, b-1) + g.edges[e].weight
		if newPath < path {
			path = newPath
		}
	}

	g.cache[b][d] = path
	return path
}

func main() {
	g, err := readGraph()
	if err != nil {
		panic("Couldn't read the graph")
	}

	// i - count of edges - 1 if on iteration i=n shortest path is less than on i=n-1 - we have a cycle
	minPath := 99999999

	//g.findShortestPathFromSource(5)

	for i, s := range g.vertices {
		//if i%10 == 0 {
		//}

		fmt.Println("started...")
		g.findShortestPathFromSource(s)
		fmt.Println(fmt.Sprintf("%v out of %v", i, len(g.vertices)))
		for k, v := range g.cache[len(g.cache)-1] {
			if g.cache[len(g.cache)-2][k] != v {
				panic("cycle found")
			}
		}

		for d, v := range g.cache[len(g.cache)-1] {
			if v < minPath && d != s {
				minPath = v
			}
		}
	}

	fmt.Println(minPath)
}

func readGraph() (*Graph, error) {
	file, err := os.Open("/Users/andrii/go/src/github.com/magento-mcom/coursera/bellman-ford/g3.txt")
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	ilr := strings.NewReader(sc.Text())
	innerSc := bufio.NewScanner(ilr)
	innerSc.Split(bufio.ScanWords)

	innerSc.Scan()
	_, err = strconv.Atoi(innerSc.Text()) // numbVert
	if err != nil {
		return nil, err
	}

	innerSc.Scan()
	numbEdges, err := strconv.Atoi(innerSc.Text())
	if err != nil {
		return nil, err
	}

	edges := make([]Edge, numbEdges)
	i := 0

	for sc.Scan() {
		lr := strings.NewReader(sc.Text())
		innerSc := bufio.NewScanner(lr)
		innerSc.Split(bufio.ScanWords)

		innerSc.Scan()
		tail, err := strconv.Atoi(innerSc.Text())
		if err != nil {
			return nil, err
		}

		innerSc.Scan()
		head, err := strconv.Atoi(innerSc.Text())
		if err != nil {
			return nil, err
		}

		innerSc.Scan()
		weight, err := strconv.Atoi(innerSc.Text())
		if err != nil {
			return nil, err
		}

		edges[i] = Edge{from: tail, to: head, weight: weight}
		i++
	}

	return newGraph(edges), nil
}
