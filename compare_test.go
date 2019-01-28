package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

type uint64arr []uint64

func (a uint64arr) Len() int           { return len(a) }
func (a uint64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a uint64arr) Less(i, j int) bool { return a[i] < a[j] }

func TestInsertionSort(t *testing.T) {
	values := make([]int64, 1000)
	for k := 0; k < 1000; k++ {
		values[k] = int64(rand.Uint64())
	}
	values2 := make([]int64, 1000)
	copy(values2, values)
	InsertionSort(values)
	sort.Sort(int64arr(values2))
	if !reflect.DeepEqual(values, values2) {
		t.Error("InsertionSort does not produce the same output as go's sort package.")
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
	sort.Sort(int64arr(values2))
	if !reflect.DeepEqual(values, values2) {
		t.Error("MergeSort does not produce the same output as go's sort package.")
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
	sort.Sort(int64arr(values2))
	if !reflect.DeepEqual(values, values2) {
		t.Error("QuickSort does not produce the same output as go's sort package.")
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
	sort.Sort(uint64arr(values2))
	if !reflect.DeepEqual(values, values2) {
		t.Error("RadixSort does not produce the same output as go's sort package.")
	}
}
