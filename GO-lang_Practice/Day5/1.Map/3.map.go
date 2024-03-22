package main

import (
	"fmt"
)

func main() {
	fmt.Println("Working with maps")
	var mp map[string]int

	// mp["Hello"] = 1   [CRASH,  not initialized]  writing will core
	fmt.Println(mp)

	k := mp["Hello"]
	fmt.Println(mp, k) // Reading is okay

	//#########################################

	fmt.Println("Deleting value from map")

	mp1 := make(map[int]int)

	//mp2 := map[int]int{0:1, 0:2}   [ERROR]

	mp1[0] = 1
	mp1[1] = 2
	mp1[3] = 3
	mp1[2] = 4

	fmt.Println(mp1) //map[0:1 1:2 2:4 3:3]

	delete(mp1, 2)
	fmt.Println(mp1) //map[0:1 1:2 3:3]

	//#########################################

	fmt.Println("Reseting Maps..")

	mp3 := map[int]int{0: 0, 1: 1, 2: 2, 3: 3}

	fmt.Println(mp3) //map[0:0 1:1 2:2 3:3]

	for k, _ := range mp3 {
		delete(mp3, k)
	}
	fmt.Println(mp3) //map[]

	mp3 = make(map[int]int) //It will also reset the map, Just create the map again
	fmt.Println(mp3)        //map[]

	//Maps are not Thread safe, If we are accessing it with 2 go routines it will crash so use mutex
	//[Cause it is always in heap, we can't access the heap simultaneously]

}
