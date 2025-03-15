package sort

import "cmp"

// InsertionSort implements insertion sort for all ordered primitive types.
// It is only good for very small slices or almost sorted data.
// For larger, unsorted values its complexity is O(n²), leading to very poor performance.
func InsertionSort[T cmp.Ordered](items []T) []T {
	for i := range items {
		for position := i; position > 0 && items[position-1] > items[position]; position-- {
			items[position], items[position-1] = items[position-1], items[position]
		}
	}
	return items
}

// InsertSorted inserts a single element into an already sorted slice.
// It is more efficient at this operation than resorting the entire array.
func InsertSorted[T cmp.Ordered](sorted []T, insert T) []T {
	out := append(sorted, insert)
	for position := len(sorted); position > 0 && out[position-1] > out[position]; position-- {
		out[position], out[position-1] = out[position-1], out[position]
	}
	return out
}
