package main

import "fmt"

func main() {
	var a int
	fmt.Print("Please Enter a number:")
	fmt.Scanf("%d", &a)

	if a < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("greater than 10")
	}

}
