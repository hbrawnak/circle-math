package helpers

import (
	"circle_calculator/calculators"
	"fmt"
)

func GetPrintResult(radius float64, calcFunction func(float64) float64) {
	result := calcFunction(radius)
	fmt.Printf("Result: %.2f\n", result)
	fmt.Println("Thank you")
}

func GetFunction(query int) func(r float64) float64 {
	queryToFunc := map[int]func(r float64) float64{
		1: calculators.CalcArea,
		2: calculators.CalcPerimeter,
		3: calculators.CalcDiameter,
	}
	return queryToFunc[query]
}
