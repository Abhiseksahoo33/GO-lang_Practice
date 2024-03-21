package main

import (
	"fmt"
	"sync"
)

/*
generate numbers from 1 to 10
calculate factorial of this ...
*/
/*
func Factorial(num int) int {
	factorial := 1

	for i := 1; i <= num; i++ {
		factorial *= i
	}
	return factorial
}

func main() {
	fmt.Println("Factorials of 1 .. 10")

	for i := 1; i <= 10; i++ {
		fmt.Println(Factorial(i))
	}
}
*/

func Factorial(num int) {
	defer wg.Done()
	factorial := 1

	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println(factorial)
}

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Factorials of 1 .. 10")

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go Factorial(i)
	}
	wg.Wait()
}

/*
o/p -> Random order
Factorials of 1 .. 10
1
2
3628800
720
5040
24
40320
120
362880
6
*/
