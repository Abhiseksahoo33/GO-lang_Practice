/*
perform multiple operation at the same time in random order.
parallelism is the ability to execute multiple computations simultaniously

OS schedules threads not processes
Threads can run concurrently or in  parrallel(2 threads can be in same core or diff core)
Max No. of threads = main memory size / size of stack

Treads communicate by sharing memory -> shared threads
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Exploring concurrency in GO..")
	go fmt.Println("This is my go routine print")

	fmt.Println("Exiting main")

	go fmt.Println("This is my second go routine..")
	time.Sleep(1 * time.Nanosecond) //This is my go routine print
}
