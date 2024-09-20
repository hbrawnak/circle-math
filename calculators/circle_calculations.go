package calculators

const Pi = 3.14

func CalcArea(r float64) float64 {
	return Pi * r * r
}

func CalcPerimeter(r float64) float64 {
	return 2 * Pi * r
}

func CalcDiameter(r float64) float64 {
	return 2 * r
}
