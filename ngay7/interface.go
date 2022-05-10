package main

import (
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int { //nếu ko phải struct thì phải dùng con trỏ
	*ic++
	return int(*ic)
}
func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World"))
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 5; i++ {
		fmt.Println(inc.Increment())
	}
}
