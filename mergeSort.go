package sort

import "cmp"

// MergeSort is an implementation of the merge sort algorithm
func MergeSort[T cmp.Ordered](set []T) []T {
	if len(set) < 2 {
		return set
	}

	// Create only one temporary array
	tmp := make([]T, len(set))

	// Copy the original data to temp
	copy(tmp, set)

	// Sort with alternating source and destination
	mergeSort(tmp, set)

	return set
}

func mergeSort[T cmp.Ordered](src, dst []T) {
	// Base case: single element or empty slice
	if len(src) == 1 {
		// For a single element, copy it to destination if needed
		dst[0] = src[0]
		return
	}

	// Find the midpoint
	mid := len(src) / 2

	// Recursively sort the two halves with swapped src and dst
	// This is the key trick: we alternate the roles of src and dst
	mergeSort(dst[:mid], src[:mid])
	mergeSort(dst[mid:], src[mid:])

	// Merge the sorted halves from src into dst
	mergeSortedSets(src[:mid], src[mid:], dst)
}

// MergeSortedSets is the merging part of the merge sort algorithm and may be used to combine two sorted sets.
// You might find it useful when combining the results of a distributed sort.
func MergeSortedSets[T cmp.Ordered](a, b []T) []T {
	l := len(a) + len(b)
	buf := make([]T, l)
	return mergeSortedSets(a, b, buf)
}

func mergeSortedSets[T cmp.Ordered](a, b []T, buf []T) []T {
	l := len(a) + len(b)
	aPos := 0
	bPos := 0
	for j := 0; j < l; j++ {
		if aPos == len(a) {
			buf[j] = b[bPos]
			bPos++
		} else if bPos == len(b) {
			buf[j] = a[aPos]
			aPos++
		} else if a[aPos] <= b[bPos] {
			buf[j] = a[aPos]
			aPos++
		} else {
			buf[j] = b[bPos]
			bPos++
		}
	}
	return buf
}
