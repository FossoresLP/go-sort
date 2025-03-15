package sort

import (
	"cmp"
	"math/bits"
	"slices"
	"unsafe"
)

// RadixSort implements radix sort using byte-by-byte sorting with 256 buckets for all integer types.
// It creates an internal copy of the supplied data, leading to one large allocation.
// The result is updated in-place and returned for convenience as well.
// Other data types such as float64, float32 and string are handled via a fallback to slices.Sort.
// Signed integers are handled by flipping the sign bit before and after sorting and treating them as unsigned integers.
// The computational complexity is O(n) with a space requirement of O(2n).
func RadixSort[T cmp.Ordered](items []T) []T {
	// No need to sort slices with less than two items
	if len(items) < 2 {
		return items
	}
	var val T
	switch any(val).(type) {
	case uint64:
		radixSortUint(any(items).([]uint64))
	case uint32:
		radixSortUint(any(items).([]uint32))
	case uint16:
		radixSortUint(any(items).([]uint16))
	case uint8:
		countingSort(any(items).([]uint8))
	case uint:
		radixSortUint(any(items).([]uint))
	case uintptr:
		radixSortUint(any(items).([]uintptr))
	case int64:
		uintslice := unsafe.Slice((*uint64)(unsafe.Pointer(unsafe.SliceData(items))), len(items))
		for i := range uintslice {
			uintslice[i] ^= 0x8000000000000000
		}
		radixSortUint(uintslice)
		for i := range uintslice {
			uintslice[i] ^= 0x8000000000000000
		}
	case int32:
		uintslice := unsafe.Slice((*uint32)(unsafe.Pointer(unsafe.SliceData(items))), len(items))
		for i := range uintslice {
			uintslice[i] ^= 0x80000000
		}
		radixSortUint(uintslice)
		for i := range uintslice {
			uintslice[i] ^= 0x80000000
		}
	case int16:
		uintslice := unsafe.Slice((*uint16)(unsafe.Pointer(unsafe.SliceData(items))), len(items))
		for i := range uintslice {
			uintslice[i] ^= 0x8000
		}
		radixSortUint(uintslice)
		for i := range uintslice {
			uintslice[i] ^= 0x8000
		}
	case int8:
		uintslice := unsafe.Slice((*uint8)(unsafe.Pointer(unsafe.SliceData(items))), len(items))
		for i := range uintslice {
			uintslice[i] ^= 0x80
		}
		countingSort(uintslice)
		for i := range uintslice {
			uintslice[i] ^= 0x80
		}
	case int:
		uintslice := unsafe.Slice((*uint)(unsafe.Pointer(unsafe.SliceData(items))), len(items))
		var mask uint = 1 << (bits.UintSize - 1)
		for i := range uintslice {
			uintslice[i] ^= mask
		}
		radixSortUint(uintslice)
		for i := range uintslice {
			uintslice[i] ^= mask
		}
	default:
		slices.Sort(items)
	}
	return items
}

// radixSortUint implements radix sort for all multi-byte unsigned integer types, adapting to their respective sizes
func radixSortUint[T uint64 | uint32 | uint16 | uint | uintptr](items []T) []T {
	tmp := make([]T, len(items))
	src := items
	dst := tmp
	var val T
	bits := int(unsafe.Sizeof(val)) * 8

	// Loop over the individual bytes of the unsigned integer type
	for shift := 0; shift < bits; shift += 8 {
		// Create buckets and count items
		bucket := [256]int{}
		for i := range items {
			bucket[int(src[i]>>shift&0xFF)]++
		}

		// Add count from previous bucket
		// The bucket values are used as indices for the sorted array and therefore have to be higher for buckets with higher sort values.
		for i := 1; i < 256; i++ {
			bucket[i] += bucket[i-1]
		}

		// Use the buckets indices when filling the sorted array
		for i := len(items) - 1; i >= 0; i-- {
			v := &bucket[int(src[i]>>shift&0xFF)]
			*v--
			dst[*v] = src[i]
		}

		// Swap source and destination for the next pass
		src, dst = dst, src
	}

	return items
}

func countingSort(items []uint8) []uint8 {
	tmp := make([]uint8, len(items))

	// Create buckets and count items
	bucket := [256]int{}
	for i := range items {
		bucket[int(items[i])]++
	}

	// Add count from previous bucket
	// The bucket values are used as indices for the sorted array and therefore have to be higher for buckets with higher sort values.
	for i := 1; i < 256; i++ {
		bucket[i] += bucket[i-1]
	}

	// Use the buckets indices when filling the sorted array
	for i := len(items) - 1; i >= 0; i-- {
		v := &bucket[int(items[i])]
		*v--
		tmp[*v] = items[i]
	}

	// Copy result back
	copy(items, tmp)

	return items
}
