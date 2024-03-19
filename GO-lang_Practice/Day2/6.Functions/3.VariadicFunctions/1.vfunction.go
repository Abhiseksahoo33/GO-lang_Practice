package main

import "fmt"

func Add(elem ...int) int {
	sum := 0
	for _, v := range elem {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("Variadic Functions")
	fmt.Println(Add(1, 2, 3, 4, 54, 5, 6, 6, 7, 6)) //94
	fmt.Println(Add(1, 2, 3))                       //6
	fmt.Println(Add(1))                             //1
	fmt.Println(Add())                              //0
}
