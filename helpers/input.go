package helpers

import (
	"errors"
	"fmt"
)

func GetRadius() (float64, error) {
	var radius float64
	fmt.Println("Enter radius of circle: ")
	if _, err := fmt.Scan(&radius); err != nil || radius <= 0 {
		return 0, errors.New("err: invalid radius. Please enter a positive number")
	}
	return radius, nil
}

func GetQuery() (int, error) {
	var query int
	fmt.Println("Enter choice: \n 1 - Area \n 2 - Perimeter \n 3 - Diameter: ")
	if _, err := fmt.Scan(&query); err != nil || query < 1 || query > 3 {
		return 0, errors.New("err: invalid choice. Please enter 1, 2, or 3")
	}
	return query, nil
}
