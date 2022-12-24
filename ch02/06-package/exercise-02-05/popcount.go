// Package popcount
package popcount

import "math/bits"

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint) int {
	return bits.OnesCount(x)
}
