package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32 = 10
	*p = i // Dereference for Value
	p = &i // Store i Address into P
	fmt.Println(*p)

	// For i := range automatically dereferences items
}
