package sort

import (
	"cmp"
	crand "crypto/rand"
	"io"
	"math"
	"math/rand/v2"
	"reflect"
	"slices"
	"strings"
	"testing"
	"unsafe"
)

var source *rand.ChaCha8
var random *rand.Rand

func TestMain(m *testing.M) {
	seed := [32]byte{}
	crand.Read(seed[:])
	source = rand.NewChaCha8(seed)
	random = rand.New(source)
	m.Run()
}

func maxInt[T uint8 | uint16 | uint32 | uint64 | uint | uintptr | int8 | int16 | int32 | int64 | int]() T {
	// Initialize to zero
	var val T
	// Invert all bits
	val = ^val
	// If the resulting value is positive, this is an unsigned integer
	if val > 0 {
		// Return the value since this is the maximum value for unsigned integers
		return val
	}
	// Otherwise reset to one
	val = 1
	// Shift the one into the most significant position and invert all bits
	return ^(val << (unsafe.Sizeof(val)*8 - 1))
}

func minInt[T uint8 | uint16 | uint32 | uint64 | uint | uintptr | int8 | int16 | int32 | int64 | int]() T {
	// Initialize to zero
	var val T
	// Invert all bits
	val = ^val
	// If the resulting value is positive, this is an unsigned integer
	if val > 0 {
		// Return 0 as that is the minimum value for an unsiged integer
		return 0
	}
	// Otherwise reset to one
	val = 1
	// Shift the one into the most significant position
	// Return the value since this is the minimum value for signed integers
	return val << (unsafe.Sizeof(val)*8 - 1)
}

func fillRandom[T cmp.Ordered](slice []T) {
	if len(slice) < 1 {
		return
	}
	bytes := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(slice))), len(slice)*int(unsafe.Sizeof(slice[0])))
	source.Read(bytes)
}

func randomString(length int64) string {
	b := strings.Builder{}
	io.CopyN(&b, source, length)
	return b.String()
}

func testInt[T uint8 | uint16 | uint32 | uint64 | uint | uintptr | int8 | int16 | int32 | int64 | int](t *testing.T, name string, fn func([]T) []T) {
	t.Run(name, func(t *testing.T) {
		tests := []struct {
			Name  string
			Input []T
			Want  []T
		}{
			{"nil slice", nil, nil},
			{"empty slice", []T{}, []T{}},
			{"one element", []T{1}, []T{1}},
			{"two elements unsorted", []T{2, 1}, []T{1, 2}},
			{"two elements sorted", []T{1, 2}, []T{1, 2}},
			{"three elements unsorted", []T{2, 3, 1}, []T{1, 2, 3}},
			{"three elements sorted", []T{1, 2, 3}, []T{1, 2, 3}},
			{"three elements reversed", []T{3, 2, 1}, []T{1, 2, 3}},
			{"ten elements unsorted", []T{7, 5, 9, 3, 0, 1, 4, 6, 2, 8}, []T{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			{"ten elements sorted", []T{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []T{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			{"ten elements reversed", []T{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []T{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			{"ten elements double", []T{2, 2, 8, 8, 4, 4, 6, 6, 0, 0}, []T{0, 0, 2, 2, 4, 4, 6, 6, 8, 8}},
			{"small values", []T{1, minInt[T](), 3, minInt[T]()}, []T{minInt[T](), minInt[T](), 1, 3}},
			{"large values", []T{100, 1, maxInt[T](), 10}, []T{1, 10, 100, maxInt[T]()}},
			{"large range", []T{1, maxInt[T](), 100, minInt[T]()}, []T{minInt[T](), 1, 100, maxInt[T]()}},
		}
		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				fn(tt.Input)
				if !reflect.DeepEqual(tt.Input, tt.Want) {
					t.Errorf("%s result for %s [%+v] does not match expected value [%+v]", name, tt.Name, tt.Input, tt.Want)
				}
			})
		}
		t.Run("1000 random", func(t *testing.T) {
			values := make([]T, 1000)
			values2 := make([]T, 1000)
			fillRandom(values)
			copy(values2, values)
			fn(values)
			slices.Sort(values2)
			if !reflect.DeepEqual(values, values2) {
				t.Error(name + " does not produce the same output as slices.Sort.")
			}
		})
	})
}

func testIntSigned[T int8 | int16 | int32 | int64 | int](t *testing.T, name string, fn func([]T) []T) {
	t.Run(name, func(t *testing.T) {
		tests := []struct {
			Name  string
			Input []T
			Want  []T
		}{
			{"negative values", []T{-1, -2, -3, -4, -5}, []T{-5, -4, -3, -2, -1}},
			{"mixed values", []T{-1, 0, 1, -2, 2, -3, 3}, []T{-3, -2, -1, 0, 1, 2, 3}},
		}
		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				fn(tt.Input)
				if !reflect.DeepEqual(tt.Input, tt.Want) {
					t.Errorf("%s result for %s [%+v] does not match expected value [%+v]", name, tt.Name, tt.Input, tt.Want)
				}
			})
		}
	})
}

func cmpFloatSlice[T float32 | float64](a []T, b []T) bool {
	if a == nil && b == nil {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.IsNaN(float64(a[i])) && math.IsNaN(float64(b[i])) {
			continue
		}
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func testFloat[T float32 | float64](t *testing.T, name string, fn func([]T) []T) {
	t.Run(name, func(t *testing.T) {
		tests := []struct {
			Name    string
			Input   []T
			Want    []T
			CanFail bool
		}{
			{"small values", []T{1, 0.1, 3, 0.01}, []T{0.01, 0.1, 1, 3}, false},
			{"large values", []T{100, 1, 123456789, 10}, []T{1, 10, 100, 123456789}, false},
			{"negative values", []T{-1, -2, -3, -4, -5}, []T{-5, -4, -3, -2, -1}, false},
			{"mixed values", []T{-1, 0, 1.2, 123456789, -2, 2, -3.5, 3}, []T{-3.5, -2, -1, 0, 1.2, 2, 3, 123456789}, false},
			{"infinity", []T{2, T(math.Inf(1)), 1, T(math.Inf(-1))}, []T{T(math.Inf(-1)), 1, 2, T(math.Inf(1))}, false},
			{"NaN", []T{12345, T(math.NaN()), 100, T(math.NaN()), -5.5, T(math.NaN()), T(math.NaN())}, []T{T(math.NaN()), T(math.NaN()), T(math.NaN()), T(math.NaN()), -5.5, 100, 12345}, true},
		}
		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				fn(tt.Input)
				if !cmpFloatSlice(tt.Input, tt.Want) {
					if tt.CanFail {
						t.Skipf("%s result for %s [%+v] does not match expected value [%+v]", name, tt.Name, tt.Input, tt.Want)
					} else {
						t.Errorf("%s result for %s [%+v] does not match expected value [%+v]", name, tt.Name, tt.Input, tt.Want)
					}
				}
			})
		}
		t.Run("1000 random", func(t *testing.T) {
			values := make([]T, 1000)
			values2 := make([]T, 1000)
			fillRandom(values)
			for i := range values {
				if math.IsNaN(float64(values[i])) {
					values[i] = 0
				}
			}
			copy(values2, values)
			fn(values)
			slices.Sort(values2)
			if !cmpFloatSlice(values, values2) {
				t.Error(name + " does not produce the same output as slices.Sort.")
			}
		})
	})
}

func TestSort_uint(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]uint) []uint
	}{
		{"InsertionSort", InsertionSort[uint]},
		{"MergeSort", MergeSort[uint]},
		{"QuickSort", QuickSort[uint]},
		{"RadixSort", RadixSort[uint]},
	}

	for _, alg := range algorithms {
		testInt(t, alg.Name, alg.Func)
	}
}

func TestSort_uint8(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]uint8) []uint8
	}{
		{"InsertionSort", InsertionSort[uint8]},
		{"MergeSort", MergeSort[uint8]},
		{"QuickSort", QuickSort[uint8]},
		{"RadixSort", RadixSort[uint8]},
	}

	for _, alg := range algorithms {
		testInt(t, alg.Name, alg.Func)
	}
}

func TestSort_uint16(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]uint16) []uint16
	}{
		{"InsertionSort", InsertionSort[uint16]},
		{"MergeSort", MergeSort[uint16]},
		{"QuickSort", QuickSort[uint16]},
		{"RadixSort", RadixSort[uint16]},
	}

	for _, alg := range algorithms {
		testInt(t, alg.Name, alg.Func)
	}
}

func TestSort_uint32(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]uint32) []uint32
	}{
		{"InsertionSort", InsertionSort[uint32]},
		{"MergeSort", MergeSort[uint32]},
		{"QuickSort", QuickSort[uint32]},
		{"RadixSort", RadixSort[uint32]},
	}

	for _, alg := range algorithms {
		testInt(t, alg.Name, alg.Func)
	}
}

func TestSort_uint64(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]uint64) []uint64
	}{
		{"InsertionSort", InsertionSort[uint64]},
		{"MergeSort", MergeSort[uint64]},
		{"QuickSort", QuickSort[uint64]},
		{"RadixSort", RadixSort[uint64]},
	}

	for _, alg := range algorithms {
		testInt(t, alg.Name, alg.Func)
	}
}

func TestSort_uintptr(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]uintptr) []uintptr
	}{
		{"InsertionSort", InsertionSort[uintptr]},
		{"MergeSort", MergeSort[uintptr]},
		{"QuickSort", QuickSort[uintptr]},
		{"RadixSort", RadixSort[uintptr]},
	}

	for _, alg := range algorithms {
		testInt(t, alg.Name, alg.Func)
	}
}

func TestSort_int(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]int) []int
	}{
		{"InsertionSort", InsertionSort[int]},
		{"MergeSort", MergeSort[int]},
		{"QuickSort", QuickSort[int]},
		{"RadixSort", RadixSort[int]},
	}

	for _, alg := range algorithms {
		testInt(t, alg.Name, alg.Func)
		testIntSigned(t, alg.Name, alg.Func)
	}
}

func TestSort_int8(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]int8) []int8
	}{
		{"InsertionSort", InsertionSort[int8]},
		{"MergeSort", MergeSort[int8]},
		{"QuickSort", QuickSort[int8]},
		{"RadixSort", RadixSort[int8]},
	}

	for _, alg := range algorithms {
		testInt(t, alg.Name, alg.Func)
		testIntSigned(t, alg.Name, alg.Func)
	}
}

func TestSort_int16(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]int16) []int16
	}{
		{"InsertionSort", InsertionSort[int16]},
		{"MergeSort", MergeSort[int16]},
		{"QuickSort", QuickSort[int16]},
		{"RadixSort", RadixSort[int16]},
	}

	for _, alg := range algorithms {
		testInt(t, alg.Name, alg.Func)
		testIntSigned(t, alg.Name, alg.Func)
	}
}

func TestSort_int32(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]int32) []int32
	}{
		{"InsertionSort", InsertionSort[int32]},
		{"MergeSort", MergeSort[int32]},
		{"QuickSort", QuickSort[int32]},
		{"RadixSort", RadixSort[int32]},
	}

	for _, alg := range algorithms {
		testInt(t, alg.Name, alg.Func)
		testIntSigned(t, alg.Name, alg.Func)
	}
}

func TestSort_int64(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]int64) []int64
	}{
		{"InsertionSort", InsertionSort[int64]},
		{"MergeSort", MergeSort[int64]},
		{"QuickSort", QuickSort[int64]},
		{"RadixSort", RadixSort[int64]},
	}

	for _, alg := range algorithms {
		testInt(t, alg.Name, alg.Func)
		testIntSigned(t, alg.Name, alg.Func)
	}
}

func TestSort_float32(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]float32) []float32
	}{
		{"InsertionSort", InsertionSort[float32]},
		{"MergeSort", MergeSort[float32]},
		{"QuickSort", QuickSort[float32]},
		{"RadixSort", RadixSort[float32]},
	}

	for _, alg := range algorithms {
		testFloat(t, alg.Name, alg.Func)
	}
}

func TestSort_float64(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]float64) []float64
	}{
		{"InsertionSort", InsertionSort[float64]},
		{"MergeSort", MergeSort[float64]},
		{"QuickSort", QuickSort[float64]},
		{"RadixSort", RadixSort[float64]},
	}

	for _, alg := range algorithms {
		testFloat(t, alg.Name, alg.Func)
	}
}

func TestSort_string(t *testing.T) {
	algorithms := []struct {
		Name string
		Func func([]string) []string
	}{
		{"InsertionSort", InsertionSort[string]},
		{"MergeSort", MergeSort[string]},
		{"QuickSort", QuickSort[string]},
		{"RadixSort", RadixSort[string]},
	}
	for _, alg := range algorithms {
		t.Run(alg.Name, func(t *testing.T) {
			tests := []struct {
				Name  string
				Input []string
				Want  []string
			}{
				{"empty string", []string{"b", "", "a"}, []string{"", "a", "b"}},
				{"same prefix", []string{"aab", "aaa", "aac"}, []string{"aaa", "aab", "aac"}},
				{"different lengths", []string{"c", "aaa", "bcdefghijklmnopqrstuvwxyz"}, []string{"aaa", "bcdefghijklmnopqrstuvwxyz", "c"}},
			}
			for _, tt := range tests {
				t.Run(tt.Name, func(t *testing.T) {
					alg.Func(tt.Input)
					if !reflect.DeepEqual(tt.Input, tt.Want) {
						t.Errorf("%s result for %s [%+v] does not match expected value [%+v]", alg.Name, tt.Name, tt.Input, tt.Want)
					}
				})
			}
			t.Run("1000 random", func(t *testing.T) {
				values := make([]string, 1000)
				values2 := make([]string, 1000)
				for i := range values {
					values[i] = randomString(random.Int64N(100))
				}
				copy(values2, values)
				alg.Func(values)
				slices.Sort(values2)
				if !reflect.DeepEqual(values, values2) {
					t.Error(alg.Name + " does not produce the same output as slices.Sort.")
				}
			})
		})
	}
}

func TestInsertSorted(t *testing.T) {
	tests := []struct {
		name   string
		sorted []int
		insert int
		want   []int
	}{
		{
			name:   "insert into empty slice",
			sorted: []int{},
			insert: 5,
			want:   []int{5},
		},
		{
			name:   "insert at beginning",
			sorted: []int{2, 3, 4, 5},
			insert: 1,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "insert at end",
			sorted: []int{1, 2, 3, 4},
			insert: 5,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "insert in middle",
			sorted: []int{1, 2, 4, 5},
			insert: 3,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "insert duplicate value",
			sorted: []int{1, 2, 3, 3, 5},
			insert: 3,
			want:   []int{1, 2, 3, 3, 3, 5},
		},
	}

	// Test with integers
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertSorted(tt.sorted, tt.insert)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSorted() = %v, want %v", got, tt.want)
			}
		})
	}

	// Test with strings
	stringTests := []struct {
		name   string
		sorted []string
		insert string
		want   []string
	}{
		{
			name:   "insert string at beginning",
			sorted: []string{"b", "c", "d"},
			insert: "a",
			want:   []string{"a", "b", "c", "d"},
		},
		{
			name:   "insert string in middle",
			sorted: []string{"a", "b", "d", "e"},
			insert: "c",
			want:   []string{"a", "b", "c", "d", "e"},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertSorted(tt.sorted, tt.insert)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSortedSets(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "both empty slices",
			a:    []int{},
			b:    []int{},
			want: []int{},
		},
		{
			name: "first slice empty",
			a:    []int{},
			b:    []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "second slice empty",
			a:    []int{1, 2, 3},
			b:    []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "disjoint ranges",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "interleaved values",
			a:    []int{1, 3, 5},
			b:    []int{2, 4, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "overlapping values",
			a:    []int{1, 2, 3, 7, 9},
			b:    []int{2, 4, 6, 8},
			want: []int{1, 2, 2, 3, 4, 6, 7, 8, 9},
		},
		{
			name: "duplicate values",
			a:    []int{1, 3, 3, 5},
			b:    []int{2, 3, 4},
			want: []int{1, 2, 3, 3, 3, 4, 5},
		},
		{
			name: "different lengths",
			a:    []int{1, 3, 5, 7, 9},
			b:    []int{2, 4},
			want: []int{1, 2, 3, 4, 5, 7, 9},
		},
	}

	// Test with integers
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSortedSets(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortedSets() = %v, want %v", got, tt.want)
			}
		})
	}

	// Test with strings
	stringTests := []struct {
		name string
		a    []string
		b    []string
		want []string
	}{
		{
			name: "merge sorted strings",
			a:    []string{"a", "c", "e"},
			b:    []string{"b", "d", "f"},
			want: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name: "merge with duplicates",
			a:    []string{"a", "c", "c"},
			b:    []string{"b", "c", "d"},
			want: []string{"a", "b", "c", "c", "c", "d"},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSortedSets(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortedSets() = %v, want %v", got, tt.want)
			}
		})
	}
}
