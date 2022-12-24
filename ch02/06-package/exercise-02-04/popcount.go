// Package popcount
package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	bits := 0
	for x != 0 {
		bit := int(x & 1)
		bits += bit
		x >>= 1
	}
	return bits
}
