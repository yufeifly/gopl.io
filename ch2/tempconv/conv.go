package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CTOF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5+32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FTOC(f Fahrenheit) Celsius {
	return Celsius((f-32)*5/9)
}