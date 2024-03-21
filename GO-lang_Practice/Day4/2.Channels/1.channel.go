package main

import (
	"fmt"
)

var ch = make(chan int)

func Factorial(num int) {
	fact := 1

	for idx := 1; idx <= num; idx++ {
		fact *= idx
	}
	fmt.Println("Calculated Factorial for ", num, " and writing this to channel")
	ch <- fact
}

func main() {
	fmt.Println("Demo: Channels -> Factorials of 1 .. 10")

	for i := 1; i <= 10; i++ {
		go Factorial(i)
	}

	for i := 0; i < 10; i++ {
		data := <-ch
		fmt.Println(data)
	}
}

/*
Demo: Channels -> Factorials of 1 .. 10
Calculated Factorial for  10  and writing this to channel
3628800
Calculated Factorial for  3  and writing this to channel
6
Calculated Factorial for  5  and writing this to channel
120
Calculated Factorial for  8  and writing this to channel
40320
Calculated Factorial for  2  and writing this to channel
2
Calculated Factorial for  4  and writing this to channel
24
Calculated Factorial for  6  and writing this to channel
720
Calculated Factorial for  1  and writing this to channel
1
Calculated Factorial for  9  and writing this to channel
362880
Calculated Factorial for  7  and writing this to channel
5040
*/
