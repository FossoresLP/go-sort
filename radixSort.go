package sort

import (
	"cmp"
	"math/bits"
	"slices"
	"unsafe"
)

// RadixSort is radix sort for 64 bit unsigned integers
// It operates byte-by-byte using 256 buckets
func RadixSort[T cmp.Ordered](items []T) []T {
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

func radixSortUint[T uint64 | uint32 | uint16 | uint | uintptr](items []T) []T {
	l := len(items)
	tmp := make([]T, l)
	src := items
	dst := tmp
	var val T
	bits := int(unsafe.Sizeof(val)) * 8

	// Loop over the 8 bytes created for each unsigned integer
	for shift := 0; shift < bits; shift += 8 {
		bucket := [256]int{}
		for i := 0; i < l; i++ {
			bucket[int(src[i]>>shift&0xFF)]++
		}

		// Add count from previous bucket
		// The bucket values are used as indices for the sorted array and therefore have to be higher for buckets with higher sort values.
		for i := 1; i < 256; i++ {
			bucket[i] += bucket[i-1]
		}

		// Use the buckets indices when filling the sorted array
		for i := l - 1; i >= 0; i-- {
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
	l := len(items)
	tmp := make([]uint8, l)

	bucket := [256]int{}
	for i := 0; i < l; i++ {
		bucket[int(items[i])]++
	}

	// Add count from previous bucket
	// The bucket values are used as indices for the sorted array and therefore have to be higher for buckets with higher sort values.
	for i := 1; i < 256; i++ {
		bucket[i] += bucket[i-1]
	}

	// Use the buckets indices when filling the sorted array
	for i := l - 1; i >= 0; i-- {
		v := &bucket[int(items[i])]
		*v--
		tmp[*v] = items[i]
	}

	copy(items, tmp)

	return items
}
