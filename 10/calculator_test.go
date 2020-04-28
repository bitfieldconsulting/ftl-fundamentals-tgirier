package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 2},
		{a: 2, b: 2, want: 4},
		{a: 6, b: 7, want: 13},
		{a: -1, b: -1, want: -2},
		{a: 0, b: 1, want: 1},
		{a: 0, b: 0, want: 0},
		{a: -1, b: 0, want: -1},
		{a: 3, b: -2, want: 1},
		{a: 6, b: 0, want: 6},
		{a: -3, b: -3, want: -6},
	}
	for _, testCase := range testCases {
		got := calculator.Add(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	var testCases = []struct {
		a, b int
		want int
	}{
		{a: 1, b: 2, want: -1},
		{a: 2, b: 2, want: 0},
		{a: -1, b: 2, want: -3},
	}

	for _, testCase := range testCases {
		want := testCase.want
		got := calculator.Substract(testCase.a, testCase.b)
		if want != got {
			t.Errorf("Substract(%d,%d): want %d, got %d", testCase.a, testCase.b, want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	var testCases = []struct {
		a, b int
		want int
	}{
		{a: 1, b: 2, want: 2},
		{a: 2, b: 2, want: 4},
		{a: -1, b: 2, want: -2},
	}

	for _, testCase := range testCases {
		want := testCase.want
		got := calculator.Multiply(testCase.a, testCase.b)
		if want != got {
			t.Errorf("Multiply(%d, %d): want %d, got %d", testCase.a, testCase.b, want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	var testCases = []struct {
		a, b        int
		want        int
		errExpected bool
	}{
		{a: 1, b: 2, want: 0, errExpected: false},
		{a: 2, b: 2, want: 1, errExpected: false},
		{a: -1, b: 2, want: 0, errExpected: false},
		{a: -2, b: 1, want: -2, errExpected: false},
		{a: 1, b: 0, want: 0, errExpected: true},
	}
	for _, testCase := range testCases {
		want := testCase.want
		got, err := calculator.Divide(testCase.a, testCase.b)
		if want != got {
			t.Errorf("Divide(%d, %d): want %d, got %d", testCase.a, testCase.b, want, got)
		}
		if err != nil && !testCase.errExpected {
			t.Errorf("Divide(%d, %d): error expected %t, error %q", testCase.a, testCase.b, testCase.errExpected, err)
		}
	}
}
