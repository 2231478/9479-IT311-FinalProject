package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a, b := 1.0, 1.0
	expected := 2.0
	result := Add(a, b)
	if result != expected {
		t.Errorf("Add(%.2f, %.2f) = %.2f; want %.2f", a, b, result, expected)
	} else {
		t.Logf("Add(%.2f, %.2f) = %.2f; PASS", a, b, result)
	}
}

func TestSubtract(t *testing.T) {
	a, b := 10.0, 4.0
	expected := 6.0
	result := Subtract(a, b)
	if result != expected {
		t.Errorf("Subtract(%.2f, %.2f) = %.2f; want %.2f", a, b, result, expected)
	} else {
		t.Logf("Subtract(%.2f, %.2f) = %.2f; PASS", a, b, result)
	}
}

func TestMultiply(t *testing.T) {
	a, b := 7.0, 6.0
	expected := 42.0
	result := Multiply(a, b)
	if result != expected {
		t.Errorf("Multiply(%.2f, %.2f) = %.2f; want %.2f", a, b, result, expected)
	} else {
		t.Logf("Multiply(%.2f, %.2f) = %.2f; PASS", a, b, result)
	}
}

func TestDivide(t *testing.T) {
	a, b := 15.0, 3.0
	expected := 5.0
	result, err := Divide(a, b)
	if err != nil {
		t.Errorf("Divide(%.2f, %.2f) returned an error: %v", a, b, err)
	} else if result != expected {
		t.Errorf("Divide(%.2f, %.2f) = %.2f; want %.2f", a, b, result, expected)
	} else {
		t.Logf("Divide(%.2f, %.2f) = %.2f; PASS", a, b, result)
	}

	// Test division by zero
	_, err = Divide(15, 0)
	if err == nil {
		t.Error("Divide(15, 0) did not return an error")
	} else {
		t.Logf("Divide(15, 0) returned expected error: %v; PASS", err)
	}
}
