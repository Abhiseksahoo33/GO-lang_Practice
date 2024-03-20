package main

import (
	"first/geo"
	mmath "first/math"        // "module_name/package_name"
	adv "first/math/advanced" // adv is the package allias of advanced
	"fmt"
	"math"

	"github.com/fatih/color"
	. "github.com/fatih/color" // Bad Practice can use [Cyan("This is from external package")]
	// _ "anypackage name"  if you dont want the package to remove while saving either give alias or place _ prefix
)

func init() {
	fmt.Println("Inside main()")
}

func main() {
	fmt.Println("This is my Package and module Demo")
	color.Cyan("This is from external package") //will print the output in Cyan coloured text
	Cyan("This is from external package")       // Bad Practice
	fmt.Println("Result from math: ", math.Sin(0.123))

	fmt.Println(mmath.Add(1, 2))
	fmt.Println(mmath.Sub(10, 20))
	fmt.Println(mmath.Mul(100, 200))
	fmt.Println(mmath.Div(20, 2))

	fmt.Println(adv.Square(10))

	fmt.Println(geo.AsianCountries())

	mmath.MethodFromGeo()

	fmt.Println(mmath.PI)
	//fmt.Println(mmath.pI) #ERROR

	fmt.Println(mmath.CircleProperties(2.4))

	color.Red("This from Exit package") //will print the output in Red coloured text
}

/*

o/p

Inside geo
Inside math
Inside advanced
Inside main()
This is my Package and module Demo
Result from math:  0.12269009002431533
3
-10
20000
10
400
[India China SriLanka Bangladesh Bhutan Nepal Afganistan]
[India China SriLanka Bangladesh Bhutan Nepal Afganistan]
3.14159
18.095558399999998 15.079631999999998

Create module->  cd folder
				go mod init <module_name> -> it will create go.mod file inside the same dir
				create the main()
				go build .
				it creates a exe file called <module_name>
				./<module_name>
*/
