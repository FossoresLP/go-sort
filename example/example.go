package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"

	localsort "github.com/fossoreslp/go-sort"
)

func main() {
	fmt.Println("Testing insert: ")
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	o := localsort.InsertSorted(s, -1)
	fmt.Println("Result: ", o, " Original: ", s)
	fmt.Println()

	testSort(10)
	fmt.Println()
	testSort(100)
	fmt.Println()
	testSort(1000)
	fmt.Println()
	testSort(10000)
	fmt.Println()
	testSort(100000)
	fmt.Println()
	testSort(1000000)
	fmt.Println()
	testSort(10000000)

	// Make sure you have enough RAM and/or time before enabling these set sizes.
	// 100 millon items does not run well with 8GB of RAM while running fine with 16GB.
	// 1 billion items does not run well with 16GB of RAM and according to used SWAP space you should have at least 48-64GB of RAM

	/*fmt.Println()
	testSort(100000000)*/
	/*fmt.Println()
	testSort(1000000000)*/
}

func testSort(amount int) {
	ints := make([]int, amount)
	for i := 0; i < amount; i++ {
		ints[i] = rand.Int()
	}
	ints2 := make([]int, amount)
	copy(ints2, ints)
	ints3 := make([]int, amount)
	copy(ints3, ints)
	ints4 := make([]int, amount)
	copy(ints4, ints)
	numbers := make([]uint64, amount)
	for i := 0; i < amount; i++ {
		numbers[i] = uint64(ints[i])
	}
	start := time.Now()
	res1 := localsort.RadixSort(numbers)
	radixResult := make([]int, amount)
	for i := 0; i < amount; i++ {
		radixResult[i] = int(res1[i])
	}
	duration := time.Since(start)
	fmt.Println("Radix sorted ", amount, " in ", duration.String())
	start = time.Now()
	sort.Ints(ints)
	duration = time.Since(start)
	fmt.Println("Sort sorted ", amount, " in ", duration.String())
	start = time.Now()
	res2 := localsort.QuickSort(ints2)
	duration = time.Since(start)
	fmt.Println("Quicksort sorted ", amount, " in ", duration.String())
	var res3 []int
	if amount < 1000000 {
		start = time.Now()
		res3 = localsort.InsertionSort(ints3)
		duration = time.Since(start)
		fmt.Println("Insertion sorted ", amount, " in ", duration.String())
	} else {
		fmt.Println("Insertion sort skipped due to high set size")
	}
	start = time.Now()
	res4 := localsort.MergeSort(ints4)
	duration = time.Since(start)
	fmt.Println("Merge sorted ", amount, " in ", duration.String())
	fmt.Println("Success:")
	fmt.Println("Radix: ", reflect.DeepEqual(ints, radixResult))
	fmt.Println("QuickSort: ", reflect.DeepEqual(ints, res2))
	if amount < 1000000 {
		fmt.Println("Insertion: ", reflect.DeepEqual(ints, res3))
	}
	fmt.Println("Merge: ", reflect.DeepEqual(ints, res4))
}
