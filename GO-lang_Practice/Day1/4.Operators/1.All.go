package main

import "fmt"

func main() {

	fmt.Println("Demo: Arithmetic Operators")
	var a, b = 100, 200
	var addition = a + b
	var sub = b - b
	var mul = a * b
	var div = a / b
	var rem = a % b
	fmt.Println("Addition, Substraction, multipication, division, reminder of a & b ", addition, sub, mul, div, rem)

	fmt.Println("Demo: Comparison Operators")
	var c, d = 10, 20
	var res = c == d
	fmt.Println(res) // false
	var res1 = c != d
	fmt.Println(res1) //true
	var res2 = c < d
	fmt.Println(res2) //true
	var res3 = c > d
	fmt.Println(res3) //false
	var res4 = c <= d
	fmt.Println(res4) //true
	var res5 = c >= d
	fmt.Println(res5) //false

	fmt.Println("Demo: Logical Operators")
	// && || !
	var e, f, g, h = 10, 20, 30, 40
	var res6 = e != f && g < h // precedence is left to right
	fmt.Println(res6)          //true [true && true = true]

	fmt.Println("Demo: Assignment Operators")
	// = += -= *= &= |= %=
	var i, j = 10, 20
	i += j
	fmt.Println(i) //30

	fmt.Println("Demo: Bitwise Operators")

	fmt.Println("Demo: short Declaration Operators")
}
