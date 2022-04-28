/*
	- 1 string là 1 slice của bytes trong Go
	- Kiểu dữ liệu string là bất biến, một khi string được tạo ra nó không thể bị thay đổi. Nếu muốn chúng ta chuyển string thành slice chứa các rune
	- Các chuỗi trong Go tuân thủ Unicode và được mã hóa UTF-8.
	Hàm  utf8.RuneCountInString(s))  thường được sử dụng để tìm độ dài của string
*/

package main

import (
	"fmt"
)

// thay đổi ký trong string theo index sử dụng rune
func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

//Truy cập từng byte của 1 string
func printByte(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("==========")
}

//Truy cập từng ký tự riêng lẻ của 1 string
func printChars(s string) {
	fmt.Print("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
	fmt.Println("==========")
}
func main() {
	str := "String 1"
	fmt.Printf("str: %s\n", str) //% s là mã định dạng để in một string
	printByte(str)
	printChars(str)

	// tạo string từ slice chứa các byte
	byteSlice := []byte{67, 97, 102, 195, 169}
	str1 := string(byteSlice)
	fmt.Println(str1)

	// Tạo string từ slice chứa các rune
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str2 := string(runeSlice)
	fmt.Println(str2)

	fmt.Println(mutate([]rune(str)))

}
