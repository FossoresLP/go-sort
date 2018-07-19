package sort

func InsertionSort(sort []int) []int {
	l := len(sort)
	sorted := make([]int, l)
	for i := 0; i < l; i++ {
		sorted[i] = sort[i]
		for position := i; position > 0 && sorted[position-1] > sorted[position]; position-- {
			sorted[position], sorted[position-1] = sorted[position-1], sorted[position]
		}
	}
	return sorted
}

func InsertSorted(sorted []int, insert int) []int {
	l := len(sorted)
	out := append(sorted, insert)
	for position := l; position > 0 && out[position-1] > out[position]; position-- {
		out[position], out[position-1] = out[position-1], out[position]
	}
	return out
}
