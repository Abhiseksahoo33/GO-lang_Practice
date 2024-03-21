package main

import (
	"fmt"
	"sync"
)

/*
http api
kafka

generator 1
				Processor Display
generator 3
*/

func Generator(wg *sync.WaitGroup, nums ...int) chan int {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}

func Square(wg *sync.WaitGroup, src1, src2 chan int) chan int {
	wg.Add(1)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		defer close(ch)
		src1State, src2state := true, true
		d := 0
		for {
			select {
			case d, src1State = <-src1:
				ch <- d * d
			case d, src2state = <-src2:
				ch <- d * d
			}
			if !src2state && !src1State {
				break
			}
		}

	}()
	return ch
}

func Display(wg *sync.WaitGroup, ch chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

func main() {
	fmt.Println("Demo Select")
	wg := sync.WaitGroup{}

	Display(&wg, Square(&wg, Generator(&wg, 1, 2, 3, 4, 5, 6, 7, 8, 10), Generator(&wg, 11, 12, 13, 14, 15, 16, 17, 18)))
	wg.Wait()
}
