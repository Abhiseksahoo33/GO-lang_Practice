package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func Generator(elements ...int) chan int {
	gch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(gch)
		for _, v := range elements {
			gch <- v
		}
	}()
	return gch
}

func Square(gch chan int) chan int {
	dch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(dch)
		for d := range gch {
			dch <- (d * d)
		}
	}()
	return dch
}

func Display(dch chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for d := range dch {
			fmt.Println(d)
		}
	}()
}

func main() {
	fmt.Println("Concurrency Pattern: Pipeline")

	gch := Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	dch := Square(gch)
	Display(dch)

	//Display(Square(Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)))

	//Display(Square(Square(Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))))
	wg.Wait()
}

/*
1
4
9
16
25
36
49
64
81
100
*/
