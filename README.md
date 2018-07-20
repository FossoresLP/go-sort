Sort
====

This package contains implementations of Radix Sort, Quicksort and Merge Sort in Golang.
They are licensed under the Boost Software License 1.0.
Benchmark results are included in benchmark.md.

Radix Sort
----------

Radix sort is an extremely fast algorithm implemented here specifically to sort unsigned integers.
It is 2-3 times faster than Quicksort for large sets of data.
Please make sure you have at least enough RAM available to fit your application and 4 times the size of the set.

```go
sort.RadixSort(set []uint64) []uint64
```

This version of Radix sort uses bytewise sorting with 256 buckets which increases memory usage due to the additional type conversions but improves performance.
The implementation is based on the design by Austin G. Walters described in [Radix Sort in Go (Golang)](https://austingwalters.com/radix-sort-in-go/)

Quicksort
---------

Quicksort is a very versatile and fast sorting algorithm implemented here to sort integers.
It is 2-3 times faster than Golang's internal sort package for large sets of data.
It uses only one copy of the set and is therefore preferred in constrained environments.

```go
sort.QuickSort(set []int) []int
```

This implementation is based on the Hoare partition scheme and has been adapted from the pseudocode on Wikipedia [Quicksort](https://en.wikipedia.org/wiki/Quicksort#Hoare_partition_scheme)

Merge Sort
----------

Merge sort is an algorithm that sorts sets by dividing them and sorting them while merging them together.
It is slower than Quicksort but can be especially useful when merging multiple sets of sorted data which may occur when distributing sorting across devices.

```go
sort.MergeSort(set []int) []int
sort.MergeSortedSets(set1 []int, set2 []int) []int
```

Merge sort is a stable sorting algorithm that handles negative values and should therefore be preferred over Quicksort when you need to preserve the order of equal elements.
Due to keeping up to 2 copies of the full set in memory, it is not recommended to use Merge sort in constrained environments.

Merge sorted sets is intended to be used to merge two sorted sets together. It is used by Merge sort internally but can also be used to merge the results of multi-threaded or distributed sorting algorithms.

Insertion Sort
--------------

Insertion sort is a sorting algorithm that works by inserting individual items at the right position, keeping the existing set sorted.
It can be used to sort a set by starting with only one value in the sorted set but is much better suited for inserting a single new element into a sorted set.

```go
sort.InsertionSort(set []int) []int
sort.InsertSorted(sortedSet []int, insert int) []int
```

Due to inferior speed compared to all other algorithms in this package I would not recommend using Insertion sort for normal applications. If you don't know about a reason to use it you probably shouldn't.

Insert is intended to be used when adding one new element as may occur when adding a new item to a database.

Example / Tests / Benchmark
---------------------------

An example can be found in example. It is also used for testing as it compares the algorithms to the sort package in Golang. It may also be consided a benchmark as it records runtime for all algorithms including Golang's sort.

To run the example just navigate to the example directory and run `go run example.go`

There are no other tests implemented for this package and I don't plan to create any in the future. In case you would like to add tests feel free to submit a PR.
