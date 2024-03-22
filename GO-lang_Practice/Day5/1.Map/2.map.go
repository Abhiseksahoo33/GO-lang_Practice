package main

import "fmt"

func main() {
	fmt.Println("Maps..")

	url_store := map[int]string{0: "https://mav.com", 1: "http://media.com"}

	fmt.Println("value for key 0 is", url_store[0])

	// str := url_store[10]
	// fmt.Println(str) //empty

	str, ok := url_store[10]
	if !ok {
		fmt.Println("Invalid")
	}
	fmt.Println(str)

	mp := make(map[int]int)
	fmt.Println(mp)

	mp[0] = 1
	mp[2] = 1000
	mp[1] = 100
	mp[100] = 1

	fmt.Println(mp) //map[0:1 1:10 2:100 100:1]

	for k, v := range mp {
		fmt.Println("Key:", k, "Value:", v)
	}

}
