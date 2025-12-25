package main

import (
	"errors"
	"fmt"
)

func main() {
	printName("Josh")
	var num1 int = 2
	var num2 int = 3
	var result int = add(num1, num2)
	var err = sampleError()
	if err != nil {
		fmt.Print("This is an error", err.Error())
	}

	fmt.Printf("The result of %v plus %v is %vc ", num1, num2, result)
}

func printName(name string) {
	fmt.Println("Hello", name)
}

// Place Return Type after Parentheses, use a bracket for multiple values (int,int)
func add(num1 int, num2 int) int {
	var sum int = num1 + num2
	if sum > 2 {
		fmt.Print("Sum is greater than 2 ")
	}
	return num1 + num2
}

func sampleError() error {
	var customErr error
	customErr = errors.New("This is an error")
	return customErr
}
