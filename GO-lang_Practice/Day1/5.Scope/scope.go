package main

import "fmt"

var global = 1000 // global variable

func main() {
	var a = 100    //Local to function
	fmt.Println(a) //local var
	fmt.Println(global)
	Test()

	{
		var c = 263
		fmt.Println(c) //Local to block
	}
}
func Test() {
	var b = 200
	fmt.Println(b)
	global = 2000
	fmt.Println(global)
}

/*
if variable starting with small letter is package label
if variable starting with capital letter accessible in all over the project

Global
package global (only package)
Local to function
Local to block

*/
