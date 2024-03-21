package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func Print(msg string) {
	fmt.Println(msg)
	wg.Done()
}

func main() {

	fmt.Println("Demo: wait groups")
	wg.Add(3)
	go Print("This is my first thread")

	go Print("This is my second thread")

	go Print("This is my third go routine")
	wg.Wait()
}
