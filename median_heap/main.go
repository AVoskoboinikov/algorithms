package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main()  {
	gen, err := numStream()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to build generator: %s", err))
	}

	hLow := MaxHeap{}
	hHigh := MinHeap{}
	smm := 0

	i := 0
	for num := range gen {
		i++
		mi := 0
		md := 0

		if len(hLow.objs) == 0 {
			hLow.Insert(num)
		} else if num <= hLow.WhatIsMax() {
			hLow.Insert(num)
		} else {
			hHigh.Insert(num)
		}

		if len(hHigh.objs) - len(hLow.objs) > 1 {
			hLow.Insert(hHigh.GetMin())
		}

		if len(hLow.objs) - len(hHigh.objs) > 1 {
			hHigh.Insert(hLow.GetMax())
		}

		if i%2 == 0 {
			mi = i/2
		}

		if i%2 == 1 {
			mi = (i+1)/2
		}

		if len(hLow.objs) >= mi {
			md = hLow.WhatIsMax()
		} else if len(hHigh.objs) >= mi {
			md = hHigh.WhatIsMin()
		}

		smm += md
	}

	//fmt.Println(hLow.objs)
	//fmt.Println(hHigh.objs)
	fmt.Println(smm%10000)
}

func numStream() (<-chan int, error) {
	file, err := os.Open("/Users/andrii/go/src/github.com/magento-mcom/coursera/median_heap/median.txt")
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
