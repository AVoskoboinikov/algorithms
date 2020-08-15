package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const MaxPath = 999999
const DefaultPath = -42

type ShortestPathsCollection struct {
	sync sync.RWMutex
	collection []int
}

func (sp *ShortestPathsCollection) Add(p int) {
	sp.sync.Lock()
	defer sp.sync.Unlock()

	sp.collection = append(sp.collection, p)
}

func (sp *ShortestPathsCollection) GetMin() int {
	min := MaxPath
	for _, p := range sp.collection {
		if p < min {
			min = p
		}
	}

	return min
}

func (sp *ShortestPathsCollection) Count() int {
	sp.sync.RLock()
	defer sp.sync.RUnlock()

	return len(sp.collection)
}

type Edge struct {
	from   int
	to     int
	weight int
}

func newGraph(edges []Edge) *Graph {
	g := Graph{
		edges:       edges,
		edgesToFrom: make(map[int][]int),
		cache:       make([][]int, 0),
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
	cache       [][]int
}

func (g *Graph) findShortestPathFromSource(s int) {
	g.cache = make([][]int, len(g.edges))
	for i := 0; i <= (len(g.edges) - 1); i++ {
		g.cache[i] = make([]int, len(g.vertices) + 1)
		for j := 1; j <= len(g.vertices); j++ {
			g.cache[i][j] = DefaultPath
		}
		g.cache[i][s] = 0
	}

	for _, v := range g.vertices {
		if s != v {
			g.cache[0][v] = MaxPath
		}
	}

	for i := 1; i <= (len(g.edges) - 1); i++ {
		if i%1000 == 0 {
			//fmt.Println(fmt.Sprintf("%v: %v out of %v", s, i, len(g.edges)))
		}

		//t := time.Now()
		for to := range g.edgesToFrom {
			g.cache[i][to] = g.findShortestPathToDestinationByBudget(to, i)
		}
		//fmt.Println(fmt.Sprintf("took: %v", time.Since(t)))
	}
}

func (g *Graph) findShortestPathToDestinationByBudget(d int, b int) int {
	if v := g.cache[b][d]; v != DefaultPath {
		return v
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


	//g.findShortestPathFromSource(1)
	//fmt.Println(g.cache)
	spc := new(ShortestPathsCollection)
	var wg sync.WaitGroup
	sema := make(chan struct{}, 4)

	for _, s := range g.vertices {
		//if i%10 == 0 {
		//}
		s := s

		wg.Add(1)
		sema <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() {<-sema}()

			g, err := readGraph()
			if err != nil {
				panic(fmt.Sprintf("Couldn't read the graph: %v", err))
			}

			minPath := MaxPath

			//fmt.Println("started...")
			g.findShortestPathFromSource(s)
			//fmt.Println(fmt.Sprintf("%v out of %v", i, len(g.vertices)))
			for k, v := range g.cache[len(g.cache)-1] {
				if g.cache[len(g.cache)-2][k] != v && v != DefaultPath {
					panic("cycle found")
				}
			}

			for d, v := range g.cache[len(g.cache)-1] {
				if d == 0 {
					continue
				}

				if v < minPath && d != s && v != DefaultPath {
					minPath = v
				}
			}

			spc.Add(minPath)
			fmt.Println(spc.Count())

		}()
	}

	wg.Wait()
	fmt.Println(spc.GetMin())
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
