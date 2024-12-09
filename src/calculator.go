package main

import (
	"errors"
	"fmt"
	"os"
	// "strconv"
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

func main() {
	// Take user input for two numbers
	var a, b float64
	var operator string

	// Get first number
	fmt.Print("Enter the first number: ")
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Invalid input for the first number:", err)
		return
	}

	// Get operator
	fmt.Print("Enter the operator (+, -, *, /): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Get second number
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Perform the chosen operation and display result
	var result float64
	switch operator {
	case "*":
		result = Multiply(a, b)
		
	case "-":
		result = Subtract(a, b)

	case "+":
		result = Add(a, b)

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
