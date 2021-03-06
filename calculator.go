// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Add takes two numbers or more and returns the result of adding them together.
func Add(a, b float64, numbers ...float64) float64 {
	sum := a + b
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Subtract takes two numbers or more and returns the result of subtracting them
// from the first.
func Subtract(a, b float64, numbers ...float64) float64 {
	result := a - b
	for _, n := range numbers {
		result -= n
	}
	return result
}

// Multiply takes two or more numbers and returns the result of multiplying them together.
func Multiply(a, b float64, numbers ...float64) float64 {
	result := a * b
	for _, n := range numbers {
		result *= n
	}
	return result
}

// Divide takes two or more numbers and returns the result of dividing the first
// with the second and the result with the others. If one number following the first one is zero it returns an error.
func Divide(a, b float64, numbers ...float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by 0")
	}
	result := a / b
	for _, n := range numbers {
		if n == 0 {
			return 0, errors.New("division by 0")
		}
		result /= n
	}
	return result, nil
}

// Sqrt takes one number and returns its square root.
// If the number is negative it returns an error.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("negative number")
	}
	return math.Sqrt(a), nil
}

// StringCalc takes a string of the form number one / operator / number two (ex: ‘12 * 3')
// and returns the resutl of the calculus.
func StringCalc(s string) (float64, error) {
	var operatorIndex int
	var operator rune

	for i, r := range s {
		if r == '+' || r == '-' || r == '*' || r == '/' {
			operator = r
			operatorIndex = i
			break
		}
	}

	if operatorIndex == 0 {
		return 0, errors.New("missing supported operator")
	}
	if operatorIndex >= len(s)-1 {
		return 0, errors.New("missing second number after operator")
	}

	s1 := strings.TrimSpace(s[:operatorIndex])
	s2 := strings.TrimSpace(s[operatorIndex+1:])

	a, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		return 0, errors.New("first number is not a floating point number")
	}

	b, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		return 0, errors.New("second number is not a floating point number")
	}

	var result float64

	switch operator {
	case '+':
		result = Add(a, b)
	case '-':
		result = Subtract(a, b)
	case '*':
		result = Multiply(a, b)
	case '/':
		result, err = Divide(a, b)
	}

	return result, err
}
