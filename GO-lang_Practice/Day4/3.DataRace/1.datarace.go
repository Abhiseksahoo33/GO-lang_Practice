package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var mutex = sync.Mutex{}

func IncrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}
func main() {
	defer func(t time.Time) {
		fmt.Println("MAin func execution time:", time.Since(t))
	}(time.Now())

	fmt.Println("demo: Data race | Mutex")

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go IncrementCounter(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
