// bản chất byte là uint8
// uint8	0 – 255
package main

import "fmt"

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func main() {
	var myByte byte = 1
	fmt.Println(myByte)
	fmt.Printf("%T", myByte)
	fmt.Println()
	var a byte = '?'
	fmt.Println(a)
	name := "Hello"
	printBytes(name)
}
