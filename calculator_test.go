package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
	"time"
)

type TestCase struct {
	name        string
	a, b        float64
	n           []float64
	want        float64
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two positive numbers", a: 2, b: 2, n: []float64{}, want: 4},
		{name: "Four positive numbers", a: 2, b: 2, n: []float64{1, 2}, want: 7},
		{name: "One negative number", a: -1, b: 0, want: -1},
		{name: "One zero number", a: 0, b: 0, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b, tc.n...)
		if tc.want != got {
			t.Errorf("%s - Add(%f, %f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.n, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < 100; i++ {
		a, b := r.Float64(), r.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Fatalf("Random Add(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two positive numbers", a: 2, b: 1, n: []float64{}, want: 1},
		{name: "Four positive numbers", a: 4, b: 1, n: []float64{1, 1}, want: 1},
		{name: "Two positive numbres which substract to a negative", a: 2, b: 3, want: -1},
		{name: "Two negative numbers which substract to a positive", a: -1, b: -3, want: 2},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b, tc.n...)
		if tc.want != got {
			t.Errorf("%s - Substract(%f, %f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.n, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two positive numbers", a: 2, b: 3, n: []float64{}, want: 6},
		{name: "Four positive numbers", a: 2, b: 3, n: []float64{2, 2}, want: 24},
		{name: "One zero number", a: 0, b: 4, want: 0},
		{name: "One negative number", a: 3, b: -1, want: -3},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b, tc.n...)
		if tc.want != got {
			t.Errorf("%s - Multiply(%f, %f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.n, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two positive numbers", a: 4, b: 2, n: []float64{}, want: 2, errExpected: false},
		{name: "Four positive numbers", a: 8, b: 2, n: []float64{2, 2}, want: 1, errExpected: false},
		{name: "Division by 0", a: 8, b: 0, want: 0, errExpected: true},
		{name: "0 divide by a number", a: 0, b: 1, want: 0, errExpected: false},
		{name: "Negative number", a: -3, b: 1, want: -3, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b, tc.n...)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s - Divide(%f, %f, %f): unexpected error status: %v", tc.name, tc.a, tc.b, tc.n, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("%s - Divide(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Positive Number", a: 144, want: 12, errExpected: false},
		{name: "Negative number", a: -2, want: 0, errExpected: true},
		{name: "Zero", a: 0, want: 0, errExpected: false},
		{name: "Non rational values", a: 2, want: 1.414, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s - Sqrt(%f): unexpected error status: %v", tc.name, tc.a, err)
		}
		if !tc.errExpected && math.Abs(tc.want-got) > 0.01 {
			t.Fatalf("%s - Sqrt(%f): want %f, got %f", tc.name, tc.a, tc.want, got)
		}
	}
}

func TestStringCalc(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		input       string
		want        float64
		errExpected bool
	}{
		{name: "Add", input: "1 + 1.5", want: 2.5, errExpected: false},
		{name: "Substract", input: "100-0.1", want: 99.9, errExpected: false},
		{name: "Multiply", input: "2*2", want: 4, errExpected: false},
		{name: "Divide", input: "18   /   6", want: 3, errExpected: false},
		{name: "Not an operation", input: "19", want: 0, errExpected: true},
		{name: "Not enough numbers", input: "1+", want: 0, errExpected: true},
		{name: "Two operators", input: "1+1/3", want: 0, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.StringCalc(tc.input)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s - StringCalc(%s): unexpected error status: %v", tc.name, tc.input, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Fatalf("%s - StringCalc(%s): want %f, got %f", tc.name, tc.input, tc.want, got)
		}
	}

}
