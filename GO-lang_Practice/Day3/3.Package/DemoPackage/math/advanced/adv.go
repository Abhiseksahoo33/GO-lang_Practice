package advanced

import (
	"first/math"
	"fmt"
)

func init() {
	fmt.Println("Inside advanced")
}
func Square(a int) int {
	a = math.Add(a, a)
	return a * a
}
