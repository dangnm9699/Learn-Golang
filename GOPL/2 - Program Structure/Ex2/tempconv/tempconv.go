package tempconv

import (
	"fmt"
)

// Celsius is a float64
type Celsius float64

// Fahrenheit is a float64
type Fahrenheit float64

// Meter is a float64
type Meter float64

// Feet is a float64
type Feet float64

// Pound is a float64
type Pound float64

// Kilogram is a float64
type Kilogram float64

func (c Celsius) String() string {
	return fmt.Sprintf("%6.2f °C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%8.2f °F", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%8.2f m", m)
}

func (ft Feet) String() string {
	return fmt.Sprintf("%8.2f ft", ft)
}

func (lb Pound) String() string {
	return fmt.Sprintf("%8.2f lb", lb)
}

func (kg Kilogram) String() string {
	return fmt.Sprintf("%8.2f kg", kg)
}
