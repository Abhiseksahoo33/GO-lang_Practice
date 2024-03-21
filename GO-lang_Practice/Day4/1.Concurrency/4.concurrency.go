package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Demo concurrency")
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
		if msg == "hello" {
			fmt.Println("Greeting...")
		} else {
			fmt.Println("Good Bye")
		}
	}("hello")
	wg.Wait()
}

/*
Demo concurrency
hello
Greeting...
*/
