package main

import "fmt"

/*
func main() {
	fmt.Println("Defer...")

	defer fmt.Println("We are calling fucntion defered") // This will be defer untill all the lines in the functions executed

	fmt.Println("After defer")
}
O/P
Defer...
After defer
We are calling fucntion defered

func main() {
	fmt.Println("Exploring Defer: ")
	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	fmt.Println("Exiting defer")
}
O/P
Exploring Defer:
Exiting defer
2nd defer
1st defer


func Test() {
	fmt.Println("Inside Test")
	defer fmt.Println("Calling defer from Test")
	fmt.Println("Exiting Test")
}

func main() {
	fmt.Println("Exploring defer")
	Test()
	fmt.Println("Exiting main")
}

 O/P
Exploring defer
Inside Test
Exiting Test
Calling defer from Test
Exiting main

*/

func Test() {
	fmt.Println("Inside Test")
	defer fmt.Println("Calling defer from Test")
	fmt.Println("Exiting Test")
}

func main() {
	fmt.Println("Exploring defer")
	defer Test()
	fmt.Println("Exiting main")
}

/*
O/p
Exploring defer
Exiting main
Inside Test
Exiting Test
Calling defer from Test
*/
