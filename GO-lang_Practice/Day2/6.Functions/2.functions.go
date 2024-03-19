package main

import "fmt"

const PI = 3.14159

func main() {
	fmt.Println("Arguments to Functions")
	Greet("Hello")
	Greet("Abhisek")

	add(2, 5)

	fmt.Println("Sum 10, 20: ", add1(10, 20)) //30
	res := add1(add1(10, 20), add1(34, 53))
	fmt.Println(res) //117

	fmt.Println("Multiple return from function")
	a, c := CircleProperty(12.3)
	fmt.Println(a, c) //475.29115110000004 77.283114

	a, _ = CircleProperty(12.3)
	fmt.Println(a) //475.29115110000004

	_, c = CircleProperty(12.3)
	fmt.Println(c) //
}

func Greet(msg string) {
	fmt.Println(msg)
}

func add(a int, b int) {
	fmt.Println(a + b) //7
}

func add1(a, b int) int { //return tyoe is int
	c := a + b
	return c
}

func CircleProperty(radius float64) (float64, float64) {
	area := PI * radius * radius
	circ := 2 * PI * radius

	return area, circ
}
