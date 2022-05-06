package main

import (
	"fmt"
)

func recoverPerson() {
	if r := recover(); r != nil {
		fmt.Println("Recover from", r) // recover khôi phục quá trình thực thi bình thường và truy xuất lỗi đến hàm panic
	}
}
func person(name *string, age *int) {
	defer recoverPerson()
	if name == nil {
		panic("runtime error: name cannot be nil")
	}
	if age == nil {
		panic("runtime error: age cannot be nil")
	}
	fmt.Printf("%s %d\n", *name, *age)
}
func main() {
	defer fmt.Println("Defer call in main")
	name := "lta"
	//age := 5
	person(&name, nil)
	fmt.Println("End main")
}
