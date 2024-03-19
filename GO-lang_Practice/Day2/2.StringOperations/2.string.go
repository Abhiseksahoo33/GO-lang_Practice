package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Working on String Functions")

	str1, str2 := "String1", "String2"

	if str1 == str2 {
		fmt.Println("String are same")
	}

	if strings.Compare(str1, str2) == 0 {
		fmt.Println("Strings are same")
	}

	fmt.Println("Strings are same: ", strings.Compare(str1, str2)) //Strings are same:  -1
	if reflect.DeepEqual(str1, str2) {
		fmt.Println("Strings are equal")
	}

	str := "1,2,3,4,5,6,7,8,9"
	value := strings.Split(str, ",") // spite make a dynamic array with splite
	for _, v := range value {
		fmt.Println(v)
	}

	str = strings.Join(value, ",")
	fmt.Println(str) //1,2,3,4,5,6,7,8,9
}
