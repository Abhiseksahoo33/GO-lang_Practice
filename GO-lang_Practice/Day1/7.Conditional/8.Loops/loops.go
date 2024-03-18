package main

import "fmt"

func main() {
	// Only for loop. There is no do while, while loop in GO
	/*
				fmt.Println("Demo: Loops in GO")
				for i := 0; i < 4; i++ {
					fmt.Printf("Abhisek\n")
				}

				var a = 0
				for {
					fmt.Println(a)
					a++

					if a == 50 {
						break
					}
				}

		outer_loop:
			for i := 1; i < 100; i++ {
				for j := 1; j <= 10; j++ {
					fmt.Println(j)
					if j%8 == 0 {
						goto outer_loop // infinite loop
					}
				}

			}

	*/
outer_loop:
	for i := 1; i < 100; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(j)
			if j%8 == 0 {
				break outer_loop
			}
		}

	}

}
