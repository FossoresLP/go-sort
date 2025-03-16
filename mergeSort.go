package sort

import (
	"cmp"
	"slices"
)

// MergeSort implements merge sort for all ordered primitive types.
// It is a stable sorting algorithm, therefore maintaining the order of elements that have the same value.
// The implementation does not operate in-place, temporarily allocating a copy of the data that needs to be sorted.
// The worst-case performance is O(n log n) with a static space requirement of O(2n)
func MergeSort[T cmp.Ordered](items []T) []T {
	if len(items) < 2 {
		return items
	}

	// Create a copy of the data since merge sort cannot easily operate in-place
	tmp := make([]T, len(items))
	copy(tmp, items)

	// Sort with alternating source and destination
	mergeSort(tmp, items)

	return items
}

// mergeSort in the recursive sorting part of merge sort and will call itself for each half of the supplied data and then mergeSortedSets to merge the results.
func mergeSort[T cmp.Ordered](src, dst []T) {
	if len(src) < 2 {
		return
	}

	// Find the midpoint
	mid := len(src) / 2

	// Recursively sort the two halves with swapped src and dst
	mergeSort(dst[:mid], src[:mid])
	mergeSort(dst[mid:], src[mid:])

	// Merge the sorted halves from src into dst
	mergeSortedSets(src[:mid], src[mid:], dst)
}

// MergeSortedSets merges two already sorted slices very efficiently.
// It is used as part of merge sort but can also be useful when inserting multiple elements into a sorted slice (though the elements to insert first have to be sorted) or when combining the results of distributed sorting algoritms.
func MergeSortedSets[T cmp.Ordered](a, b []T) []T {
	if len(a) == 0 {
		return slices.Clone(b)
	}
	if len(b) == 0 {
		return slices.Clone(a)
	}
	buf := make([]T, len(a)+len(b))
	mergeSortedSets(a, b, buf)
	return buf
}

// mergeSortedSets implements the actual sorting logic but requires a target buffer to be supplied, making it unsuitable as the public interface.
// It is used by mergeSort and wrapped by MergeSortedSets for external use.
func mergeSortedSets[T cmp.Ordered](a, b []T, buf []T) {
	length := len(a) + len(b)
	aPos := 0
	bPos := 0
	for i := range length {
		if a[aPos] <= b[bPos] {
			buf[i] = a[aPos]
			aPos++
			if aPos == len(a) {
				copy(buf[i+1:], b[bPos:])
				return
			}
		} else {
			buf[i] = b[bPos]
			bPos++
			if bPos == len(b) {
				copy(buf[i+1:], a[aPos:])
				return
			}
		}
	}
}
