package sort

import (
	"math/rand"
	"reflect"
	"slices"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	values := make([]int64, 1000)
	for k := 0; k < 1000; k++ {
		values[k] = int64(rand.Uint64())
	}
	values2 := make([]int64, 1000)
	copy(values2, values)
	InsertionSort(values)
	slices.Sort(values2)
	if !reflect.DeepEqual(values, values2) {
		t.Error("InsertionSort does not produce the same output as slices.Sort.")
	}
}

func TestMergeSort(t *testing.T) {
	values := make([]int64, 1000)
	for k := 0; k < 1000; k++ {
		values[k] = int64(rand.Uint64())
	}
	values2 := make([]int64, 1000)
	copy(values2, values)
	values = MergeSort(values)
	slices.Sort(values2)
	if !reflect.DeepEqual(values, values2) {
		t.Error("MergeSort does not produce the same output as slices.Sort.")
	}
}

func TestQuickSort(t *testing.T) {
	values := make([]int64, 1000)
	for k := 0; k < 1000; k++ {
		values[k] = int64(rand.Uint64())
	}
	values2 := make([]int64, 1000)
	copy(values2, values)
	QuickSort(values)
	slices.Sort(values2)
	if !reflect.DeepEqual(values, values2) {
		t.Error("QuickSort does not produce the same output as slices.Sort.")
	}
}

func TestRadixSort(t *testing.T) {
	values := make([]uint64, 1000)
	for k := 0; k < 1000; k++ {
		values[k] = rand.Uint64()
	}
	values2 := make([]uint64, 1000)
	copy(values2, values)
	values = RadixSort(values)
	slices.Sort(values2)
	if !reflect.DeepEqual(values, values2) {
		t.Error("RadixSort does not produce the same output as slices.Sort.")
	}
}
