package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	fmt.Printf("Absolute zero in Celsius is %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("0°C = %s\n", tempconv.CToF(tempconv.Celsius(0)).String())
	fmt.Printf("0°F = %s\n", tempconv.FToC(tempconv.Fahrenheit(0)).String())
}
