package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Exploring strings...")
	var testString string = "Wlcome to th GO lang"
	fmt.Println("TestString contains", testString)

	fmt.Println("length of string", len(testString)) // length of string 20

	fmt.Println("Declare string using SDO")
	myString := "Welcome Abhisek..."

	for i := 0; i < len(myString); i++ {
		fmt.Println(myString[i])                 // ASCII value of char
		fmt.Println(reflect.TypeOf(myString[i])) //uint8
	}

	for _, v := range myString {
		fmt.Println(v)
	}

	for i, _ := range myString {
		fmt.Println(i)
	}

	fmt.Println("This is my print:", myString[0])
	//myString[0] = 102      ERROR cannot assign to myString[0] (neither addressable nor a map index expression

	var testString1 = "Welcome to Golang training, \nWe are discussing strings"
	fmt.Println(testString1)

	fmt.Println("Exploring Strings: rawstring")
	var myString2 = `This is my test string,   //used for comments, documentation, http response 
	we are discussing rawstring here`
	fmt.Println(myString2)

}
