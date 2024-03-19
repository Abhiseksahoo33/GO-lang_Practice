package main

import "fmt"

/*
func main() {
	fmt.Println("Announymous with defer")

	i := 0

	defer func(in int) {
		fmt.Println(in) //0
	}(i)

	i = 100

	fmt.Println("We are exiting main")
}
*/

func main() {
	fmt.Println("Announymous with defer")

	i := 0

	defer func(in *int) {
		fmt.Println(*in) //100
	}(&i)

	i = 100

	fmt.Println("We are exiting main")
}
