package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time handling")

	//current time
	t := time.Now()
	fmt.Println(t) //2024-03-20 14:26:48.9173515 +0530 IST m=+0.014713401

	//constructing time
	t = time.Date(2024, 3, 20, 14, 35, 0, 0, &time.Location{})
	fmt.Println(t) //2024-03-20 14:35:00 +0000 UTC

	//Formatting time to string & string to time
	// yy/mm/dd hh:mi:ss
	layout := "2006/01/02 15:04:05"
	str := t.Format(layout)
	fmt.Println(str) //2024/03/20 14:35:00

	t1, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t1) //2024-03-20 14:35:00 +0000 UTC
	}

	//epoc linux time
	utime := t.Unix()
	fmt.Println(utime) //1710945300
	t2 := time.Unix(utime, 0)
	fmt.Println(t2) //2024-03-20 20:05:00 +0530 IST

	//Time addition and substraction
	t3 := t.Add(10 * time.Minute)
	fmt.Println(t3) //2024-03-20 14:45:00 +0000 UTC

	//t.Sub()
	//time.Since()

}
