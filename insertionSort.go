package sort

import "cmp"

// InsertionSort is an implementation of the Insertion sort algorithm.
// It is extremely slow compared to other algorithms like Quicksort, Merge sort and Radix sort and should therefore only be used for very small sets.
func InsertionSort[T cmp.Ordered](sort []T) []T {
	for i := 0; i < len(sort); i++ {
		for position := i; position > 0 && sort[position-1] > sort[position]; position-- {
			sort[position], sort[position-1] = sort[position-1], sort[position]
		}
	}
	return sort
}

// InsertSorted inserts a single element into a sorted set.
// This can be useful when inserting a new element into a database.
func InsertSorted[T cmp.Ordered](sorted []T, insert T) []T {
	l := len(sorted)
	out := append(sorted, insert)
	for position := l; position > 0 && out[position-1] > out[position]; position-- {
		out[position], out[position-1] = out[position-1], out[position]
	}
	return out
}
