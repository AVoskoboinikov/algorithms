package sort

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	//fmt.Println(basicSort([]int{4, 1, 2, 5, 3, 7, 6, 9, 8}))
	//fmt.Println(mergeSort([]int{4, 1, 2, 5, 3, 7, 6, 9, 8}))
	//fmt.Println(len(readArr())) // 100 000
	//fmt.Println(mergeSort(readArr()))
	fmt.Println("measureBasics")
	measureBasics()

	fmt.Println("measureMerges")
	measureMerges()
	//compareBasicAndMerge()
}

func compareBasicAndMerge()  {
	arr := readArr()
	arr = append(arr, readArr()...)
	arr = append(arr, readArr()...)

	fmt.Println(fmt.Sprintf("total lenght is %v", len(arr)))

	basic := basicSort(arr)
	merge := mergeSort(arr)

	for i := 0; i < len(arr); i++ {
		if basic[i] != merge[i] {
			fmt.Println(fmt.Sprintf("basic %v and rerge %v are not equal", basic[i], merge[i]))
		}
	}

	fmt.Println("done")
}

func measureBasics() {
	arr := readArr()
	arr = append(arr, readArr()...)
	arr = append(arr, readArr()...)

	measureBasic(arr, 100, 10)
	measureBasic(arr, 200, 10)
	measureBasic(arr, 300, 10)
	measureBasic(arr, 400, 10)
	measureBasic(arr, 500, 10)
	measureBasic(arr, 600, 10)
	measureBasic(arr, 700, 10)
	measureBasic(arr, 800, 10)
	measureBasic(arr, 900, 10)
	measureBasic(arr, 1000, 10)
	measureBasic(arr, 1100, 10)
	measureBasic(arr, 1200, 10)
	measureBasic(arr, 1300, 10)
	measureBasic(arr, 1400, 10)
	measureBasic(arr, 1500, 10)
	measureBasic(arr, 1600, 10)
	measureBasic(arr, 1700, 10)
	measureBasic(arr, 1800, 10)
	measureBasic(arr, 1900, 10)
	measureBasic(arr, 2000, 10)
	measureBasic(arr, 2100, 10)
	measureBasic(arr, 2200, 10)
	measureBasic(arr, 2300, 10)
	measureBasic(arr, 2400, 10)
	measureBasic(arr, 2500, 10)
	measureBasic(arr, 2600, 10)
	measureBasic(arr, 2700, 10)
	measureBasic(arr, 2800, 10)
	measureBasic(arr, 2900, 10)
	measureBasic(arr, 3000, 10)
}


func measureMerges() {
	arr := readArr()
	arr = append(arr, readArr()...)
	arr = append(arr, readArr()...)

	measureMerge(arr, 100, 10)
	measureMerge(arr, 200, 10)
	measureMerge(arr, 300, 10)
	measureMerge(arr, 400, 10)
	measureMerge(arr, 500, 10)
	measureMerge(arr, 600, 10)
	measureMerge(arr, 700, 10)
	measureMerge(arr, 800, 10)
	measureMerge(arr, 900, 10)
	measureMerge(arr, 1000, 10)
	measureMerge(arr, 1100, 10)
	measureMerge(arr, 1200, 10)
	measureMerge(arr, 1300, 10)
	measureMerge(arr, 1400, 10)
	measureMerge(arr, 1500, 10)
	measureMerge(arr, 1600, 10)
	measureMerge(arr, 1700, 10)
	measureMerge(arr, 1800, 10)
	measureMerge(arr, 1900, 10)
	measureMerge(arr, 2000, 10)
	measureMerge(arr, 2100, 10)
	measureMerge(arr, 2200, 10)
	measureMerge(arr, 2300, 10)
	measureMerge(arr, 2400, 10)
	measureMerge(arr, 2500, 10)
	measureMerge(arr, 2600, 10)
	measureMerge(arr, 2700, 10)
	measureMerge(arr, 2800, 10)
	measureMerge(arr, 2900, 10)
	measureMerge(arr, 3000, 10)
}

func measureBasic(arr []int, length int, repetions int) {
	start := time.Now()
	for i := 0; i < repetions; i++ {
		_ = basicSort(arr[:length])
	}
	fmt.Println(fmt.Sprintf("Sorting of %v elements takes %v nanoseconds.", length, int(time.Since(start))/repetions))
}

func measureMerge(arr []int, length int, repetions int) {
	start := time.Now()
	for i := 0; i < repetions; i++ {
		_ = mergeSort(arr[:length])
	}
	fmt.Println(fmt.Sprintf("Sorting of %v elements takes %v nanoseconds.", length, int(time.Since(start))/repetions))
}

func BasicSort(arr []int) []int {
	return basicSort(arr)
}

func basicSort(arr []int) []int {
	var tmp int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}

	return arr
}

//func mergeSort(arr []int) []int {
//	divider := len(arr) / 2
//
//	sortedLeft := mergeSort(arr[:divider])
//	sortedRight := mergeSort(arr[divider:])
//	sortedArr := combine(sortedLeft, sortedRight)
//
//	return sortedArr
//}
//
//func combine(left []int, right []int) []int {
//	sortedArr := make([]int, len(left) + len(right))
//	il := 0
//	ir := 0
//
//	for i := range sortedArr {
//		if left[il] > right[ir] {
//			sortedArr[i] = right[ir]
//			ir++
//		} else {
//			sortedArr[i] = left[il]
//			il++
//		}
//	}
//
//	return sortedArr
//}

func MergeSort(arr []int) []int {
	return mergeSort(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	sortedArr := make([]int, len(arr))
	divider := len(arr) / 2

	sortedLeft := mergeSort(arr[:divider])
	sortedRight := mergeSort(arr[divider:])

	il := 0
	ir := 0
	for i := range sortedArr {
		if (il + 1) > len(sortedLeft) {
			sortedArr[i] = sortedRight[ir]
			ir++

			continue
		}

		if (ir + 1) > len(sortedRight) {
			sortedArr[i] = sortedLeft[il]
			il++

			continue
		}

		if sortedLeft[il] > sortedRight[ir] {
			sortedArr[i] = sortedRight[ir]
			ir++
		} else if sortedLeft[il] < sortedRight[ir] {
			sortedArr[i] = sortedLeft[il]
			il++
		} else if sortedLeft[il] == sortedRight[ir] {
			sortedArr[i] = sortedRight[ir]
			il++
		}
	}

	return sortedArr
}

func mergeSortAndInvertions(arr []int) ([]int, int) {
	invertions := 0

	if len(arr) == 1 {
		return arr, invertions
	}

	sortedArr := make([]int, len(arr))
	divider := len(arr) / 2

	sortedLeft, invLeft := mergeSortAndInvertions(arr[:divider])
	sortedRight, invRight := mergeSortAndInvertions(arr[divider:])

	invertions += invLeft
	invertions += invRight

	il := 0
	ir := 0
	for i := range sortedArr {
		if (il + 1) > len(sortedLeft) {
			sortedArr[i] = sortedRight[ir]
			ir++

			continue
		}

		if (ir + 1) > len(sortedRight) {
			sortedArr[i] = sortedLeft[il]
			il++

			continue
		}

		if sortedLeft[il] > sortedRight[ir] {
			sortedArr[i] = sortedRight[ir]
			ir++

			invertions += len(sortedLeft) - il
		} else if sortedLeft[il] < sortedRight[ir] {
			sortedArr[i] = sortedLeft[il]
			il++
		} else if sortedLeft[il] == sortedRight[ir] {
			sortedArr[i] = sortedRight[ir]
			il++
		}
	}

	return sortedArr, invertions
}

func readArr() []int {
	readFile, err := os.Open("integerArray.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var arr []int

	for fileScanner.Scan() {
		el, _ := strconv.Atoi(fileScanner.Text())
		arr = append(arr, el)
	}

	readFile.Close()

	return arr
}
