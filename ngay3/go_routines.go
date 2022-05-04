package main

import (
	"fmt"
	"time"
)

func printNum() {
	for i := 0; i <= 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func printChars() {
	for i := 'a'; i < 'f'; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	go printNum()
	go printChars()
	// time.Sleep(3000 * time.Millisecond)
	// fmt.Println("End main")
}

/*
num:			0		1	2	3		4	  5
millisecond:	200	300	400	600	800 900	1000 1200 1500	3000 end main
char:				a		b		c		  d		e

*/
