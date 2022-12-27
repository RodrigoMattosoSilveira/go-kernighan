package popcount

import (
	popLoop "github.com/RodrigoMattosoSilveira/go-kernighan/ch02/06-package/exercise-02-03"
	popTable "github.com/RodrigoMattosoSilveira/go-kernighan/ch02/06-package/popcount"
	"testing"
)

// popcount
func BenchmarkTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popTable.PopCount(1023)
	}
}

// exercise_02_03
func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popLoop.PopCount(1023)
	}
}

// exercise_02_04
func BenchmarkShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1023)
	}
}
