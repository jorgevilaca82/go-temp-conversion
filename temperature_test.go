package main

import (
	"testing"
)

func TestTemperatureConversions(t *testing.T) {
	// Test cases for Celsius to Fahrenheit conversion
	celsiusToFahrenheitTests := []struct {
		input    Celsius
		expected Fahrenheit
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
	}

	for _, test := range celsiusToFahrenheitTests {
		result := ConvertToFahrenheit(test.input)
		if result != test.expected {
			t.Errorf("ConvertToFahrenheit(%v) = %v; want %v", test.input, result, test.expected)
		}
	}

	// Test cases for Fahrenheit to Celsius conversion
	fahrenheitToCelsiusTests := []struct {
		input    Fahrenheit
		expected Celsius
	}{
		{32, 0},
		{212, 100},
		{-40, -40},
	}

	for _, test := range fahrenheitToCelsiusTests {
		result := ConvertToCelsius(test.input)
		if result != test.expected {
			t.Errorf("ConvertToCelsius(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}