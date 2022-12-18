package popcount

import (
	"testing"
)

// TestPopCount_0 validates that PopCount(0) is 1
func TestPopCount_0(t *testing.T) {
	testFor := uint64(0)
	expect := 0
	got := PopCount(testFor)
	if got != expect {
		t.Errorf("PopCount(%d) = %d; expect %d", testFor, got, expect)
	}
}

// TestPopCount_3 validates that PopCount(3) is 2
func TestPopCount_3(t *testing.T) {
	testFor := uint64(3)
	expect := 2
	got := PopCount(testFor)
	if got != expect {
		t.Errorf("PopCount(%d) = %d; expect %d", testFor, got, expect)
	}
}
