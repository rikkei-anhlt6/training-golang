//Recover chỉ hđ khi nó dc gọi từ cùng go routines đang panic, không thể phục hồi từ panic đang xảy ra trong goroutine khác
package main

import (
	"fmt"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	defer recovery() // ko thể rocovery vì panic xảy ra ở goroutine divide
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(5, 0, done)
	<-done
}

func divide(a int, b int, done chan bool) {
	fmt.Printf("%d / %d = %d", a, b, a/b)
	done <- true

}

func main() {
	sum(5, 0)
	//ch := make(chan bool)

	fmt.Println("normally returned from main")
}
