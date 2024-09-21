package helpers

import (
	"os"
	"testing"
)

func simulateInput(input string) func() {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	w.Write([]byte(input))
	w.Close()
	os.Stdin = r
	return func() {
		os.Stdin = oldStdin
	}
}

func TestGetRadius(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		hasError bool
	}{
		{"5\n", 5.0, false},
		{"-1\n", 0.0, true},
		{"0\n", 0.0, true},
		{"invalid\n", 0.0, true},
	}

	for _, test := range tests {
		resetInput := simulateInput(test.input)
		defer resetInput()

		result, err := GetRadius()

		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for input %q, got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("Expected %v, got %v for input %q", test.expected, result, test.input)
			}
		}
	}
}

func TestGetQuery(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"1\n", 1, false},
		{"2\n", 2, false},
		{"3\n", 3, false},
		{"4\n", 0, true},
		{"invalid\n", 0, true},
	}

	for _, test := range tests {
		resetInput := simulateInput(test.input)
		defer resetInput()

		result, err := GetQuery()

		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for input %q, got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("Expected %d, got %d for input %q", test.expected, result, test.input)
			}
		}
	}
}
