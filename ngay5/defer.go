// Khi có nhiều defer sẽ thự hiện theo LIFO, defer nào được gọi cuối cùng sẽ được thực hiện trước
package main

import (
	"fmt"
)

func finish() {
	fmt.Println("Finished")
}

type person struct {
	name string
	age  int
}

func (p person) info() {
	fmt.Println(p.name, p.age)
}
func start() int {
	defer finish() // hàm finish sẽ được gọi trước khi hàm start trả về //defer function
	fmt.Println("Started")
	a := 5
	fmt.Println("trong start", a)
	return a
}
func main() {
	fmt.Println(start())
	p := person{"lta", 23}
	defer p.info() // defer method
	fmt.Println("End Main")

	b := 50
	defer fmt.Println(b) // giá trị của b là 50, được tính toán lúc hàm được gọi chứ ko phải lúc hàm được thực thi
	b = 100
	fmt.Println(b)

}
