package main

import (
	"bufio"
	"fmt"
	"github.com/magento-mcom/coursera/sort"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
)


type Node struct {
	from 	int
	to	 	int
	weight	int
}

type NodeList []Node
type WeightedNodeList map[int]NodeList
type HammingNode []int
type HammingNodeList []int


func main() {
	//uf := UnionFind{
	//	nodes: []int{0,0,1,2,3,5,5,6,7,8,9},
	//}
	//p := uf.GetParent(0)
	//p = uf.GetParent(1)
	//p = uf.GetParent(3)
	//p = uf.GetParent(3)
	//p = uf.GetParent(2)
	//p = uf.GetParent(9)
	//p = uf.GetParent(8)
	//fmt.Println(p)
	maxSpacingClustering()
	//maxClustersWithSpacing()
}

func maxClustersWithSpacing()  {
	hnl, err := readHammingNodes()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to build nodes: %s", err))
	}

	_ = computeDistances(hnl)
}

func maxSpacingClustering()  {
	nl, uf, err := readNodes()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to build nodes: %s", err))
	}

	wnl := nodeLisToWeighted(nl)

	weights := make([]int, 0)
	for k := range wnl {
		weights = append(weights, k)
	}

	sweights := sort.MergeSort(weights)
	lastWeight := map[int]int{}
	totalClusters := 1

	for _, w := range sweights {
		for _, n := range wnl[w] {
			fmt.Println(uf.clustersCount)

			uf.Merge(n.from, n.to)

			if uf.clustersCount == totalClusters {
				break
			}
		}

		lastWeight[w] = uf.clustersCount
		if uf.clustersCount == totalClusters {
			break
		}
	}

	fmt.Println(lastWeight)
	fmt.Println(uf.clustersCount)
}

func nodeLisToWeighted(nl NodeList) WeightedNodeList {
	wnl := make(WeightedNodeList)

	for _, n := range nl {
		if _, ok := wnl[n.weight]; !ok {
			wnl[n.weight] = NodeList{}
		}

		wnl[n.weight] = append(wnl[n.weight], n)
	}

	return wnl
}

func readNodes() (NodeList, *UnionFind, error) {
	//file, err := os.Open("/Users/andrii/go/src/github.com/magento-mcom/coursera/clustering/clustering1.txt")
	file, err := os.Open("/Users/andrii/go/src/github.com/magento-mcom/coursera/clustering/clustering_big1_res.txt")
	if err != nil {
		return nil, nil, err
	}

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	size, err := strconv.Atoi(sc.Text())

	nl := NodeList{}
	uf := UnionFind{nodes:make([]int, size+1), clustersCount:size}

	for sc.Scan() {
		lr := strings.NewReader(sc.Text())
		innerSc := bufio.NewScanner(lr)
		innerSc.Split(bufio.ScanWords)

		innerSc.Scan()
		node1, err := strconv.Atoi(innerSc.Text())
		if err != nil {
			return nil, nil, err
		}

		innerSc.Scan()
		node2, err := strconv.Atoi(innerSc.Text())
		if err != nil {
			return nil, nil, err
		}

		innerSc.Scan()
		weight, err := strconv.Atoi(innerSc.Text())
		if err != nil {
			return nil, nil, err
		}

		nl = append(nl, Node{from:node1, to:node2, weight:weight})
		uf.nodes[node1] = node1
		uf.nodes[node2] = node2
	}

	return nl, &uf, nil
}

func computeDistances(hnl HammingNodeList) NodeList {
	weightLimit := 2
	hl := NodeList{}

	fmt.Println(200000)

	for i, from := range hnl {
		//fmt.Println(fmt.Sprintf("computing for %v, len: %v", i, len(hl)))
		for j, to := range hnl {
			if i == j {
				continue
			}

			d := computeHammingDistance(from, to)
			if d > weightLimit {
				continue
			}

			fmt.Println(fmt.Sprintf("%v %v %v", i+1, j+1, d))
			//hl = append(hl, Node{from:i+1,to:j+1,weight:d})
			//hl[cnt] = Node{from:i+1,to:j+1,weight:d}
			//cnt++
		}
	}

	return hl
}

func computeHammingDistance(n1 int, n2 int) int {
	return bits.OnesCount(uint(n1 ^ n2))
}

func readHammingNodes() (HammingNodeList, error) {
	file, err := os.Open("/Users/andrii/go/src/github.com/magento-mcom/coursera/clustering/clustering_big1.txt")
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	sc.Scan()

	lr := strings.NewReader(sc.Text())
	innerSc := bufio.NewScanner(lr)
	innerSc.Split(bufio.ScanWords)

	innerSc.Scan()
	size, err := strconv.Atoi(innerSc.Text())

	innerSc.Scan()
	bsize, err := strconv.Atoi(innerSc.Text())

	hnl := make(HammingNodeList, size)
	cnt := 0

	for sc.Scan() {
		lr := strings.NewReader(sc.Text())
		innerSc := bufio.NewScanner(lr)
		innerSc.Split(bufio.ScanWords)

		bts := ""

		for i := 1; i <= bsize; i++ {
			innerSc.Scan()
			bit, err := strconv.Atoi(innerSc.Text())
			if err != nil {
				return nil, err
			}

			bts += fmt.Sprintf("%v", bit)
		}

		parsed, _ := strconv.ParseInt(bts, 2, 64)
		hnl[cnt] = int(parsed)
		cnt++
	}

	return hnl, nil
}