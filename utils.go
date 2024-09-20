package main

import "fmt"

func printResult(radius float64, calcFunction func(float64) float64) {
	result := calcFunction(radius)
	fmt.Printf("Result: %.2f\n", result)
	fmt.Println("Thank you")
}

func getFunction(query int) func(r float64) float64 {
	queryToFunc := map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return queryToFunc[query]
}
