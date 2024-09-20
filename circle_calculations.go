package main

const Pi = 3.14

func calcArea(r float64) float64 {
	return Pi * r * r
}

func calcPerimeter(r float64) float64 {
	return 2 * Pi * r
}

func calcDiameter(r float64) float64 {
	return 2 * r
}
