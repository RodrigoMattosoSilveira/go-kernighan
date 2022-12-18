package tempconv

import (
	"fmt"
	"testing"
)

// TestCtoF_BoilingC validates that CtoF(100) is 212
func TestCtoF_BoilingC(t *testing.T) {
	testFor := BoilingC
	expect := Fahrenheit(212.0)
	got := CtoF(testFor)
	if got != expect {
		t.Errorf("CtoF(%f) = %f; expect %f", testFor, got, expect)
	}
}

// TestCtoF_FreezingC validates that CtoF(100) is 212
func TestCtoF_FreezingC(t *testing.T) {
	testFor := FreezingC
	expect := Fahrenheit(32).Round(2)
	got := CtoF(testFor).Round(2)
	if got != expect {
		t.Errorf("CtoF(%f) = %f; expect %f", testFor, got, expect)
	}
}

// TestCtoF_AbsoluteC validates that CtoF(100) is 212
func TestCtoF_AbsoluteC(t *testing.T) {
	testFor := AbsoluteZeroC
	expect := Fahrenheit(-459.67).Round(2)
	got := CtoF(testFor).Round(2)
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

// TestKtoC_500 validates that KtoC(500) is 226.85
func TestKtoC_500(t *testing.T) {
	testFor := Kelvin(500)
	expect := Celsius(226.85).Round(2)
	got := KtoC(testFor).Round(2)
	if got != expect {
		t.Errorf("KtoC(%f) = %f; expect %f", testFor, got, expect)
	}
}

// TestCtoK_100 validates that CtoK(100) is 0
func TestCtoK_100(t *testing.T) {
	testFor := Celsius(100)
	expect := Kelvin(326.85).Round(2)
	got := CtoK(testFor).Round(2)
	if got != expect {
		t.Errorf("CtoK(%f) = %f; expect %f", testFor, got, expect)
	}
}
