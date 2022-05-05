package main

import (
	"fmt"
	"sync"
)

func printNumber(wg *sync.WaitGroup) {
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d: ", i)
	}
	wg.Done()
}
func printChar(wg *sync.WaitGroup) {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c: ", i)
	}
	wg.Done()
}
func main() {
	//wg := sync.WaitGroup{}
	//var wg = sync.WaitGroup{} //contro
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumber(&wg)
	go printChar(&wg)
	wg.Wait() // Locking việc thực thi code ở hàm mial lại cho đến khi wg.Done() dc gọi ở Go routines
	fmt.Println("main finish")
}
