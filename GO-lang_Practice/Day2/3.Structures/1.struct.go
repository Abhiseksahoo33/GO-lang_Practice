package main

import "fmt"

type Student struct { // allocate memory in stack
	Id    int
	Name  string
	Class string
	Dept  string
}

func main() {
	var std Student
	fmt.Println(std) //{0   }

	std = Student{Id: 1, Name: "Abhisek Sahoo", Class: "MCA", Dept: "CSC"}
	fmt.Println(std) //{1 Abhisek Sahoo MCA CSC}

	fmt.Println(std.Id)
	fmt.Println(std.Name)
	fmt.Println(std.Class)
	fmt.Println(std.Dept)

	std.Id = 101
	fmt.Println(std.Id) //101
}
