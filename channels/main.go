package main

import "fmt"

// Channel is used to store values while being thread safe
// Unbuffered channels can only store one value at a time, whereas buffered channels can hold more

func main() {
	var c = make(chan int, 5) // Create Channel (Add Size parameter for buffered)
	go process(c)
	// Wait for c to get input and spit out as necessary
	for i := range c {
		fmt.Println(i)
	}
}

func process(c chan int) {
	// Inserts Value 0 to 4, One At A Time
	// Insert 0, Wait for Main to Spit Out Current Value in C, then Insert 1
	for i := range 5 {
		c <- i
	}
	close(c) // Close Channel on Completion
}
