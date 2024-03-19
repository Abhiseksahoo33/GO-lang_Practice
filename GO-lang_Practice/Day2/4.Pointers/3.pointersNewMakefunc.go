package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
	Class  string
}

func main() {

	v := Student{RollNo: 1, Name: "Abhisek", Class: "xyz"} // stored in stack  we can allocate only 8mb in stack
	fmt.Println(v)                                         //{1 Abhisek xyz}

	p := new(Student) // new() used to allocate memory in the heap only for user defined types [for system defind type use make]
	p.RollNo = 1
	p.Class = "Test"
	p.Name = "Best"
	fmt.Println(*p) //{1 Best Test}

	fmt.Println("Exploring make ")
	// make only workes with container types

	intPtr := make([]int, 10)
	fmt.Println(intPtr) //[0 0 0 0 0 0 0 0 0 0]

	intPtr[0] = 100
	fmt.Println(intPtr) //[100 0 0 0 0 0 0 0 0 0]

	fmt.Println("Demo: Make: Array of pointers")

	v2 := make([]Student, 10)
	fmt.Println(v2) //[{0  } {0  } {0  } {0  } {0  } {0  } {0  } {0  } {0  } {0  }]

}
