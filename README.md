Sort
====

This package contains implementations of Radix Sort, Quicksort, Merge Sort and Insertion Sort in Go.

In addition, it provides the functions `InsertSorted` and `MergeSortedSets`.

All implementations use generics and can operate on ordered primitive types as defined by `cmp.Ordered`.

Sorting is performed in-place where possible and the input slice is always updated. As a convenience, it is returned as well.

Benchmarks for all supported functions can be found in [benchmark.md](./benchmark.md).

Radix Sort
----------

Radix Sort is a sorting algorithm that can operate in linear time O(n) for certain inputs.

This implementation currently only supports integer types (signed and unsigned), falling back to `slices.Sort` automatically for all other types of data.

Since Radix Sort is very difficult to implement efficiently in-place, this implementation creates a copy of the data.

Sorting is performed using 256 buckets which means a single byte of the input items is processed at a time.

```go
sort.RadixSort[T cmp.Ordered](items []T) []T
```

The implementation is based on the design by Austin G. Walters described in [Radix Sort in Go (Golang)](https://austingwalters.com/radix-sort-in-go/)

Quicksort
---------

Quicksort is a very versatile and fast sorting algorithm.

A more optimized version is part of the standard library as `slices.Sort` and should be used instead of this implementation.

In-place sorting means that no additional memory is allocated.

While the average complexity is O(n log n), Quicksort has a worst-case complexity of O(n²).

The standard library implementation gets around this by falling back to Heap Sort in certain cases.

```go
sort.QuickSort[T cmp.Ordered](items []T) []T
```

This implementation is based on the Hoare partition scheme and has been adapted from the pseudocode on Wikipedia [Quicksort](https://en.wikipedia.org/wiki/Quicksort#Hoare_partition_scheme)

Merge Sort
----------

Merge Sort is a sorting algorithm that works by recursively merging sorted slices, starting from a size of 1.

It is a stable sorting algorithm meaning items with the same value will retain their original order.

The worst-case complexity is O(n log n), but it generally performs slightly worse than a well-optimized Quicksort.

An in-place implementation is not possible, meaning a copy of the items is created in the process.

```go
sort.MergeSort[T cmp.Ordered](items []T) []T
```

A part of merge sort, the function `MergeSortedSets` is exposed as well.

It efficiently combines two already sorted sets.

This can be useful for combining the results of distributed (or multithreaded) sorting algorithms as well as inserting multiple elements (which can quickly be sorted e.g. using insertion sort) into a large, already sorted slice.

```go
sort.MergeSortedSets[T cmp.Ordered](a []T, b []T) []T
```

Insertion Sort
--------------

Insertion Sort is a very simply but inefficient sorting algorithm that works by starting with a single value and then adding one item after the other at the right position to keep the result sorted.

It operates in-place and can be very efficient for small sets since it does not depend on recursion, but with a complexity of O(n²), it does not scale well.

```go
sort.InsertionSort[T cmp.Ordered](items []T) []T

```

The insertion process of Insertion Sort can be used individually to efficiently insert a single element into an already sorted slice.

```go
sort.InsertSorted[T cmp.Ordered](sorted []T, insert T) []T
```
