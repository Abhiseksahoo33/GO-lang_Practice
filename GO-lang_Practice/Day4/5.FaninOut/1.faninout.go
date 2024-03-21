package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const FanoutCount = 3

func Generator(list ...string) chan string {
	gch := make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(gch)
		for _, v := range list {
			gch <- v
		}
	}()
	return gch
}

func Process(gch chan string) chan string {
	dch := make(chan string)
	fwg := sync.WaitGroup{}
	wg.Add(1)

	for i := 0; i < FanoutCount; i++ {
		fwg.Add(1)
		go func() {
			defer fwg.Done()
			// defer wg.Done()
			// defer close(dch)
			for str := range gch {
				rev := ""
				for _, c := range str {
					rev = string(c) + rev
				}
				for i := 0; i < math.MaxUint32; i++ {

				}
				dch <- rev
			}
		}()
	}
	go func() {
		defer wg.Done()
		defer close(dch)
		fwg.Wait()
	}()

	return dch
}

func Display(dch chan string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for str := range dch {
			fmt.Println(str)
		}
	}()
}

func main() {
	defer func(t time.Time) {
		fmt.Println("Execution time", time.Since(t))
	}(time.Now())

	fmt.Println("Demo: Fanout FanIn")
	Display(Process(Generator("HII", "Hello", "Good Day", "Abhisek")))
	wg.Wait()
}
