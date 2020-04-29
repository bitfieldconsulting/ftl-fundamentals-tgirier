// Package calculator implements calculation functionalities.
package calculator

import "errors"

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Substract returns the result of a minus b.
func Substract(a, b int) int {
	return a - b
}

// Multiply returns the result of a times b.
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the result of a divide b.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("input invalid - division by zero not possible")
	}
	return a / b, nil
}
