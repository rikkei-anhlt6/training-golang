package main

import (
	"fmt"
)

//struct lồng nhau (Nested structs)

type Adrress struct {
	city  string
	state string
}
type Person struct {
	name   string
	age    int
	adress Adrress
}
type Person2 struct {
	name string
	age  int
	Adrress
}

func main() {
	p := Person{
		name: "lta",
		age:  18,
		adress: Adrress{
			city:  "a",
			state: "b",
		},
	}
	p4 := Person{
		name: "lta",
		age:  18,
		adress: Adrress{
			city:  "a",
			state: "b",
		},
	}
	p5 := Person{
		name: "lta",
		age:  18,
	}
	p1 := Person2{
		name: "lta2",
		age:  20,
		Adrress: Adrress{
			city:  "abc",
			state: "vcb",
		}, //Promoted
	}
	p3 := &Person{
		name: "lta3",
		age:  23,
		adress: Adrress{
			city:  "av",
			state: "bd",
		},
	}
	fmt.Println(p.adress)
	fmt.Println(p.adress.city)
	fmt.Println(p1.city) // city is promoted field

	fmt.Println((*p3).age) // Pointers to a struct

	/*Cấu trúc là các kiểu giá trị và có thể so sánh được nếu từng trường của chúng có thể so sánh được.
	Hai biến Struct được coi là bằng nhau nếu các trường tương ứng của chúng bằng nhau. */
	if p == p4 { //so sánh 2 struct có cùng kiểu
		fmt.Println("p và p4 có thể so sánh")
	}
	if p5 == p {
		fmt.Println("p và p5 có thể so sánh")
	} else {
		fmt.Println("p và p5 ko thể so sánh")
	}
}
