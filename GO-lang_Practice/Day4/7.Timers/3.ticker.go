package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Demo: ticker...")

	tkr := time.NewTicker(1 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		//defer wg.Done()
		for {
			_, b := <-tkr.C
			if !b {
				wg.Done()
				break
			}
			fmt.Println("Ticker expired", time.Now())
		}
	}()
	time.AfterFunc(3*time.Second, func() {
		tkr.Stop()
	})
	wg.Wait()
}
