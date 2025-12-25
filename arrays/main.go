package main

import "fmt"

func main() {
	// Initialise the Array
	var intArr [3]int32 = [3]int32{1, 2, 3}
	intArr[0] = 100
	fmt.Println(intArr[0])
	fmt.Println(&intArr[0])

	//Slices - Lists
	var intSlice []int32 = []int32{0, 1, 2}
	fmt.Println(intSlice[1])
	fmt.Println(append(intSlice, 67))

	// Hashmaps
	// String Key with Int value
	var myMap map[string]uint8 = map[string]uint8{"John": 67, "Jane": 68, "Josh": 69}
	delete(myMap, "Jane")

	// For Loops (Use := to assign new variables)
	for key, pair := range myMap {
		println(key, pair)
	}

	for i := 0; i <= 10; i++ {
		println(i)
	}
}
