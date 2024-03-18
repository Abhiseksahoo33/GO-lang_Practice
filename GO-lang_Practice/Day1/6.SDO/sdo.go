package main

import (
	"fmt"
	"reflect"
)

//x := 100  ERROR

func main() {
	fmt.Println("DEMO: Short declartion operators") // only local variables not global
	var a, b = 10, 20
	res := a + b
	fmt.Println(res)

	add, sub, mul, div := a+b, a-b, a*b, a/b
	fmt.Println(add, sub, mul, div) //30 -10 200 0

	fmt.Println("DEMO2: SDO")
	c, d := 1, "hi"
	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(d)) //int string
}
