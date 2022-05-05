package main

import (
	"fmt"
)

type Info struct {
	name string
	age  int
}

// Anonymous fields
type Person struct {
	string
	int
}

func main() {
	// anonymous structs
	struct1 := struct {
		name string
		age  int
	}{
		name: "lta",
		age:  23,
	}
	fmt.Println(struct1)
	info1 := Info{"lta1", 23} // cách 1
	info2 := Info{
		age:  23,
		name: "lta2",
	}
	fmt.Println(info1.name) // truy cập các trường riêng lẻ trong struct
	fmt.Println(info2)

	p1 := Person{
		string: "person1",
		int:    12,
	}
	fmt.Println(p1.string, p1.int)
}
