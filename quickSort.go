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
	p := 0
	for {
		for items[i] < pivot {
			i++
		}
		for items[j] > pivot {
			j--
		}
		if i >= j {
			p = j
			break
		}
		items[i], items[j] = items[j], items[i]
	}
	QuickSort(items[:p])
	QuickSort(items[p+1:])
	return items
}
