package main

import "fmt"

func main() {
	fmt.Println("Working on String Operations")

	a, b := "First String", " Second String"
	//a, b := "First String", ` Second String`  // Allowed
	c := a + b
	fmt.Println(c) //First String Second String

	c = fmt.Sprint(a, b)
	fmt.Println(c) //First String Second String

	e, f := "First String", "Second String"
	g := fmt.Sprintln(e, f)
	fmt.Println(g) //First String Second String  # Add space between strings

}
