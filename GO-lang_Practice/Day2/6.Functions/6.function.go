package main

import "fmt"

func HigherOrderFucntion(str string, fn func(bt rune)) {
	for _, v := range str {
		fn(v)
	}
}

func RuneReader(r rune) {
	fmt.Println(r)
}

func main() {
	fmt.Println("Higher order functions")

	var str = "This example is from Go training"
	HigherOrderFucntion(str, RuneReader)
	anfunc := func(r rune) {
		fmt.Println("Received", r)
	}

	HigherOrderFucntion(str, anfunc)

	HigherOrderFucntion(str, func(r rune) {
		fmt.Printf("We just got %c \n", r)
	})
}

/*
Higher order functions
84
104
105
115
32
101
120
97
109
112
108
101
32
105
115
32
102
114
111
109
32
71
111
32
116
114
97
105
110
105
110
103
Received 84
Received 104
Received 105
Received 115
Received 32
Received 101
Received 120
Received 97
Received 109
Received 112
Received 108
Received 101
Received 32
Received 105
Received 115
Received 32
Received 102
Received 114
Received 111
Received 109
Received 32
Received 71
Received 111
Received 32
Received 116
Received 114
Received 97
Received 105
Received 110
Received 105
Received 110
Received 103
We just got T
We just got h
We just got i
We just got s
We just got
We just got e
We just got x
We just got a
We just got m
We just got p
We just got l
We just got e
We just got
We just got i
We just got s
We just got
We just got f
We just got r
We just got o
We just got m
We just got
We just got G
We just got o
We just got
We just got t
We just got r
We just got a
We just got i
We just got n
We just got i
We just got n
We just got g
*/
