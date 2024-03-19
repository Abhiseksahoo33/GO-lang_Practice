package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Sort packages")

	v := []int{1, 6, 4, 3, 6, 2}

	fmt.Println(v) //[1 6 4 3 6 2]
	sort.Ints(v)
	fmt.Println(v) //[1 2 3 4 6 6]

	a := sort.SearchInts(v, 6)
	fmt.Println(a) //4th index
	b := sort.SearchInts(v, 100)
	fmt.Println(b) //if not found returns length

}
