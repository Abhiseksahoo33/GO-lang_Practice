package main

import "fmt"

func HOFucntion() func() {

	a := func() {
		fmt.Println("This is from function retured from HOF")
	}
	return a
}

func main() {
	fmt.Println("Higher Order Functions: ")

	fn := HOFucntion()

	fn()
	fn()
}
