package helpers

import (
	"circle-math/calculators"
	"fmt"
)

// GetPrintResult calculates the result using the provided calculation function
// and prints the formatted result along with a thank you message.
// It takes the radius of the circle and a function that performs the calculation
// as parameters.
func GetPrintResult(radius float64, calcFunction func(float64) float64) {
	result := calcFunction(radius)
	fmt.Printf("Result: %.2f\n", result)
	fmt.Println("Thank you")
}

// GetFunction returns the appropriate calculation function based on the user's query.
// It takes an integer query (1, 2, or 3) and returns a function that calculates
// the area, perimeter, or diameter of a circle. If the query is invalid, it returns nil.
func GetFunction(query int) func(r float64) float64 {
	queryToFunc := map[int]func(r float64) float64{
		1: calculators.Area,
		2: calculators.Perimeter,
		3: calculators.Diameter,
	}
	return queryToFunc[query]
}
