package main

import (
	"errors"
	"testing"
)

func Addition(a, b float64) float64 {
	return a + b
}

func Subtraction(a, b float64) float64 {
	return a - b
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func Multiplication(a, b float64) float64 {
	return a * b
}

func TestAddition(t *testing.T) {
	result := Addition(3, 4)
	if result != 7 {
		t.Errorf("Addition(3,4) = %f; expected 7", result)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(7, 4)
	if result != 3 {
		t.Errorf("Subtraction(7,4) = %f; expected 3", result)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(6, 3)
	if err != nil || result != 2 {
		t.Errorf("Division(6,3) = %f, %v; expected 2, <nil>", result, err)
	}

	_, err = Division(6, 0)
	if err == nil {
		t.Error("Division(6,0) did not return an error")
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(3, 4)
	if result != 12 {
		t.Errorf("Multiplication(3,4) = %f; expected 12", result)
	}
}
