package main

import (
	"circle_calculator/helpers"
	"fmt"
)

func main() {
	radius, err := helpers.GetRadius()
	if err != nil {
		fmt.Println(err)
		return
	}

	query, err := helpers.GetQuery()
	if err != nil {
		fmt.Println(err)
		return
	}

	calcFunc := helpers.GetFunction(query)
	if calcFunc == nil {
		fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		return
	}

	helpers.GetPrintResult(radius, calcFunc)
}
