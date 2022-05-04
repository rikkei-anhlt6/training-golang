package main

import (
	"fmt"
)

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func main() {
	var num int = 589
	ch := make(chan int)
	go digits(num, ch)
	for v := range ch {
		fmt.Println(v)
	}
}
