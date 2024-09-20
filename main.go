package main

import (
	"fmt"
)

func main() {
	radius, err := getRadius()
	if err != nil {
		fmt.Println(err)
		return
	}

	query, err := getQuery()
	if err != nil {
		fmt.Println(err)
		return
	}

	calcFunc := getFunction(query)
	if calcFunc == nil {
		fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		return
	}

	printResult(radius, calcFunc)
}
