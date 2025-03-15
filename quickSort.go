package sort

import "cmp"

// QuickSort sorts the supplied integer array using quicksort
func QuickSort[T cmp.Ordered](s []T) []T {
	quicksort(&s, 0, len(s)-1)
	return s
}

func quicksort[T cmp.Ordered](s *[]T, lo, hi int) {
	if lo < hi {
		p := partition(s, lo, hi)
		quicksort(s, lo, p)
		quicksort(s, p+1, hi)
	}
}

func partition[T cmp.Ordered](s *[]T, lo, hi int) int {
	pivot := (*s)[lo]
	i := lo - 1
	if i == -1 {
		i++
	}
	j := hi + 1
	if j == len(*s) {
		j--
	}
	for {
		for (*s)[i] < pivot {
			i++
		}
		for (*s)[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
