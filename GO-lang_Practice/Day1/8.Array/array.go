package main

import "fmt"

func main() {
	fmt.Println("Demo: Array")
	var array [10]int
	fmt.Println(array) // [0 0 0 0 0 0 0 0 0 0]
	array[1] = 100
	fmt.Println(array) //[0 100 0 0 0 0 0 0 0 0]

	var array1 [5]int = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	array2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array2) //[1 2 3 4 5 6 7 8 9 10]

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array3)
	fmt.Println(len(array3)) //10

	fmt.Println("Demo: Array initialize at specific index")
	array4 := [...]int{0: 100, 4: 100, 9: 1000}
	fmt.Println(array4)      //[100 0 0 0 100 0 0 0 0 1000]
	fmt.Println(len(array4)) //10

	fmt.Println("Demo: Array range operator")
	array5 := [...]int{1: 100, 50: 10000, 3: 40, 10: 12345}
	for idx, val := range array5 {
		fmt.Println(idx, val)
	}

	fmt.Println("Demo2: Array range operator")
	array6 := [...]int{1: 100, 50: 10000, 3: 40, 10: 12345}
	for _, val := range array6 { //can use _ for idx
		fmt.Println(val)
	}

	fmt.Println("Demo2: Array filter")
	array7 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array7[:5])  //[1 2 3 4 5]
	fmt.Println(array7[0:3]) //[1 2 3]
	fmt.Println(array7[3:7]) //[4 5 6 7]
	fmt.Println(array7[7:])  //[8 9 10]  7th index to last index

	fmt.Println("Demo2: Array String")
}
