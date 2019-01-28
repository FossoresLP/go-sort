package sort

// MergeSort is an implementation of the merge sort algorithm
func MergeSort(set []int64) []int64 {
	l := len(set)
	if l < 2 {
		return set
	}

	pivot := l / 2
	return MergeSortedSets(MergeSort(set[:pivot]), MergeSort(set[pivot:]))
}

// MergeSortedSets is the merging part of the merge sort algorithm and may be used to combine two sorted sets.
// You might find it useful when combining the results of a distributed sort.
func MergeSortedSets(a, b []int64) []int64 {
	l := len(a) + len(b)
	group := make([]int64, l)
	aPos := 0
	bPos := 0
	for j := 0; j < l; j++ {
		if aPos == len(a) {
			group[j] = b[bPos]
			bPos++
		} else if bPos == len(b) {
			group[j] = a[aPos]
			aPos++
		} else if a[aPos] <= b[bPos] {
			group[j] = a[aPos]
			aPos++
		} else {
			group[j] = b[bPos]
			bPos++
		}
	}
	return group
}
