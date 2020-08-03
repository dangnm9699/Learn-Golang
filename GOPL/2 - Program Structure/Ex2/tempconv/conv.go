package tempconv

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FtToM converts Feet to Meters
func FtToM(ft Feet) Meter {
	return Meter(ft / 3.2808)
}

// MToFt converts Meters to Feet
func MToFt(m Meter) Feet {
	return Feet(m * 3.2808)
}

// LbToKg converts Pounds to Kilograms
func LbToKg(lb Pound) Kilogram {
	return Kilogram(lb * 0.45359237)
}

// KgToLb converts Kilograms to Pounds
func KgToLb(kg Kilogram) Pound {
	return Pound(kg / 0.45359237)
}
