package main

import (
	"flag"
	"fmt"
)

var temp = CelsiusFlag("temp", 20, "температура")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}


// Определение своего флага
type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = fahrenheitToCelsius(Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("incorrect temp %s", unit)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
