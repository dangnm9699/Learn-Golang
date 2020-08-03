package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"./tempconv"
)

var usage = `Unit:
	Celsius 	C
	Fahrenheit 	F
	Meter		m
	Feet		ft
	Kilogram	kg
	Pound		lb
	`

func main() {
	r := regexp.MustCompile("(-?[0-9]+)([a-z]+|[A-Z])")
	if len(os.Args) == 1 {
		log.Fatal("Command-line needs more than one argument\n", usage)
	} else {
		for _, arg := range os.Args[1:] {
			unitConversion(arg, r)
		}
	}
}

func unitConversion(arg string, r *regexp.Regexp) {
	result := r.FindStringSubmatch(arg)
	if len(result) != 3 {
		fmt.Println("An error has occurred")
	} else {
		value, err := strconv.ParseFloat(result[1], 64)
		if err != nil {
			fmt.Println("An error has occurred while parsing float")
			return
		}
		switch result[2] {
		case "C":
			fmt.Println(tempconv.CToF(tempconv.Celsius(value)).String())
			break
		case "F":
			fmt.Println(tempconv.FToC(tempconv.Fahrenheit(value)).String())
			break
		case "m":
			fmt.Println(tempconv.MToFt(tempconv.Meter(value)).String())
			break
		case "ft":
			fmt.Println(tempconv.FtToM(tempconv.Feet(value)).String())
			break
		case "kg":
			fmt.Println(tempconv.KgToLb(tempconv.Kilogram(value)).String())
			break
		case "lb":
			fmt.Println(tempconv.LbToKg(tempconv.Pound(value)).String())
			break
		default:
			fmt.Fprintf(os.Stderr, usage)
			return
		}
	}
}
