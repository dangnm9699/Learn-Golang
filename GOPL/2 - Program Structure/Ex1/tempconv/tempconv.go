package tempconv

import (
	"fmt"
)

// Celsius is a float64
type Celsius float64

// Fahrenheit is a float64
type Fahrenheit float64

// Kelvin is a float64
type Kelvin float64

const (
	// AbsoluteZeroC is Absolute Zero in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is Freezing temperature in Celsius
	FreezingC Celsius = 0
	// BoilingC is Boiling temperature in Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%6.2f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%6.2f°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}
