package main

import "fmt"

type Celsius float64
type Fahrenheit float64

// Implementing the Stringer interface for both types
func (c Celsius) String() string {
	return fmt.Sprintf("%.2f °C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f °F", f)
}

// Conversion functions
func ConvertToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func ConvertToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Main function to demonstrate the usage of the types and conversion functions
func main() {
	var tempC Celsius
	fmt.Print("Enter temperature in Celsius: ")
	fmt.Scanln(&tempC)

	fahrenheit := ConvertToFahrenheit(tempC)
	fmt.Printf("Temperature in Fahrenheit: %s\n", fahrenheit)

	// Example of converting back to Celsius
	var tempF Fahrenheit
	fmt.Print("Enter temperature in Fahrenheit: ")
	fmt.Scanln(&tempF)

	celsius := ConvertToCelsius(tempF)
	fmt.Printf("Temperature in Celsius: %s\n", celsius)
}
