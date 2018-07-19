package sort

import (
	"encoding/binary"
)

// RadixSort is radix sort for 64 bit unsigned integers implemented using binary data.
// This algorithm does not compare the data and due to the binary implementation also does not perform any calculations with it.
// It is faster than quicksort but cannot handle negative values.
func RadixSort(items []uint64) []uint64 {
	// Create variables
	size := len(items)
	sort := make([][8]byte, size)
	sorted := make([][8]byte, size)

	// Convert the unsigned integers to binary data
	for i := 0; i < size; i++ {
		binary.LittleEndian.PutUint64(sort[i][:], items[i])
	}

	// Loop over the 8 bytes created for each unsigned integer
	for e := 0; e < 8; e++ {
		bucket := [256]int{0}
		for i := 0; i < size; i++ {
			bucket[sort[i][e]]++
		}

		// Add count from previous bucket
		// The bucket values are used as indices for the sorted array and therefore have to be higher for buckets with higher sort values.
		for i := 1; i < 256; i++ {
			bucket[i] += bucket[i-1]
		}

		// Use the buckets indices when filling the sorted array
		for i := size - 1; i >= 0; i-- {
			bucket[sort[i][e]]--
			sorted[bucket[sort[i][e]]] = sort[i]
		}

		// Copy the values from sorted back to sort for the next run of the loop
		copy(sort, sorted)
	}

	// Convert the bytes back to unsigned integers and return those
	out := make([]uint64, size)
	for i := 0; i < size; i++ {
		out[i] = binary.LittleEndian.Uint64(sorted[i][:])
	}
	return out
}
