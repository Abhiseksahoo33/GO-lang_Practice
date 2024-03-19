package main

import "fmt"

/*
func main() {
	fmt.Println("Working with structs")
	type Employee struct {
		EId        int
		Name, Dept string
	}

	var emp Employee = Employee{EId: 101, Name: "Abhisek", Dept: "AMF"}
	fmt.Println(emp) //{101 Abhisek AMF}
}


type PromptedFields struct {
	int
	uint8
	string
	float32
	bool
}

func main() {
	fmt.Println("Prompted fields..")

	var pf = PromptedFields{int: 1, uint8: 2, string: "Hello", float32: 3.2323, bool: true}
	fmt.Println(pf) //{1 2 Hello 3.2323 true}
}
*/

func main() {

	fmt.Println("Exploring annonymous structs")

	as := struct {
		int
		float32
		string
	}{int: 0, float32: 3.343, string: "hi"}
	fmt.Println(as) //{0 3.343 hi}

	a := as
	fmt.Println(as, a) //{0 3.343 hi}

}
