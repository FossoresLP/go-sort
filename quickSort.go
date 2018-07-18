package sort

// QuickSort sorts the supplied integer array using quicksort
func QuickSort(s []int) []int {
	quicksort(&s, 0, len(s)-1)
	return s
}

func quicksort(s *[]int, lo, hi int) {
	if lo < hi {
		p := partition(s, lo, hi)
		quicksort(s, lo, p)
		quicksort(s, p+1, hi)
	}
}

func partition(s *[]int, lo, hi int) int {
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
