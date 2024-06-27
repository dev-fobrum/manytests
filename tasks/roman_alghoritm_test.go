package tasks

import (
	"fmt"
	"testing"
)

type isValidateRomanProps struct {
	input    string
	expected bool
}

type executeRomanAlgorithmProps struct {
	input       string
	expected    int
	expectError bool
}

func TestIsValidRoman(t *testing.T) {
	tests := []isValidateRomanProps{
		{"III", true},
		{"IV", true},
		{"IX", true},
		{"LVIII", true},
		{"MCMXCIV", true},
		{"IIII", false},
		{"VV", false},
		{"IC", false},
		{"IM", false},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("For %s expect %t", test.input, test.expected)

		t.Run(testname, func(t *testing.T) {
			result := isValidRoman(test.input)
			if result != test.expected {
				t.Errorf("isValidRoman(%s) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestExecuteRomanAlgorithm(t *testing.T) {
	tests := []executeRomanAlgorithmProps{
		{"III", 3, false},
		{"IV", 4, false},
		{"IX", 9, false},
		{"LVIII", 58, false},
		{"MCMXCIV", 1994, false},
		{"IIII", 0, true},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("For %s expecting %d", test.input, test.expected)

		t.Run(testname, func(t *testing.T) {
			result, err := executeRomanAlghoritm(test.input)
			if test.expectError {
				if err == nil {
					t.Errorf("executeRomanAlghoritm(%s) expected error but got none", test.input)
				}
			} else {
				if err != nil {
					t.Errorf("executeRomanAlghoritm(%s) returned error: %v", test.input, err)
				}
				if result != test.expected {
					t.Errorf("executeRomanAlghoritm(%s) = %d; want %d", test.input, result, test.expected)
				}
			}
		})

	}
}
