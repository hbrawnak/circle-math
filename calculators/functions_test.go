package calculators

import "testing"

func TestArea(t *testing.T) {
	radius := 4.0
	expected := Pi * radius * radius
	result := Area(radius)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestPerimeter(t *testing.T) {
	radius := 5.0
	expected := 2 * Pi * radius
	result := Perimeter(radius)

	if result != expected {
		t.Errorf("calcPerimeter(%v) = %v; want %v", radius, result, expected)
	}
}

func TestDiameter(t *testing.T) {
	radius := 5.0
	expected := 2 * radius
	result := Diameter(radius)

	if result != expected {
		t.Errorf("calcDiameter(%v) = %v; want %v", radius, result, expected)
	}
}
