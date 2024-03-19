package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Pointers..")

	i := 100
	var ptrInt *int = &i
	fmt.Println("value of i: ", i)                       //100
	fmt.Println("Address of i:", &i)                     //0xc000096068
	fmt.Println("contains of i:", ptrInt)                //0xc000096068
	fmt.Println("contains of i using pointer:", *ptrInt) //100

	fmt.Println("Pointer to Array")
	var ar [3]string = [3]string{"Hi", "Hello", "Abhisek"}

	ptrArray := &ar // ptrArray is pointer to string array
	ptrEle1 := &ar[0]

	fmt.Println(ptrArray)  //&[Hi Hello Abhisek] [Referenced value]
	fmt.Println(*ptrArray) //[Hi Hello Abhisek] [Value]
	fmt.Println(ptrEle1)   //0xc0000220f0

	*ptrEle1 = "This is modified"
	fmt.Println(ar) //[This is modified Hello Abhisek]

	fmt.Println(reflect.TypeOf(ptrEle1))  //*string
	fmt.Println(reflect.TypeOf(ptrArray)) //*[3]string
}
