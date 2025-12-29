package main

import (
	"fmt"
	"sync"
	"time"
)

// Sample Data
var list []int16 = []int16{1, 2, 3, 4, 5}
var copiedList []int16 = []int16{}

// Mutex
var m = sync.Mutex{}

// Wait Group (Await All) -> Add to It, Then Wait for All to Finish
var waitGroup = sync.WaitGroup{}

// Concurrency CAN be parallel execution, but it refers to moving on to other tasks while waiting for one to finish
func main() {
	var currentTime = time.Now()
	for i := range len(list) {
		waitGroup.Add(1)
		go retrieveNumber(i) // Sets To Concurrent, But Does Not Wait
	}
	waitGroup.Wait()
	fmt.Println("Final copied list is ", copiedList)
	fmt.Println("Time taken in total is ", time.Since(currentTime))

}

func retrieveNumber(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The number in question is ", list[i])
	m.Lock()
	copiedList = append(copiedList, list[i])
	m.Unlock()
	waitGroup.Done() // Minus One Index from Weight Group
}
