package main

import (
	"fmt"
)

func main() {
	first := "anh"
	last := "lt"
	name := first + " " + last
	fmt.Println("My name is", name)

	//  string literals
	a := `Say "hello" to Go!`
	fmt.Println(a)

	b := `Say "hello" to Go!\n`
	fmt.Println(b)

	// khai báo nhiều dòng string
	d := `This string is on 
	multiple lines
	within a single back 
	quote on either side.`
	fmt.Println(d)

	// in ra xâu có chứa " "
	f := "Say \"hello\" to Go!"
	fmt.Println(f)

	var string1 string
	string1 = "string 1"
	fmt.Println(string1)
}
