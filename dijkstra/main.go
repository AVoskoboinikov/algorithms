package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	label int
	edges map[int]int
}

type Graph map[int]Node

type ShortPath map[int]int

func main() {
	g, err := buildGraph()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to build graph: %s", err))
	}

	sp := getShortestPath(1, g)

	// 7,37,59,82,99,115,133,165,188,197
	fmt.Println(fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v", sp[7], sp[37], sp[59], sp[82], sp[99], sp[115], sp[133], sp[165], sp[188], sp[197]))
	//fmt.Println(fmt.Sprintf("%v,%v,%v,%v,%v,%v", sp[1], sp[2], sp[3], sp[4], sp[5], sp[6]))
}

func getShortestPath(source int, graph Graph) ShortPath {
	shp := make(ShortPath)
	shp[source] = 0

	for len(shp) != len(graph) {
		min := 1000000
		minV := 0

		for tl, p := range shp {
			for hd, w := range graph[tl].edges {
				if _, ok := shp[hd]; ok {
					continue
				}

				score := p + w
				if score < min {
					min = score
					minV = hd
				}
			}
		}

		if minV == 0 {
			break
		}

		shp[minV] = min
	}

	for v, _ := range graph {
		if _, ok := shp[v]; !ok {
			shp[v] = 1000000
		}
	}

	return shp
}

func buildGraph() (Graph, error) {
	file, err := os.Open("/Users/andrii/go/src/github.com/magento-mcom/coursera/dijkstra/dijkstraData.txt")
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

		graph[lbl] = Node{label:lbl, edges:map[int]int{}}

		for innerSc.Scan() {
			slice := strings.Split(innerSc.Text(), ",")

			head, err := strconv.Atoi(slice[0])
			if err != nil {
				return nil, err
			}

			weight, err := strconv.Atoi(slice[1])
			if err != nil {
				return nil, err
			}

			graph[lbl].edges[head] = weight
		}
	}

	return graph, nil
}