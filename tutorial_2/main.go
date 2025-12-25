package main

import "fmt"

func main() {
	printName("Josh")
	var num1 int = 2
	var num2 int = 3
	var result int = add(num1, num2)
	fmt.Print("The result of 2 plus 3 is", result)
}

func printName(name string) {
	fmt.Println("Hello", name)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}
