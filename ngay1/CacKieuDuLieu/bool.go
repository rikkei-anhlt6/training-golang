package main

import (
	"fmt"
)

// Bool là kiểu dữ liệu chỉ nhận 2 giá trị hoặc true hoặc false

func main() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
}
