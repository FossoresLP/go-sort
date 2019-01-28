package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkInsertionSort(b *testing.B) {
	for i := 100; ; i *= 10 {
		maxn := 0
		b.Run(fmt.Sprintf("InsertionSort %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				b.StopTimer()
				values := make([]int64, i)
				for k := 0; k < i; k++ {
					values[k] = int64(rand.Uint64())
				}
				b.StartTimer()
				InsertionSort(values)
			}
			maxn = b.N
		})
		if maxn == 1 {
			break
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 100; ; i *= 10 {
		maxn := 0
		b.Run(fmt.Sprintf("MergeSort %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				b.StopTimer()
				values := make([]int64, i)
				for k := 0; k < i; k++ {
					values[k] = int64(rand.Uint64())
				}
				b.StartTimer()
				MergeSort(values)
			}
			maxn = b.N
		})
		if maxn == 1 {
			break
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 100; ; i *= 10 {
		maxn := 0
		b.Run(fmt.Sprintf("QuickSort %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				b.StopTimer()
				values := make([]int64, i)
				for k := 0; k < i; k++ {
					values[k] = int64(rand.Uint64())
				}
				b.StartTimer()
				QuickSort(values)
			}
			maxn = b.N
		})
		if maxn == 1 {
			break
		}
	}
}

func BenchmarkRadixSort(b *testing.B) {
	for i := 100; ; i *= 10 {
		maxn := 0
		b.Run(fmt.Sprintf("RadixSort %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				b.StopTimer()
				values := make([]uint64, i)
				for k := 0; k < i; k++ {
					values[k] = rand.Uint64()
				}
				b.StartTimer()
				RadixSort(values)
			}
			maxn = b.N
		})
		if maxn == 1 {
			break
		}
	}
}
