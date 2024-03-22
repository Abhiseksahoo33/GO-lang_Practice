package main

import (
	"fmt"
)

func main() {
	fmt.Println("Working with maps")
	var Map map[string]string = map[string]string{"WRG_URL": "https://wrg.mav.com",
		"MSTORE_URL": "https://mstore.mav.com", "MEDIA_URL": "https://media.mav.com"}

	for key, value := range Map {
		fmt.Println("Key: ", key, "\t Value:", value)
	}
}

/*
Working with maps
Key:  MSTORE_URL         Value: https://mstore.mav.com
Key:  MEDIA_URL          Value: https://media.mav.com
Key:  WRG_URL    Value: https://wrg.mav.com
*/
