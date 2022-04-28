/*
	rune bản chất là int32
	Không quan trọng điểm mã chiếm bao nhiêu byte, nó có thể được biểu thị bằng rune
	Run đại diện unicode code point trong golang
	Rune dùng trong string là một dạng unicode
	Trong mã hóa UTF-8, một code point có thể chiếm nhiều hơn 1 byte
	UTF-8 mã hóa tất cả Unicode trong phạm vi từ 1 đến 4 byte
*/
package main

import (
	"fmt"
)

func main() {
	var str1 string = "Trường" //6 ký tự, 9 byte

	for index, c := range str1 {
		fmt.Printf("%c start at byte %d\n", c, index)
	} // c kiểu rune
	fmt.Println()
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c", str1[i]) // in ra TrÆ°á»ng
	}
	fmt.Println(len(str1))

	//Convert string qua Rune
	runes := []rune(str1)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i]) // in ra Trường
	}
}
