package main

import (
	"fmt"
	"sync"
)

/*
Example:1

var wg sync.WaitGroup = sync.WaitGroup{}

func Print(msg string) {
	wg.Add(1)
	fmt.Println(msg)
	wg.Done()
}

func main() {
	fmt.Println("Demo: wait groups")
	//wg.Add(1)
	go Print("This is my first thread")
	//wg.Add(1)
	go Print("This is my second thread")
	//wg.Add(1)
	go Print("This is my third go routine")
	wg.Wait()
}
*/

// Example 2:

var wg sync.WaitGroup = sync.WaitGroup{}

func Print(msg string) {
	fmt.Println(msg)
	wg.Done()
}

func SchedulePrint(msg string) {
	wg.Add(1)
	go Print(msg)
}

func main() {
	fmt.Println("Demo: wait groups")
	SchedulePrint("This is my first thread")
	SchedulePrint("This is my second thread")
	SchedulePrint("This is my third go routine")
	wg.Wait()
}

/*
Demo: wait groups
This is my third go routine
This is my first thread
This is my second thread
*/
