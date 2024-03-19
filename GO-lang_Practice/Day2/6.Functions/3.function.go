package main

import "fmt"

func RectangleProperty(a, b float64) (area float64, cir float64) {
	area = a * b
	cir = 2 * (a + b)
	return
}

func main() {
	fmt.Println("Named return: ")

	a, c := RectangleProperty(1.0, 3.0)
	fmt.Println(a, c) // 3 8
}
