package tempconv

import (
	"fmt"
	"testing"
)

// TestCtoF_0 validates that CtoF(100) is 212
func TestCtoF_0(t *testing.T) {
	testFor := BoilingC
	expect := Fahrenheit(212.0)
	got := CtoF(testFor)
	if got != expect {
		t.Errorf("CtoF(%f) = %f; expect %f", testFor, got, expect)
	}
}

// TestCtoF_100 validates that CtoF(100) is 212
func TestCtoF_100(t *testing.T) {
	testFor := FreezingC
	expect := Fahrenheit(32)
	got := CtoF(testFor)
	if got != expect {
		t.Errorf("CtoF(%f) = %f; expect %f", testFor, got, expect)
	}
}

// TestCToString validates the string conversion works
func TestCelsius_String(t *testing.T) {
	testFor := Celsius(20.0)
	expect := fmt.Sprintf("%g°C", testFor)
	got := testFor.String()
	if got != expect {
		t.Errorf("CToString(%f) = %s; expect %s", testFor, got, expect)
	}
}

// TestFtoC_100 validates that FtoC(32) is 0
func TestFtoC_32(t *testing.T) {
	testFor := Fahrenheit(32)
	expect := Celsius(0)
	got := FtoC(testFor)
	if got != expect {
		t.Errorf("FtoC(%f) = %f; expect %f", testFor, got, expect)
	}
}

// TestFtoC_212 validates that FtoC(32) is 0
func TestFtoC_0(t *testing.T) {
	testFor := Fahrenheit(212)
	expect := Celsius(100)
	got := FtoC(testFor)
	if got != expect {
		t.Errorf("FtoC(%f) = %f; expect %f", testFor, got, expect)
	}
}