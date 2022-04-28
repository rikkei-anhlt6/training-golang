// Trong Go có 2 kiểu số thực là float32 và float64.
package main

import (
	"fmt"
)

func main() {
	a, b := 5.67, 8.97
	var c, d float32
	c, d = 5.67, 8.97
	sum := a + b
	sub := a - b
	fmt.Println("Tổng a và b:", sum)
	fmt.Println("Hiệu a và b:", sub)
	sum2 := c + d
	sub2 := c - d
	fmt.Println("Sum2:", sum2)
	fmt.Println("Sub2:", sub2)
}
