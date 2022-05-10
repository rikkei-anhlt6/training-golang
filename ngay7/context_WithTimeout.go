package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	canceledChannel := make(chan bool)
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			canceledChannel <- true
			return
		}
	}()
	isCanceledChannel := <-canceledChannel
	if isCanceledChannel {
		close(canceledChannel)
		return
	}
	time.Sleep(time.Second * 10)
	fmt.Println("end")
}
func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	doSomething(ctx)
	/*
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		time.AfterFunc(time.Second, func() {
			cancel() // sau 1s tự động cancel, chủ động hơn trong cancel con text
		})
		defer cancel() // đảm bảo khi 1 func kết thúc thì call cancel
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // sau 3s in ra context deadline exceeded
		}
	*/
}
