package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int, 10)
var wg = sync.WaitGroup{}

func Factorial(num int) {
	defer wg.Done()

	fact := 1
	for idx := 1; idx <= num; idx++ {
		fact *= idx
	}
	ch <- fact
}

func main() {

	defer func(t time.Time) {
		fmt.Println("Took ", time.Since(t))
	}(time.Now())

	fmt.Println("Demo: Channels -> Factorials of 1 .. 10")

	go func() {
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go Factorial(i)
		}
		wg.Wait()
		close(ch)
	}()

	for i := 0; i < 10; i++ {
		data, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(data)
	}

	// for v := range ch {
	// 	fmt.Println(v)
	// }
}
