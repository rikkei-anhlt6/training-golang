package main

import (
	"fmt"
)

func number(channel chan int) {
	for i := 0; i <= 5; i++ {
		channel <- i
	}
	close(channel) // Close sẽ ko có thêm dữ liệu nào dc gửi trên kênh
}

type chuoi []int

func main() {
	ch := make(chan int)
	go number(ch)
	for {
		// dùng biến ok trong khi nhận dữ liệu để xem dữ liệu để kiểm tra kênh đã bị đóng hay chưa
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
	// nhận dữ liệu từ kênh ch cho đến khi nó bị đóng
	ch1 := make(chan int)
	go number(ch1)
	for v := range ch1 {
		fmt.Println("Received ", v)
	}
}
