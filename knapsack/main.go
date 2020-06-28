package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	value  int
	weight int
}

type Items []Item

func main() {
	size, itms, err := readFile()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to read file: %s", err))
	}

	//fmt.Println(size, itms)

	//values := make(map[string]int, (len(itms)+1)*(size+1))
	values := map[string]int{}
	for x := 0; x < size+1; x++ {
		values[key(0, x)] = 0
	}

	j := 1
	for i := 1; i < len(itms)+1; i++ {
		fmt.Println(fmt.Sprintf("[proc] %v out of %v", j/(((len(itms)+1)*(size+1))/100), 100))

		for x := 0; x < size+1; x++ {
			j++

			if itms[i-1].weight > x {
				values[key(i,x)] = values[key(i-1, x)]
				continue
			}

			values[key(i,x)] = int(
				math.Max(
					float64(values[key(i-1, x)]),
					float64(values[key(i-1, x-itms[i-1].weight)]+itms[i-1].value),
				),
			)
		}
	}

	k := fmt.Sprintf("%v:%v",len(itms),size)
	fmt.Println(values[k])
	//fmt.Println(values)
}

func key(i int, j int) string {
	//return fmt.Sprintf("%v:%v",i,j)
	return fmt.Sprint(i,":",j)
}

func readFile() (int, Items, error) {
	// 2493893
	//file, err := os.Open("/Users/andrii/go/src/github.com/magento-mcom/coursera/knapsack/knapsack.txt")
	file, err := os.Open("/Users/andrii/go/src/github.com/magento-mcom/coursera/knapsack/knapsack_big.txt")
	if err != nil {
		return 0, nil, err
	}

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	sc.Scan()

	lr := strings.NewReader(sc.Text())
	innerSc := bufio.NewScanner(lr)
	innerSc.Split(bufio.ScanWords)

	innerSc.Scan()
	size, err := strconv.Atoi(innerSc.Text())
	if err != nil {
		return 0, nil, err
	}

	innerSc.Scan()
	cnt, err := strconv.Atoi(innerSc.Text())
	if err != nil {
		return 0, nil, err
	}

	itms := make(Items, cnt)

	i := 0
	for sc.Scan() {
		lr := strings.NewReader(sc.Text())
		innerSc := bufio.NewScanner(lr)
		innerSc.Split(bufio.ScanWords)

		innerSc.Scan()
		value, err := strconv.Atoi(innerSc.Text())
		if err != nil {
			return 0, nil, err
		}

		innerSc.Scan()
		weight, err := strconv.Atoi(innerSc.Text())
		if err != nil {
			return 0, nil, err
		}

		itms[i] = Item{value: value, weight: weight}
		i++
	}

	return size, itms, nil
}
