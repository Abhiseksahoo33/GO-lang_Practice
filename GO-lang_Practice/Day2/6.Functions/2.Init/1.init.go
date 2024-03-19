package main

import "fmt"

func init() {
	fmt.Println("We are init") // first will run , it comes in run time
}

func init() {
	fmt.Println("We are init2") // 2nd
}

func init() {
	fmt.Println("We are init3") // 3rd
}

func init() {
	fmt.Println("We are init4") // 4th
}

func main() {
	fmt.Println("Working with init") // init executes much before main
}

func init() {
	fmt.Println("We are init4") // 5th
}

/*
We are init
We are init2
We are init3
We are init4
We are init4
Working with init
*/
