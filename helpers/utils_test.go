package helpers

import (
	"bytes"
	"circle_calculator/calculators"
	"fmt"
	"testing"
)

func TestGetPrintResult(t *testing.T) {
	var buf bytes.Buffer

	printResult := func(radius float64, calcFunction func(float64) float64) {
		result := calcFunction(radius)
		buf.WriteString(fmt.Sprintf("Result: %.2f\n", result))
		buf.WriteString("Thank you\n")
	}

	radius := 5.0
	expectedOutput := "Result: 78.50\nThank you\n"
	printResult(radius, calculators.Area)

	if buf.String() != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, buf.String())
	}

	buf.Reset()

	expectedOutput = "Result: 31.40\nThank you\n"
	printResult(radius, calculators.Perimeter)

	if buf.String() != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, buf.String())
	}
}

func TestGetFunction(t *testing.T) {
	tests := []struct {
		query    int
		expected func(float64) float64
	}{
		{1, calculators.Area},
		{2, calculators.Perimeter},
		{3, calculators.Diameter},
	}

	for _, test := range tests {
		result := GetFunction(test.query)
		if result == nil {
			t.Errorf("Expected a valid function for query %d, but got nil", test.query)
		}
	}

	invalidQuery := 4
	result := GetFunction(invalidQuery)
	if result != nil {
		t.Errorf("Expected nil for invalid query %d, but got a function", invalidQuery)
	}
}
