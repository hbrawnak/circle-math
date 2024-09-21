package helpers

import (
	"errors"
	"fmt"
)

// GetRadius prompts the user to enter the radius of a circle.
// It reads the input and checks if it's a valid positive number.
// Returns the radius as a float64 and an error if the input is invalid.
func GetRadius() (float64, error) {
	var radius float64
	fmt.Println("Enter radius of circle: ")
	if _, err := fmt.Scan(&radius); err != nil || radius <= 0 {
		return 0, errors.New("err: invalid radius. Please enter a positive number")
	}
	return radius, nil
}

// GetQuery prompts the user to enter a choice for the calculation type.
// It reads the input and checks if it's a valid option (1, 2, or 3).
// Returns the choice as an int and an error if the input is invalid.
func GetQuery() (int, error) {
	var query int
	fmt.Println("Enter choice: \n 1 - Area \n 2 - Perimeter \n 3 - Diameter: ")
	if _, err := fmt.Scan(&query); err != nil || query < 1 || query > 3 {
		return 0, errors.New("err: invalid choice. Please enter 1, 2, or 3")
	}
	return query, nil
}
