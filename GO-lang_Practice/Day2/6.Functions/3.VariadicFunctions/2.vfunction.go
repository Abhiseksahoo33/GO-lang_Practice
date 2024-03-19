package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int, elem ...int) int {
	sum := a + b
	fmt.Println(reflect.TypeOf(elem)) //[]int
	for _, v := range elem {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("Variadic fucntions")

	fmt.Println(Add(1, 3, 4, 5, 6, 6, 6))
	fmt.Println(Add(1, 2))
	//fmt.Println(Add(1))  not enough arguments in call to Add
	//fmt.Println(Add())  not enough arguments in call to Add
}
