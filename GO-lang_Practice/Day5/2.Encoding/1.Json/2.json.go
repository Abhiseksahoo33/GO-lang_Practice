package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	RollNo int    `json:"roll_no"` //json tag
	Name   string `json:"name"`
	Class  string `json:"class"`
}

func main() {
	fmt.Println("Json encoding..")

	jsondata := `{

        "roll_no": 1,

        "name": "Abhisek",

        "class": "MCA"

}`

	fmt.Println(jsondata)
	data := []byte(jsondata)
	fmt.Println(data)

	std := Student{}
	//err := json.Unmarshal(data, std) ERROR [json: Unmarshal(non-pointer main.Student)] Pass by value
	err := json.Unmarshal(data, &std)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(std) //{1 Abhisek MCA}
	}
}
