package tempconv

// FtoC converts Fahrenheit to Celsius
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KtoC converts Kelvin to Celsius
func KtoC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// CtoF converts Celsius to Fahrenheit
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CtoK converts Celsius to Kelvin
func CtoK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
