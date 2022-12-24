// Package popcount
package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	bits := 0
	for temp := x; temp != 0; temp = temp >> 1 {
		bit := int(temp % 2)
		bits += bit
	}
	return bits
}
