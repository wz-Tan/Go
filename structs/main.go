package main

import "fmt"

type gasEngine struct {
	gallons uint8
	kph     uint8
	owner
}

type electricEngine struct {
	kilowatts uint8
	kph       uint8
}

type owner struct {
	name string
}

// Interface (Classifies Structs Based off Common Functions)
// Generalisation
type engine interface {
	isTankEmpty() bool
}

func (ge gasEngine) isTankEmpty() bool {
	return ge.gallons <= 0
}

func (ee electricEngine) isTankEmpty() bool {
	return ee.kilowatts <= 0
}

func canReach(e engine) {
	if e.isTankEmpty() {
		fmt.Print("Tank is empty")
	} else {
		fmt.Print("Tank Not Empty")
	}
}

// Assigning Methods to Structs
func (engine gasEngine) distanceLeft() uint8 {
	return engine.gallons * engine.kph
}

func main() {
	var myEngine gasEngine = gasEngine{gallons: 0, kph: 67, owner: owner{"Smith"}}
	fmt.Println("The gas is ", myEngine.gallons)
	fmt.Println("The owner is", myEngine.name) // You Can Directly Access the Struct within Struct

	fmt.Println("the distance left is ", myEngine.distanceLeft())
	canReach(myEngine)
}
