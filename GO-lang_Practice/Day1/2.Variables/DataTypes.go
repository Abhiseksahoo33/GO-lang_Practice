package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo: Variables")
	// var <variable_name> <type> = <value>   (type or value is optional one of these is mandatory)  default value = empty

	var var1 int8 = 123
	//var var1 int8 = 200 cannot use 200 (untyped int constant) as int8 value in variable declaration (overflows)
	fmt.Println("var1: ", var1)

	var var2 uint32 = 337
	fmt.Println("Var2: ", var2)

	var var3, var4 int = 337256, 2323
	fmt.Println("Var3: ", var3, ", var4: ", var4)

	var emptyVar int
	fmt.Println("empty Var: ", emptyVar)

	var c, d = 10, 20
	fmt.Println("Value of c and d", c, d)
	fmt.Printf("Type of c and d are %T %T \n", c, d)  //Type of c and d are int int
	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(d)) //int int

	var a = 'a'
	fmt.Println(a, reflect.TypeOf(a)) //97 int32

	var integer, strings = 100, "hello"
	fmt.Println(integer, strings) //100 hello

}

/*
Variables
~~~~~~~~~
1. Basic->	 			int(int8, int16, int32, int64, uint8, uint16, uint32, uint64, int, uint [32/64bits], rune(32 bit integer), byute,uintptr (32/64)),
						Float(float32, float64),
						Complex(complex64[realpart imaginary part], complex128),
						Boolean(8 bit),
						String(Imutable characters)
2. Aggregate Type->		Array,
						Struct
3. Reference Type->		Pointer,
						Maps(key-value pair),
						Slices,
						Function,
						Channel(Full-duplex, Read-write)
4. Interface Type->







*/
