package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"log"
	"sync/atomic"
	"sync"
)

func main() {
	gen, err := numStream()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to build generator: %s", err))
	}

	hash := map[int]interface{}{}
	for num := range gen {
		hash[num] = nil
	}

	var cnt int32
	var wg sync.WaitGroup

	for t := -10000; t <= 10000; t++ {
		fmt.Println(fmt.Sprintf("calculating for %v", t))

		wg.Add(1)
		go func() {
			exist := compute2Sum(t, hash)
			if exist == 1 {
				atomic.AddInt32(&cnt, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(cnt)
}

func compute2Sum(t int, h map[int]interface{}) int {
	//cnt := 0
	for v1 := range h {
		v2 := t - v1

		if v1 == v2 {
			continue
		}

		if _, ok := h[v2]; ok {
			//cnt++
			return 1
		}
	}

	return 0
}

func numStream() (<-chan int, error) {
	file, err := os.Open("/Users/andrii/go/src/github.com/magento-mcom/coursera/2sum/2sum.txt")
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	ch := make(chan int)

	go func() {
		for sc.Scan() {
			val, err := strconv.Atoi(sc.Text())
			if err != nil {
				continue
			}

			ch <- val
		}

		close(ch)
	}()

	return ch, nil
}
