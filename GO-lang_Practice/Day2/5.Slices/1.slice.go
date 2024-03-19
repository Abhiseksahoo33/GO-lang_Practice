package main

import "fmt"

// type Slice struct {
// 	ptr      *[cap]array // expandable & sinkable // size of array is capacity
// 	len, cap int
// }

func main() {
	fmt.Println("Exploring Slice: ")          // basically a structure
	var slice []int = []int{1, 2, 3, 4, 5, 6} // if size is not given it is a slice
	fmt.Println(slice)                        //[1 2 3 4 5 6]

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println("Length of slice: ", len(slice)) //6

	fmt.Println("Demo Slice: ")
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice1)
	fmt.Println(len(slice1)) //7
	fmt.Println(cap(slice1)) //7

	slice1 = append(slice1, 100)
	fmt.Println(slice1)      //[1 2 3 4 5 6 7 100]
	fmt.Println(len(slice1)) //8
	fmt.Println(cap(slice1)) //14 (double of size7 )
	fmt.Println(slice1)

}
