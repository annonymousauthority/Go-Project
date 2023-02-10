package arithmetic

import (
	"errors"
	"log"
	"math"
)

func PerformArithmetic(a int32, b int32) (int32, int32, int32, int32) {
	/*
		the goal of this function is to perform different
		arithmetic operations ==>
		Addition(+)
		Subtraction (-)
		Division (/)
		Multiplication (*)
	*/

	var sum, mul, sub int32
	log.SetPrefix("Result: ")
	log.SetFlags(0)
	sum = addition(a, b)
	sub = subtraction(a, b)
	div, err := division(a, b)
	mul = multiply(a, b)
	if err != nil {
		div = 0
	}
	return sum, mul, div, sub
}

func addition(a int32, b int32) int32 {
	sum := a + b
	return sum
}

func subtraction(a int32, b int32) int32 {
	sub := a - b
	return sub
}

func division(a int32, b int32) (int32, error) {
	// Assumption is b is a non-zero number
	// Update will be a checker to make sure division is carried out when b is non-zero.
	switch b {
	case 0:
		return b, errors.New("cannot divide by zero")
	case math.MaxInt32:
		return b, errors.New("Number exceeds maximum int32 values")
	case math.MinInt32:
		return b, errors.New("Number is less than the minimum int32 values")
	}

	switch a {
	case math.MaxInt32:
		return a, errors.New("Number exceeds maximum int32 values")
	case math.MinInt32:
		return a, errors.New("Number is less than the minimum int32 values")
	}
	divide := a / b
	return divide, nil
}

func multiply(a int32, b int32) int32 {
	mul := a * b
	return mul
}
