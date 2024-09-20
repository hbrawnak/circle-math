package calculators

const Pi = 3.14

func Area(r float64) float64 {
	return Pi * r * r
}

func Perimeter(r float64) float64 {
	return 2 * Pi * r
}

func Diameter(r float64) float64 {
	return 2 * r
}
