package sort

// RadixSort is radix sort for 64 bit unsigned integers
// It operates byte-by-byte using 256 buckets
func RadixSort(items []uint64) []uint64 {
	l := len(items)
	tmp := make([]uint64, l)
	src := items
	dst := tmp

	// Loop over the 8 bytes created for each unsigned integer
	for shift := 0; shift < 64; shift += 8 {
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
