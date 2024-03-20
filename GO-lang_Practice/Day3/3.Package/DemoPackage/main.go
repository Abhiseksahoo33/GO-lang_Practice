package main

import (
	"first/geo"
	mmath "first/math"        // "module_name/package_name"
	adv "first/math/advanced" // adv is the package allias of advanced
	"fmt"
	"math"
	// _ "anypackage name"  if you dont want the package to remove while saving either give alias or place _ prefix
)

func init() {
	fmt.Println("Inside main()")
}

func main() {
	fmt.Println("This is my Package and module Demo")

	fmt.Println("Result from math: ", math.Sin(0.123))

	fmt.Println(mmath.Add(1, 2))
	fmt.Println(mmath.Sub(10, 20))
	fmt.Println(mmath.Mul(100, 200))
	fmt.Println(mmath.Div(20, 2))

	fmt.Println(adv.Square(10))

	fmt.Println(geo.AsianCountries())

	mmath.MethodFromGeo()
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


Create module->  cd folder
				go mod init <module_name> -> it will create go.mod file inside the same dir
				create the main()
				go build .
				it creates a exe file called <module_name>
				./<module_name>
*/
