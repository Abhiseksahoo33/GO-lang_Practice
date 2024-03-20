package math

import (
	"first/geo"
	"fmt"
)

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
