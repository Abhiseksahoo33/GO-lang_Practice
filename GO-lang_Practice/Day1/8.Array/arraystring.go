package main

import "fmt"

func main() {

	fmt.Println("Demo2: Array String")
	var array [7]string = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for idx, val := range array {
		fmt.Println(idx, val)
	}
}

/*
Demo2: Array String
0 Sunday
1 Monday
2 Tuesday
3 Wednesday
4 Thursday
5 Friday
6 Saturday
*/
