/*
	declare channels: 	nameChannel := make(chan Type)
	data := <- a // read from channela a, đọc dữ liệu từ channel gán dữ liệu vào data, nhận
	a <- data // write to channel a, ghi dữ liệu data vào channel a, gửi
*/

package main

import (
	"fmt"
)

func number(result chan int) {
	result <- 10
}

func sendDataOnly(sendch chan<- int) {
	fmt.Println("sendonly")
	sendch <- 10
}

func receiveDataOnly(receivech <-chan int) {

}
func main() {
	rs := make(chan int)
	go number(rs)
	//a := <-rs
	fmt.Println(<-rs)

	//deadlock vì khi 1 goroutine gửi dữ liệu trên 1 kênh thì phải có ít nhất 1 go routine khác nhận dữ liệu trên kênh đấy
	ch := make(chan int)
	ch <- 5

	//Unidirectional channels: channels 1 chiều
	// chỉ truyền vào
	ch1 := make(chan<- int)
	go sendDataOnly(ch1)

	//chỉ đọc/ lấy ra
	ch2 := make(chan int, 1)
	ch2 <- 5 // write, gửi
	go receiveDataOnly(ch2)
	fmt.Println(<-ch2)
	fmt.Println("main function")

}
