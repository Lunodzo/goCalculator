/*
This is a simple GO calculator to operate on two numbers.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func main() {
	// Read the input from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the first number: ")
	number1, _ := reader.ReadString('\n')
	fmt.Println("Enter the second number: ")
	number2, _ := reader.ReadString('\n')

	// Convert the input strings to integers
	num1, _ := strconv.Atoi(strings.TrimSpace(number1))
	num2, _ := strconv.Atoi(strings.TrimSpace(number2))

	// Choose the operation
	fmt.Println("Choose the operation: ")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	operation, _ := reader.ReadString('\n')
	op, _ := strconv.Atoi(strings.TrimSpace(operation))

	// Perform the operation
	// You can choose to loop this part to allow the user to perform multiple operations
	switch op {
	case 1:
		fmt.Println("The sum is: ", add(num1, num2))
	case 2:
		fmt.Println("The difference is: ", sub(num1, num2))
	case 3:
		fmt.Println("The product is: ", mul(num1, num2))
	case 4:
		fmt.Println("The quotient is: ", div(num1, num2))
	default:
		fmt.Println("Invalid operation")
	}
}
