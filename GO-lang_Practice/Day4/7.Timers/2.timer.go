package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Demo: Timer", time.Now())
	wg := sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(5*time.Second, func() {
		defer wg.Done()
		fmt.Println("Timer is expired..")
	})
	//time.NewTimer() // don't prefer
	wg.Wait()
}
