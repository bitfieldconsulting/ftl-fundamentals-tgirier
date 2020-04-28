package calculator_test

import (
	"calculator"
	"testing"
)

var testCases = []struct {
	a, b int
	sub  int
	mult int
	div  float64
}{
	{a: 1, b: 2, sub: -1, mult: 2, div: 0.5},
	{a: 2, b: 2, sub: 0, mult: 4, div: 1},
	{a: -1, b: 2, sub: -3, mult: -2, div: -0.5},
}

func TestSubtract(t *testing.T) {
	for _, testCase := range testCases {
		want := testCase.sub
		got := calculator.Substract(testCase.a, testCase.b)
		if want != got {
			t.Errorf("Substract(%d,%d): want %d, got %d", testCase.a, testCase.b, want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	for _, testCase := range testCases {
		want := testCase.mult
		got := calculator.Multiply(testCase.a, testCase.b)
		if want != got {
			t.Errorf("Multiply(%d, %d): want %d, got %d", testCase.a, testCase.b, want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	for _, testCase := range testCases {
		want := testCase.div
		got := calculator.Divide(testCase.a, testCase.b)
		if want != got {
			t.Errorf("Divide(%d, %d): want %g, got %g", testCase.a, testCase.b, want, got)
		}
	}
}
