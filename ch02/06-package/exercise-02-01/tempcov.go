package tempconv

import (
	"fmt"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
)

// String converts a Celsius value to a string with °C appended to it
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// Round converts a Celsius value into a float, with the specified precision
func (c Celsius) Round(precision uint) float64 {
	_fmt := "." + strconv.Itoa(int(precision)) + "f"
	s := fmt.Sprintf(_fmt, c)
	v, _ := strconv.ParseFloat(s, 8)
	return v
}

// String converts a Fahrenheit value to a string with °F appended to it
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// Round converts a Fahrenheit value into a float, with the specified precision
func (f Fahrenheit) Round(precision uint) float64 {
	_fmt := "." + strconv.Itoa(int(precision)) + "f"
	s := fmt.Sprintf(_fmt, f)
	v, _ := strconv.ParseFloat(s, 8)
	return v
}

// String converts a Kelvin value to a string with °k appended to it
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

// Round converts a Fahrenheit Kelvin into a float, with the specified precision
func (k Kelvin) Round(precision uint) float64 {
	_fmt := "." + strconv.Itoa(int(precision)) + "f"
	s := fmt.Sprintf(_fmt, k)
	v, _ := strconv.ParseFloat(s, 8)
	return v
}
