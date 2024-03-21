package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func Generator(gch chan int, elements ...int) {
	defer wg.Done()
	defer close(gch)
	for _, v := range elements {
		gch <- v
	}
}

func Square(gch, dch chan int) {
	defer wg.Done()
	defer close(dch)
	for d := range gch {
		dch <- (d * d)
	}

}

func Display(dch chan int) {
	defer wg.Done()
	for d := range dch {
		fmt.Println(d)
	}

}

func main() {
	fmt.Println("Concurrency Pattern: Pipeline")

	gch := make(chan int)
	dch := make(chan int)
	wg.Add(3)
	go Generator(gch, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	go Square(gch, dch)
	go Display(dch)
	wg.Wait()
}
