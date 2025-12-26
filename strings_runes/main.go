package main

import "fmt"

func main() {
	var myString string = "résumé"
	fmt.Println(myString)
	for i, v := range myString {
		// Notice How Special Characters Take 2 Bits
		println(i, v)
	}
	println("The second item is", myString[1])

	// Runes (Unicode)
	var myRune []rune = []rune("résumé")

	// This is Also OK
	// var myCustomRune = []rune("hello")

	// Here it just takes the 2 bits taken by said special character
	println("The second item is", myRune[1])

}
