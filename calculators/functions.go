package calculators

// Pi is the mathematical constant π (pi) used in calculations.
const Pi = 3.14

// Area calculates the area of a circle given its radius.
// It takes the radius (r) as input and returns the area as a float64.
func Area(r float64) float64 {
	return Pi * r * r
}

// Perimeter calculates the perimeter (circumference) of a circle given its radius.
// It takes the radius (r) as input and returns the perimeter as a float64.
func Perimeter(r float64) float64 {
	return 2 * Pi * r
}

// Diameter calculates the diameter of a circle given its radius.
// It takes the radius (r) as input and returns the diameter as a float64.
func Diameter(r float64) float64 {
	return 2 * r
}
