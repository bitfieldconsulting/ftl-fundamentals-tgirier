// Package calculator implements calculation functionalities.
package calculator

// Substract returns the result of a minus b.
func Substract(a int, b int) int {
	return a - b
}

// Multiply returns the result of a times b.
func Multiply(a int, b int) int {
	return a * b
}

// Divide returns the result of a divide b.
func Divide(a int, b int) float64 {
	return float64(a) / float64(b)
}
