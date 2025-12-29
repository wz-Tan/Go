package main

import "fmt"

type gasEngine struct {
	gallons int
}

type electricEngine struct {
	kilowatt int
}

type car[T gasEngine | electricEngine] struct {
	name   string
	engine T
}

func main() {
	var intSlice []int = []int{1, 2, 3}
	var sum = sumOfSlice(intSlice)
	fmt.Println(sum)

	var e gasEngine = gasEngine{10}

	// Declare T as gasEngine
	var car1 car[gasEngine] = car[gasEngine]{"Car 1", e}
	fmt.Print(car1)
}

// We Can Take in Variables of Different Types, and Return the Same Type Back
// Works with Struct Type as well
func sumOfSlice[T int | float32 | float64](sliceObject []T) T {
	var sum T
	for _, v := range sliceObject {
		sum += v
	}
	return sum
}
