package math

import (
	"first/geo"
	"fmt"
)

const PI = 3.14159 //exported that started with upper case it can accessible in outside of package [package global]

const pI = 3.14159 // it can accessible in inside of package [example in circle.go] [global]

func init() {
	fmt.Println("Inside math")
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func MethodFromGeo() {
	fmt.Println(geo.AsianCountries()) //access method of geo from math package
}

func main() {
	fmt.Println("This is my Package and modules demo...")

}
