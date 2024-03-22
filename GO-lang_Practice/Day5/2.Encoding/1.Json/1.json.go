package main

import (
	"encoding/json"
	"fmt"
)

/*
JSON(REQUEST) -> BINARY -> JSON
XML -> BINARY -> XML  etc
*/

type Student struct {
	RollNo int    `json:"roll_no"` //json tag
	Name   string `json:"name"`
	Class  string `json:"class"`
}

func main() {
	fmt.Println("JSON encoding in GO")

	std := Student{RollNo: 1, Name: "Abhisek", Class: "MCA"}

	fmt.Println(std) //{1 Abhisek MCA}

	data, err := json.Marshal(std)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
		//[123 34 114 111 108 108 95 110 111 34 58 49 44 34 110 97 109 101 34 58 34 65 98 104 105 115 101 107 34 44 34 99 108 97 115 115 34 58 34 77 67 65 34 125]
		//{"roll_no":1,"name":"Abhisek","class":"MCA"} fmt.Println(string(data))
	}

	data, err = json.MarshalIndent(std, "\n", "\t")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

	/*
		{

		        "roll_no": 1,

		        "name": "Abhisek",

		        "class": "MCA"

		}  // Size is more than data
	*/
}
