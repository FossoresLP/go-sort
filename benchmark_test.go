package sort

import (
	"cmp"
	"fmt"
	"math"
	"math/rand"
	"slices"
	"testing"
)

func slicesSort[T cmp.Ordered](items []T) []T {
	slices.Sort(items)
	return items
}

func BenchmarkSort(b *testing.B) {
	tests := []struct {
		Name string
		Func func([]uint64) []uint64
	}{
		{"InsertionSort", InsertionSort[uint64]},
		{"MergeSort", MergeSort[uint64]},
		{"QuickSort", QuickSort[uint64]},
		{"RadixSort", RadixSort[uint64]},
		{"slices.Sort", slicesSort[uint64]},
	}
	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			runs := math.MaxInt
			for i := 100; i <= 1000000000; i *= 10 {
				b.Run(fmt.Sprintf("items=%d", i), func(b *testing.B) {
					values := make([]uint64, i)
					data := make([]uint64, i)
					for k := range len(values) {
						values[k] = rand.Uint64()
					}
					for b.Loop() {
						copy(data, values)
						tt.Func(data)
					}
					runs = b.N
				})
				if runs == 1 {
					break
				}
			}
		})
	}
}
