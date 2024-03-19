package main

import "fmt"

/*
func main() {
	fmt.Println("Annonymous Functions")

	func() { // function block
		fmt.Println("Inside annonymous")
	}() // function block

	fmt.Println("Exiting Functions")
}

o/p
Annonymous Functions
Inside annonymous
Exiting Functions



func main() {
	fmt.Println("Demo: Annonymous Func")

	func(msg string) {
		fmt.Println("Printing msg", msg)
	}("Hello")

	fmt.Println("Exiting main")
}

o/p
Demo: Annonymous Func
Printing msg Hello
Exiting main

*/

var (
	a = func(msg string) {
		fmt.Println(msg)
	}
)

func main() {
	fmt.Println("Demo:")

	fn := func(msg string) {
		fmt.Println(msg)
	}

	fn("Hi")
	fn("Hello")
	a("Abhisek")
	fmt.Println("Exit")
}

/*
Demo:
Hi
Hello
Abhisek
Exit
*/
