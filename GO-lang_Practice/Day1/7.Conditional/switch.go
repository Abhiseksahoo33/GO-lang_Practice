package main

import "fmt"

const (
	Sunday    int = iota // will start with 0
	Monday               //1
	Tuesday              //2
	Wednesday            //3..... so on
	Thursday
	Friday
	Saturday
)

func main() {
	/*
		var wday int
		fmt.Println("DEMO: Switch case")
		fmt.Println("Enter a day:")
		fmt.Scanf("%d", &wday)

		switch wday {
		case Sunday:
			fmt.Println("Sunday")
			fallthrough // will execute the nxt line also not break
		case Monday:
			fmt.Println("Monday")
		case Tuesday:
			fmt.Println("Tuesday")
		case Wednesday:
			fmt.Println("Wednesday")
		case Thursday, Friday, Saturday:
			fmt.Println("End of week ")
		default:
			fmt.Printf("NOT valid")
		}
	*/

	fmt.Println("DEMO2: Switch case")
	var Digit int
	fmt.Println("Enter a Digit:")
	fmt.Scanf("%d", &Digit)
	switch {
	case Digit <= 9:
		fmt.Println("Single digit")
	case Digit > 9 && Digit <= 99:
		fmt.Println("Tow digit")
	default:
		fmt.Println("Invalid")
	}
}
