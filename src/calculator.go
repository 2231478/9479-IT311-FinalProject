package main

import (
	"errors"
	"fmt"
)

// Addition
func Add(a, b float64) float64 {
	return a + b
}

// Subtraction
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiplication
func Multiply(a, b float64) float64 {
	return a * b
}

// Division
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// printResult prints the result of the calculation
func printResult(a float64, operator string, b float64, result float64) {
	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", a, operator, b, result)
}

func main() {
	// Take user input for two numbers
	var a, b float64
	var operator string

	// Get first number
	fmt.Print("Enter the first number: ")
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Error: Invalid input for the first number.")
		return
	}

	// Get operator
	fmt.Print("Enter the operator (+, -, *, /): ")
	if _, err := fmt.Scanln(&operator); err != nil {
		fmt.Println("Error: Invalid input for the operator.")
		return
	}

	// Get second number
	fmt.Print("Enter the second number: ")
	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println("Error: Invalid input for the second number.")
		return
	}

	// Perform the chosen operation and display result
	var result float64
	var err error
	switch operator {
	case "+":
		result = Add(a, b)
	case "-":
		result = Subtract(a, b)
	case "*":
		result = Multiply(a, b)
	case "/":
		result, err = Divide(a, b)
	default:
		fmt.Println("Error: Invalid operator. Use one of +, -, *, /.")
		return
	}

	// Check for division error
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	printResult(a, operator, b, result)
}
