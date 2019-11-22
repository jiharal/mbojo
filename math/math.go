package math

// Create a Rectangle Object
// This function will return value of `area` and `perimeter`
func Rectangle(a, b float64) (float64, float64) {
	area := a * b
	perimeter := 2 * (a + b)
	return area, perimeter
}
