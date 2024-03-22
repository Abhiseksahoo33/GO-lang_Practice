package main

import (
	"encoding/xml"
	"fmt"
)

type Marks struct {
	S1    int     `xml:"Maths"` //xml tag
	S2    int     `xml:"Science"`
	S3    int     `xml:"ComputerScience"`
	Total int     `xml:"Total"`
	Avg   float64 `xml:"Average"`
}

type Student struct {
	RollNo int    `xml:"roll_no"`
	Name   string `xml:"name"`
	Class  string `xml:"class"`
	M      Marks  `xml:"marks"` // Sub type
}

func (m *Marks) SetMarks(s1, s2, s3 int) {
	m.S1 = s1
	m.S2 = s2
	m.S3 = s3
	m.Total = m.S1 + m.S2 + m.S3
	m.Avg = float64(m.Total) / 3
}

func main() {
	fmt.Println("XML encoding..")

	std := new(Student)
	std.RollNo = 1
	std.Name = "Abhisek"
	std.Class = "MCA"
	std.M.SetMarks(99, 100, 89)

	fmt.Println(std)

	data, err := xml.Marshal(std)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

	data, err = xml.MarshalIndent(std, "\n", "\t")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}

/*
XML encoding..
&{1 Abhisek MCA {99 100 89 288 96}}
<Student><roll_no>1</roll_no><name>Abhisek</name><class>MCA</class><marks><Maths>99</Maths><Science>100</Science><ComputerScience>89</ComputerScience><Total>288</Total><Average>96</Average></marks></Student>

<Student>

        <roll_no>1</roll_no>

        <name>Abhisek</name>

        <class>MCA</class>

        <marks>

                <Maths>99</Maths>

                <Science>100</Science>

                <ComputerScience>89</ComputerScience>

                <Total>288</Total>

                <Average>96</Average>

        </marks>

</Student>
*/
