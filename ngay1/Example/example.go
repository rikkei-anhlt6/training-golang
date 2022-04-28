/* với các ký tự đặc biệt như ký tự quốc tế là 世, UTF-8 sẽ là bắt buộc. Go sử dụng kiểu bí danh rune cho dữ liệu UTF-8
/ Hàm utf8.RuneCountInString(s string) (n int) của gói UTF8 thường được sử dụng để tìm độ dài của string.
 Phương thức này coi string như một đối số và trả về số lượng rune trong đó. */
package main

import (
	"fmt"
	"unicode/utf8"
)

// Sử dụng index để lấy ra substring sử dụng rune
func printCharsAndByte(s string) {
	fmt.Printf("Characters: ")
	for index, runes := range s {
		fmt.Printf("%c starts at byte %d\n", runes, index)
	}
}
func main() {
	var str1 string = "ｿｹｯﾄを作成する" // len = 27, dùng UTF-8 để biểu diễn
	var str2 string = "abcdefghi" // len = 9, dùng ASCII để biểu diễn
	fmt.Println(len(str1))        // tính theo số byte sử dụng
	fmt.Println(len(str2))

	// tính độ dài string
	fmt.Println("Độ dài str1: ", utf8.RuneCountInString(str1)) // Trả về số lượng rune sử dụng
	fmt.Println("Độ dài str2: ", utf8.RuneCountInString(str2))
	fmt.Println()
	//In từng ký tự theo rune sử dụng
	printCharsAndByte(str1)
	fmt.Println()
	printCharsAndByte(str2)

}
