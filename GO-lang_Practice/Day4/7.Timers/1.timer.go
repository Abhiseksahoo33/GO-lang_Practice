package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Demo: Timer", time.Now())

	ch := time.After(5 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ch
		fmt.Println("Timer expired", time.Now())
	}()
	fmt.Println("Timer is expired..", time.Now())
	wg.Wait()
}
