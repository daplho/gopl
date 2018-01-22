package popcount3

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	popCount := 0
	for x != 0 {
		x = x & (x - 1)
		popCount++
	}
	return popCount
}
