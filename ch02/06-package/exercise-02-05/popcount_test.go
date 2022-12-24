package popcount

import (
	"testing"
)

// TestPopCount_0 validates that PopCount(0) is 1
func TestPopCount_0(t *testing.T) {
	testFor := uint(0)
	expect := 0
	got := PopCount(testFor)
	if got != expect {
		t.Errorf("PopCount(%d) = %d; expect %d", testFor, got, expect)
	}
}

// TestPopCount_3 validates that PopCount(3) is 2
func TestPopCount_31(t *testing.T) {
	testFor := uint(31)
	expect := 5
	got := PopCount(testFor)
	if got != expect {
		t.Errorf("PopCount(%d) = %d; expect %d", testFor, got, expect)
	}
}

// TestPopCount_1023 validates that PopCount(1023) is 2
func TestPopCount_1023(t *testing.T) {
	testFor := uint(1023)
	expect := 10
	got := PopCount(testFor)
	if got != expect {
		t.Errorf("PopCount(%d) = %d; expect %d", testFor, got, expect)
	}
}
