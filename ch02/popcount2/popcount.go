package popcount2

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	popCount := 0
	for i := uint(0); i < 64; i++ {
		if (x>>i)&1 != 0 {
			popCount++
		}
	}
	return popCount
}
