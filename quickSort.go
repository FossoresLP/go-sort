package sort

import "cmp"

// QuickSort implements the quicksort algorithm for all ordered primitive types.
// It operates in-place without additional memory allocations.
//
// Deprecated: use slices.Sort instead which provides a more optimized implementation of quicksort.
func QuickSort[T cmp.Ordered](items []T) []T {
	if len(items) < 2 {
		return items
	}
	pivot := items[0]
	i := 0
	j := len(items) - 1
	for {
		for items[i] < pivot {
			i++
		}
		for items[j] > pivot {
			j--
		}
		if i >= j {
			QuickSort(items[:j+1])
			QuickSort(items[j+1:])
			return items
		}
		items[i], items[j] = items[j], items[i]
		i++
		j--
	}
}
