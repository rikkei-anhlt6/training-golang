package main

import (
	"fmt"
	"unicode/utf8"
)

func catchuoi(start int, end int, s string) string {
	if start > utf8.RuneCountInString(s) || end > utf8.RuneCountInString(s) {
		return "ngoai pham vi"
	} else {
		str := []rune(s)
		slice_str := str[start : end+1]
		return string(slice_str)

	}

}
func main() {
	var s string = "a~bcd"
	var s1 string = "ｿｹｯﾄを作成する"
	fmt.Println(catchuoi(0, 10, s))
	fmt.Println(catchuoi(0, 7, s1))
}
