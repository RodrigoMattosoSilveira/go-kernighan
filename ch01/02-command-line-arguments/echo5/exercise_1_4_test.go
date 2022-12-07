// Benchmark for and Join ... I need to get to chapter 11!
package main

//!+bench
import (
	"testing"
)

var args = []string{"command", "arg1", "arg2"}

// !+bench
func BenchmarkForLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ForLoop(args)
	}
}

func BenchmarkJoinLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinLoop(args)
	}
}
