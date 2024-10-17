package romanconv

import (
	"fmt"
)

type decimalToRoman struct {
	decimal int
	roman   string
}

// Decimal to Roman numeral equivalence for 10's, 5's and 4's
var decimalToRomanTable = []decimalToRoman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// IntToRoman maps a decimal value to its equivalent Roman numeral string
func IntToRoman(value int) (string, error) {
	var roman string

	if value <= 0 || value > 3999 {
		return "", fmt.Errorf("Input out of range")
	}

	// Find the next candidate value
	for i := 0; value > 0 && i < len(decimalToRomanTable); i++ {
		decimalPart := decimalToRomanTable[i].decimal

		// Loop until value being converted is less than the candidate decimalPart
		for decimalPart <= value {
			value -= decimalPart
			roman += decimalToRomanTable[i].roman
		}
	}
	return roman, nil
}
