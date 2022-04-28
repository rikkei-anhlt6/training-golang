/*
	Trong Go, có 2 kiểu số phức là complex64 và complex128
	Complex64: Số phức có phần thực float32 và phần ảo
	Complex128: Số phức có phần thực float64 và phần ảo
	Khởi tạo
		func complex(r, i FloatType) ComplexType
		hoặc	c := a + bi
*/

package main

import (
	"fmt"
)

func main() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	c_add := c1 + c2
	fmt.Println("Tổng 2 số phức c1 và c2:", c_add)
	c_mul := c1 * c2
	fmt.Println("Tích 2 số phức c1 và c2:", c_mul)
}
