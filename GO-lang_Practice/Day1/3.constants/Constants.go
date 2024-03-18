package main

import "fmt"

func main() {
	fmt.Println("Demo: COnstants")
	// const <name> <type> = <vaule>  type is optional

	const c1 = 7
	fmt.Println(c1)

	const s, m, t, w, th, f, st = 1, 2, 3, 4, 5, 6, 7

	fmt.Println("Enter a week number")
	var week = 7

	fmt.Println("Week day is: ", week, week == s, week == st) //Week day is:  7 false true

	// const a,b int  ERROR
	const a, b int = 10, 20
	fmt.Println(a, b)

}
