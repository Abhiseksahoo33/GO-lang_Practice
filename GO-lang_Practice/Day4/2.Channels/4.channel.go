package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int, 3)
var wg = sync.WaitGroup{}
var wg1 = sync.WaitGroup{}

func Factorial(num int) {
	defer wg.Done()

	fact := 1
	for idx := 1; idx <= num; idx++ {
		fact *= idx
	}
	ch <- fact
}

func ChannelMonitor() {
	go func() {
		defer wg1.Done()
		for i := 0; i < 100; i++ {
			fmt.Println(cap(ch), len(ch))
			time.Sleep(1 * time.Microsecond)
		}
	}()
}

func main() {

	defer func(t time.Time) {
		fmt.Println("Took ", time.Since(t))
	}(time.Now())

	fmt.Println("Demo: Channels -> Factorials of 1 .. 10")
	wg1.Add(1)
	ChannelMonitor()
	go func() {
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go Factorial(i)
		}
		wg.Wait()
		close(ch)
	}()

	wg1.Wait()
	for i := 0; i < 10; i++ {
		_, ok := <-ch
		if !ok {
			break
		}
		//fmt.Println(data)
	}

}
